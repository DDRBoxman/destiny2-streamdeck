package bungieclient

import (
	"fmt"
	"archive/zip"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"errors"
)

func DownloadContentDB(path, dst string) error {
	url := fmt.Sprintf("https://bungie.net%s", path)

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Couldn't download DB, http status: %d", resp.StatusCode)
	}

	tmpfile, err := ioutil.TempFile("", "worldContent.content")
	if err != nil  {
		return err
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil  {
		return err
	}

	zipReader, err := zip.OpenReader(tmpfile.Name())
	if err != nil {
		return err
	}
	defer zipReader.Close()

	if len(zipReader.File) < 1 {
		return errors.New("World Content DB didn't have any files in it.")
	}

	sqliteFile := zipReader.File[0]
	sqliteFileReader, err := sqliteFile.Open()
	if err != nil {
		return err
	}
	defer sqliteFileReader.Close()

	out, err := os.Create(dst)
	if err != nil  {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, sqliteFileReader)
	if err != nil  {
		return err
	}

	return nil
}