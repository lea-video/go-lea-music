package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lea-video/go-lea-music/db"
	"github.com/lea-video/go-lea-music/model"
)

const (
	ExpandMedia  = "media"
	ExpandTracks = "tracks"
	ExpandArtist = "artists"
)

func doExpansions(c *fiber.Ctx, db db.GenericDB, resp model.ResponseObject) (model.ResponseObject, error) {
	// collect ids to expand
	toResolve := collectExpansionIDs(c, resp)

	// filter out those already present
	toResolve[ExpandMedia] = filterUniqueness(toResolve[ExpandMedia], resp.Media)
	toResolve[ExpandTracks] = filterUniqueness(toResolve[ExpandTracks], resp.MediaTrack)
	toResolve[ExpandArtist] = filterUniqueness(toResolve[ExpandArtist], resp.Artists)

	// do expansions
	resp, err := doExpansion(resp, db, toResolve)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func collectExpansionIDs(c *fiber.Ctx, resp model.ResponseObject) map[string][]int {
	toResolve := make(map[string][]int)

	mediaIDs := make([]int, 0)
	if hasExpand(c, ExpandMedia) {
		for _, m := range resp.MediaTrack {
			mediaIDs = append(mediaIDs, m.ParentMedia)
		}
		for _, t := range resp.TMPFile {
			mediaIDs = append(mediaIDs, t.ParentMedia)
		}
	}
	toResolve[ExpandMedia] = mediaIDs

	trackIDs := make([]int, 0)
	if hasExpand(c, ExpandTracks) {
		for _, m := range resp.Media {
			trackIDs = append(trackIDs, m.Tracks...)
		}
	}
	toResolve[ExpandTracks] = trackIDs

	artistIDs := make([]int, 0)
	if hasExpand(c, ExpandArtist) {
		for _, m := range resp.Artists {
			trackIDs = append(artistIDs, m.Members...)
		}
	}
	toResolve[ExpandArtist] = artistIDs

	return toResolve
}

func doExpansion(resp model.ResponseObject, db db.GenericDB, toResolve map[string][]int) (model.ResponseObject, error) {
	if len(toResolve[ExpandMedia]) > 0 {
		medias, err := db.GetMediaByID(toResolve[ExpandMedia])
		if err != nil {
			return resp, err
		}
		resp.Media = mergeMaps(resp.Media, medias)
	}

	if len(toResolve[ExpandTracks]) > 0 {
		tracks, err := db.GetMediaTracksByID(toResolve[ExpandTracks])
		if err != nil {
			return resp, err
		}
		resp.MediaTrack = mergeMaps(resp.MediaTrack, tracks)
	}

	if len(toResolve[ExpandArtist]) > 0 {
		artists, err := db.GetArtistsByID(toResolve[ExpandArtist])
		if err != nil {
			return resp, err
		}
		resp.Artists = mergeMaps(resp.Artists, artists)
	}

	return resp, nil
}

func filterUniqueness[T any](toResolve []int, present map[int]T) []int {
	// Create a map to store the present values as a set
	presentSet := make(map[int]bool)
	for key := range present {
		presentSet[key] = true
	}

	// Second lookup-table to ensure uniqueness
	seen := make(map[int]bool)

	// Filter the toResolve list
	filtered := make([]int, 0)
	for _, value := range toResolve {
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
