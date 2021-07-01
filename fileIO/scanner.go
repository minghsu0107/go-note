package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
func (s *Scanner) Bytes() []byte
func (s *Scanner) Err() error
func (s *Scanner) Scan() bool
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Text() string
*/
func main() {
	var filename string
	fmt.Print("檔案名稱：")
	fmt.Scanf("%s", &filename)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // return string
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
