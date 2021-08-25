package utils

import (
	"testing"
)

func TestNew (t *testing.T) {

	{
		trackA, er := New(1, "artist_a", "album_a", "track_a")
		if er != nil || trackA.TrackId != 1 || trackA.Artist != "artist_a" || trackA.Album != "album_a" || trackA.TrackName != "track_a" {
			t.Errorf("Wronng Track created")
		}
	}
	{
		trackAA, er := New(2, "artist_a", "album_a", "track_aa")
		if er != nil || trackAA.TrackId != 2 || trackAA.Artist != "artist_a" || trackAA.Album != "album_a" || trackAA.TrackName != "track_aa" {
			t.Errorf("Wronng Track created")
		}
	}
	{
		trackAAA, er := New(3, "artist_a", "album_aa", "track_aa")
		if er != nil || trackAAA.TrackId != 3 || trackAAA.Artist != "artist_a" || trackAAA.Album != "album_aa" || trackAAA.TrackName != "track_aa" {
			t.Errorf("Wronng Track created")
		}
	}
	{
		trackB, _ := New(4, "artist_b", "album_b", "track_b")
		if trackB.TrackId != 4 || trackB.Artist != "artist_b" || trackB.Album != "album_b" || trackB.TrackName != "track_b" {
			t.Errorf("Wronng Track created")
		}
	}

	_, err := New (1, "artist_a","album_a", "track_a")
		if err == nil {
			t.Errorf("Track cannot be created twice")
		}
}

func createLibrary() {
	 New(1, "artist_a", "album_a", "track_a")
	 New(2, "artist_a", "album_a", "track_aa")
	 New(3, "artist_a", "album_aa", "track_aa")
	 New(4, "artist_b", "album_b", "track_b")

}

func exists (tracks *[]Track, track *Track) bool{
	for _, curTrack := range *tracks{
		//if curTrack.TrackName == track.TrackName && curTrack.Artist == track.Artist && curTrack.Album == track.Album
		if curTrack == *track{return true}
	}
	return false
}

func TestGetArtistTracks(t *testing.T) {

	createLibrary()
	{
		var tracksOfArtistA = GetArtistTracks("artist_a")
		if len (tracksOfArtistA) != 3 {
			t.Errorf("invalid amount tracks of artist_a")
		}
		if (!exists(&tracksOfArtistA, &Track{1, "track_a", "album_a", "artist_a"})) {
			t.Errorf("cant find track a")
		}
		if (!exists(&tracksOfArtistA, &Track{2, "track_aa", "album_a", "artist_a"})) {
			t.Errorf("cant find track aa")
		}
		if (!exists(&tracksOfArtistA, &Track{3, "track_aa", "album_aa", "artist_a"})) {
			t.Errorf("cant find track aa. album aa.")
		}
		var tracksOfArtistB = GetArtistTracks("artist_b")
		if len (tracksOfArtistB) != 1 {
			t.Errorf("invalid amount tracks of artist_b")
		}
		if (!exists(&tracksOfArtistB, &Track{4, "track_b", "album_b", "artist_b"})) {
			t.Errorf("cant find track b")
		}
	}
}

func TestGetArtistAlbumTracks(t *testing.T) {

	createLibrary()
	{
		tracksOfArtistAAlbumA := GetArtistAlbumTracks("artist_a", "album_a")
		if len (tracksOfArtistAAlbumA) != 2 {
			t.Errorf("invalid amount tracks of artist_a album_a")
		}
		if (!exists(&tracksOfArtistAAlbumA, &Track{1, "track_a", "album_a", "artist_a"})) {
			t.Errorf("cant find track a")
		}
		if (!exists(&tracksOfArtistAAlbumA, &Track{2, "track_aa", "album_a", "artist_a"})) {
			t.Errorf("cant find track aa")
		}
	}
	{
		tracksOfArtistAAlbumAA := GetArtistAlbumTracks("artist_a", "album_aa")
		if len (tracksOfArtistAAlbumAA) != 1 {
			t.Errorf("invalid amount tracks of artist_a album_aa")
		}
		if (!exists(&tracksOfArtistAAlbumAA, &Track{3, "track_aa", "album_aa", "artist_a"})) {
			t.Errorf("cant find track a")
		}
	}
	{
		tracksOfArtistBAlbumB := GetArtistAlbumTracks("artist_b", "album_b")
		if len (tracksOfArtistBAlbumB) != 1 {
			t.Errorf("invalid amount tracks of artist_b album_b")
		}
		if (!exists(&tracksOfArtistBAlbumB, &Track{4, "track_b", "album_b", "artist_b"})) {
			t.Errorf("cant find track b")
		}
	}
}

func TestInitLibraryFromArray(t *testing.T) {

	{
		resetLibrary()
		lines := make([]string, 0)
		lines = append(lines, string("artist_a, album_a, track_a"))
		lines = append(lines, string("artist_a, album_a, track_aa"))

		tracks := initLibraryFromArray(lines)
		if tracks == nil || len (tracks) != 2 {
			t.Errorf("lines are not parsed")
		}
	}
}
