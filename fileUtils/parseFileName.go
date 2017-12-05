package fileUtils

import "path/filepath"
		
func parseFileName(filePath string) string{
	return filepath.Base(filePath)
}
