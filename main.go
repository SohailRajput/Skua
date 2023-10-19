package main

import (
	"fmt"

	"github.com/peterh/liner"
	"github.com/skua/hashtable"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	fmt.Println("Skua v0.0.1 - Key-Value Store cli")
	fmt.Printf("Max Key length: 256 bytes\nMax Value length: 1MB\nUsage:\n\t add <key> <value>\n")
	fmt.Println("\t update <key> <value>")
	fmt.Println("\t del <key>")
	fmt.Println("\t get <key>")
	fmt.Printf("Type 'exit' to quit\n\n")
	ht := hashtable.NewHashTable(16)
	for {
		cmd, err := line.Prompt("Skua> ")
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		switch cmd {
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			processCommand(cmd, ht)
		}
	}
}

func processCommand(cmd string, ht *hashtable.HashTable) {
	args := splitCommand(cmd)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "add":
		if len(args) == 3 {
			key := args[1]
			value := args[2]
			ht.Insert(key, value)
		} else {
			fmt.Println("Skua> Usage: add key value")
		}
	case "update":
		if len(args) == 3 {
			key := args[1]
			value := args[2]
			ht.Update(key, value)
		} else {
			fmt.Println("Skua> Usage: update key value")
		}
	case "get":
		if len(args) == 2 {
			key := args[1]
			v, e := ht.Get(key)
			if e {
				fmt.Printf("Skua> %s not found\n", key)
			}
			fmt.Println(v)
		} else {
			fmt.Println("Skua> Usage: get key")
		}
	case "del":
		if len(args) == 2 {
			key := args[1]
			ht.Delete(key)
		} else {
			fmt.Println("Skua> Usage: del key")
		}
	default:
		fmt.Printf("Skua> Unrecognized command: %s\n", args[0])
	}
}

func splitCommand(cmd string) []string {
	args := make([]string, 0)
	inQuotes := false
	currentArg := ""

	for _, char := range cmd {
		switch char {
		case ' ':
			if !inQuotes {
				args = append(args, currentArg)
				currentArg = ""
			} else {
				currentArg += string(char)
			}
		case '"':
			inQuotes = !inQuotes
		default:
			currentArg += string(char)
		}
	}

	if currentArg != "" {
		args = append(args, currentArg)
	}

	return args
}
