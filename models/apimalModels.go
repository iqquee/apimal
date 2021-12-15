package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Animal struct {
	ID                  primitive.ObjectID `bson:"_id"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	Image               string             `json:"image"`
	Habitat             []string           `json:"habitat"`
	Domain              string             `json:"domain"`
	Kingdom             string             `json:"kingdom"`
	Phylum              string             `json:"phylum"`
	Class               string             `json:"class"`
	Order               string             `json:"order"`
	Family              string             `json:"family"`
	Genus               string             `json:"genus"`
	Specie              string             `json:"specie"`
	Color               string             `json:"color"`
	Predator            []string           `json:"predator"`
	FoodType            []string           `json:"food_type"`
	OvuationPeriod      string             `json:"ovulation_period"`
	GestationPeriod     string             `json:"gestation_period"`
	EstimatedPopulation int                `json:"extimated_population"`
	ExtinctionStatus    string             `json:"extinction_status"`
	Reproduction        string             `json:"reproduction"`
	Motility            string             `json:"motility"`
	MatingSeason        string             `json:"mating_season"`
	ModeOfBirth         string             `json:"mode_of_birth"`
	Animal_id           string             `json:"animal_id"`
}
