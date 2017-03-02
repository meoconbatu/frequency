package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(words []string) (m FreqMap) {
	m = FreqMap{}
	lenWords := len(words)
	resultChan := make(chan FreqMap, lenWords)
	for _, word := range words {
		go func(w string, channel chan FreqMap) {
			resultChan <- Frequency(w)
		}(word, resultChan)
	}
	for i := 0; i < lenWords; i++ {
		result := <-resultChan
		for r, f := range result {
			m[r] += f
		}
	}
	return

}
