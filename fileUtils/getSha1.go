package fileUtils

import ("crypto/sha1"
		"log"
		"os"
		"io"
)

func GetSha1(filePath string) []byte{
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}

