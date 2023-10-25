package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	screenWidth, screenHeight := robotgo.GetScreenSize()

	for {
		x := rand.Intn(screenWidth)
		y := rand.Intn(screenHeight)

		robotgo.Move(x, y)

		sleepTime := time.Duration(rand.Intn(10)+5) * time.Second
		fmt.Printf("Moved mouse to (%d, %d). Sleeping for %s...\n", x, y, sleepTime)
		time.Sleep(sleepTime)
	}
}
