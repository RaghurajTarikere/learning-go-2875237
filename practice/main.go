package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter data to be written in file")
	content, err := reader.ReadString('\n')
	checkErr(err)
	file, err := os.Create("./FromString.txt")
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Printf("Wrote %v characters to a file \n", length)
	defer file.Close()
	defer readFile("./FromString.txt")
}

func readFile(path string) {
	data, err := os.ReadFile(path)
	checkErr(err)
	fmt.Println("Read data is:", string(data))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
