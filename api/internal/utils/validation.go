package utils

import (
	"mime/multipart"
	"strings"
)

func ValidateFile(file *multipart.FileHeader) bool {
	extension := strings.Split(file.Filename, ".")
	return extension[1] == "csv"
}

// func ValidateHeaders(header []string) error {

// 	isExist := contains(header, "unix")
// 	isExist = contains(header, "symbol")
// 	isExist = contains(header, "open")
// 	isExist = contains(header, "high")
// 	isExist = contains(header, "low")
// 	isExist = contains(header, "close")

// 	return nil
// }

func contains(s []string, str string) bool {
	for _, v := range s {
		if strings.ToLower(v) == str {
			return true
		}
	}

	return false
}
