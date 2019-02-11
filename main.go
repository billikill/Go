package main

import (
"fmt"
"os"
"os/user"
)

func main() {
	userName, _ := user.Current()
	fmt.Println("Привет ",userName.Name)
	name, _ := os.Hostname()
	fmt.Println("Имя Хоста:", name)
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
}
