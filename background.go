package main

import (
	"fmt"
)

type Background struct {
  width  int
  height int
  speed  int
  score  int
}

func moveBackground(background *Background) {

}

func draw(background *Background) string {
	var buffer bytes.Buffer
	var ptsstr string = fmt.Sprintf("Points: %v \n", pts)

	buffer.Write([]byte(ptsstr))

	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.get(x, y) > 0 {
				if gameover {
					buffer.Write(as.Bytes(ansi.Color("█", ansi.Red)))
				} else {
					buffer.Write(as.Bytes(ansi.Color("█", ansi.Green)))
				}
			} else if field.get(x, y) < 0 {
				buffer.Write(as.Bytes(ansi.Color("█", ansi.Blue)))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}