package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadMoveInput() (string, string, string, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nYour move (format: <player> <from> <to>, e.g. white e2 e4): ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "exit" {
		return "", "", "", false
	}

	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		fmt.Println("Invalid input. Use format: <player> <from> <to>")
		return "", "", "", true
	}

	return parts[0], parts[1], parts[2], true
}
