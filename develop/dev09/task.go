package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const(
	Dir = "pages"
)

func setUrl(sl []string) []string  {
	m := make(map[string]int)
	set := make([]string, 0)
	for _, v := range sl {
		m[v] += 1
	}

	for k := range m {
		if m[k] < 2 {
			set = append(set, k)
		}
	}

	return set
}

func GetWithClient(client *http.Client) func(string) error {
	return func(url string) error {
		r, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		w, err := client.Do(r)
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}(w.Body)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return err
		}
		paths := strings.Split(strings.Trim(url, " "), "/")
		fileName := strings.Join(setUrl(paths[1:]),"_")

		err = write(fileName, p)
		if err != nil {
			return err
		}

		return nil
	}
}

func write(fileName string, p []byte) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(dir, Dir, fileName+".html"), os.O_CREATE|os.O_WRONLY, fs.ModePerm)
	defer f.Close()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.Name(), p, fs.ModePerm)
}

func main()  {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	DownPage := GetWithClient(client)
	err := DownPage("https://seasonkrasoty.ru/product/legkiy_omolazhivayushchiy_krem/")
	if err != nil {
		fmt.Println(err.Error())
	}
}
