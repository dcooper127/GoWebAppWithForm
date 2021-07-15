package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type myValues struct {
	Input1 string
	Input2 string
	Result string
}

//Uses mapping for an 0(n) solution to see if two words are anagrams of each other.
//Strategy: uses (character, occurrence) mapping to count the occurrence of each character in the first word and subtracts
//the occurrence of each character in the second word. If any character in the map has a non 0 value then
//it appeared in one word and not the other, therefore the words are not anagrams of each other.
func isAnagram(word1 string, word2 string) bool {
	var letterCounts map[string]int = make(map[string]int)

	//lowercase both words
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	//

	//itterate over word 1
	for _, char := range word1 {
		charCount, present := letterCounts[string(char)]
		if !(present) {
			letterCounts[string(char)] = 1
		} else {
			letterCounts[string(char)] = charCount + 1
		}
	}

	for _, char := range word2 {
		charCount, present := letterCounts[string(char)]
		if !(present) {
			letterCounts[string(char)] = -1
		} else {
			letterCounts[string(char)] = charCount - 1
		}
	}

	//itterate over word 2
	for _, v := range letterCounts {
		if v != 0 {
			return false
		}
	}

	return true
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		word1 := ""
		word2 := ""
		result := ""

		values := myValues{Input1: word1, Input2: word2, Result: result}

		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, values)
	} else {
		r.ParseForm()

		//word1 := r.Form["ip1"][0]
		//word2 := r.Form["ip2"][0]

		word1 := "test"
		word2 := "best"

		output := isAnagram(word1, word2)

		outputString := ""

		if output {
			outputString = word1 + " and " + word2 + " are anagrams."
		} else {
			outputString = word1 + " and " + word2 + "are not anagrams."
		}

		values := myValues{Input1: word1, Input2: word2, Result: outputString}

		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, values)
	}

}

func main() {
	http.HandleFunc("/", testFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
