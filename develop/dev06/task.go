package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cuter struct {
	t  string
	sl []string
	f  []string
	d  string
	s  bool
	total string
}

func (c *Cuter) split() []string {
	return strings.Split(c.t, c.d)
}

func (c *Cuter) Cut() string {
	c.sl = c.split()

	// Если не нашлись разделители
	if len(c.sl) <= 1 {

		// Не выводим строки если нет разделителя
		if c.s {
			return ""
		}
		c.total = c.sl[0]
		return c.total
	}

	for _, v := range c.f {
		j, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}

		if len(c.sl)-1 > j {
			j -= 1
			if j < 0 {
				j = 0
			}
			c.total += c.sl[j] + " "
		}
	}

	return c.total
}

var fields = flag.String("f", "", "выбрать поля (колонки)")
var delimiter = flag.String("d", "\t", "использовать другой разделитель")
var separated = flag.Bool("s", false, "только строки с разделителем")

func main() {
	flag.Parse()
	text := flag.Arg(0)
	fmt.Println("flag: ", text)
	c := Cuter{t: text, f: strings.Split(*fields, ","), d: *delimiter, s: *separated}
	fmt.Fprintf(os.Stdout, "%s", c.Cut())
}
