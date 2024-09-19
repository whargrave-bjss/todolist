const TaskList = ({tasks, handleUpdateTask, handleDeleteTask}) => {
    return (
        <ul className="divide-y divide-gray-200">
            {Array.isArray(tasks) && tasks.map(task => (
              <li key={task.ID} className="py-4">
                <div className="flex items-center justify-between">
                  <span className={`text-lg ${task.Done ? 'line-through text-gray-500' : 'text-gray-800'}`}>
                    {task.Item}
                  </span>
                  <div>
                    <button 
                      onClick={() => handleUpdateTask(task.ID, task.Done)}
                      className={`mr-2 px-3 py-1 rounded ${task.Done ? 'bg-yellow-500 hover:bg-yellow-600' : 'bg-green-500 hover:bg-green-600'} text-white`}
                    >
                      {task.Done ? 'Undo' : 'Done'}
                    </button>
                    <button 
                      onClick={() => handleDeleteTask(task.ID)}
                      className="px-3 py-1 bg-red-500 hover:bg-red-600 text-white rounded"
                    >
                      Delete
                    </button>
                  </div>
                </div>
              </li>
            ))}
          </ul>
    )
}

export default TaskList