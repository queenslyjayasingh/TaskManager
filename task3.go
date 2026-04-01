package main
import "fmt"

type Task struct {
    name string
    priority int
}

func main() {
    tasks := []Task{
        {"Study", 2},
        {"Project", 1},
    }

    for _, t := range tasks {
        fmt.Println(t.name, "Priority:", t.priority)
    }
}
