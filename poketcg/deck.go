package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func InitSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type Card struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Details     []string `json:"details"`
	Attachments *Deck    `json:"-"`
}

func (c *Card) String() string {
	return c.Name
}

func (c *Card) ShortString() string {
	return c.Name
}

func (c *Card) LongString() string {
	return fmt.Sprintf("ID: %s\nName: %s\nDetails: %s\nAttachments: %s", c.Id, c.Name, strings.Join(c.Details, "\n"), c.Attachments.List())
}

type Deck struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Cards []*Card `json:"cards"`
}

func (d *Deck) Shuffle() {
	N := len(d.Cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
	}
}

func (d *Deck) String() string {
	return fmt.Sprintf("ID: %s\nName: %s\nCards: %s", d.Id, d.Name, d.List())
}

func (d *Deck) List() string {
	ret := ""
	for idx, c := range d.Cards {
		ret = fmt.Sprintf("%s[%d] %s\t", ret, idx, c.ShortString())
	}
	return ret
}

func (d *Deck) Move(idx int, d2 *Deck) {
	newCards := []*Card{}
	for i, card := range d.Cards {
		if i == idx {
			d2.Append(card)
			continue
		}
		newCards = append(newCards, card)
	}
	d.Cards = newCards
}

func (d *Deck) Append(c *Card) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Prepend(c *Card) {
	d.Cards = append([]*Card{c}, d.Cards...)
}
