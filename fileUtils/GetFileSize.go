package fileUtils

import ("log"
		"os"
)

func GetFileSize(filePath string) int{
	file, err := os.Open( filePath ) 
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fi, err := file.Stat()
    
    if err != nil {
    log.Fatal(err)
    }
	return fi.Size()
}
