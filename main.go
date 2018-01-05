package main

//create new functions in seperate files under fileUtils
//fileUtils/reverse.go is to get an idea of how to create a new function
//Test output of functions in here
import (
	"fmt"
	"./fileUtils"
)

func main() {	
	fmt.Println("file name: " + fileUtils.ParseFileName("./testfile.txt"))
	fmt.Printf("SHA1: %x\n", fileUtils.GetSha1("./testfile.txt"))
	fmt.Printf("MD5 %x\n", fileUtils.GetMd5("./testfile.txt"))
}