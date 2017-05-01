package main

import "fmt"

type HandValue uint16
	
type Hand struct {
	Cards [5]Card
}

func findfast(p uint32) uint32 {
    var a, b, r uint32
    p += 0xe91aaa35
    p ^= p >> 16
    p += p << 8
    p ^= p >> 4
    b  = (p >> 8) & 0x1ff
    a  = (p + (p << 2)) >> 19
    r  = a ^ hashAdjust[b]
    return r
}

func NewHand(cards [5]Card) *Hand {
	return &Hand{Cards: cards}
}

func (h *Hand) String() string{
	return fmt.Sprintf("[ %s %s %s %s %s ]", h.Cards[0], h.Cards[1], h.Cards[2], h.Cards[3], h.Cards[4])
}

func (h *Hand) SameSuit() bool {
	return h.Cards[0] & h.Cards[1] & h.Cards[2] & h.Cards[3] & h.Cards[4] & 0xf000 !=0
}

func (h *Hand) PrimesProduct() uint32 {
	return uint32((h.Cards[0] & 0xff) * (h.Cards[1] & 0xff) * (h.Cards[2] & 0xff) * (h.Cards[3] & 0xff) * (h.Cards[4] & 0xff))
}

func (h *Hand) Bitwise() int {
	return int(h.Cards[0] | h.Cards[1] | h.Cards[2] | h.Cards[3] | h.Cards[4]) >> 16
}

func (h *Hand) EvalFast() HandValue {

	bits := h.Bitwise()
	if h.SameSuit() {
		return Flushes[bits]
	}

	if s := Unique5[bits]; s != 0 {
        		return s
    	}
    	
    	return hashValues[findfast(h.PrimesProduct())]
}

