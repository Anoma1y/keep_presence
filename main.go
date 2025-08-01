package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	duration := flag.Int("duration", 60, "Duration in seconds to run the mouse mover")
	flag.Parse()

	rand.New(rand.NewSource(time.Now().UnixNano()))

	screenWidth, screenHeight := robotgo.GetScreenSize()

	for {
		x := rand.Intn(screenWidth)
		y := rand.Intn(screenHeight)

		robotgo.Move(x, y)

		sleepTime := time.Duration(rand.Intn(10) + *duration) * time.Second
		fmt.Printf("Moved mouse to (%d, %d). Sleeping for %s...\n", x, y, sleepTime)
		time.Sleep(sleepTime)
	}
}
