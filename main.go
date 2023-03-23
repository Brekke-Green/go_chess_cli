package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/coral"
    "github.com/Brekke-Green/go_chess_cli/game"
)

var (
	Version    = ""
	CommitSHA  = ""

	rootCmd = &coral.Command{
		Use:	"chess",
		Short: 	"Play chess against Stockfish",
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) == 0 {
				startPos, _ := readStdin()

				debug := os.Getenv("DEBUG")
				if debug != "" {
					f, err := tea.LogToFile(debug, "")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()
				}

				p := tea.NewProgram(
					game.NewGameStart(startPos),
					tea.WithAltScreen(),
					tea.WithMouseCellMotion(),
				)

				_, err := p.Run()
				return err
			}

			return cmd.Help()
		},
		DisableFlagsInUseLine: true,
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}	
}

func readStdin() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if stat.Mode()&os.ModeNamedPipe == 0 && stat.Size() == 0 {
		return "", errors.New("No starting position?")
	}

	reader := bufio.NewReader(os.Stdin)
	var b strings.Builder

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		_, err = b.WriteRune(r)
		if err != nil {
			return "", err
		}
	}

	return b.String (), nil
}

