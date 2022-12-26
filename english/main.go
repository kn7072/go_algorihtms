package main

import (
	"fmt"

	"english/config"
)

func main() {
	// viper := config.ConfReader.GetViper()
	all_words := getAllWords()
	fmt.Printf("%T\n", all_words)

	sentencesMap := getSentence()
	// fmt.Println(sentences)

	createFileForAddSentence(all_words, sentencesMap)
	rootDir := config.ConfReader.GetString("dir_all_word_files")
	ch := getAllPathes(rootDir)
	<-ch 
}
