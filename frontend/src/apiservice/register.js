export const register = async (username, password) => {
    try {
        const response = await fetch('http://localhost:8080/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ Username: username, Password: password }),
            credentials: 'include'
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const user = await response.json();
        return user
    } catch (error) {
        console.error('Error adding task:', error);
        throw error;
    }
};