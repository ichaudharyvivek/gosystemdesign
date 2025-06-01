// These are the complex subclasses that becomes a headache if used directly
package facade

import "fmt"

type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is ON")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Println("Playing movie:", movie)
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is OFF")
}

type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is ON")
}

func (p *Projector) SetInput(source string) {
	fmt.Println("Projector input set to:", source)
}

func (p *Projector) Off() {
	fmt.Println("Projector is OFF")
}

type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound System is ON")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Println("Volume set to", level)
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound System is OFF")
}
