package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(b string) *Pet
	SetMinWeight(w int) *Pet
	SetMaxWeight(w int) *Pet
	SetWeight(w int) *Pet
	SetDescription(d string) *Pet
	SetLifespan(l int) *Pet
	SetGeographicOrigin(g string) *Pet
	SetColor(c string) *Pet
	SetAge(a int) *Pet
	SetAgeEstimated(b bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() *Pet {
	return &Pet{}
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(b string) *Pet {
	p.Breed = b
	return p
}

func (p *Pet) SetMinWeight(w int) *Pet {
	p.MinWeight = w
	return p
}

func (p *Pet) SetMaxWeight(w int) *Pet {
	p.MaxWeight = w
	return p
}

func (p *Pet) SetWeight(w int) *Pet {
	p.Weight = w
	return p
}

func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}

func (p *Pet) SetLifespan(l int) *Pet {
	p.Lifespan = l
	return p
}

func (p *Pet) SetGeographicOrigin(g string) *Pet {
	p.GeographicOrigin = g
	return p
}

func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}

func (p *Pet) SetAge(a int) *Pet {
	p.Age = a
	return p
}

func (p *Pet) SetAgeEstimated(e bool) *Pet {
	p.AgeEstimated = e
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("minimum weight must be less than maximum weight")
	}
	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2
	return p, nil
}
