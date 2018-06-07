package handlers

import (
	"net/http"
	"strconv"

	"github.com/nylo-andry/search-service/repositories"
	"github.com/nylo-andry/search-service/requests"
)

func Populate(w http.ResponseWriter, req *http.Request) {
	numberArr, ok := req.URL.Query()["number"]
	if !ok {
		requests.RespondWithError(w, http.StatusBadRequest, "Attach proper parameters")
		return
	}
	numberStr := numberArr[0]
	number, err := strconv.Atoi(numberStr)

	if err != nil {
		requests.RespondWithError(w, http.StatusBadRequest, "Attach proper parameters")
		return
	}

	err = repositories.Populate(number)
	if err != nil {
		requests.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
}

func Search(w http.ResponseWriter, req *http.Request) {
	term, from, size, ok := getQueryParams(req)
	if !ok {
		requests.RespondWithError(w, http.StatusBadRequest, "Attach proper parametersssss")
		return
	}

	res, err := repositories.Search(term, from, size)
	if err != nil {
		requests.RespondWithError(w, http.StatusInternalServerError, "An error occured")
		return
	}

	requests.RespondWithJSON(w, http.StatusOK, res)
}
