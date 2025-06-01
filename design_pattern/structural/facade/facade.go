/*
Facade is simply a wrapper to hide difficult information from client.
Assume if you have multiple small classes working together, its complex to handle all.
Create a central class that handles all for you. You don't deal with subclasses anymore.
Diagram:

	+---------------------+
	|   HomeTheaterFacade |
	+---------------------+
	| - dvd: DVDPlayer    |
	| - projector: Projector |
	| - sound: SoundSystem |
	+---------------------+
	| +WatchMovie(movie)  |
	| +EndMovie()         |
	+---------------------+

		       |
		       | uses
		-----------------------
		|          |          |
		v          v          v

	+---------+ +-----------+ +-------------+
	| DVDPlayer| | Projector | | SoundSystem|
	+---------+ +-----------+ +-------------+
	| +On()   | | +On()     | | +On()       |
	| +Play() | | +SetInput()| | +SetVolume()|
	| +Off()  | | +Off()     | | +Off()      |
	+---------+ +-----------+ +-------------+
*/
package facade

import "fmt"

// One class that has all
// Now you deal with this class only
type HomeTheaterFacade struct {
	dvd       *DVDPlayer
	projector *Projector
	sound     *SoundSystem
}

func NewHomeTheaterFacade(dvd *DVDPlayer, projector *Projector, sound *SoundSystem) *HomeTheaterFacade {
	return &HomeTheaterFacade{dvd, projector, sound}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.projector.On()
	h.projector.SetInput("DVD")
	h.sound.On()
	h.sound.SetVolume(5)
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.dvd.Off()
	h.sound.Off()
	h.projector.Off()
}
