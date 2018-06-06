package handlers

import (
	"net/http"
	"strconv"
)

func getQueryParams(req *http.Request) (string, int, int, bool) {
	termArr, ok := req.URL.Query()["q"]
	if !ok {
		return "", 0, 0, false
	}
	term := termArr[0]

	fromArr, ok := req.URL.Query()["from"]
	if !ok {
		return "", 0, 0, false
	}
	fromStr := fromArr[0]
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		return "", 0, 0, false
	}

	sizeArr, ok := req.URL.Query()["size"]
	if !ok {
		return "", 0, 0, false
	}
	sizeStr := sizeArr[0]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return "", 0, 0, false
	}

	return term, from, size, true
}
