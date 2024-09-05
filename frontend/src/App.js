import React, { useState, useEffect } from 'react';
import { addTask } from './apiservice/addTask';
import { deleteTask } from './apiservice/deleteTask';
import { updateTask } from './apiservice/updateTask';
import './App.css';
const App = () => {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState('');

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/tasks');
      if (!response.ok) {
        throw new Error('Failed to fetch tasks');
      }
      const data = await response.json();
      setTasks(data);
    } catch (error) {
      console.error('Error fetching tasks:', error);
    }
  };

  const handleAddTask = async (e) => {
    e.preventDefault();
    if (newTask.trim()) {
      try {
        const addedTask = await addTask(newTask);
        setTasks([...tasks, addedTask]);
        setNewTask('');
      } catch (error) {
        console.error('Error adding task:', error);
      }
    }
  }

  const handleDeleteTask = async (id) => {
    try {
      const result = await deleteTask(id);
      console.log('Delete task result:', result);
      setTasks(tasks.filter(task => task.ID !== id));
    } catch (error) {
      console.error('Error deleting task:', error.message);
    }
  }

  const handleUpdateTask = async (id, currentDoneStatus) => {
    try {
      const updatedTask = await updateTask(id, !currentDoneStatus);
      setTasks(tasks.map(task => 
        task.ID === id ? { ...task, Done: updatedTask.Done } : task
      ));
    } catch (error) {
      console.error('Error updating task:', error);
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-400 via-pink-500 to-red-500 py-8">
      <div className="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl">
        <div className="p-8">
          <h1 className="text-3xl font-bold text-center text-gray-800 mb-8">My Todo List</h1>
          
          <form onSubmit={handleAddTask} className="mb-6">
            <div className="flex items-center border-b border-teal-500 py-2">
              <input 
                className="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
                type="text" 
                value={newTask}
                onChange={(e) => setNewTask(e.target.value)}
                placeholder="Enter a new task"
              />
              <button 
                className="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
                type="submit"
              >
                Add Task
              </button>
            </div>
          </form>

          <ul className="divide-y divide-gray-200">
            {Array.isArray(tasks) && tasks.map(task => (
              <li key={task.Id} className="py-4">
                <div className="flex items-center justify-between">
                  <span className={`text-lg ${task.Done ? 'line-through text-gray-500' : 'text-gray-800'}`}>
                    {task.Item}
                  </span>
                  <div>
                    <button 
                      onClick={() => handleUpdateTask(task.ID, task.Done)}
                      className={`mr-2 px-3 py-1 rounded ${task.Done ? 'bg-yellow-500 hover:bg-yellow-600' : 'bg-green-500 hover:bg-green-600'} text-white`}
                    >
                      {task.Done ? 'Undo' : 'Done'}
                    </button>
                    <button 
                      onClick={() => handleDeleteTask(task.ID)}
                      className="px-3 py-1 bg-red-500 hover:bg-red-600 text-white rounded"
                    >
                      Delete
                    </button>
                  </div>
                </div>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default App;