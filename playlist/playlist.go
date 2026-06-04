package main

import (
	"fmt"

	"github.com/Sheikh-Fahad-Ahmed/go-review/playlist/songs"
)

type playlist struct {
	songs        []*songs.Song
	currentTrack *songs.Song
}

func (p *playlist) Add(s *songs.Song) {
	if len(p.songs) > 0 {
		previousSong := p.songs[len(p.songs)-1]
		previousSong.Next = s
	}
	p.songs = append(p.songs, s)
}

func (p *playlist) Play() {
	if len(p.songs) == 0 {
		fmt.Println("Playlist is empty")
	}
	p.currentTrack = p.songs[0]
	fmt.Printf("Now playing: %s by %s\n", p.currentTrack.Title, p.currentTrack.Artist)
}

func (p *playlist) Next() {
	if p.currentTrack == nil {
		fmt.Println("There is no track currently playing, Use Play() to start playing songs")
	}

	if p.currentTrack.Next != nil {
		p.currentTrack = p.currentTrack.Next
		fmt.Printf("Next song is %s by %s\n", p.currentTrack.Title, p.currentTrack.Artist)
	} else {
		fmt.Println("You are at the end of the playlist")
		p.currentTrack = nil
	}

}
