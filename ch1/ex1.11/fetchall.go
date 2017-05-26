package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// read in the site list
	fin, err := os.Open("./sites.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
	}
	defer fin.Close()

	sites := processSites(fin)

	// setup output file
	fout, err := os.Create("./tmp.out")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: opening file: %v\n", err)
		os.Exit(1)
	}
	defer fout.Close()

	start := time.Now()

	ch := make(chan string)
	for _, url := range sites {
		go fetch(url, ch) // start a goroutine
	}
	for range sites {
		fout.WriteString(<-ch + "\n") // receive from channel ch
	}
	s := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	fout.WriteString(s)
}

func processSites(fin *os.File) []string {
	var list []string
	input := bufio.NewScanner(fin)
	for input.Scan() {
		list = append(list, strings.Split(input.Text(), ",")[1])
	}
	return list
}

func fetch(url string, ch chan<- string) {
	url = "http://" + url[1:len(url)-1]
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	//fmt.Printf("%.2fs  %7d  %s", secs, nbytes, url)
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
