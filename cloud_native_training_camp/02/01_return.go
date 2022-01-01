package _2

import "fmt"

func Return1() error {
	err := new(error)
	return *err
}

func Return2() (err error) {
	err = fmt.Errorf("this is an error")
	return err
}
