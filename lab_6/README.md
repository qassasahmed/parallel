
# Lab 6: WorkerPool Part II
Made with ❤️ and ☕
---

## `Task` Struct
The `Task` struct represents a unit of work to be processed. It contains the following fields:
- `ID` (int): A unique identifier for the task.
- `ProcessTime` (int): The amount of time (in milliseconds) the task takes to process.

### Code:
```go
// Define Task struct
type Task struct {
    ID          int
    ProcessTime int
}

// Define a method of the Task struct
func (t *Task) Process() {
    fmt.Printf("Processing Task %d...\n", t.ID)
    time.Sleep(time.Duration(t.ProcessTime)*time.Second)
}
```

### Role:
The `Task` struct encapsulates the data and behavior for individual units of work. Its `Process` method simulates processing the task by sleeping for the specified `ProcessTime`.

---

## `WorkerPool` Struct

The `WorkerPool` struct manages the concurrent processing of tasks. It contains the following fields:
- `Tasks` ([]Task): A slice of tasks to be processed.
- `NumberOfWorkers` (int): The number of workers available for processing.
- `TasksChannel` (chan Task): A channel for distributing tasks to workers.
- `WG` (sync.WaitGroup): A synchronization primitive to wait for all workers to complete their tasks.

### Code:
```go
// Define WP Struct
type WorkerPool struct {
    Tasks           []Task
    NumberOfWorkers int
    TasksChannel    chan Task
    WG              sync.WaitGroup
}

// Define a method on the WorkerPool struct for workers
func (w *WorkerPool) Worker() {
    for task := range w.TasksChannel {
        task.Process()
        w.WG.Done()
    }
}

// Start the worker pool and distribute tasks
func (w *WorkerPool) Pool() {
    for i := 0; i < w.NumberOfWorkers; i++ {
        go w.Worker()
    }

    for _, task := range w.Tasks {
        w.WG.Add(1)
        w.TasksChannel <- task
    }

    close(w.TasksChannel)
    w.WG.Wait()
}
```

### Role:
The `WorkerPool` struct is responsible for managing concurrent task processing:
- The `Worker` method processes tasks received from the `TasksChannel`.
- The `Pool` method initializes workers, assigns tasks, and waits for completion.

---


## The `main()` function
```go
func main(){

	tasks := make([]Task, 20)
	for i:=0; i < cap(tasks); i++{
		tasks[i] = Task{
			ID: i+1,
			ProcessTime: 1,
            }
	}

	wp := WorkerPool{
		Tasks: tasks,
		NumberOfWorkers: 5,
	}

	wp.Pool()
    
	fmt.Println("All tasks completed")
}

```