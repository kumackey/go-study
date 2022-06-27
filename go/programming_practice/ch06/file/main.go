package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const filename = "programming_practice/ch06/file/data1"
const filename2 = "programming_practice/ch06/file/data2"

func main() {
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile(filename)
	fmt.Print(string(read1))

	file1, _ := os.Create(filename2)
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open(filename2)
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)

	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))

}
