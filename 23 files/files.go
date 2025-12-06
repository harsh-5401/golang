package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo , err :=f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())

	// // read file

	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// // we must always close a file
	// defer f.Close()

	// buf := make([]byte, 15)

	// // always read and save it to buffer temp memeory
	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// println("data ", string(buf) , d)

	// read file  simplest(only for small file)

	// f, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// println("data ", string(f))

	// read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileinfo, err := dir.ReadDir(-1)

	// for _, fi := range fileinfo {
	// 	fmt.Println(fi.Name() , fi.IsDir())
	// }

	// create a file

	// file, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// // file.WriteString("nice language")  // default append mode

	// bytes := []byte("hello golang")
	// file.Write(bytes)

	// transfer data of one file to another and also use (streaming fashion)

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	// FIX: Open destination file in write mode
	destFile, err := os.Create("example2.txt") // <--- IMPORTANT
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		if err := writer.WriteByte(b); err != nil {
			panic(err)
		}
	}

	writer.Flush()
	fmt.Println("return to new file successfully")

	os.Remove("example2.txt")  // to delete the file

}
