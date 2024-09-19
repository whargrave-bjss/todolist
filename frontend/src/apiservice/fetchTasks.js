
export const fetchTasks = async () => {
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        throw new Error('No authentication token found :(');
      }
      const response = await fetch('http://localhost:8080/api/tasks', {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        },
        credentials: 'include'
      });
      if (!response.ok) {
        throw new Error('Failed to fetch tasks');
      }
      const tasks = await response.json();
      return tasks;
  } catch (error) {
    console.error('Error fetching tasks:', error);
    throw error 
  }
};