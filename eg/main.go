package main

import (
	"fmt"

	cowsay "github.com/usrbinkat/Neo-cowsay"
)

func main() {
	if false {
		simple()
	} else {
		complex()
	}
}

func simple() {
	say, err := cowsay.Say(
		cowsay.Phrase("Hello"),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}

func complex() {
	cow, err := cowsay.NewCow(
		cowsay.Phrase("Hello"),
		cowsay.BallonWidth(40),
		//cowsay.Thinking(),
		cowsay.Random(),
	)
	if err != nil {
		panic(err)
	}
	say, err := cow.Say()
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
