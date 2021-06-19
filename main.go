package main

import "backend/di"

func main() {
	event := di.InitializeEvent()

	event.Start()
}
