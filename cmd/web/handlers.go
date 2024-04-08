package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ildx/breeders/models"
	"github.com/ildx/breeders/pets"
	"github.com/tsawler/toolbox"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.html", nil)
}

func (app *application) Page(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.html", page), nil)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	dom, _ := app.App.Models.Dog.GetDogOfMonthByID(1)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			DogName:          "Sam",
			BreedID:          breed.ID,
			Color:            "Black & Tan",
			DateOfBirth:      dob,
			SpayedOrNeutered: 0,
			Description:      "Sam is a very good boy.",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	data := make(map[string]any)
	data["dog"] = dog

	app.render(w, "dog-of-month.page.html", &templateData{
		Data: data,
	})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.html", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsHandler(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogs, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dogs)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("A mixed breed of unknown origin. Propably has some German Shepherd heritage.").
		SetGeographicOrigin("United States").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("felis silvertris house cat").
		SetWeight(4).
		SetDescription("A beautiful house cat.").
		SetGeographicOrigin("Canada").
		SetColor("Calico").
		SetAge(1).
		SetAgeEstimated(true).
		Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	b, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, b)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	s := chi.URLParam(r, "species")
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	pet, err := pets.NewPetWithBreedFromAbstractFactory(s, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, pet)
}
