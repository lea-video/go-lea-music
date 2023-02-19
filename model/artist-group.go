package model

type ArtistGroup struct {
	*Artist `gorm:"embedded"`
	Members []*Artist `gorm:"many2many:link_artist_solo_to_artist_group;"`
}
