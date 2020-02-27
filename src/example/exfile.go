package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// File Example
func main() {
	exosfile()
	exioutilfile()
	excreatecsv()
	exreadcsv()
}

// os file read , write
func exosfile() {
	fi, err := os.Open("readme.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fo, err := os.Create("readme_copy.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}

// ioutil file read write
func exioutilfile() {
	bytes, err := ioutil.ReadFile("readme.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("readme_copy_ioutil.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}

func excreatecsv() {
	file, err := os.Create("temp.csv")
	if err != nil {
		panic(err)
	}

	wr := csv.NewWriter(bufio.NewWriter(file))
	wr.Write([]string{"A", "1"})
	wr.Write([]string{"B", "2"})
	wr.Write([]string{"C", "3"})
	wr.Write([]string{"D", "4"})
	wr.Flush()
}

func exreadcsv() {
	file, _ := os.Open("temp.csv")

	rdr := csv.NewReader(bufio.NewReader(file))

	rows, _ := rdr.ReadAll()

	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s", rows[i][j])
		}
		fmt.Println()
	}
}
