package main

import "github.com/Sheikh-Fahad-Ahmed/go-review/playlist/songs"

func main() {

	newPlaylist := playlist{}

	song1 := songs.Create("Yellow", "Coldplay")
	song2 := songs.Create("Sweden", "C418")
	song3 := songs.Create("Drag Path", "Twenty one Pilots")

	newPlaylist.Add(song1)
	newPlaylist.Add(song2)
	newPlaylist.Add(song3)

	newPlaylist.Play()
	newPlaylist.Next()
	newPlaylist.Next()
	newPlaylist.Next()

}
