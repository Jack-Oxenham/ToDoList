package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tasks []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewTasks(tasks)
		case 2:
			fmt.Print("Enter a new task: ")
			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Task added.")
		case 3:
			viewTasks(tasks)
			if len(tasks) == 0 {
				break
			}
			fmt.Print("Enter task number to delete: ")
			var index int
			fmt.Scanln(&index)
			if index > 0 && index <= len(tasks) {
				tasks = append(tasks[:index-1], tasks[index:]...)
				fmt.Println("Task deleted.")
			} else {
				fmt.Println("Invalid task number.")
			}
		case 4:
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func viewTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
