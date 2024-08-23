import axios from 'axios'

document.addEventListener('DOMContentLoaded', function() {
    fetch('/tasks.json')
    .then(response => response.json())
    .then(data => {
        const taskList = document.getElementById('task-list')
        data.Tasks.forEach(task => {
            const li = document.createElement('li')
            li.textContent = task.Item
            if (task.Done) {
                li.style.textDecoration = 'line-through'
            }
            taskList.appendChild(li)
        })
    })
    .catch(error => console.error('Error loading tasks:', error))
})

document.addEventListener('DOMContentLoaded', function() {
    const addTaskForm = document.getElementById('add-task-form');
    const newTaskInput = document.getElementById('new-task');

    const addTask = async (task) => {
        try {
            const response = await fetch('http://localhost:3000/add-task', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ Item: task, Done: false })
            });
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return await response.json();
        } catch (error) {
            console.error('Error adding task:', error);
            throw error;
        }
    };

    addTaskForm.addEventListener('submit', async function(event) {
        event.preventDefault();
        const task = newTaskInput.value.trim();
        if (task) {
            try {
                await addTask(task);
                newTaskInput.value = '';
                loadTasks();
            } catch (error) {
                console.error('Error adding task:', error);
            }
        }
    });

    function loadTasks() {
        fetch('/tasks.json')
            .then(response => response.json())
            .then(data => {
                const taskList = document.getElementById('task-list');
                taskList.innerHTML = ''; // Clear existing tasks
                data.Tasks.forEach(task => {
                    const li = document.createElement('li');
                    li.textContent = task.Item;
                    if (task.Done) {
                        li.style.textDecoration = 'line-through';
                    }
                    taskList.appendChild(li);
                });
            })
            .catch(error => console.error('Error loading tasks:', error));
    }
    loadTasks();
});