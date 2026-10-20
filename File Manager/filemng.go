package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n================================")
		fmt.Println("    ***** FILE MANAGER *****    ")
		fmt.Println("================================")
		fmt.Println("1. Create")
		fmt.Println("2. List")
		fmt.Println("3. Read")
		fmt.Println("4. Update")
		fmt.Println("5. Delete")
		fmt.Println("6. Exit")
		fmt.Println("------------------------------")

		var ch int
		fmt.Print("Enter choice: ")
		fmt.Fscan(r, &ch)
		r.ReadString('\n')

		switch ch {

		case 1:
			fmt.Print("Enter file name: ")
			name, _ := r.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter content: ")
			content, _ := r.ReadString('\n')

			os.WriteFile(name+".txt", []byte(content), 0644)
			fmt.Println("***** File Created *****")

		case 2:
			fmt.Println("----- FILES -----")
			files, _ := os.ReadDir(".")

			found := false
			for _, f := range files {
				if !f.IsDir() && filepath.Ext(f.Name()) == ".txt" {
					fmt.Println(f.Name())
					found = true
				}
			}

			if !found {
				fmt.Println("There are no files")
			}

		case 3:
			fmt.Print("Enter file name: ")
			name, _ := r.ReadString('\n')
			name = strings.TrimSpace(name)

			data, err := os.ReadFile(name + ".txt")
			if err != nil {
				fmt.Println("File not found!")
			} else {
				fmt.Println("----- Content -----")
				fmt.Print(string(data))
			}

		case 4:
			fmt.Print("Enter old file name: ")
			old, _ := r.ReadString('\n')
			old = strings.TrimSpace(old)

			fmt.Print("Enter new file name: ")
			newName, _ := r.ReadString('\n')
			newName = strings.TrimSpace(newName)

			fmt.Print("Enter new content: ")
			content, _ := r.ReadString('\n')

			os.Rename(old+".txt", newName+".txt")
			os.WriteFile(newName+".txt", []byte(content), 0644)
			fmt.Println("***** File Updated *****")

		case 5:
			fmt.Print("Enter file name: ")
			name, _ := r.ReadString('\n')
			name = strings.TrimSpace(name)

			os.Remove(name + ".txt")
			fmt.Println("***** File Deleted *****")

		case 6:
			fmt.Println("===== Thank You =====")
			return

		default:
			fmt.Println("----- Invalid Choice -----")
		}
	}
}
