package model

// NOTE: unused
type Song struct {
	ID             int
	Name           string
	Artists        []int
	PreferredAudio int
	PreferredVideo int
}

// NOTE: unused
type Performance struct {
	ID      int
	Song    int
	Artists []int
	Media   []int
}
