package utils

import (
	"errors"
	"mime/multipart"
	"ohlc-data-api/api/dto"
	"strings"
)

func ValidateFile(file *multipart.FileHeader) bool {
	extension := strings.Split(file.Filename, ".")
	return extension[1] == "csv"
}

func ValidateCSVHeaders(header []string) error {
	approvedHeaders := []string{"unix", "high", "low", "symbol", "close", "open"}

	if len(header) != 6 {
		return errors.New("file should have 6 columns")
	}

	// unix
	if !contains(approvedHeaders, strings.ToLower(header[0])) {
		return errors.New("unix field is missing")
	}

	// symbol
	if !contains(approvedHeaders, strings.ToLower(header[1])) {
		return errors.New("symbol field is missing")
	}

	// open
	if !contains(approvedHeaders, strings.ToLower(header[2])) {
		return errors.New("open field is missing")
	}

	// high
	if !contains(approvedHeaders, strings.ToLower(header[3])) {
		return errors.New("high field is missing")
	}

	// low
	if !contains(approvedHeaders, strings.ToLower(header[4])) {
		return errors.New("low field is missing")
	}

	// close
	if !contains(approvedHeaders, strings.ToLower(header[5])) {
		return errors.New("close field is missing")
	}

	return nil
}

func SearchValidation(params *dto.SearchParams) error {
	defaultPageSize := 100
	if params.PageSize > defaultPageSize {
		return errors.New("page size entered is over the limit")
	}
	err := sortStringCheck(params.Sort)
	if err != nil {
		return err
	}

	return nil
}

func sortStringCheck(str string) error {
	approvedFields := []string{"created_at", "unix", "high", "low", "symbol", "close", "open"}

	field := strings.Split(str, " ")
	if !contains(approvedFields, field[0]) {
		return errors.New("sort field is not allowed")
	}

	if !(strings.ToLower(field[1]) == "desc") || (strings.ToLower(field[1]) == "asc") {
		return errors.New("order can only be done by desc or asc")
	}

	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
