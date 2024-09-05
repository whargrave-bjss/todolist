export const deleteTask = async (id) => {
    const response = await fetch(`http://localhost:8080/api/delete-task/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include',
    });

    if (!response.ok) {
        if (response.status === 400) {
            const errorData = await response.json();
            throw new Error(`Failed to delete task: ${errorData.error}`);
        }
        throw new Error(`Failed to delete task: ${response.statusText}`);
    }

    const result = await response.json();
    return result;
};

