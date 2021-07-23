package pyscripts

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	_ "embed"
)

//	I'll be using Python to convert images

//go:embed scripts/ascii-converter.py
var ascii string

func RunScript(script string) {
	fmt.Println("taocs")
	if script == "convert" {
		cmd := exec.Command("python", "-c", ascii, "./tacos.png")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Println(cmd.Run())
	}
}

const (
	pastebinDevKey = "b09173a266b328998674c8f48d1fc605"
)

var (
	// ErrPutFailed is returned when a paste could not be uploaded to pastebin.
	ErrPutFailed = errors.New("pastebin put failed")
	// ErrGetFailed is returned when a paste could not be fetched from pastebin.
	ErrGetFailed = errors.New("pastebin get failed")
)

// pastebin stuff
// Pastebin represents an instance of the pastebin service.
type Pastebin struct{}

// Put uploads text to Pastebin with optional title returning the ID or an error.
func (p Pastebin) Put(text, title string) (id string, err error) {
	data := url.Values{}
	// Required values.
	data.Set("api_dev_key", pastebinDevKey)
	data.Set("api_option", "paste") // Create a paste.
	data.Set("api_paste_code", text)
	// Optional values.
	data.Set("api_paste_name", title)      // The paste should have title "title".
	data.Set("api_paste_private", "0")     // Create a public paste.
	data.Set("api_paste_expire_date", "N") // The paste should never expire.

	resp, err := http.PostForm("https://pastebin.com/api/api_post.php", data)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", ErrPutFailed
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
