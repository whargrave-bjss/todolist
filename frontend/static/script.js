import {addTask} from './apiservice/addTask.js';
import {deleteTask} from './apiservice/deleteTask.js';
import {updateTask} from './apiservice/updateTask.js';

document.addEventListener('DOMContentLoaded', function() {
    const addTaskForm = document.getElementById('add-task-form');
    const newTaskInput = document.getElementById('new-task');
    const taskList = document.getElementById('task-list');
    

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

    
    taskList.addEventListener('click', async function(event) {
        if (event.target.classList.contains('delete-button')) {
            const li = event.target.closest('li');
            console.log('Li element:', li); 
            const taskId = li.getAttribute('data-id'); 
            console.log('Task ID:', taskId); 
    
            if (taskId) {
                try {
                    await deleteTask(taskId);
                    loadTasks();
                } catch (error) {
                    console.error('Error deleting task:', error);
                }
            } else {
                console.error('Task ID not found');
            }
        }
    });

    taskList.addEventListener('click', async function(event) {
        if (event.target.classList.contains('update-button')) {
            const li = event.target.closest('li');
            console.log('Li element:', li); 
            const taskId = li.getAttribute('data-id'); 
            console.log('Task ID:', taskId); 
    
            if (taskId) {
                try {
                    await updateTask(taskId);
                    loadTasks();
                } catch (error) {
                    console.error('Error deleting task:', error);
                }
            } else {
                console.error('Task ID not found');
            }
        }
    });


function loadTasks() {
    fetch('/tasks.json')
        .then(response => response.json())
        .then(data => {
            const taskList = document.getElementById('task-list');
            taskList.innerHTML = '';
            data.Tasks.forEach(task => {
                const li = document.createElement('li');
                li.setAttribute('data-task-id', task.ID); 
                const taskText = document.createElement('span');
                taskText.textContent = task.Item;
                if (task.Done) {
                    taskText.style.textDecoration = 'line-through';
                }
                li.appendChild(taskText);

                const deleteButton = document.createElement('button');
                deleteButton.textContent = 'Delete';
                deleteButton.className = 'delete-button';
                li.appendChild(deleteButton);
                const updateButton = document.createElement('button');
                updateButton.textContent = 'Done';
                updateButton.className = 'update-button';
                li.appendChild(updateButton);

                taskList.appendChild(li);
            });
        })
        .catch(error => console.error('Error loading tasks:', error));
}
});