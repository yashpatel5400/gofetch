package main

import (
	"fmt"
	"bytes"
	"github.com/SimonWaldherr/golibs/ansi"
	"github.com/SimonWaldherr/golibs/as"
)

// global variables used for rendering different "objects" in scene
var SKY    = 0
var GROUND = 1
var CLOUD  = 2
var PLAYER = 3
var ENEMY  = 4

// variables for adjusting window size
var WIDTH  = 100
var HEIGHT = 30

// global variables used for background "sections" -- marks height btween which
// certain background characteristics occur
// ---------------------------
// -----       SKY       -----
// -----      CLOUD      -----
// ---------------------------
// -----     GROUND      -----
// ---------------------------
var CLOUD_LEVEL  = []int{20,30}
var GROUND_LEVEL = []int{0,10}

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
	board := [][]int{}

	for i := 0; i < HEIGHT; i++ {
		row := make([]int, WIDTH)
		if i >= GROUND_LEVEL[0] && i < GROUND_LEVEL[1] {
			for j := 0; j < WIDTH; j++ {
				row[j] = GROUND
			}
		} 
		board = append(board, row)
	}

	return &Background{
		width:  WIDTH,
		height: HEIGHT,
		board:  board,

		speed: 1,
		score: 0,
		gameover: false,	
	}
}

func insertOnBoard(background *Background, positions [][]int, id string) *Background {
	var identifier int
	switch id {
	case "sky":
		identifier = SKY
	case "tree":
		identifier = GROUND
	case "cloud":
		identifier = CLOUD
	case "player":
		identifier = PLAYER
	case "enemy":
		identifier = ENEMY
	default:
		fmt.Println("Error: Not a valid object type!")
		return background
	}

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

func moveBackground(background *Background) {
	for y := 0; y < background.height; y++ {
		// old part of background that's shifting/moving at set speed
		for x := 0; x < background.width - background.speed; x++ {
			background.board[x][y] = background.board[x + background.speed][y]
		}

		// new background being generated to replace the old background
		for x := background.width - background.speed; x < background.width; x++ {
			background.board[x][y] = SKY
		}
	}
}

func insertClouds(background *Background) {
	return
}

func render(background *Background) {
	var buffer bytes.Buffer
	for y := background.height - 1; y >= 0; y-- {
		for x := 0; x < background.width; x++ {
			if background.board[y][x] == SKY {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else if background.board[y][x] == GROUND {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Green)))
			} else if background.board[y][x] == CLOUD {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.White)))
			} else if background.board[y][x] == PLAYER {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else if background.board[y][x] == ENEMY {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Red)))
			} 
		}
		buffer.WriteByte('\n')
	}
	fmt.Println(buffer.String())
}