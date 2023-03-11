package main

import "fmt"

type Cell int

const (
	Empty Cell = iota
	X
	O
)

func (c Cell) String() string {
	switch c {
	case Empty:
		return "[   ]"
	case X:
		return "[ X ]"
	case O:
		return "[ O ]"
	default:
		panic("Failed")
	}
}

type Board struct {
	game_board [][]Tile
}

type Tile struct {
	is_marked   bool
	player_mark Cell
}

func (b Board) Win_Check() string {
	x_win := Board_Validator(b, X)
	o_win := Board_Validator(b, O)

	switch {
	case x_win && !o_win:
		return "X won."
	case o_win && !x_win:
		return "O won."
	default:
		return "Game in progress."
	}
}

func (b Board) Place_Mark(cell_loc int, is_player_one bool) bool {

	var mark Cell

	if is_player_one {
		mark = X
	} else {
		mark = O
	}

	if cell_loc < 1 || cell_loc > 9 {
		fmt.Println("Non-valid tile choice")
		return false
	} else {
		if cell_loc >= 1 && cell_loc <= 3 {
			if b.game_board[0][cell_loc-1].is_marked {
				fmt.Println("Tile already marked, pick another one.")
				return false
			} else {
				b.game_board[0][cell_loc-1] = Tile{
					is_marked:   true,
					player_mark: mark,
				}
				return true
			}
		} else if cell_loc >= 4 && cell_loc <= 6 {
			offset := 4

			if b.game_board[1][cell_loc-offset].is_marked {
				fmt.Println("Tile already marked, pick another one.")
				return false
			} else {
				b.game_board[1][cell_loc-offset] = Tile{
					is_marked:   true,
					player_mark: mark,
				}
				return true
			}
		} else {
			offset := 7

			if b.game_board[2][cell_loc-offset].is_marked {
				fmt.Println("Tile already marked, pick another one.")
				return false
			} else {
				b.game_board[2][cell_loc-offset] = Tile{
					is_marked:   true,
					player_mark: mark,
				}
				return true
			}
		}
	}
}

func Create_board() *Board {
	board := &Board{
		game_board: make([][]Tile, 3),
	}

	for i := range board.game_board {
		board.game_board[i] = make([]Tile, 3)
		board.game_board[i] = []Tile{
			{is_marked: false, player_mark: Empty},
			{is_marked: false, player_mark: Empty},
			{is_marked: false, player_mark: Empty},
		}
	}
	return board
}

func (b Board) Print_board() {
	fmt.Println("--------")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(b.game_board[i][j].player_mark)
		}
		fmt.Println()
	}
	fmt.Println("--------")
}

func (b Board) Is_full() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.game_board[i][j].player_mark == Empty {
				return false
			}
		}
	}
	return true
}
