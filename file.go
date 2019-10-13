package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var fileName string

	if len(os.Args) == 1 {
		fmt.Println("Filename not provided")
		return
	}

	fileName = os.Args[1]

	path := "./" + fileName

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := []byte("Hello world writing to Files!")

	n, err := f.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The number of bytes written %v\n", n)

	time.Sleep(2 * time.Second)
	fmt.Println("Start reading")

	readFile, err := os.Open(path)

	fileInfo, err := readFile.Stat()
	if err != nil {
		panic(err)
	}

	buf := make([]byte, fileInfo.Size())
	for {
		n, err := readFile.Read(buf)

		// break if all content has been read
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf[:n]))
	}
}
