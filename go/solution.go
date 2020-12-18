package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"reflect"

	"runtime"
	"strings"

	prose "github.com/jdkato/prose/v2"
	porterstemmer "github.com/reiver/go-porterstemmer"
)

type PlagiarismChecker struct {
	HashTable map[string][]int
	KGram     int
}

func NewPlagiarismChecker(fileA, fileB string) *PlagiarismChecker {
	checker := &PlagiarismChecker{
		KGram:     5,
		HashTable: make(map[string][]int),
	}
	checker.GenerateFileHash(checker.GetFileContent(fileA), "a")
	checker.GenerateFileHash(checker.GetFileContent(fileB), "b")

	return checker
}

func (pc PlagiarismChecker) PrepareContent(content string) string {
	doc, err := prose.NewDocument(content)
	if err != nil {
		log.Fatal(err)
	}
	data := []string{}
	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		if !Contains(stopwords, tok.Text) {
			stem := porterstemmer.StemString(tok.Text)
			data = append(data, stem)
		}
	}

	return strings.Join(data[:], "")
}

func (pc PlagiarismChecker) GenerateFileHash(content, docType string) {
	text := pc.PrepareContent(content)
	fmt.Println(text)

	textRolling := NewRabinKarp(strings.ToLower(text), pc.KGram)

	for i := 0; i <= len(text)-pc.KGram+1; i++ {
		if len(pc.HashTable[docType]) == 0 {
			pc.HashTable[docType] = []int{textRolling.Hash}
		} else {
			pc.HashTable[docType] = append(pc.HashTable[docType], textRolling.Hash)
		}
		if textRolling.NextWindow() == false {
			break
		}
	}

	fmt.Println(pc.HashTable)
}

func (pc PlagiarismChecker) GetRate() float64 {
	return pc.CalculatePlagiarismRate()
}

func (pc PlagiarismChecker) GetFileContent(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (pc PlagiarismChecker) CalculatePlagiarismRate() float64 {

	THA := len(pc.HashTable["a"])
	THB := len(pc.HashTable["b"])
	intercet := Intersect(pc.HashTable["a"], pc.HashTable["b"])
	SH := reflect.ValueOf(intercet).Len()

	// Formular for plagiarism rate
	// P = (2 * SH / THA * THB ) 100%
	p := float64(2 * SH)/ float64(THA * THB)
	return float64(p* 100)
}

func main() {
	_, b, _, _ := runtime.Caller(0)
	//get root
	Root := filepath.Join(filepath.Dir(b), "../")
	docsPath := path.Join(Root, "/docs/")
	checker := NewPlagiarismChecker(
		path.Join(docsPath, "/document_a.txt"),
		path.Join(docsPath, "/document_b.txt"),
	)
	fmt.Printf("The percentage of plagiarism held by both documents is  %f", checker.GetRate())
}
