export const addTask = async (task) => {
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