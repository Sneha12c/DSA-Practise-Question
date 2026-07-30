package main

import (
	"bufio"
	"os"
)

func main() {
	// fs, err := os.Open("examples.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := fs.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(fileInfo.Name(), fileInfo.Mode())

	// defer fs.Close()

	// buf := make([]byte, 10)

	// d, err := fs.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("examples.txt") // better for small files as it load allcontent into memory
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data, string(data))

	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(9)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("hi long")
	// f.WriteString("bnjcn")
	// bytes := []byte("ndjfksnkjn")
	// f.WriteAt(bytes, 6)

	sourceFile, err := os.Open("examples.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
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
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

}
