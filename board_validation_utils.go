package main

func Board_Validator(b Board, player_mark Cell) bool {
	if (b.game_board[0][0].player_mark == player_mark && b.game_board[0][1].player_mark == player_mark && b.game_board[0][2].player_mark == player_mark) ||
		(b.game_board[1][0].player_mark == player_mark && b.game_board[1][1].player_mark == player_mark && b.game_board[1][2].player_mark == player_mark) ||
		(b.game_board[2][0].player_mark == player_mark && b.game_board[2][1].player_mark == player_mark && b.game_board[2][2].player_mark == player_mark) ||
		(b.game_board[0][0].player_mark == player_mark && b.game_board[1][0].player_mark == player_mark && b.game_board[2][0].player_mark == player_mark) ||
		(b.game_board[0][1].player_mark == player_mark && b.game_board[1][1].player_mark == player_mark && b.game_board[2][1].player_mark == player_mark) ||
		(b.game_board[0][2].player_mark == player_mark && b.game_board[1][2].player_mark == player_mark && b.game_board[2][2].player_mark == player_mark) ||
		(b.game_board[0][0].player_mark == player_mark && b.game_board[1][1].player_mark == player_mark && b.game_board[2][2].player_mark == player_mark) ||
		(b.game_board[0][2].player_mark == player_mark && b.game_board[1][1].player_mark == player_mark && b.game_board[2][0].player_mark == player_mark) {
		return true
	}
	return false
}
