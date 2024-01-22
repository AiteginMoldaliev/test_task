package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AiteginMoldaliev/test-task/api/contracts"
	db "github.com/AiteginMoldaliev/test-task/db/sqlc"
	"github.com/gin-gonic/gin"
)

const (
	PersonAgeAPI    = "https://api.agify.io/?name="
	PersonGenderAPI = "https://api.genderize.io/?name="
	PersonNationAPI = "https://api.nationalize.io/?name="
)

func (server *Server) CreatPerson(ctx *gin.Context) {
	var req contracts.CreatePersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var (
		genderChan = make(chan contracts.PersonGenderResponse)
		ageChan    = make(chan contracts.PersonAgeResponse)
		nationChan = make(chan contracts.PersonNationResponse)
		errorChan  = make(chan error, 3)
	)

	// add [gender, age, nation]
	// get Gender
	go func() {
		body, err := GetRequest(PersonGenderAPI + req.Firstname)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get gender: %v", err)
			return
		}

		var genderResp contracts.PersonGenderResponse
		err = json.Unmarshal(body, &genderResp)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get gender: %v", err)
			return
		}

		genderChan <- genderResp
	}()

	// get Age
	go func() {
		body, err := GetRequest(PersonAgeAPI + req.Firstname)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get age: %v", err)
			return
		}

		var ageResp contracts.PersonAgeResponse
		err = json.Unmarshal(body, &ageResp)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get age: %v", err)
			return
		}

		ageChan <- ageResp
	}()

	// get Nation
	go func() {
		body, err := GetRequest(PersonNationAPI + req.Firstname)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get nation: %v", err)
			return
		}

		var nationResp contracts.PersonNationResponse
		err = json.Unmarshal(body, &nationResp)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get nation: %v", err)
			return
		}

		nationChan <- nationResp
	}()

	arg := db.CreatePersonParams{
		Firstname:  req.Firstname,
		Surname:    req.Surname,
		Patronymic: sql.NullString{String: req.Patronymic, Valid: true},
	}

	for i := 0; i < 3; i++ {
		select {
		case genderResp := <-genderChan:
			arg.Gender = genderResp.Gender
		case ageResp := <-ageChan:
			arg.Age = int64(ageResp.Age)
		case nationResp := <-nationChan:
			var maxProbability = 0.0
			var maxProbabilityCountryID string

			for _, country := range nationResp.Country {
				if country.Probability > maxProbability {
					maxProbability = country.Probability
					maxProbabilityCountryID = country.CountryID
				}
			}
			arg.Nationality = maxProbabilityCountryID
		case err := <-errorChan:
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
	}

	person, err := server.store.CreatePerson(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	resp := contracts.ToPersonResponse(person)

	ctx.JSON(http.StatusOK, resp)
}

func GetRequest(url string) (body []byte, err error) {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to send GET request to address %v: %v", url, err)
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response's body of address %v: %v", url, err)
		return
	}

	return
}

func (server *Server) UpdatePerson(ctx *gin.Context) {
	var req contracts.UpdatePersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var (
		genderChan = make(chan contracts.PersonGenderResponse)
		ageChan    = make(chan contracts.PersonAgeResponse)
		nationChan = make(chan contracts.PersonNationResponse)
		errorChan  = make(chan error, 3)
	)

	// add [gender, age, nation]
	// get Gender
	go func() {
		body, err := GetRequest(PersonGenderAPI + req.Firstname)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get gender: %v", err)
			return
		}

		var genderResp contracts.PersonGenderResponse
		err = json.Unmarshal(body, &genderResp)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get gender: %v", err)
			return
		}

		genderChan <- genderResp
	}()

	// get Age
	go func() {
		body, err := GetRequest(PersonAgeAPI + req.Firstname)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get age: %v", err)
			return
		}

		var ageResp contracts.PersonAgeResponse
		err = json.Unmarshal(body, &ageResp)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get age: %v", err)
			return
		}

		ageChan <- ageResp
	}()

	// get Nation
	go func() {
		body, err := GetRequest(PersonNationAPI + req.Firstname)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get nation: %v", err)
			return
		}

		var nationResp contracts.PersonNationResponse
		err = json.Unmarshal(body, &nationResp)
		if err != nil {
			errorChan <- fmt.Errorf("Failed to get nation: %v", err)
			return
		}

		nationChan <- nationResp
	}()

	arg := db.UpdatePersonParams{
		ID:         req.Id,
		Firstname:  req.Firstname,
		Surname:    req.Surname,
		Patronymic: sql.NullString{String: req.Patronymic, Valid: true},
	}

	for i := 0; i < 3; i++ {
		select {
		case genderResp := <-genderChan:
			arg.Gender = genderResp.Gender
		case ageResp := <-ageChan:
			arg.Age = int64(ageResp.Age)
		case nationResp := <-nationChan:
			var maxProbability = 0.0
			var maxProbabilityCountryID string

			for _, country := range nationResp.Country {
				if country.Probability > maxProbability {
					maxProbability = country.Probability
					maxProbabilityCountryID = country.CountryID
				}
			}
			arg.Nationality = maxProbabilityCountryID
		case err := <-errorChan:
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
	}

	person, err := server.store.UpdatePerson(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	resp := contracts.ToPersonResponse(person)

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) GetPerson(ctx *gin.Context) {
	var req contracts.GetPersonRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	person, err := server.store.GetPerson(ctx, int64(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	resp := contracts.ToPersonResponse(person)

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) DeletePerson(ctx *gin.Context) {
	var req contracts.GetPersonRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := server.store.DeletePerson(ctx, int64(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "person was deleted"})
}

func (server *Server) QueryPerson(ctx *gin.Context) {
	var req contracts.QueryPersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	filter := contracts.ToQueryPersonFilter(req)

	persons, err := server.store.GetPersonsByFilter(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	var resp []contracts.PersonResponse
	for _, person := range persons {
		resp = append(resp, contracts.ToPersonResponse(person))
	}

	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) GetPersonsList(ctx *gin.Context) {
	var req contracts.GetPersonsListRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	pagination := db.GetPersonsListParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	persons, err := server.store.GetPersonsList(ctx, pagination)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	var resp []contracts.PersonResponse
	for _, person := range persons {
		resp = append(resp, contracts.ToPersonResponse(person))
	}

	ctx.JSON(http.StatusOK, resp)
}
