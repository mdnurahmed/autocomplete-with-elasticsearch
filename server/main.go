package main

import (
	"autocomplete/app/router"
)

func main() {
	r := router.InitializeApp()
	r.Run()
}
