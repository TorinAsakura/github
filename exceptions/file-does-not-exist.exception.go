package exceptions

import "fmt"

type FileDoesNotExistException struct {
	FilePath string
}

func (e *FileDoesNotExistException) Error() string {
	return fmt.Sprintf("File does not exist at path: %s", e.FilePath)
}
