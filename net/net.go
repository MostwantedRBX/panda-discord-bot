package net

import (
	"errors"
	"io"
	"net/http"
	"os"
)

//	Hoping to be able to convert images in the near future.
func DownloadFile(address string, filename string) error {
	response, err := http.Get(address)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//	check the status code from the URL
	if response.StatusCode != 200 {
		return errors.New("status code from address isn't right... are you sure the address is correct?")
	}

	//	create the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	//	fill the file with the data from the interwebs
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
