package game

import (
//    "fmt"
    "strings"

    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    dt "github.com/dylhunn/dragontoothmg"
)

type Game struct {
    board       *dt.Board
    moves       []dt.Move
    pieceMoves  []dt.Move
    selected    string
    buffer      string
    err         error
    textInput   textinput.Model
}

type (
    errMsg error
) 

func NewGame() *Game {
    return NewGameStart(dt.Startpos)
}

func NewGameStart(position string) *Game {
    m := &Game{}

//    if !fen.IsValid(position) {
//        position = dt.Startpos
//    }

    if position == position {
        return m
    }

//    board := dt.ParseFen(position)
//    m.board = &board
//    m.moves = m.board.GenerateLegalMoves()
    
    return m
}

func (m *Game) Init() tea.Cmd {
    return nil
}

func (m *Game) View() string {
    var s strings.Builder
    grid := "  ┌───┬───┬───┬───┬───┬───┬───┬───┐\n8 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n7 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n6 │   │   │   │   │   │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n5 │   │   │   │   │   │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n4 │   │   │   │   │ . │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n3 │   │   │   │   │ . │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n2 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n1 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │\n  └───┴───┴───┴───┴───┴───┴───┴───┘\n"
    s.WriteString(grid)
    return s.String()
}

func (m *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
            return m, tea.Quit
        }

    case errMsg:
        m.err = msg
        return m, nil
    }

    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}
