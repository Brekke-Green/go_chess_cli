package fen

import (
    "fmt"
    "strings"
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

func Fields(fen string) []string {
    return strings.Split(fen, " ")
}

func Ranks(fen string) []string {
    return strings.Split(Fields(fen)[0], "/")
}

func Grid(fen string) [8][8]string {
    var grid [8][8]string
    for r, rank := range Ranks(fen) { 
        var row [8]string
        c := 0
        for _, col := range rank {
            skip := 1
            if isNumeric(col) {
                skip = runToInt(col)
            } else {
                row[c] = fmt.Sprintf("%c", col)
            }
            c += skip
        }
        grid[r] = row
    }
    return grid
}

func isNumeric(r rune) bool {
    return r >= '0' && r <= '9'
}

func runToInt(r rune) int {
    return int(r - '0')
}
