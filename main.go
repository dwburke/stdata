package main

import (
    "github.com/dwburke/atexit"

    "github.com/dwburke/stdata/cmd"
)

func main() {
    defer atexit.AtExit()

	cmd.Execute()
}
