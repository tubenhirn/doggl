package main

import (
	"github.com/tubenhirn/doggl/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
