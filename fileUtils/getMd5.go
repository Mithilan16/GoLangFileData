package fileUtils

import ("crypto/md5"
		"log"
		"os"
		"io"
)

func GetMd5(filePath string) []byte{
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}