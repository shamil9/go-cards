package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	stringBytes := []byte(d.toString())

	return ioutil.WriteFile(filename, stringBytes, 0666)
}

func newDeckFromFile(filename string) deck {
	stringBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	stringSlices := strings.Split(string(stringBytes), ",")

	return deck(stringSlices)
}

func (d deck) shuffle() {
	r := setRandomSeed()

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func setRandomSeed() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())

	return rand.New(seed)
}
