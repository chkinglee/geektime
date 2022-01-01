package _1

func ModifyStringSlice(srcSlice, dstSlice []string) []string {
	for index := range srcSlice {
		srcSlice[index] = dstSlice[index]
	}
	return srcSlice
}
