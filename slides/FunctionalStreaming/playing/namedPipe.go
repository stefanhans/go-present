package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Try to receive a string from a named pipe (mkfifo) /tmp/myPipe")

	strchan := make(chan string)
	bchan := make(chan bool)

	go func() {
		for {
			b, err := ioutil.ReadFile("/tmp/myPipe")
			if err != nil {
				return
			}

			if string(b) == "\n" {
				bchan <- true
			}
			strchan <- string(b)
		}
	}()

	for {
		select {
		case str := <-strchan:
			fmt.Printf("%s", str)
		case <-bchan:
			return
		}
	}

/*
$ mkfifo /tmp/myPipe
$ echo "test" > /tmp/myPipe
$ echo "test" > /tmp/myPipe
$ echo "" > /tmp/myPipe
*/
}
