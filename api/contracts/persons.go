package contracts

import (
	"time"

	db "github.com/AiteginMoldaliev/test-task/db/sqlc"
)

type CreatePersonRequest struct {
	Firstname  string `json:"firstname" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type PersonResponse struct {
	Id          int64     `json:"id"`
	Firstname   string    `json:"firstname"`
	Surname     string    `json:"surname"`
	Patronymic  string    `json:"patronymic,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Age         int64     `json:"age,omitempty"`
	Nationality string    `json:"nationality,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func ToPersonResponse(person db.Person) PersonResponse {
	return PersonResponse{
		Id:          person.ID,
		Firstname:   person.Firstname,
		Surname:     person.Surname,
		Patronymic:  person.Patronymic.String,
		Gender:      person.Gender,
		Age:         person.Age,
		Nationality: person.Nationality,
		CreatedAt:   person.CreatedAt,
	}
}

type UpdatePersonRequest struct {
	Id         int64  `json:"id" binding:"required"`
	Firstname  string `json:"firstname" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic,omitempty"`
}

type QueryPersonRequest struct {
	Gender      string `json:"gender"`
	MaxAge      int64  `json:"max_age"`
	MinAge      int64  `json:"min_age"`
	Nationality string `json:"nationality"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

func ToQueryPersonFilter(req QueryPersonRequest) db.GetPersonsByFilterParams {
	return db.GetPersonsByFilterParams{
		Gender:      req.Gender,
		Age:         req.MaxAge,
		Age_2:       req.MinAge,
		Nationality: req.Nationality,
		Limit:       req.Limit,
		Offset:      req.Offset,
	}
}

type GetPersonsListRequest struct {
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

type GetPersonRequest struct {
	Id int `uri:"id" binding:"required,min=1"`
}

type PersonAgeResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type PersonGenderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type CountryInfo struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type PersonNationResponse struct {
	Count   int           `json:"count"`
	Name    string        `json:"name"`
	Country []CountryInfo `json:"country"`
}
