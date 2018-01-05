package fileUtils

import "path/filepath"
		
func ParseFileName(filePath string) string{
	return filepath.Base(filePath)
}
