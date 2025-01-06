package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

type TapeRecoder struct {
	Microphones int
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecoder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecoder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecoder) Stop() {
	fmt.Println("Stopped!")
}
