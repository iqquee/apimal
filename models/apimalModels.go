package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Animal struct {
	ID                  primitive.ObjectID `bson:"_id"`
	Name                string             `json:"name" validate:"required,min=2"`
	Description         string             `json:"desc" validate:"required,min=20"`
	Image               string             `json:"image" validate:"required"`
	Habitat             []string           `json:"habitat" validate:"required"`
	Domain              string             `json:"domain" validate:"required,min=2"`
	Kingdom             string             `json:"kingdom" validate:"required,min=2"`
	Phylum              string             `json:"phylum" validate:"required,min=2"`
	Class               string             `json:"class" validate:"required,min=2"`
	Order               string             `json:"order" validate:"required,min=2"`
	Family              string             `json:"family" validate:"required,min=2"`
	Genus               string             `json:"genus" validate:"required,min=2"`
	Specie              string             `json:"specie" validate:"required,min=2"`
	Color               []string           `json:"color" validate:"required"`
	Predator            []string           `json:"predator" validate:"required"`
	FoodType            []string           `json:"food_type" validate:"required" `
	OvuationPeriod      string             `json:"ovulation_period" validate:"required,min=2"`
	GestationPeriod     string             `json:"gestation_period" validate:"required,min=2"`
	EstimatedPopulation int                `json:"extimated_population" validate:"required,min=2"`
	ExtinctionStatus    string             `json:"extinction_status" validate:"required,min=2"`
	Reproduction        string             `json:"reproduction" validate:"required,min=2"`
	Motility            string             `json:"motility" validate:"required,min=2"`
	MatingSeason        string             `json:"mating_season" validate:"required,min=2"`
	ModeOfBirth         string             `json:"mode_of_birth" validate:"required,min=2"`
	Animal_id           string             `json:"animal_id"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
}
