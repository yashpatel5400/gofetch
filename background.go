package main

import (
	"fmt"
	"bytes"
	"time"
	"math/rand"
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
// -------------------------------
// | Score |                 -----
// ---------       SKY       -----
// ---------      CLOUD      -----
// -------------------------------
// ---------     GROUND      -----
// -------------------------------
var SCORE_BOX    = &BoundingBox{bottom:25, top:35, left:5, right:6}
var GROUND_LEVEL = &BoundingBox{bottom:0,  top:10, left:0, right:WIDTH}
var CLOUD_LEVEL  = &BoundingBox{bottom:20, top:30, left:0, right:WIDTH}

// object used to represent the boundaries of where object is/can be
type BoundingBox struct {
	top int
	bottom int
	right int
	left int
}

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

func inside(x int, y int, box *BoundingBox) bool {
	return (x >= box.left && x < box.right &&
			y >= box.bottom && y < box.top)
}

func initBackground() *Background {
	board := [][]int{}

	for i := 0; i < HEIGHT; i++ {
		row := make([]int, WIDTH)
		if i >= GROUND_LEVEL.bottom && i < GROUND_LEVEL.top {
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

func insertOnBoard(background *Background, box *BoundingBox, id string) *Background {
	var identifier int
	switch id {
	case "sky":
		identifier = SKY
	case "ground":
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

	for y := box.top; y < box.bottom; y++ {
		for x := box.left; x < box.right; x++ {
			background.board[y][x] = identifier
		}
	}
	return background
} 

func moveBackground(background *Background) *Background {
	for y := 0; y < background.height; y++ {
		var DEFAULT int
		if y >= GROUND_LEVEL.bottom && y < GROUND_LEVEL.top {
			DEFAULT = GROUND
		} else {
			DEFAULT = SKY
		}

		for x := 0; x < background.width; x++ {
			// old part of background that's shifting/moving at set speed
			if x < background.width - background.speed {
				background.board[y][x] = background.board[y][x + background.speed]
			} else { // new background being generated to replace the old background
				background.board[y][x] = DEFAULT
			}			
		}
	}
	return background
}

func insertCloud(background *Background) *Background {
	CLOUD_RANGE := CLOUD_LEVEL.top - CLOUD_LEVEL.bottom
	rand.Seed(time.Now().UnixNano())
	randLocation := rand.Intn(CLOUD_RANGE)
	cloudY       := CLOUD_LEVEL.bottom + randLocation
	cloudBlock   := &BoundingBox{
		left: background.width-2, right: background.width-1, 
		bottom: cloudY-1, top: cloudY}
	return insertOnBoard(background, cloudBlock, "cloud")
}

func render(background *Background) {
	var buffer bytes.Buffer
	for y := background.height - 1; y >= 0; y-- {
		for x := 0; x < background.width; x++ {
			curPixel := background.board[y][x]
			if curPixel == SKY {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else if curPixel == GROUND {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Green)))
			} else if curPixel == CLOUD {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.White)))
			} else if curPixel == PLAYER {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Yellow)))
			} else if curPixel == ENEMY {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Red)))
			} else {
				buffer.Write(as.Bytes(curPixel))
			}
		}
		buffer.WriteByte('\n')
	}
	fmt.Println(buffer.String())
}