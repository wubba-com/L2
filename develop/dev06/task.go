package main

import (
	"flag"
	"fmt"
	"strings"
)

type Cuter struct {
	t  string
	sl []string
	f  int
	d  string
	s  bool
}

func (c *Cuter) Cut() {
	c.sl = strings.Split(c.t, c.d)
}

var fields = flag.Int("f", 0, "выбрать поля (колонки)")
var delimiter = flag.String("d", "\t", "использовать другой разделитель")
var separated = flag.Bool("s", false, "только строки с разделителем")

func main() {
	flag.Parse()
	text := flag.Arg(0)
	c := Cuter{t: text, f: *fields, d: *delimiter, s: *separated}
	fmt.Println(c)
}
