package fileUtils
import "fmt"
		

func ExampleParseFileName(){
	fmt.Println(parseFileName("../testfile.txt"))
	//Output:
	//testfile.txt
}

func ExampleGetSha1(){
	fmt.Printf("%x",getSha1("../testfile.txt"))
	//Output:
	//2fd4e1c67a2d28fced849ee1bb76e7391b93eb12	
}

