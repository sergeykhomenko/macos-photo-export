package main

import "io/ioutil"

func main() {
	files, err := ioutil.ReadDir("~/Pictures/Photos\\ Library.photoslibrary/originals/0")
	if err != nil {
		log.Panic(err)
	}

	for _, f := range files {
		fmt.Prinln(f.Name())
	}
}
