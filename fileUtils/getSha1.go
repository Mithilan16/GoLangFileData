package fileUtils

import ("crypto/sha1"
		"log"
		"io/ioutil"
)

func getSha1(filePath string) []byte{
	dat, err := ioutil.ReadFile(filePath)
	
	if err != nil {
    	log.Fatal(err)
    }
    defer file.Close()
    
    h:=sha1.New();
	h.Write([]byte(dat))
		
    return h.Sum(nil)
}

