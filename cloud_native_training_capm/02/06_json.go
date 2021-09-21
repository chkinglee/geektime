package _2

import "encoding/json"

type dog struct {
	Name   string
	Gender int
}

func unmarshal2Struct(dogStr string) dog {
	d := dog{}
	err := json.Unmarshal([]byte(dogStr), &d)
	if err != nil {
		print(err)
	}
	return d
}

func marshal2String(d dog) string {
	b, err := json.Marshal(&d)
	if err != nil {
		print(err)
	}
	return string(b)
}
