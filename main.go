package main

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/gookit/color"
)

type UserInput struct {
	File string `arg:"-f,required"`
}

func main() {

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
