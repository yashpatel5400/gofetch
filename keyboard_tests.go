package main

import (
	"fmt"
)

/*****************************************************************************/
/* Checks to ensure that the arrow keys are being properly read from keyboard*/
/* Successful run will show that the up/down keys are properly being read in */
/*****************************************************************************/
func test_arrows() {
	keyboard := initKeys()
	for {
		keyboard = readKeys(keyboard)
		// fmt.Println(keyboard.upPushed)
		// fmt.Println(keyboard.downPushed)
	}
}

/*****************************************************************************/
/* Runs all the tests written above -- used only in production compiles      */
/*****************************************************************************/
func main() {
	fmt.Println("Testing keyboard...")
	test_arrows()
}