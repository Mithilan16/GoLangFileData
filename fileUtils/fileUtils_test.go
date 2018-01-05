package fileUtils
import "fmt"
		

func ExampleParseFileName(){
	fmt.Println(ParseFileName("../testfile.txt"))
	//Output:
	//testfile.txt
}

func ExampleGetSha1(){
	fmt.Printf("%x",GetSha1("../testfile.txt"))
	//Output:
	//2fd4e1c67a2d28fced849ee1bb76e7391b93eb12	
}

func ExampleGetMd5(){
	fmt.Printf("%x",GetMd5("../testfile.txt"))
	//Output:
	//9e107d9d372bb6826bd81d3542a419d6
}