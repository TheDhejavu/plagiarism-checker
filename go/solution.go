package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type PlagiarismChecker struct {
	HashTable map[string][]int
	KGram     int
}

func NewPlagiarismChecker(fileA, fileB string) *PlagiarismChecker {
	checker := &PlagiarismChecker{
		KGram: 2,
	}
	checker.GenerateFileHash(checker.GetFileContent(fileA), "a")
	checker.GenerateFileHash(checker.GetFileContent(fileB), "b")

	return checker
}

func (pc PlagiarismChecker) PrepareContent(content string) string {
	return content
}

func (pc PlagiarismChecker) GenerateFileHash(content, docType string) {
	text := pc.PrepareContent(content)

	textRolling := NewRabinKarp(strings.ToLower(text), pc.KGram)
	pc.HashTable = make(map[string][]int)

	for i := 0; i <= len(text)-pc.KGram+1; i++ {
		if len(pc.HashTable[docType]) == 0 {
			pc.HashTable[docType] = []int{textRolling.Hash}
		} else {
			pc.HashTable[docType] = append(pc.HashTable[docType], textRolling.Hash)
		}
		fmt.Println(textRolling.Hash)
		if textRolling.NextWindow() == false {
			break
		}
	}

	fmt.Println(pc.HashTable)
}

func (pc PlagiarismChecker) GetRate() float64 {
	return 0.0
}

func (pc PlagiarismChecker) GetFileContent(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (pc PlagiarismChecker) CalculatePlagiarismRate() {

}

func main() {
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../")
	docsPath := path.Join(Root, "/docs/")
	checker := NewPlagiarismChecker(
		path.Join(docsPath, "/document_a.txt"),
		path.Join(docsPath, "/document_b.txt"),
	)
	fmt.Printf("The percentage of plagiarism held by both documents is  %f", checker.GetRate())
}
