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
var SCORE  = 5

// variables for adjusting window size
var WIDTH  = 100
var HEIGHT = 30

// global variables used for background "sections" -- marks height btween which
// certain background characteristics occur
// -------------------------------
// ------     ENEMY/CLOUD    -----
// -------------------------------
// ---------     GROUND      -----
// | Score |                 -----
// -------------------------------
var SCORE_BOX    = &BoundingBox{bottom:5, top:6, left:5, right:6}
var GROUND_LEVEL = &BoundingBox{bottom:0,  top:10, left:0, right:WIDTH}
var CLOUD_LEVEL  = &BoundingBox{bottom:20, top:30, left:0, right:WIDTH}
var ENEMY_LEVEL  = &BoundingBox{bottom:13, top:30, left:0, right:WIDTH}

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

/*****************************************************************************/
/* Initializes the background struct object and returns a pointer to it      */
/*****************************************************************************/
func initBackground() *Background {
	board := [][]int{}
	for y := 0; y < HEIGHT; y++ {
		row := make([]int, WIDTH)
		if y >= GROUND_LEVEL.bottom && y < GROUND_LEVEL.top {
			for x := 0; x < WIDTH; x++ {
				row[x] = GROUND
			}
		} else if y >= SCORE_BOX.bottom && y < SCORE_BOX.top {
			row[SCORE_BOX.left] = SCORE
		} 

		board = append(board, row)
	}

	return &Background{
		width:  WIDTH,
		height: HEIGHT,
		board:  board,

		speed: 2,
		score: 0,
		gameover: false,	
	}
}

/*****************************************************************************/
/* Given a background object, a valid box corresponding to a position on the */
/* board, and a string of what the box corresponds to (i.e. "player" or      */
/* "enemy"), returns background with object in position                      */
/*****************************************************************************/
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

	for y := box.bottom; y < box.top; y++ {
		for x := box.left; x < box.right; x++ {
			background.board[y][x] = identifier
		}
	}
	return background
} 

/*****************************************************************************/
/* Given the background, returns a new background where everything has been  */
/* shifted according to the current speed (does not move the player)         */
/*****************************************************************************/
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
				// these do not "move" with the rest of the background
				if background.board[y][x] == SCORE {
					continue
				}

				if  background.board[y][x + background.speed] == SCORE || 
					background.board[y][x + background.speed] == PLAYER {

					background.board[y][x] = SKY
				} else {
					background.board[y][x] = background.board[y][x + background.speed]
				}
			} else { // new background being generated to replace the old background
				background.board[y][x] = DEFAULT
			}			
		}
	}
	background.score += background.speed
	return background
}

/*****************************************************************************/
/* Given the background, a bounding box that corresponds to the level (i.e a */
/* region of the background where particular objects can spawn), the height  */
/* and width of such objects, and the object type, randomly creates an       */
/* of the specified type somewhere in the level and returns background w/ obj*/
/*****************************************************************************/
func insertObj(background *Background, level *BoundingBox, 
	width int, height int, objType string) *Background {

	levelRange := level.top - level.bottom
	rand.Seed(time.Now().UnixNano())
	randLocation := rand.Intn(levelRange)
	objY       := level.bottom + randLocation
	objBlock   := &BoundingBox{
		left: background.width-width, right: background.width-1, 
		bottom: objY-height, top: objY}
	return insertOnBoard(background, objBlock, objType)
}

/*****************************************************************************/
/* Given the background, return new background with cloud inserted (one)     */
/*****************************************************************************/
func insertCloud(background *Background) *Background {
	CLOUD_WIDTH  := 5
	CLOUD_HEIGHT := 3
	return insertObj(background, CLOUD_LEVEL, CLOUD_WIDTH, CLOUD_HEIGHT, "cloud")
}

/*****************************************************************************/
/* Given the background, return new background with enemy inserted (one)     */
/*****************************************************************************/
func insertEnemy(background *Background) *Background {
	ENEMY_WIDTH := 3
	ENEMY_HEIGHT := 3
	return insertObj(background, ENEMY_LEVEL, ENEMY_WIDTH, ENEMY_HEIGHT, "enemy")
}

/*****************************************************************************/
/* Given the background, outputs the drawing on command line/prompt being run*/
/*****************************************************************************/
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
				buffer.Write(as.Bytes(background.score))
			}
		}
		buffer.WriteByte('\n')
	}
	fmt.Println(buffer.String())
}