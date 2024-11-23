
# Parallel Computing Lab 5 Notes - Autumn 2024-25

This document contains the concepts and examples from the four previous labs.

## Review of Previous Labs

### Lab 1: Basics of Golang and Concurrency

#### Example: Running Functions Sequentially and Concurrently
This example demonstrates the difference between running functions sequentially and concurrently, while also measuring execution time.

```go
package main

import (
    "fmt"
    "time"
)

func printNumbers(num int) {
    for i := 1; i <= num; i++ {
        fmt.Printf("%d ", i)
        time.Sleep(1000 * time.Millisecond)
    }
}

func printLetters(char rune) {
    for i := 'A'; i <= char; i++ {
        fmt.Printf("%c ", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    now := time.Now()
    defer func() {
        fmt.Println("Execution time =", time.Since(now))
    }()
    printNumbers(5)
    printLetters('E')
}
```

---

### Lab 2: Channels and Buffered Channels

#### Example: Preventing Application Closure with Channels
This program calculates the factorial of a number while simulating a file copy process in another goroutine.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch_signal := make(chan bool)
    go copy_simulation(20, ch_signal)

    var input_num int
    fmt.Printf("Enter Number: ")
    fmt.Scan(&input_num) // Use error handling

    answer := factorial(input_num)
    fmt.Printf("%d! = %d\n", input_num, answer)

    <-ch_signal
}

func copy_simulation(num int, done chan bool) {
    for i := 1; i <= num; i++ {
        fmt.Printf("-")
        time.Sleep(1000 * time.Millisecond)
    }
    done <- true
}

func factorial(num int) int {
    fact := 1
    for i := 2; i <= num; i++ {
        fact *= i
    }
    return fact
}
```

---

### Lab 3: Sending Data Between Goroutines

#### Example: Concurrent Loops with Channel Communication
The following program uses a channel to communicate between goroutines and implements pseudo-random number generation.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func getMoney(ch chan int) {
    var trials int
    fmt.Println("Number of trials: ")
    if _, err := fmt.Scan(&trials); err != nil {
        fmt.Println("Please Provide an Integer number")
        return
    }

    rand.Seed(time.Now().UnixNano())
    for i := 1; i <= trials; i++ {
        amount := rand.Intn(500)
        ch <- amount
    }
    close(ch)
}

func main() {
    channel := make(chan int)
    go getMoney(channel)

    for msg := range channel {
        fmt.Printf("You've won: %d$\n", msg)
    }
}
```

---

### Lab 4: Waitgroups, Slices, and Structs

#### Example 1: Synchronizing Goroutines with Waitgroups

```go
package main

import (
    "fmt"
    "sync"
)

func printNumber(num int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println(num)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go printNumber(i, &wg)
    }
    wg.Wait()
}
```

#### Example 2: Slices and Arrays

```go
package main

import (
    "fmt"
)

func main() {
    number := [5]int{1, 2, 3, 4, 5}
    fmt.Println(number) // [1, 2, 3, 4, 5]

    slice := number[2:4] // len = 2, cap = 3
    fmt.Println("Cap:", cap(slice))

    slice[0] = -10
    fmt.Println("new Slice", slice)  // [-10, 4]
    fmt.Println("new Array", number) // [1, 2, -10, 4, 5]
}
```

#### Example 3: Structs

```go
package main

import (
    "fmt"
)

func main() {
    type Student struct {
        courses [5]string
        grades  [5]float64
    }

    qassas := Student{
        courses: [5]string{"DB", "IT", "AI", "ML", "CN"},
        grades:  [5]float64{85.0, 90.0, 92.0, 95.5, 100.0},
    }

    fmt.Println("course:", qassas.courses)
}
```
---

# Worker Pool in Go

A **worker pool** is a concurrency pattern used to manage and distribute tasks among multiple worker goroutines efficiently. It helps optimize resource usage by limiting the number of active workers while processing a potentially large number of tasks.  

## Why Use a Worker Pool?
1. **Resource Management**: Prevents overwhelming the system by limiting the number of goroutines.
2. **Parallelism**: Efficiently processes tasks in parallel.
3. **Control**: Enables better control over task execution and resource utilization.  

## How a Worker Pool Works
1. **Task Queue**: A channel or buffer holds the tasks to be processed.
2. **Workers**: Goroutines (workers) are launched to pull tasks from the queue.
3. **Results**: Optionally, results can be sent back via another channel.
  

## Example: Basic Worker Pool

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Worker function
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for task := range tasks {
        fmt.Printf("Worker %d processing task %d\n", id, task)
        time.Sleep(time.Second) // Simulate task processing time
    }
}

func main() {
    const numWorkers = 3
    const numTasks = 10

    tasks := make(chan int, numTasks)
    var wg sync.WaitGroup

    // Start workers
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, &wg)
    }

    // Send tasks to the task queue
    for i := 1; i <= numTasks; i++ {
        tasks <- i
    }
    close(tasks) // Close the task channel to signal no more tasks

    // Wait for all workers to finish
    wg.Wait()

    fmt.Println("All tasks processed.")
}
```

## Output
```
Worker 1 processing task 1
Worker 2 processing task 2
Worker 3 processing task 3
Worker 1 processing task 4
Worker 2 processing task 5
Worker 3 processing task 6
Worker 1 processing task 7
Worker 2 processing task 8
Worker 3 processing task 9
Worker 1 processing task 10
All tasks processed.
```

## Key Points in the Example
1. **Task Channel**: Tasks are sent into the `tasks` channel.
2. **Workers**: Three workers process tasks concurrently.
3. **Synchronization**: The `sync.WaitGroup` ensures all workers complete before exiting the program.


## Use Cases for Worker Pools
1. **Web Servers**: Handling multiple client requests.
2. **Task Queues**: Processing jobs from a queue in parallel.
3. **File Processing**: Reading or writing files concurrently.
4. **Batch Operations**: Running computational tasks or API calls in parallel.

---

Thank You ☕