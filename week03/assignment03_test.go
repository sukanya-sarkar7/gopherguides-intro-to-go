package main

import (
	"fmt"
	"testing"
)

func TestPlays(t *testing.T) {
	m := Movie{
		Length:  3,
		Name:    "Titanic",
		rating:  1.2,
		play:    1,
		viewers: 50,
	}
	//viewers := 50
	expPlays := (Movie).Plays(m)
	fmt.Println(expPlays)
	if expPlays != 1 {
		t.Errorf("expected plays doesn't match with the value of actual plays")
	}
}

func TestViewers(t *testing.T) {
	m := Movie{
		Length:  3,
		Name:    "Titanic",
		rating:  1.2,
		play:    0,
		viewers: 50,
	}
	//viewers := 50
	expViewer := (Movie).Viewers(m)
	fmt.Println(expViewer)
	if expViewer != 50 {
		t.Errorf("expected viewers doesn't match with the value of actual viewers")
	}

}
func TestRating(t *testing.T) {
	m := Movie{
		Length:  3,
		Name:    "Titanic",
		rating:  1.2,
		play:    1,
		viewers: 50,
	}
	totalRating := (Movie).Rating(m)
	fmt.Println(totalRating)
}

func TestCritiqueFn(t *testing.T) {
	m := Movie{
		Length:  3,
		Name:    "Titanic",
		rating:  1.2,
		play:    0,
		viewers: 0,
	}

	res, err := CritiqueFn(&m)
	fmt.Println(res)
	fmt.Println(err)
	if res != 6.89 {
		t.Errorf("can't review a movie without watching it first")

	}

}
