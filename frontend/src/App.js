import React, { useState, useEffect } from 'react';
import { addTask } from './apiservice/addTask';
import { deleteTask } from './apiservice/deleteTask';
import { updateTask } from './apiservice/updateTask';
import { fetchTasks } from './apiservice/fetchTasks';
import AddTaskForm from './components/AddTaskForm';
import LoginForm from './components/LoginForm';
import TaskList from './components/TaskList';
import './App.css';
const App = () => {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentUser, setCurrentUser] = useState(null);



  useEffect(() => {
    const loadTasks = async () => {
      try {
        const tasksData = await fetchTasks();
        setTasks(tasksData);
      } catch (error) {
        console.error('Error fetching tasks:', error);
      }
    }
    loadTasks();
  }, );

  useEffect(() => {
    if (currentUser) {
      fetchTasks().then(tasks => setTasks(tasks));
    }
  }, [currentUser]);

  const handleAddTask = async (e) => {
    e.preventDefault();
    if (newTask.trim()) {
      try {
        const addedTask = await addTask(newTask, currentUser.id);
        setTasks(prevTasks => {
        return Array.isArray(prevTasks) ? [...prevTasks, addedTask] : [addTask]; });
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

  const handleUpdateTask = async (id, currentStatus) => {
    try {
      const newStatus = !currentStatus;
      await updateTask(id, newStatus);
      setTasks(tasks.map(task => 
        task.ID === id ? { ...task, Done: newStatus } : task
      ));
    } catch (error) {
      console.error('Error updating task:', error);
    }
  }

  return (
    (isLoggedIn ? ( 
    <div className="min-h-screen bg-gradient-to-br from-purple-400 via-pink-500 to-red-500 py-8">
      <div className="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl">
        <div className="p-8">
          <h1 className="text-3xl font-bold text-center text-gray-800 mb-8">My Todo List</h1>
          <AddTaskForm handleAddTask={handleAddTask} newTask={newTask} setNewTask={setNewTask}/>
          <TaskList tasks={tasks} handleDeleteTask={handleDeleteTask} handleUpdateTask={handleUpdateTask}/>
        </div>
      </div>
    </div>
    ) : (
      <LoginForm setIsLoggedIn={setIsLoggedIn} setCurrentUser={setCurrentUser} />     
    ))
  );
}

export default App;