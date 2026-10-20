package main

import "fmt"

func main() {
	todos := []string{}

	for {
		fmt.Println("\n===============================")
		fmt.Println("     ***** TODO LIST *****     ")
		fmt.Println("===============================")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Exit")
		fmt.Println("------------------------------")

		var ch, n int
		fmt.Print("Enter choice: ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			var task string
			fmt.Print("Enter task : ")
			fmt.Scan(&task)
			todos = append(todos, task)
			fmt.Println("***** Task Added *****")

		case 2:
			fmt.Println("----- TODO LIST -----")

			if len(todos) > 0 {
				for i, t := range todos {
					fmt.Println(i+1, ".", t)
				}
			} else {
				fmt.Println("Your Task is empty")
			}

		case 3:
			fmt.Print("Enter task number: ")
			fmt.Scan(&n)
			if n > 0 && n <= len(todos) {
				fmt.Print("Enter new task: ")
				fmt.Scan(&todos[n-1])
				fmt.Println("***** Task Updated *****")
			}

		case 4:
			fmt.Print("Enter task number: ")
			fmt.Scan(&n)
			if n > 0 && n <= len(todos) {
				todos = append(todos[:n-1], todos[n:]...)
				fmt.Println("***** Task Deleted *****")
			}

		case 5:
			fmt.Println("===== Thank You =====")
			return

		default:
			fmt.Println("----- Invalid Choice -----")
		}
	}
}
