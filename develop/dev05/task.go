package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

/**
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
*/

func writeFile(name string, b []byte) {
	err := ioutil.WriteFile(name, b, fs.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[err] %s", err.Error())
		return
	}
}

func readFile(scan *bufio.Scanner) []string {
	s := make([]string, 0)

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

func After(after int, lines []string, i int) string {
	s := fmt.Sprintf("%s\n", lines[i])
	var j = 1
	for after > 0 {
		if i+j < len(lines) {
			s += fmt.Sprintf("%s\n", lines[i+j])
		}

		after--
		j++
	}

	return s
}

func Before(before int, lines []string, i int) string {
	var s string
	var k = 1
	for before > 0 {
		if i-before > -1 {
			s += fmt.Sprintf("%s\n", lines[i-before])
		}
		before--
		k++
	}
	s += fmt.Sprintf("%s\n", lines[i])

	return s
}

func Grep(lines []string, sub string, ic bool) string {
	var search string
	var i int
	if ic {
		sub = strings.ToLower(sub)
	}
	for _, l := range lines {
		if i = strings.Index(l, sub); i > -1 {
			i -= 1
			if i < 0 {
				i = 0
			}
			break
		}
	}

	if *after > 0 {
		search = fmt.Sprintf("%s\n", lines[i])
		search += After(*after, lines, i)
	}

	if *before > 0 {
		search = Before(*before, lines, i)
		search += fmt.Sprintf("%s\n", lines[i])
	}

	if *context > 0 {
		search = Before(*context, lines, i)
		search += fmt.Sprintf("%s\n", lines[i])
		search += After(*context, lines, i)
	}

	search = strings.Trim(search, "\n")
	if *count {
		c := len(strings.Split(search, "\n"))
		fmt.Println(c)
	}

	if *line {
		search = fmt.Sprintf("%d: %s", i, lines[i])
	}

	if *fixed {
		l := Fixed(lines, sub)
		return lines[l]
	}
	
	if *invert {
		Invert(lines, i)
	}
	fmt.Println(search)
	return fmt.Sprintf("%s", lines[i-1])
}

func Invert(lines []string, i int) string {
	s := make([]string, 0)

	s = append(s, lines[:i]...)
	if i+1 < len(lines) {
		s = append(s, lines[i+1:]...)
	}

	return strings.Join(s, "\n")
}

func Fixed(lines []string, sub string) int {
	for i, l := range lines {
		if l == sub {
			i -= 1
			if i < 0 {
				i = 0
			}
			return i
		}
	}
	return len(lines)
}

var after = flag.Int("A", 0, "печатать +N строк после совпадения")
var before = flag.Int("B", 0, "печатать +N строк до совпадения")
var context = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
var count = flag.Bool("c", false, "количество строк")
var ignoreCase = flag.Bool("i", false, "игнорировать регистр")
var invert = flag.Bool("v", false, "вместо совпадения, исключать")
var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
var line = flag.Bool("n", false, "напечатать номер строки")

var fileName string
var sl []string

func main() {
	flag.Parse()
	fileName = flag.Arg(0)
	fmt.Println(flag.Args())
	r, err := os.Open(fileName)
	defer r.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[err] %s", err.Error())
		return
	}
	sc := bufio.NewScanner(r)
	sl = readFile(sc)
	Grep(sl, `addEventListener`, *ignoreCase)
}
