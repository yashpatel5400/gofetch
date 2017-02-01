package main

import (
	"fmt"
)

type Background struct {
	// size of the background
	width  int
	height int

	// current speed at which scene is moving -- increases w/ run length
	speed  int

	// score of the player (increases based on speed)
	score  int

	// boolean to indicate whether the run has ended or not
	gameover bool
}

func initBackground() *Background {
	return &Background{
		width:  600,
		height: 400,
		
		speed: 1,
		score: 0,
		gameover: false,	
	}
}

func draw(background *Background) string {
	var buffer bytes.Buffer
	var ptsstr string = fmt.Sprintf("Points: %v \n", pts)

	buffer.Write([]byte(ptsstr))

	for y := 0; y < background.height; y++ {
		for x := 0; x < background.width; x++ {
			if background.get(x, y) > 0 {
				if gameover {
					buffer.Write(as.Bytes(ansi.Color("█", ansi.Red)))
				} else {
					buffer.Write(as.Bytes(ansi.Color("█", ansi.Green)))
				}
			} else if background.get(x, y) < 0 {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}