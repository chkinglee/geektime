package _2

import "testing"

func TestCat_Human(t *testing.T) {
	c := cat{}
	c.Eat()
	c.Talk()
	c.Hunt()

	h := human{}
	h.Eat()
	h.Talk()
	h.Hunt()
}
