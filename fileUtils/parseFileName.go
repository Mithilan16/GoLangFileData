package fileUtils

import "path/filepath"
		
func parseFileName(filePath string){
	return filepath.Base(filePath)
}
