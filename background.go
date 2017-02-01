package main

import (
	"fmt"
	"bytes"
	"github.com/SimonWaldherr/golibs/ansi"
	"github.com/SimonWaldherr/golibs/as"
)

// object containing the background of the game -- also contains
// variables about what is in each position on the board
type Background struct {
	// size of the background
	width  int
	height int

	// board as array -- each point corresponds an entry
	board [][]int

	// current speed at which scene is moving -- increases w/ run length
	speed  int

	// score of the player (increases based on speed)
	score  int

	// boolean to indicate whether the run has ended or not
	gameover bool
}

func initBackground() *Background {
	width := 100
	height:= 30
	board := [][]int{}

	for i := 0; i < width; i++ {
		row := make([]int, height)
		board = append(board, row)
	}

	return &Background{
		width:  width,
		height: height,
		board: board,

		speed: 1,
		score: 0,
		gameover: false,	
	}
}

func insertOnBoard(background *Background, positions [][]int, identifier int) *Background {
	PAIR_SIZE := 2

	for i := 0; i < len(positions); i++ {
		curPosition := positions[i]
		if len(curPosition) != PAIR_SIZE {
			fmt.Println("Attempting to insert non-pair formatted position!!")
			return background
		}

		background.board[curPosition[0]][curPosition[1]] = identifier
	}
	return background
} 

func render(background *Background) string {
	SKY    := 0
	TREE   := 1
	CLOUD  := 2
	PLAYER := 3
	ENEMY  := 4

	var buffer bytes.Buffer
	for y := 0; y < background.height; y++ {
		for x := 0; x < background.width; x++ {
			if background.board[x][y] == SKY {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else if background.board[x][y] == TREE {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Red)))
			} else if background.board[x][y] == CLOUD {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.White)))
			} else if background.board[x][y] == PLAYER {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else if background.board[x][y] == ENEMY {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Red)))
			} 
		}
		buffer.WriteByte('\n')
	}
	fmt.Println(buffer.String())
}