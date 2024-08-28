export const updateTask = async (taskId) => {
    try {
        const response = await fetch(`http://localhost:3000/update-task/${taskId}`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Done: true })
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return await response.json();
    } catch (error) {
        console.error('Error updating task:', error);
        throw error;
    }
};
