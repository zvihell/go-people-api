package models

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type AgifyResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GenderizeResponse struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type NationalizeResponse struct {
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

type Country struct {
	CountryID string `json:"country_id"`
}
