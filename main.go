package main

import (
	"fmt"
	"os"
	"time"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
        fmt.Println("Clicker started.")
	fmt.Println("Press + (key: 43) – ON/OFF clicker.")
	fmt.Println("Press - (key: 45) – save/cancel fixed position.")
	fmt.Println("Press ESC – Shutdown aplication.")

	clicking := false
	useFixedPosition := false
	var fixedX, fixedY int


	go func() {
		for ev := range hook.Start() {
			if ev.Kind == hook.KeyDown {
				switch ev.Keychar {
				case 43: // +
					clicking = !clicking
					if clicking {
						fmt.Println("Clicking: ON")
					} else {
						fmt.Println("Clicking: OFF")
					}
				case 27: // ESC
					fmt.Println("Aplication down")
					os.Exit(0)
				case 45: // -
					if clicking {
						clicking = !clicking
						fmt.Println("Clicking: OFF")
					}

					if !useFixedPosition {
						fixedX, fixedY = robotgo.GetMousePos()
						useFixedPosition = true
						fmt.Printf("Fixed position saved: (%d, %d)\n", fixedX, fixedY)
					} else {
						useFixedPosition = false
						fmt.Println("Fixed position cleared. Using live cursor.")
					}
					
				default: // Debug
					 fmt.Printf("KeyDown: %v\n", ev.Keychar)
				}
			}
		}
	}()

	for  {
		if clicking {
			var x, y int
			if useFixedPosition {
				robotgo.MoveMouse(fixedX, fixedY)
				x, y = fixedX, fixedY
			} else {
				x, y = robotgo.GetMousePos()
			}
			robotgo.Click("left", false)
			fmt.Printf("Click: (%d, %d)\n", x, y)
			time.Sleep(40 * time.Millisecond)
		} else {
			time.Sleep(100 * time.Millisecond)
		}

	}
}

