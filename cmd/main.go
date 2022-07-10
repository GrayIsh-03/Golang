package main

import (
	"golang/internal/deftype"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var player Player = deftype.TapePlayer{}
	mixtape := []string{"Jessie's Girls", "Whip it", "9 to 5"}
	playList(player, mixtape)
	player = deftype.TapeRecorder{}
	playList(player, mixtape)

}
