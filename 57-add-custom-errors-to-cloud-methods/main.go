package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type File interface {
	Load(string) (io.ReadCloser, error)
	Save(String, io.ReadSeeker) error
}

type LocalFile struct {
	Base string
}

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

//func (l LocalFile) Save(path string, body io.ReadSeeker) error {
//	p := filepath.Join(l.Base, path)
//	d := filepath.Dir(p)
//	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
//	if err != nil {
//		return err
//	}
//
//	f, err := os.Create(p)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	_, err = io.Copy(f, body)
//	return err
//}

var (
	ErrFileNotFound   = errors.New("File Not Found")
	ErrCannotLoadFile = errors.New("Unable to load file")
	ErrCannotSaveFile = errors.New("Unable to save file")
)

func (l LocalFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	var oerr error
	o, err := os.Open(p)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Unable to find %s", path)
		oerr = ErrFileNotFound
	} else if err != nil {
		log.Printf("Error loading file %s, err: %s", path, err)
		oerr = ErrCannotLoadFile
	}
	return o, oerr
}

// ABOVE IS THE IMPLEMENTATION. BELOW THE USAGE

func main() {
	content := `test abc`
	body := bytes.NewReader([]byte(content))

	store, err := fileStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Storing content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Retrieving content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err := ioutil.ReadAll(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(o))
}

// AND THIS IS THE CLOUD IMPLEMENTATION OF THE FILE STORAGE.
// here it's a local file, but could be changed to a cloud connection

func fileStore() (File, error) {
	return &LocalFile{Base: "."}, nil
}
