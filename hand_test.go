package main

import (
	"testing"
)

var handtests = []struct {
	in  *Hand
	value HandValue
}{
	//Royal Flush
	{
		NewHand([5]Card{
			NewCard("As"),
			NewCard("Ks"),
			NewCard("Qs"),
			NewCard("Js"),
			NewCard("Ts"),
		}), 
		1,
	},
	//Straight Flush
	{
		NewHand([5]Card{
			NewCard("Ks"),
			NewCard("Qs"),
			NewCard("Js"),
			NewCard("Ts"),
			NewCard("9s"),
		}), 
		2,
	},
	//4 of a kind
	{
		NewHand([5]Card{
			NewCard("As"),
			NewCard("Ah"),
			NewCard("Ad"),
			NewCard("Ac"),
			NewCard("Ks"),
		}), 
		11,
	},
	//Full House
	{
		NewHand([5]Card{
			NewCard("As"),
			NewCard("Ah"),
			NewCard("Ad"),
			NewCard("Kc"),
			NewCard("Ks"),
		}), 
		167,
	},
	//Flush
	{
		NewHand([5]Card{
			NewCard("As"),
			NewCard("Ks"),
			NewCard("Qs"),
			NewCard("Js"),
			NewCard("9s"),
		}), 
		323,
	},
}

func TestHandValue(t *testing.T) {

	for _, hand := range handtests {
		value := hand.in.EvalFast()
		if value != hand.value {
			t.Fatalf("%s hand value %d does not match expected value %d", hand, hand.value, value)
		}
	}
}

func TestSameSuit(t *testing.T) {
	suitedHand := NewHand([5]Card{
			NewCard("As"),
			NewCard("Ks"),
			NewCard("Qs"),
			NewCard("Js"),
			NewCard("9s"),
	})

	unsuitedHand := NewHand([5]Card{
			NewCard("Ad"),
			NewCard("Ks"),
			NewCard("Qs"),
			NewCard("Js"),
			NewCard("9s"),
	})

	if ! suitedHand.SameSuit() {
		t.Fatalf("%s hand is suited, wrongly classified as unsuited", suitedHand)
	}

	if unsuitedHand.SameSuit() {
		t.Fatalf("%s hand is unsuited, wrongly classified as unsuited", unsuitedHand)
	}
}
