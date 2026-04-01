package main
import "fmt"

func main() {
    tasks := []string{"Study", "Code"}
    done := make(map[int]bool)

    done[1] = true

    for i, t := range tasks {
        if done[i] {
            fmt.Println(t, "(Done)")
        } else {
            fmt.Println(t)
        }
    }
}
