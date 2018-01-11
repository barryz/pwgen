package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	DefaultPassLength = 16
)

var (
	chars = []rune("QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm0987654321!@#$%^&*()_+=-")

	// Flags
	size = flag.Int("s", DefaultPassLength, "password length, default is 16")
)

func init() {
	flag.Usage = usage
	flag.Parse()
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
	if *size > 0 {
		randomPassword = genPasswordWithLength(*size)
	} else {
		randomPassword = genPasswordWithLength(DefaultPassLength)
	}

	fmt.Printf("%s\n", randomPassword)
}

func genPasswordWithLength(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	charsArray := make([]rune, length)
	for i := range charsArray {
		charsArray[i] = chars[next(0, len(chars))]
	}
	return string(charsArray)
}

func next(min, max int) int {
	return min + rand.Intn(max-min)
}
