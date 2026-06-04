package songs

type Song struct {
	Title  string
	Artist string
	Next   *Song
}

func Create(title string, artist string) *Song {
	return &Song{
		Title:  title,
		Artist: artist,
		Next:   nil,
	}
}
