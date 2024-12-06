package main

import(
	"fmt"
	"time"
	"sync"
)

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
    w.TasksChannel = make(chan Task, len(w.Tasks))
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