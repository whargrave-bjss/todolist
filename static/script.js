import { addTask } from "../apiservice/addTask";
import { deleteTask } from "../apiservice/deleteTask";

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
        if  (event.target.classList.contains('delete-button')) {
            const taskId = event.target.dataset.taskId;
            try {
                await deleteTask(taskId);
                loadTasks();
            } catch (error) {
                console.error('Error deleting task:', error);
            }
        }
    });
    loadTasks(); 
});


    function loadTasks() {
        fetch('/tasks.json')
            .then(response => response.json())
            .then(data => {
                const taskList = document.getElementById('task-list');
                taskList.innerHTML = ''; // Clear existing tasks
                data.Tasks.forEach(task => {
                    const li = document.createElement('li');
                    const taskText = document.createElement('span');
                    taskText.textContent = task.Item;
                    if (task.Done) {
                        li.style.textDecoration = 'line-through';
                    }
                    li.appendChild(taskText);

                    const deleteButton = document.createElement('button');
                    deleteButton.textContent = 'Delete';
                    deleteButton.className = 'delete-button';
                    deleteButton.dataset.taskId = task.ID;
                    li.appendChild(deleteButton);
                    taskList.appendChild(li);
                });
            })
            .catch(error => console.error('Error loading tasks:', error));
    }