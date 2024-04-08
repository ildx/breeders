package models

import (
	"database/sql"
	"time"
)

var repo Repository

type Models struct {
	DogBreed DogBreed
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMysqlRepository(conn)
	} else {
		repo = newTestRepository(nil)
	}
	return &Models{
		DogBreed: DogBreed{},
	}
}

type Breeder struct {
	ID          int         `json:"id"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	ZipCode     string      `json:"zip_code"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	Lifespan    int    `json:"lifespan"`
}

type Dog struct {
	ID               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	AverageLifespan  int    `json:"average_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return repo.GetAllDogBreeds()
}

type Cat struct {
	ID               int       `json:"id"`
	CatName          string    `json:"cat_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type CatBreed struct {
	ID               int    `json:"id" xml:"id"`
	Breed            string `json:"breed" xml:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs" xml:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs" xml:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight" xml:"average_weight"`
	AverageLifespan  int    `json:"average_lifespan" xml:"average_lifespan"`
	Details          string `json:"details" xml:"details"`
	AlternateNames   string `json:"alternate_names" xml:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin" xml:"geographic_origin"`
}
