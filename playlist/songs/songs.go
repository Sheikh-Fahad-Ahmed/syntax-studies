package songs

type Song struct {
	Title  string
	Artist string
	Next   *string
}

func (s *Song) Create(title string, artist string) Song {
	return Song{
		Title:  title,
		Artist: artist,
	}
}
