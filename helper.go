package go_helper

import (
	"net/http"
	"os"
)

// TODO: Please implement this with more HTTP Header
type HeaderParams struct {
	AccessControlAllowOrigin  string
	AccessControlAllowMethods string
	AccessControlAllowHeaders string
	ContentType               string
}

func GetENV(key string, default_value string) (ENV_VALUE string) {
	ENV_VALUE = os.Getenv(key)
	if ENV_VALUE == "" {
		ENV_VALUE = default_value
	}
	return
}

// TODO: Please implement this with more HTTP Header
func SetHeader(w http.ResponseWriter, params *HeaderParams) {

	// set default value
	if params.AccessControlAllowOrigin == "" {
		params.AccessControlAllowOrigin = "*"
	}
	if params.AccessControlAllowMethods == "" {
		params.AccessControlAllowMethods = "*"
	}
	if params.AccessControlAllowHeaders == "" {
		params.AccessControlAllowHeaders = "*"
	}
	if params.ContentType == "" {
		params.ContentType = "application/json"
	}

	// to set which origin can access this rest api
	w.Header().Set("Access-Control-Allow-Origin", params.AccessControlAllowOrigin)

	// to set which methods is allowed to access this rest api
	w.Header().Set("Access-Control-Allow-Methods", params.AccessControlAllowMethods)

	// to set which headers is allowed to access this rest api
	w.Header().Set("Access-Control-Allow-Headers", params.AccessControlAllowHeaders)

	// to set content type of header
	w.Header().Set("Content-Type", params.ContentType)
}
