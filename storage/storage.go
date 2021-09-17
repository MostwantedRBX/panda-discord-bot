package storage

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/mostwantedrbx/discord-go/config"
)

func ReturnRandomPokemon() string {
	//	TODO: find a better way to store this, seems like a mundane use of resources to have a package dedicated to this
	pokemon := [6]string{
		"Whose that pokemon!",
		"Its Pikachu!",
		"I choose you, Team Rocket!",
		"I'm pretty sure Ash has been 10 years old for 20 years now...",
		"My favourite Pokemon is Snorlax",
		"I heard there was a panda Pokemon...",
	}
	return pokemon[rand.Intn(5)]
}

var (
	// errPutFailed is returned when a paste could not be uploaded to pastebin.
	errPutFailed = errors.New("pastebin put failed")
)

//	pastebin stuff
//	Pastebin represents an instance of the pastebin service.
type Pastebin struct{}

//	Put uploads text to Pastebin with optional title returning the ID or an error.
func (p Pastebin) Put(text, title string) (id string, err error) {
	data := url.Values{}

	data.Set("api_dev_key", config.PastebinToken) //	dev key
	data.Set("api_option", "paste")               //	create a paste.
	data.Set("api_paste_code", text)              //	the content of the paste
	data.Set("api_paste_name", title)             //	the paste should have title "title".
	data.Set("api_paste_private", "0")            //	create a public paste.
	data.Set("api_paste_expire_date", "N")        //	the paste should never expire.

	resp, err := http.PostForm("https://pastebin.com/api/api_post.php", data)
	if err != nil {
		time.Sleep(time.Second)
		return "", err
	}

	if resp.StatusCode != 200 {
		time.Sleep(time.Second)
		return "", errPutFailed
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		time.Sleep(time.Second)
		return "", err
	}

	return string(respBody), nil
}
