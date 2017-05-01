package main

func makeRanks(min, max Rank) []Rank {
    a := make([]Rank, max-min+1)
    for i := range a {
        a[i] = min + Rank(i)
    }
    return a
}

func isRedSuit(a Suit, redSuits [2]Suit) bool {
    for _, red := range redSuits {
        if red == a {
            return true
        }
    }
    return false
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

