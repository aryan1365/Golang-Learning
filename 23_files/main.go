package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }
	// defer f.Close()
	// fileInfo, er := f.Stat()
	// if er != nil {
	// 	//log the error
	// 	panic(er)
	// } else {
	// 	//do something with the file info
	// 	fmt.Println(fileInfo.Name())
	// 	fmt.Println(fileInfo.IsDir())
	// 	fmt.Println(fileInfo.Size())
	// 	fmt.Println(fileInfo.Mode())
	// 	fmt.Println(fileInfo.ModTime())
	// 	fmt.Println(fileInfo.Name())

	// 	//  //reading file
	// 	buf := make([]byte,8)
	// 	d, error := f.Read(buf)
	// 	if error != nil {
	// 		panic(error)
	// 	}
	// 	for i := 0; i < len(buf); i++ {
	// 		println("data", d, string(buf[i]))

	// 	}

	// }

	//loads complete file into our memory once

	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	//read folder
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// fileInfo, err := dir.ReadDir(0)
	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())

	// }

	//create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// // f.WriteString("hello golangg")
	// bytes := []byte("He he")
	// f.Write(bytes)

	//read and write data from one file to another
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()
	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("Done")

	//delete a file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted ")
}
