package handlers

import (
    "fmt"
    "net/http"
	
    "time"
    "go-security-app/utils"
    "github.com/gorilla/mux"
	"encoding/json"

	"github.com/gorilla/csrf"
    "go-security-app/utils"
	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
	
   

)

var requestCount = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Number of HTTP requests",
    },
    []string{"path"},
)

var (
    httpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path"},
    )
)

func SetupRoutes() *mux.Router {
    prometheus.MustRegister(httpRequests)

    router := mux.NewRouter()
    router.Use(RequestLogger)
    router.HandleFunc("/", HomeHandler).Methods("GET")
    router.Handle("/metrics", promhttp.Handler()).Methods("GET") // Prometheus endpoint
    return router
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Secure Go Web App!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is a demo application.")
}


type UserInput struct {
    Name  string `json:"name" validate:"required,min=3,max=50"`
    Email string `json:"email" validate:"required,email"`
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
    var input UserInput

   
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := utils.ValidateStruct(&input); err != nil {
        http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Input is valid!"))
}


func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        claims, err := auth.ValidateJWT(tokenStr)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Attach user info to the context
        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}



func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value("user").(*auth.Claims)
    w.Write([]byte("Welcome " + user.Username))
}


// func SetupRoutes() *mux.Router {
//     csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
//     router := mux.NewRouter()
//     router.HandleFunc("/", HomeHandler).Methods("GET")
//     router.Use(csrfMiddleware)
//     return router
// }
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Security-Policy", "default-src 'self'")
        w.Header().Set("X-Frame-Options", "DENY")
        next.ServeHTTP(w, r)
    })
}
var requestCount = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Number of HTTP requests",
    },
    []string{"path"},
)
func ErrorHandler(w http.ResponseWriter, err error) {
    utils.Logger.Error("Error occurred", zap.Error(err))
    http.Error(w, "Something went wrong", http.StatusInternalServerError)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    utils.Logger.Info("Home handler accessed")
    w.Write([]byte("Welcome to the secure Go app!"))
}

func RequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)

        httpRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

        utils.Logger.Info("HTTP Request",
            zap.String("method", r.Method),
            zap.String("path", r.URL.Path),
            zap.Int("status", http.StatusOK),
            zap.Duration("duration", duration),
        )
    })
}

