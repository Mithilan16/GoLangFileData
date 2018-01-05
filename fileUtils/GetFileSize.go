package fileUtils

func GetFileSize(filePath string) []byte{
	file, err := os.Open( filePath ) 
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fi, err := file.Stat()
    
    if err != nil {
    log.Fatal(err)
    }
	return fi.size()
}