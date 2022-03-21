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

func writeFile(name string, b []byte)  {
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

func Grep(lines []string, sub string)  {

	for _, line := range lines {
		if i := strings.Index(line, sub); i > -1 {
			fmt.Println(lines[i-1], i)
		}
	}
}

var after = flag.Int("A", 0, "печатать +N строк после совпадения")
//var before = flag.Int("B", 0, "печатать +N строк до совпадения")
//var context = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
//var count = flag.Int("c", 0, "количество строк")
//var ignoreCase = flag.Bool("i", false, "игнорировать регистр")
//var invert = flag.Bool("v", false, "вместо совпадения, исключать")
//var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
//var line = flag.Bool("n", false, "напечатать номер строки")
var fileName string
var sl []string

func main()  {
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
	Grep(sl, `addEventListener`)
}
