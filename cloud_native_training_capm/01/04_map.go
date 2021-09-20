package _1

func Map() map[string]interface{} {
	myMap := make(map[string]interface{}, 10)
	myMap["string"] = "abc"
	myMap["int"] = 123
	myMap["bool"] = true
	return myMap
}
