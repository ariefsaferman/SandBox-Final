package main

import (
	"fmt"
	"strings"
)

type Board struct {
	board string
}

type Player struct {
	shape string
	win   bool
}

func checkBoard(board Board) bool {
	if (strings.Contains(board.board, "X") && strings.Contains(board.board, "O")) || strings.Contains(board.board, "-") {
		return true
	}
	return false
}

func checkInProgress(board Board) bool {
	if strings.Contains(board.board, "-") {
		return true
	}
	return false
}

func checkCond(player Player, board Board) bool {
	if (strings.Compare(player.shape, string(board.board[0])) == 0) &&
		(strings.Compare(player.shape, string(board.board[1])) == 0) &&
		(strings.Compare(player.shape, string(board.board[2])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[3])) == 0) &&
		(strings.Compare(player.shape, string(board.board[4])) == 0) &&
		(strings.Compare(player.shape, string(board.board[5])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[6])) == 0) &&
		(strings.Compare(player.shape, string(board.board[7])) == 0) &&
		(strings.Compare(player.shape, string(board.board[8])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[0])) == 0) &&
		(strings.Compare(player.shape, string(board.board[3])) == 0) &&
		(strings.Compare(player.shape, string(board.board[6])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[1])) == 0) &&
		(strings.Compare(player.shape, string(board.board[4])) == 0) &&
		(strings.Compare(player.shape, string(board.board[7])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[2])) == 0) &&
		(strings.Compare(player.shape, string(board.board[5])) == 0) &&
		(strings.Compare(player.shape, string(board.board[8])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[0])) == 0) &&
		(strings.Compare(player.shape, string(board.board[4])) == 0) &&
		(strings.Compare(player.shape, string(board.board[8])) == 0) {
		return true
	} else if (strings.Compare(player.shape, string(board.board[2])) == 0) &&
		(strings.Compare(player.shape, string(board.board[4])) == 0) &&
		(strings.Compare(player.shape, string(board.board[6])) == 0) {
		return true
	}
	return false
}

func CheckWin(playerX, playerO Player, board Board) {
	xWin := checkCond(playerX, board)
	oWin := checkCond(playerO, board)
	isProgress := checkInProgress(board)
	isValidBoad := checkBoard(board)

	if xWin && oWin {
		fmt.Println("Invalid game board")
	} else if xWin {
		fmt.Println("x Wins!")
	} else if oWin {
		fmt.Println("O Wins!")
	} else if !xWin && !oWin && isProgress {
		fmt.Println("Game still in progress!")
	} else {
		fmt.Println("Its a draw!")
	}
}

func playGame() {
	var playerX = Player{"X", false}
	var playerO = Player{"O", false}
	var board = Board{"XXXOOOXXO"}

	CheckWin(playerX, playerO, board)
}

func main() {
	playGame()
}
