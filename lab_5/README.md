
# Parallel Computing Lab 5 Notes - Autumn 2024-25

This document consolidates the concepts and examples from Lab 5 and previous labs in the Parallel Computing course.

## Review of Previous Labs

### Lab 1: Basics of Golang and Concurrency

#### Example: Running Functions Sequentially and Concurrently
This example demonstrates the difference between running functions sequentially and concurrently, while also measuring execution time.

```go
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

Thank you for reviewing the content from Lab 5 and previous labs.
