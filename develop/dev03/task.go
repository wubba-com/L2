package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func UtilSort(f *os.File, k int)  {
	fscan := bufio.NewScanner(f)
	s := make([][]string, 0)
	var i int
	for fscan.Scan() {
		words := strings.Split(fscan.Text(), " ")
		slW := make([]string, 0)
		slW = append(slW, words...)
		s = append(s, slW)
		i++
	}
	sort.SliceStable(s, func(i, j int) bool {
		fmt.Println("r", s[i])
		return s[i][k] < s[j][k]
	})

	fmt.Println(s)
	//var str string
	//for _, v := range s{
	//	fmt.Println(v)
	//	str += strings.Join(v, "\n")
	//}
	//err := ioutil.WriteFile(f.Name(), []byte(strings.Join(s)), fs.ModePerm)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
}

var fileName string
var column int
func main()  {
	flag.StringVar(&fileName, "f", "", "file name")
	flag.IntVar(&column, "k", 2, "указание колонки для сортировки")
	flag.Parse()

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	UtilSort(f, column)

}