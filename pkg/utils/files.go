package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteFilesByExtension(dirPath, ext string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			fmt.Printf("remove file: %s\n", path)
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Printf("file %s remove from dir %s\n", ext, dirPath)
	return nil
}
