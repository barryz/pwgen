package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	DefaultPassLength = 16
)

var (
	chars = []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0987654321")

	// Flags
	size  = flag.Int("s", DefaultPassLength, "password length, default is 16")
	count = flag.Int("c", 1, "numbers of password to generation")

	wg sync.WaitGroup
)

func init() {
	flag.Usage = usage
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
}

func usage() {
	fmt.Fprintf(os.Stderr, "pwgen is CLI tool for generate random password\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "pwgen [option]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var randomPassword string

	for i := 1; i <= *count; i++ {
		wg.Add(1)
		randomPassword = ""

		go func(i int) {
			defer wg.Done()
			if *size > 0 {
				randomPassword = genPasswordWithLength(*size)
			} else {
				randomPassword = genPasswordWithLength(DefaultPassLength)
			}
			fmt.Printf("%s\n", randomPassword)
		}(i)
	}

	wg.Wait()
}

func genPasswordWithLength(length int) string {
	charsArray := make([]rune, length)
	for i := range charsArray {
		charsArray[i] = chars[next(0, len(chars))]
	}
	return string(charsArray)
}

func next(min, max int) int {
	return min + rand.Intn(max-min)
}
