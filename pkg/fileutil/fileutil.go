package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilePaths(dirPath string) ([]string, error) {
	sourceFileNames := make([]string, 0)

	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(path, ".c") || strings.HasSuffix(path, ".h") {
				sourceFileNames = append(sourceFileNames, path)
			}

			return nil
		})

	if err != nil {
		//log.Println(err)
	}

	return sourceFileNames, err
}
