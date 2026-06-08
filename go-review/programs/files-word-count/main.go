package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/Sheikh-Fahad-Ahmed/syntax-studies/go-review/programs/files-word-count/helper"
)

func worker(i int, wg *sync.WaitGroup, wordCounts chan<- int) {
	defer wg.Done()

	file, err := os.Open(fmt.Sprintf("file%d.txt", i))
	if err != nil {
		fmt.Printf("error opening file%d: %v\n", i, err)
		return
	}
	defer file.Close()
	fmt.Printf("file%d opened....counting words now..\n", i)

	scanner := bufio.NewScanner(file)
	count := helper.Read(scanner)

	fmt.Printf("file%d closing...total words: %d\n", i, count)
	wordCounts <- count
}

func main() {
	wordCountChan := make(chan int, 4)
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)

		go worker(i, &wg, wordCountChan)
	}

	wg.Wait()
	close(wordCountChan)

	sum := 0
	for count := range wordCountChan {
		sum += count

	}
	fmt.Println("all files opened and closed")
	fmt.Printf("\n\nThe sum of words in all the files: %d", sum)

}
