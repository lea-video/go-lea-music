package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/db"
	"github.com/lea-video/go-lea-music/model"
)

const (
	EXPAND_MEDIA  = "media"
	EXPAND_TRACKS = "tracks"
	EXPAND_ARTIST = "artists"
)

func doExpansions(c *fiber.Ctx, db db.GenericDB, resp model.ResponseObject) (model.ResponseObject, error) {
	// collect ids to expand
	to_resolve := collectExpansionIDs(c, resp)

	// filter out those already present
	to_resolve[EXPAND_MEDIA] = filterUniqueness(to_resolve[EXPAND_MEDIA], resp.Media)
	to_resolve[EXPAND_TRACKS] = filterUniqueness(to_resolve[EXPAND_TRACKS], resp.MediaTrack)
	to_resolve[EXPAND_ARTIST] = filterUniqueness(to_resolve[EXPAND_ARTIST], resp.Artists)

	// do expansions
	resp, err := doExpansion(resp, db, to_resolve)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func collectExpansionIDs(c *fiber.Ctx, resp model.ResponseObject) (to_resolve map[string][]int) {
	mediaIDs := make([]int, 0)
	if hasExpand(c, EXPAND_MEDIA) {
		for _, m := range resp.MediaTrack {
			mediaIDs = append(mediaIDs, m.ParentMedia)
		}
		for _, t := range resp.TMPFile {
			mediaIDs = append(mediaIDs, t.ParentMedia)
		}
	}
	to_resolve[EXPAND_MEDIA] = mediaIDs

	trackIDs := make([]int, 0)
	if hasExpand(c, EXPAND_TRACKS) {
		for _, m := range resp.Media {
			trackIDs = append(trackIDs, m.Tracks...)
		}
	}
	to_resolve[EXPAND_TRACKS] = trackIDs

	artistIDs := make([]int, 0)
	if hasExpand(c, EXPAND_ARTIST) {
		for _, m := range resp.Artists {
			trackIDs = append(artistIDs, m.Members...)
		}
	}
	to_resolve[EXPAND_ARTIST] = artistIDs

	return
}

func doExpansion(resp model.ResponseObject, db db.GenericDB, to_resolve map[string][]int) (model.ResponseObject, error) {
	if len(to_resolve[EXPAND_MEDIA]) > 0 {
		medias, err := db.GetMediaByID(to_resolve[EXPAND_MEDIA])
		if err != nil {
			return resp, err
		}
		resp.Media = mergeMaps(resp.Media, medias)
	}

	if len(to_resolve[EXPAND_TRACKS]) > 0 {
		tracks, err := db.GetMediaTracksByID(to_resolve[EXPAND_TRACKS])
		if err != nil {
			return resp, err
		}
		resp.MediaTrack = mergeMaps(resp.MediaTrack, tracks)
	}

	if len(to_resolve[EXPAND_ARTIST]) > 0 {
		artists, err := db.GetArtistsByID(to_resolve[EXPAND_ARTIST])
		if err != nil {
			return resp, err
		}
		resp.Artists = mergeMaps(resp.Artists, artists)
	}

	return resp, nil
}

func filterUniqueness[T any](to_resolve []int, present map[int]T) []int {
	// Create a map to store the present values as a set
	presentSet := make(map[int]bool)
	for key := range present {
		presentSet[key] = true
	}

	// Second lookup-table to ensure uniqueness
	seen := make(map[int]bool)

	// Filter the to_resolve list
	filtered := make([]int, 0)
	for _, value := range to_resolve {
		if !presentSet[value] && !seen[value] {
			filtered = append(filtered, value)
			seen[value] = true
		}
	}

	return filtered
}

func mergeMaps[T any](map1, map2 map[int]T) map[int]T {
	merged := make(map[int]T)

	for key, value := range map1 {
		merged[key] = value
	}

	for key, value := range map2 {
		merged[key] = value
	}

	return merged
}

func hasExpand(c *fiber.Ctx, tag string) bool {
	args := strings.Split(c.Query("expand"), ",")

	for _, a := range args {
		if strings.EqualFold(a, tag) {
			return true
		}
	}
	return false
}
