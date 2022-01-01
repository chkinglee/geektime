package _2

import "fmt"

type LifeSkills interface {
	Eat()
	Talk()
}

type AnimalWorkSkills interface {
	Hunt()
}

type HumanWorkSkills interface {
	Build()
}

type WorkSkills interface {
	AnimalWorkSkills
	HumanWorkSkills
}

type cat struct {
}
type human struct {
}

func (c *cat) Eat() {
	fmt.Println("我吃猫粮")
}

func (c *cat) Talk() {
	fmt.Println("喵喵喵")
}

func (c *cat) Hunt() {
	fmt.Println("我可以抓老鼠")
}

func (h *human) Eat() {
	fmt.Println("我不吃猫粮")
}

func (h *human) Talk() {
	fmt.Println("嘿嘿嘿")
}

func (h *human) Hunt() {
	fmt.Println("我可以抓坏人")
}
