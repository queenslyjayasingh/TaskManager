package main
import "fmt"

func main() {
    var tasks []string
    var choice int
    var task string

    fmt.Println("1.Add Task 2.View Tasks")
    fmt.Scan(&choice)

    if choice == 1 {
        fmt.Scan(&task)
        tasks = append(tasks, task)
        fmt.Println("Task Added")
    } else {
        for _, t := range tasks {
            fmt.Println(t)
        }
    }
}
