package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	title, _ := askQuestion("What is the title")
	description, _ := askQuestion("What is the description")

	fmt.Println(title)
	fmt.Println(description)
}

func askQuestion(question string) (string, error) {
	fmt.Printf("%s?\n", question)

	reader := bufio.NewReader(os.Stdin)

	readString, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	return strings.TrimSpace(readString), nil
}
