package main

func medgeBanner(base priority map[rune][]string)map[rune][]string{
	result := map[rune][]string{}
	
	for key, value := range base{
		result[key] = value
	}
	for key, value := range priority{
		result[key] = value
	}
	return result
}
