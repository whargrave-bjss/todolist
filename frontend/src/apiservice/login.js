export const login = async (username, password) => {
    try {
        const response = await fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
            credentials: 'include'
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();

        localStorage.setItem('token', data.token);
        localStorage.setItem('userId', data.id)
        localStorage.setItem('username', data.Username)
        return data
    } catch (error) {
        console.error('Error adding task:', error);
        throw error;
    }
};