export const deleteTask = async (taskId) => {
    try {
        const response = await fetch(`http://localhost:3000/delete-task/${taskId}`, {
            method: 'DELETE',
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return await response.json();
    } catch (error) {
        console.error('Error deleting task:', error);
        throw error;
    }
};
    