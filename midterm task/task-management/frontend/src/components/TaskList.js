// src/components/TaskList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';

const TaskList = () => {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        const fetchTasks = async () => {
            const response = await axios.get('http://localhost:8000/tasks');
            setTasks(response.data);
        };

        fetchTasks();
    }, []);

    const deleteTask = async (id) => {
        await axios.delete(`http://localhost:8000/tasks/${id}`);
        setTasks(tasks.filter(task => task.id !== id));
    };

    return (
        <div>
            <h2>Task List</h2>
            <ul>
                {tasks.map(task => (
                    <li key={task.id}>
                        {task.title} - {task.done ? 'Done' : 'Pending'}
                        <button onClick={() => deleteTask(task.id)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TaskList;
