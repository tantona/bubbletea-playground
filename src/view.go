package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"golang.org/x/term"
)

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	value int
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	return nil
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.MouseMsg:
		e := msg.(tea.MouseMsg)
		if e.Type != tea.MouseLeft {
			return m, nil
		}

		if zone.Get("ok").InBounds(e) {
			return m, okClicked

		} else if zone.Get("cancel").InBounds(e) {
			return m, cancelClicked
		}
	case sayMsg:
		fmt.Println("GOT SAY MESSAGE")

		return m, nil

	case tea.KeyMsg:
		if msg.(tea.KeyMsg).Type == tea.KeyEsc {
			return m, tea.Quit
		}
	}

	return m, nil
}

// Views return a string based on data in the model. That string which will be
// rendered to the terminal.
func (m model) View() string {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	docStyle := lipgloss.NewStyle().Padding(1, 2, 1, 2)
	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginTop(1)

	activeButtonStyle := buttonStyle.Copy().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")).
		MarginRight(2).
		Underline(true)
	okButton := activeButtonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Maybe")
	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(fmt.Sprintf("Hi. This program will exit in %d seconds. To quit sooner press any key.", m.value))
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, zone.Mark("ok", okButton), zone.Mark("cancel", cancelButton))
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		docStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
	)

	return zone.Scan(dialog)

}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
type tickMsg time.Time

type sayMsg int

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}

func okClicked() tea.Msg {
	exec.Command("say", "OK").Output()
	time.Sleep(3 * time.Second)

	return sayMsg(0)
}

func cancelClicked() tea.Msg {
	exec.Command("say", "CANCEL").Output()
	return sayMsg(0)
}
