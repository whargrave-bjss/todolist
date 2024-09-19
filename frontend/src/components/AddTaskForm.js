const AddTaskForm = ({handleAddTask, newTask, setNewTask}) => {
    return (
    <form onSubmit={handleAddTask} className="mb-6">
    <div className="flex items-center border-b border-teal-500 py-2">
      <input 
        className="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
        type="text" 
        value={newTask}
        onChange={(e) => setNewTask(e.target.value)}
        placeholder="Enter a new task"
      />
      <button 
        className="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
        type="submit"
      >
        Add Task
      </button>
    </div>
  </form>
    )
}

export default AddTaskForm