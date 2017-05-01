package main

import (
	"fmt"

	"github.com/fatih/color"
)
/*Static class that handles cards. We represent cards as 32-bit integers, so 
there is no object instantiation - they are just ints. Most of the bits are 
used, and have a specific meaning. See below: 

								Card:

					  bitrank     suit rank   prime
				+--------+--------+--------+--------+
				|xxxbbbbb|bbbbbbbb|cdhsrrrr|xxpppppp|
				+--------+--------+--------+--------+

	1) p = prime number of rank (deuce=2,trey=3,four=5,...,ace=41)
	2) r = rank of card (deuce=0,trey=1,four=2,five=3,...,ace=12)
	3) cdhs = suit of card (bit turned on based on suit of card)
	4) b = bit turned on depending on rank of card
	5) x = unused

This representation will allow us to do very important things like:
- Make a unique prime prodcut for each hand
- Detect flushes
- Detect straights

and is also quite performant.*/

type csuit rune
type crank rune
type Card int32
type Suit uint16
type Rank uint16

const (
	Club    Suit = 0x8000
	Diamond      = 0x4000
	Heart        = 0x2000
	Spade        = 0x1000
)

const (
	Deuce Rank = iota
	Trey
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	strRanks = "23456789TJQKA"
)

var IntRanks = makeRanks(1, 13)
var Primes = [13]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}

var CharSuitToIntSuit = map[csuit]Suit{
		's' : Spade, // spades
		'h' : Heart, // hearts
		'd' : Diamond, // diamonds
		'c' : Club, // clubs
}

var CharRankToIntRank = map[crank]Rank{
	'2': Deuce,
	'3': Trey,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Jack,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

var strSuits = map[Suit]string {
		Spade : "s", // spades
		Heart : "h", // hearts
		Diamond : "d", // diamonds
		Club : "c", // clubs
}
//for pretty printing
var PrettySuits = map[Suit]string {
		Spade : "\u2660", // spades
		Heart : "\u2764", // hearts
		Diamond : "\u2666", // diamonds
		Club : "\u2663", // clubs
}

//hearts and diamonds
var PrettyReds = [2]Suit{Heart, Diamond}

/*Converts Card string to binary integer representation of card, inspired by:
http://www.suffecool.net/poker/evaluator.html*/
func NewCard(value string) Card {
	runes := []rune(value)
	rankChar := crank(runes[0])
	suitChar := csuit(runes[1])
	rank := CharRankToIntRank[rankChar]
	suit := CharSuitToIntSuit[suitChar]

	bitrank := 1 << uint(16 + rank)


	n := Primes[rank] | (int(rank) << 8) | int(suit) | bitrank
	return Card(n)
}

func (c Card) String() string {
	return string(strRanks[c.Rank()]) + strSuits[c.Suit()]
}

func (c Card) Rank() Rank {
	return Rank((c >> 8) & 0x0f)
}

func (c Card) Suit() Suit {
	return Suit(c & 0xf000)
}

func (c Card) BitRank() int {
	return int(c) >> 16
}

func (c Card) Prime() int {
	return int(c) & 0xff
}

//Expects a list of cards in integer form. 
/*func primeProductFromHand(cards []Card) (product int) {
	product = 1
	for _ , c := range(cards) {
		product *= (c & 0xFF)
	}
	return
}
*/

//Prints a single card
func (c Card) PrettyStr() string {

	colorEnabled := true

	//suit and rank
	rank := c.Rank()
	suit := c.Suit()

	//if we need to color red
	s := PrettySuits[suit]
	red := color.New(color.FgRed).SprintFunc()
	if (colorEnabled && isRedSuit(suit, PrettyReds)) {
		s = red(s)
	}

	r := strRanks[rank]

	return " [ " + string(r) + " " + string(s) + " ] "
}

// Expects a list of cards in integer form.
func PrintPrettyCards(cards []Card) string {
	output := " "
	for index, c := range(cards) {
		if index != len(cards) - 1 {
			output += c.PrettyStr() + ","
		} else {
			output += c.PrettyStr() + " "
		}
	}

	return fmt.Sprintf("%s", output)
}
