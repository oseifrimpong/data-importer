package utils

// Custom error codes and string
const (
	//message
	SUCCESSFUL         = "successful"
	CREATION_FAILED    = "failed to create data from csv"
	NOT_SUPPORTED_FILE = "file not supported, upload a csv file"
	UPLOAD_FAILED      = "failed to upload file, try again"
	RETRIEVE_FAILED    = "failed to retrieve data"

	//code
	SUCCESSFUL_CODE      = 2000
	CREATION_FAILED_CODE = 4000
	RETRIEVE_FAILED_CODE = 4004
)
