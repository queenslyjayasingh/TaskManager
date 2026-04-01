package main

import "fmt"

// Struct for Level 3 (priority)
type Task struct {
	name     string
	priority int
}

func main() {
	var choice int

	// Data structures
	var tasks []Task
	done := make(map[int]bool)

	for {
		fmt.Println("\n==== TASK MANAGER ====")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")

		fmt.Scan(&choice)

		switch choice {

		// 🔹 LEVEL 1 FEATURE: Add Task
		case 1:
			var name string
			var priority int

			fmt.Print("Enter task name: ")
			fmt.Scan(&name)

			fmt.Print("Enter priority (1-High, 2-Medium, 3-Low): ")
			fmt.Scan(&priority)

			tasks = append(tasks, Task{name, priority})
			fmt.Println("✅ Task Added")

		// 🔹 LEVEL 1 + LEVEL 3: View Tasks with Priority
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks available")
				continue
			}

			fmt.Println("\n--- Task List ---")
			for i, t := range tasks {
				status := ""
				if done[i] {
					status = "(Done)"
				}

				fmt.Printf("%d. %s | Priority: %d %s\n", i, t.name, t.priority, status)
			}

		// 🔹 LEVEL 2 FEATURE: Mark Task as Done
		case 3:
			var index int
			fmt.Print("Enter task index to mark as done: ")
			fmt.Scan(&index)

			if index >= 0 && index < len(tasks) {
				done[index] = true
				fmt.Println("✔ Task marked as done")
			} else {
				fmt.Println("Invalid index")
			}

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
