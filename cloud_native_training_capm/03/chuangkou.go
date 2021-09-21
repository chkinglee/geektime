package _3

import (
	"fmt"
)

func chuangkou() {
	a := "abced"
	b := "aefabwefwfabcwbcedzcabcex"

	windowSize := 0
	matchString := ""
	for startIndexA := 0; startIndexA < len(a); startIndexA++ {
		for endIndexA := startIndexA + 1; endIndexA <= len(a); endIndexA++ {
			//fmt.Println(a[startIndexA:endIndexA])
			//if strings.Contains(b, a[startIndexA:endIndexA]) {
			//	if len(a[startIndexA:endIndexA]) > windowSize {
			//		windowSize = len(a[startIndexA:endIndexA])
			//		matchString = a[startIndexA:endIndexA]
			//	}
			//}
			for startIndexB := 0; startIndexB+endIndexA-startIndexA <= len(b); startIndexB++ {
				if a[startIndexA:endIndexA] == b[startIndexB:startIndexB+endIndexA-startIndexA] {
					if len(a[startIndexA:endIndexA]) > windowSize {
						windowSize = len(a[startIndexA:endIndexA])
						matchString = a[startIndexA:endIndexA]
					}
				}
			}
		}
	}
	fmt.Println(windowSize, " ", matchString)
}
