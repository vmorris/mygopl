package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const descriptionURL = "https://xkcd.com/###/info.0.json"

// ComicDescriptions contains all the results and the last downloaded comic number
type ComicDescriptions struct {
	LastDownloaded int
	Results        []ComicDescription
}

// ComicDescription contains the decoded JSON result from xkcd info
type ComicDescription struct {
	Num        int
	Link       string
	Year       string
	Day        string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
}

// GetDescription downloads the JSON for a specific xkcd comic and returns
// the decoded result
func GetDescription(num int) (*ComicDescription, error) {
	downloadURL := strings.Replace(descriptionURL, "###", strconv.Itoa(num), 1)
	resp, err := http.Get(downloadURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result ComicDescription
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

/* example JSON results from https://xkcd.com/571/info.0.json
{"month": "4",
"num": 571,
"link": "",
"year": "2009",
"news": "",
"safe_title": "Can't Sleep",
"transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}",
"alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.",
"img": "https://imgs.xkcd.com/comics/cant_sleep.png",
"title": "Can't Sleep",
"day": "20"}
*/
