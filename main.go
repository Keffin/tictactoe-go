package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fetch_position() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	pos, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("something went wrong when reading input")
	}
	return pos, nil
}

func main() {
	fmt.Println("Enter a value (1-9)")

	b := Create_board()

	b.Print_board()

	is_player_one := true

	for {
		if b.Is_full() {
			fmt.Println("Game ended. Full board.")
			return
		}

		if is_player_one {
			fmt.Println("Current player is X.")
		} else {
			fmt.Println("Current player is O.")
		}

		pos, err := fetch_position()
		if err != nil {
			fmt.Println(err)
			return
		}

		// To have only one table being rendered at time.
		fmt.Print("\033[H\033[2J")

		if is_player_one {
			valid_mark := b.Place_Mark(pos, is_player_one)
			if valid_mark {
				is_player_one = !is_player_one
			}
		} else {
			valid_mark := b.Place_Mark(pos, is_player_one)
			if valid_mark {
				is_player_one = !is_player_one
			}
		}

		b.Print_board()
		winner := b.Win_Check()
		if winner == "X won." {
			fmt.Println(winner)
			return
		} else if winner == "O won." {
			fmt.Println(winner)
			return
		} else {
			continue
		}
	}

}
