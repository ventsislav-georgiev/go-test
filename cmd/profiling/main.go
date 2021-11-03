package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

var (
	buf = make([]byte, 1)
)

func readbyte(file io.Reader) (rune, error) {
	_, err := file.Read(buf[:])
	return rune(buf[0]), err
}

func profileIfRequested() func() {
	prof := flag.String("profile", "", "enable profiling (cpu, mem, trace)")
	flag.CommandLine.Parse(os.Args[2:])

	var profileFunc func(p *profile.Profile) = nil
	switch *prof {
	case "cpu":
		profileFunc = profile.CPUProfile
	case "mem":
		profileFunc = profile.MemProfileRate(1)
	case "trace":
		profileFunc = profile.TraceProfile
	}

	if profileFunc == nil {
		return nil
	}

	return profile.Start(profileFunc, profile.ProfilePath(".")).Stop
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("input file path expected!")
	}

	if prof := profileIfRequested(); prof != nil {
		defer prof()
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open file: %q: %v", filename, err)
	}
	b := bufio.NewReader(f)

	words := 0
	inword := false

	for {
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("cannot open read: %q: %v", filename, err)
		}

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}

		inword = unicode.IsLetter(r)
	}

	fmt.Printf("%q: %d words\n", filename, words)
}
