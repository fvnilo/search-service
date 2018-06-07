package handlers

import (
	"errors"
	"net/http"
	"strconv"
)

func getQueryParams(req *http.Request) (string, int, int, bool) {
	term, err := getStringParam(req, "q")
	if err != nil {
		return "", 0, 0, false
	}

	from, err := getNumberParam(req, "from")
	if err != nil {
		return "", 0, 0, false
	}

	size, err := getNumberParam(req, "size")
	if err != nil {
		return "", 0, 0, false
	}

	return term, from, size, true
}

func getNumberParam(req *http.Request, paramName string) (int, error) {
	param, ok := req.URL.Query()[paramName]
	if !ok {
		return 0, errors.New("An error occured while reading query param: " + paramName)
	}

	paramValue, err := strconv.Atoi(param[0])
	if err != nil {
		return 0, err
	}

	return paramValue, nil
}

func getStringParam(req *http.Request, paramName string) (string, error) {
	param, ok := req.URL.Query()[paramName]

	if !ok {
		return "", errors.New("An error occured while reading query param: " + paramName)
	}

	return param[0], nil
}
