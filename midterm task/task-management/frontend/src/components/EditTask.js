import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';

const EditTask = () => {
  const { id } = useParams();
  const [task, setTask] = useState('');
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchTask = async () => {
      try {
        const response = await axios.get(`/api/tasks/${id}`);
        setTask(response.data.name);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching task:', error);
        alert('Failed to fetch task.');
      }
    };

    fetchTask();
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.put(`/api/tasks/${id}`, { name: task });
      alert('Task updated successfully!');
    } catch (error) {
      console.error('Error updating task:', error);
      alert('Failed to update task.');
    }
  };

  if (loading) return <div>Loading...</div>;

  return (
    <div>
      <h2>Edit Task</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={task}
          onChange={(e) => setTask(e.target.value)}
          placeholder="Task name"
          required
        />
        <button type="submit">Update Task</button>
      </form>
    </div>
  );
};

export default EditTask;
