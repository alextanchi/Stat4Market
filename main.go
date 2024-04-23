package main

import (
	"Stat4Market/cmd/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

}
