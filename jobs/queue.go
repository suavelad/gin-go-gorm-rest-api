package jobs

type TaskFunc func(interface{}) interface{}

// Task struct to hold information about the background task
type Task struct {
	Function TaskFunc
	Input    interface{}
	Result   chan interface{}
	Error    chan error
}

// Function to execute the task in the background

func ExecuteTask(task Task) {
	go task.Function(task.Input)

}
