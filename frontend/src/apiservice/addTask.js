export const addTask = async (task, userID) => {
    try {
        const response = await fetch('http://localhost:8080/api/add-task', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            },
            body: JSON.stringify({ Item: task, Done: false, UserID: userID }),
            credentials: 'include'
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const newTask = await response.json();
        return newTask
    } catch (error) {
        console.error('Error adding task:', error);
        throw error;
    }
};

