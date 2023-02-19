package model

type ArtistSolo struct {
	*Artist `gorm:"embedded"`
	Groups  []*ArtistGroup `gorm:"many2many:link_artist_solo_to_artist_group;"`
}
