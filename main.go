package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type text struct {
	typed    string
	complete string
}

func calculateSpeed(sampleText text, timeElapsed time.Duration) int {
	correct := 0
	for i := 0; i < len(sampleText.complete); i++ {
		if sampleText.complete[i] == sampleText.typed[i] {
			correct++
		}
	}
	speed := int(float64(correct) / (timeElapsed.Minutes() * 5))
	return speed
}

func main() {
	sampleText := text{typed: "", complete: "hello world"}
	textSpacing := 20
	textSize := 30
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	startingTime := time.Now()
	textXPos := 50
	finished := false
	timeElapsed := time.Since(startingTime)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		if !finished {
			for i := 0; i < len(sampleText.complete); i++ {
				if i < len(sampleText.typed) {
					if sampleText.typed[i] == sampleText.complete[i] {
						rl.DrawText(string(sampleText.typed[i]), int32(i)*int32(textSpacing), int32(textXPos), int32(textSize), rl.Green)
					} else {
						rl.DrawText(string(sampleText.typed[i]), int32(i)*int32(textSpacing), int32(textXPos), int32(textSize), rl.Red)
					}
				} else {
					rl.DrawText(string(sampleText.complete[i]), int32(i)*int32(textSpacing), int32(textXPos), int32(textSize), rl.Black)
				}
			}
			buttonPressed := rl.GetCharPressed()
			if buttonPressed != 0 {
				charPressed := rune(buttonPressed)
				sampleText.typed = fmt.Sprintf("%s%c", sampleText.typed, charPressed)
			}
			if rl.IsKeyPressed(rl.KeyBackspace) {
				if len(sampleText.typed) > 0 {
					sampleText.typed = sampleText.typed[:len(sampleText.typed)-1]
				}
			}
			// strings.Split(timeElapsed.String(), ".")[0]
			timeElapsed = time.Since(startingTime)
			timeStr := fmt.Sprint(int(timeElapsed.Seconds()))
			rl.DrawText(timeStr, 0, 0, 30, rl.Black)
		}
		if finished || len(sampleText.typed) == len(sampleText.complete) {
			finished = true
			finishMessage := fmt.Sprintf("%s %d WPM", "speed: ", calculateSpeed(sampleText, timeElapsed))
			rl.DrawText(finishMessage, 0, 0, 50, rl.Black)
		}

		rl.EndDrawing()
	}
}
