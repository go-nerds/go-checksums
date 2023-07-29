package main

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

type UserInput struct {
	File string `arg:"-f,required"`
}

func main() {

	info := color.FgLightCyan.Render
	pterm.DefaultBox.
		WithRightPadding(10).
		WithLeftPadding(10).
		WithTopPadding(2).
		WithBottomPadding(2).
		Println(info("Go check-sum created by AAVision"))

	userInput := UserInput{}
	arg.MustParse(&userInput)

	color.Yellowln("Processing...")
	start := time.Now()

	err := generateHashes(userInput.File)

	if err != nil {
		fmt.Println(err)
	}

	duration := time.Since(start)
	color.Green.Println("Finished in:", duration)

}
