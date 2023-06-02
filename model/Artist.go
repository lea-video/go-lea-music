package model

type OneOfArtist struct {
	ArtistType string `json:"type"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Members    []int  `json:"members,omitempty"`
}

type ArtistSolo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (as ArtistSolo) ToOneOf() *OneOfArtist {
	return &OneOfArtist{
		ArtistType: "solo",
		ID:         as.ID,
		Name:       as.Name,
		Members:    nil,
	}
}

type ArtistGroup struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Members []int  `json:"members"`
}

func (ag ArtistGroup) ToOneOf() *OneOfArtist {
	return &OneOfArtist{
		ArtistType: "group",
		ID:         ag.ID,
		Name:       ag.Name,
		Members:    ag.Members,
	}
}
