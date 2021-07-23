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
	fmt.Println("Running a Python script")
	if script == "convert" {
		cmd := exec.Command("python", "-c", ascii, "./tacos.png") //	TODO: don't hardcode the filename.
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//	run and print the results and output of the python script
		log.Println(cmd.Run())
	}
}

const (
	//	TODO: add this to the config
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

	data.Set("api_dev_key", pastebinDevKey) //	dev key
	data.Set("api_option", "paste")         //	create a paste.
	data.Set("api_paste_code", text)        //	the content of the paste
	data.Set("api_paste_name", title)       //	the paste should have title "title".
	data.Set("api_paste_private", "0")      //	create a public paste.
	data.Set("api_paste_expire_date", "N")  //	the paste should never expire.

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
