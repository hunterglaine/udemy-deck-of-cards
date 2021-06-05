package main

import "testing"

func TestNewDeck(t *testing.T) {
	//t is our testing handler - who we tell something went wrong
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[len(d)-1])
	}
}
