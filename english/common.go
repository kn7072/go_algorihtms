package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"english/config"
)

func getAllWords() []string {
	pathToAllWords := config.ConfReader.GetString("path.path_to_all_words")
	tempSlice := make([]string, 0, 6000)

	fmt.Printf("File to read %v\n", pathToAllWords)
	if file, err := os.OpenFile(pathToAllWords, os.O_RDONLY, 0755); err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			line, errRead := reader.ReadString('\n')
			if errRead == nil {
				sliceLine := strings.Split(line, ";")
				line = sliceLine[0]
				tempSlice = append(tempSlice, strings.ReplaceAll(line, "\n", ""))
				// fmt.Println(strings.ReplaceAll(line, "\n", ""))
			} else if errRead == io.EOF {
				break
			} else {
				panic(errRead)
			}
		}
	} else {
		panic(err)
	}

	return tempSlice
}

func getSentence() map[string][]string {
	result := make(map[string][]string)

	pathToFileSentence := config.ConfReader.GetString("path.path_to_sentence")
	if file, errFile := os.OpenFile(pathToFileSentence, os.O_RDONLY, 0o755); errFile == nil {
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			line, errLine := reader.ReadString('\n')
			switch errLine {
			case nil:
				sliceW := strings.Split(line, ";")
				sentence := strings.TrimSpace(sliceW[1])
				translate := strings.TrimSpace(sliceW[2])
				tempWordsSlice := strings.Split(sliceW[0], ",")

				for _, wordI := range tempWordsSlice {
					wordI = strings.TrimSpace(wordI)
					if _, ok := result[wordI]; !ok {
						result[wordI] = make([]string, 0)
					}

					result[wordI] = append(result[wordI], fmt.Sprintf("%s;%s", sentence, translate))
				}
			case io.EOF:
				goto exit
			default:
				panic(errLine)
			}
		}
	} else {
		panic(errFile)
	}

exit:
	fmt.Println("Всего уникальных слов с примерами", len(result))

	return result
}

func createFileForAddSentence(words []string, wordsSentence map[string][]string) {
	pathToFile := config.ConfReader.GetString("path.file_to_add_sentence")
	countSentenceForWord := config.ConfReader.GetInt("path.count_sentence_for_word")

	if file, errFile := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY, 0o755); errFile == nil {
		defer file.Close()

		for _, wordI := range words {
			if sentence, ok := wordsSentence[wordI]; ok {
				if len(sentence) < countSentenceForWord {
					if _, err := file.WriteString(wordI + "\n"); err != nil {
						panic(err.Error())
					}
				}
			} else {
				if _, err := file.WriteString(wordI + "\n"); err != nil {
					panic(err.Error())
				}
			}
		}
	} else {
		panic(errFile.Error())
	}
}

func createFileLearnSentence() {

}

func getAllPathes(pathes ){
	ch := make(chan string)
	

	go func() {
		defer close(ch)
		files, err := os.ReadDir(rootDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, fileI := range files {

		}

	}()

}