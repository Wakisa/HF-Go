package main

import "wakisa.com/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {

	for _, song := range songs { //loop over each song
		device.Play(song) // paly current song
	}

	device.Stop() // stop the player once we are done
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{} // Update the variable to hold any Player
	playList(player, mixtape)               // Pass a TapePlayer variable to hold to plyaList
	player = gadget.TapeRecoder{}
	playList(player, mixtape)
	recorder, ok := player.(gadget.TapeRecoder)
	if ok {
		recorder.Record()
	}
	recorder.Stop()

}
