package main

func main() {
	words := []string{
		"helloiEAthura",
		"kura fpl",
		"fpl kura fpl",
		"kura fpl",
	}

	for _, word := range words {
		result := rotateVowels(word)
		println(word)
		println(result)
		println()
	}
}

func rotateVowels(s string) string {
	sb := []byte(s)

	mid := len(sb) / 2
	if len(s)%2 == 1 {
		mid++
	}
	index := len(sb) - 1

	for i := 0; i <= index; i++ {
		if isVowel(sb[i]) {
			for j := index; j > i; j-- {
				if isVowel(sb[j]) {
					sb[i], sb[j] = sb[j], sb[i]
					index = j - 1
					break
				}
			}
		}
	}
	return string(sb)
}

func isVowel(ch byte) (isvowel bool) {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		isvowel = true
	}
	return
}
