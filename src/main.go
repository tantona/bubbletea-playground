package main

// A simple program that counts down from 5 and then exits.

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	zone.NewGlobal()
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		if _, err := tea.LogToFile(logfilePath, "simple"); err != nil {
			log.Fatal(err)
		}
	}

	// Initialize our program
	p := tea.NewProgram(model{value: 5}, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
