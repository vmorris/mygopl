// xkcdsearch
// searches the descriptions of all xkcd comics and prints out the URL and
// transcript of each matching result.
// CLI switch -u downloads the json results from xkcd.com and updates our offline index

package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/vmorris/mygopl/ch4/ex4.12/xkcd"
)

var comics xkcd.ComicDescriptions

const file = "./descriptions.gob"

func main() {
	downloadPtr := flag.Bool("u", false, "update descriptions and exit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [-u] term1 term2 ... termN\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Paramters:")
		flag.PrintDefaults()
	}
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()

	if *downloadPtr {
		err := Load(file, &comics)
		check(err)

		fmt.Println("Downloading descriptions...")
		download()
		fmt.Println("Downloaded all descriptions.")
		fmt.Println("Saving downloaded results...")
		err = Save(file, comics)
		check(err)
		fmt.Printf("Results saved to %v. Exiting.\n", file)
		os.Exit(0)
	}

	err := Load(file, &comics)
	check(err)

	// build search term array
	searchTerms := os.Args[1:]

	for _, comic := range comics.Results {
		for _, term := range searchTerms {
			if strings.Contains(comic.Transcript, term) {
				fmt.Printf("---------------------\nRESULT FOUND: %v\n---------------------\n", comic.Num)
				fmt.Printf("%v\n%v\n", comic.Img, comic.Transcript)
			}
		}
	}

}

func download() {
	lastDownloaded := comics.LastDownloaded
	var i int
	if lastDownloaded == 0 {
		i = 1
	} else {
		i = lastDownloaded
	}
	for {
		if i == 404 {
			i++
			continue // Randall is pretty funny
		}
		result, err := xkcd.GetDescription(i)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcdgetter: %v\n", err)
			break
		}
		fmt.Printf("%v -- ", result.Num)
		comics.LastDownloaded = result.Num
		comics.Results = append(comics.Results, *result)
		i++
	}
}

// Save encoding to gob file
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	defer file.Close()
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	return err
}

// Load encoding from gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	defer file.Close()
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	return err
}

func check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%v\t%v\n%v\n", line, file, e)
		os.Exit(1)
	}
}
