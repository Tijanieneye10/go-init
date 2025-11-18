package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (n *Note) Save() error {

	data, err := json.Marshal(n)

	if err != nil {
		return fmt.Errorf("failed to marshal note: %w", err)
	}

	err = os.WriteFile("database.json", data, 0644)

	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func ReadNote(filename string) (*Note, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var note Note

	err = json.Unmarshal(data, &note)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal note: %w", err)
	}

	return &note, nil
}

func main() {

	fmt.Println(ReadNote("database.json"))
	title, _ := askQuestion("What is the title")
	description, _ := askQuestion("What is the description")

	note := &Note{Title: title, Description: description, Date: time.Now()}

	err := note.Save()
	if err != nil {
		return
	}
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
