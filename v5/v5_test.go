package openrec_test

import (
	"github.com/KuroseNil/OPENREC-PublicAPI-v5-go/v5"
	"testing"
)

func TestGetMovies(t *testing.T) {
	movie, err := openrec.GetMovies(true, openrec.SORT_LIVE_VIEWS, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Get %d Movies...\n", len(movie))
}

func TestGetMoviesByChannelID(t *testing.T) {
	chanID := "shizurin23"
	movie, err := openrec.GetMoviesByChannelID(chanID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Get %d Movies...\n", len(movie))
}
