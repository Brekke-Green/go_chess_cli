package game

import (
//    "fmt"
    "strings"

    "github.com/charmbracelet/bubbles/textinput"
    //"github.com/charmbracelet/lipgloss"
    tea "github.com/charmbracelet/bubbletea"
    dt "github.com/dylhunn/dragontoothmg"
    "github.com/Brekke-Green/go_chess_cli/fen"
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

type ScoreSheet struct {
    sheet       []string
}

type (
    errMsg error
) 

var translate = map[string]string {
    "":  " ",
	"B": "♝",
	"K": "♚",
    "N": "♞",
	"P": "♟",
	"Q": "♛",
	"R": "♜",
	"b": "♗",
	"k": "♔",
	"n": "♘",
	"p": "♙",
	"q": "♕",
	"r": "♖",
}

func NewGame() *Game {
    return NewGameStart(dt.Startpos)
}

func NewGameStart(position string) *Game {
    m := &Game{}

//    if !fen.IsValid(position) {
//        position = dt.Startpos
//    }

    board := dt.ParseFen(dt.Startpos)
    m.board = &board
    m.moves = m.board.GenerateLegalMoves()
    
    
    return m
}

func (m *Game) Init() tea.Cmd {
    return nil
}

func (m *Game) View() string {
    var s strings.Builder
    // grid := "  ┌───┬───┬───┬───┬───┬───┬───┬───┐\n8 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n7 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n6 │   │   │   │   │   │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n5 │   │   │   │   │   │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n4 │   │   │   │   │ . │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n3 │   │   │   │   │ . │   │   │   │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n2 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │\n  ├───┼───┼───┼───┼───┼───┼───┼───┤\n1 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │\n  └───┴───┴───┴───┴───┴───┴───┴───┘\n"
    var grid = fen.Grid(m.board.ToFen())
    for _, row := range grid {
        for _, letter := range row {
            s.WriteString(translate[letter])
        }
        s.WriteString("\n")
    }
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
