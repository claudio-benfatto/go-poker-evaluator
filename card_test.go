package main

import (
	"testing"
)

var cardtests = []struct {
	in  string
	prime int
	rank Rank
	suit Suit
}{
	{"2s", 2, Deuce, Spade},
	{"3h", 3, Trey, Heart},
	{"4d", 5, Four, Diamond},
	{"5c", 7, Five, Club},
	{"6s", 11, Six, Spade},
	{"7h", 13, Seven, Heart},
	{"8d", 17, Eight, Diamond},
	{"9c", 19, Nine, Club},
	{"Ts", 23, Ten, Spade},
	{"Jh", 29, Jack, Heart},
	{"Qd", 31, Queen, Diamond},
	{"Kc", 37, King, Club},
}

func TestNewCard(t *testing.T) {
	cards := map[string]Card{
		"Kd": 134236965,
		"5s": 529159,
		"Jc": 33589533,
	}
	for card, value := range cards {
		c := NewCard(card)
		if c != value {
			t.Fatalf("%s card value %d does not match expected value %d", card, c, value)
		}
	}
}


func TestPrime(t *testing.T) {
	for _, tc := range cardtests {
		c := NewCard(tc.in)
		if c.Prime() != tc.prime {
			t.Fatalf("%s prime %d does not match expected prime %d", tc.in, c.Prime(), tc.prime)
		}
	}
}

func TestRank(t *testing.T) {
	for _, tc := range cardtests {
		c := NewCard(tc.in)
		if c.Rank() != tc.rank {
			t.Fatalf("%s rank %d does not match expected rank %d", tc.in, c.Rank(), tc.rank)
		}
	}
}

func TestBitRank(t *testing.T) {
	for _, tc := range cardtests {
		c := NewCard(tc.in)
		if c.BitRank() != int(1 << tc.rank) {
			t.Fatalf("%s bitrank %b does not match expected bitrank %b", tc.in, c.BitRank(), uint( 1 << tc.rank))
		}
	}
}

func TestSuit(t *testing.T) {
	for _, tc := range cardtests {
		c := NewCard(tc.in)
		if c.Suit() != tc.suit {
			t.Fatalf("%s suit %d does not match expected suit %d", tc.in, c.Suit(), tc.suit)
		}
	}
}
