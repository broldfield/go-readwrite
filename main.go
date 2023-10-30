package main

import (
	"fmt"

	"brodieoldfield.com/go/readwrite/fileutils"
)

// instead of ./test.txt, can use os.Getwd, and append before test.txt
// rootpath, _ := os.Getwd()
// rootpath + "text.txt"
func main() {
	contents, err := fileutils.ReadTextFile("./test.txt")
	if err != nil {
		fmt.Printf("Error Occurred %v", err)
	} else {
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", contents, contents, contents)
		fileutils.WriteToFile("./text.txt", newContent)
		fmt.Println(contents)
	}
}
