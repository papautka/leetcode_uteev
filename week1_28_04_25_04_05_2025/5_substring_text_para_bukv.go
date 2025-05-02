package week1_28_04_25_04_05_2025

import "fmt"

//
//func FindSubstrig(text string) string {
//	paraSymbol := ""
//	for _, charch := range text {
//		if strings.Contains(rune(charch), substring)
//	}
//	return substring
//}

// 3.
func FinalBreaks(s string) string {
	arrWord := BreaksTextOnWord(s)
	sliceSlog := make([]string, 0)
	for _, word := range arrWord {
		nowSlice := BreaksWordOnSlog(word)
		sliceSlog = append(sliceSlog, nowSlice...)
	}
	fmt.Println(sliceSlog) // "MAMA VSKF" -> [MA AM MA VS SK KF]

	// ищем уникальную подстроку
	counts := make(map[string]int)

	for _, word := range sliceSlog {
		counts[word]++
	}

	maxCount := 0
	result := ""

	for pair, count := range counts {
		if count > maxCount || (count == maxCount && pair > result) {
			maxCount = count
			result = pair
		}
	}

	return result
}

// 2. Разбивает слово на подстроки где каждая состоит из 1-2-x букв
// A -> A
// VI -> VI
// МАМ -> MA AM
// MAMA -> MA AM MA
func BreaksWordOnSlog(word string) []string {
	arrOfSubstring := make([]string, 0)
	if len(word) == 1 || len(word) == 2 {
		arrOfSubstring = append(arrOfSubstring, word)
		return arrOfSubstring
	}
	lastSymb := word[0]
	emptyWord := string(lastSymb)
	for i := 1; i < len(word); i++ {
		ch := word[i]
		emptyWord += string(ch)
		if len(emptyWord) == 2 {
			arrOfSubstring = append(arrOfSubstring, emptyWord)
			lastSymb = word[i]
			emptyWord = string(lastSymb)
		}
	}
	// добавляем последний слог
	if len(emptyWord) == 2 {
		arrOfSubstring = append(arrOfSubstring, emptyWord)
	}

	return arrOfSubstring
}

// 1. разбивает текст на слова и возвращает массив слов
func BreaksTextOnWord(text string) []string {
	arrayOfWords := make([]string, 0)
	emptyWord := ""
	for i := 0; i < len(text); i++ {
		ch := text[i]
		if ch == ' ' || ch == '\t' {
			if len(emptyWord) > 0 {
				arrayOfWords = append(arrayOfWords, emptyWord)
				emptyWord = ""
			}
		} else {
			emptyWord += string(ch)
		}
	}

	// добавляем последнее слово
	arrayOfWords = append(arrayOfWords, emptyWord)

	return arrayOfWords
}
