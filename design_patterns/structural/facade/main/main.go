package main

import (
	f "design-patterns/structural/facade"
	"fmt"
)

func main() {
	dvd := &f.DVDPlayer{}
	proj := &f.Projector{}
	sd := &f.SoundSystem{}

	facade := f.NewHomeTheaterFacade(dvd, proj, sd)
	facade.WatchMovie("Happy World!")
	fmt.Println()
	facade.EndMovie()

}
