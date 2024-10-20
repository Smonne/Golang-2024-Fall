// src/App.js
import React from 'react';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
    return (
        <Router>
            <div className="container">
                <h1>Task Management Application</h1>
                <Switch>
                    <Route path="/" exact>
                        <TaskForm fetchTasks={() => window.location.reload()} />
                        <TaskList />
                    </Route>
                </Switch>
            </div>
            <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/add" element={<AddTask />} />
            <Route path="/edit/:id" element={<EditTask />} />
            </Routes>
        </Router>
    );
}

export default App;
