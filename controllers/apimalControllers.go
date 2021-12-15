package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/hisyntax/apimal/database"
	"github.com/hisyntax/apimal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var animalCollection *mongo.Collection = database.OpenCollection(database.Client, "animal")
var validate = validator.New()

func CreateAnimalHandler(c *gin.Context) {
	//open up a dataase connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection on the insertion is done
	defer cancel()

	//declear a variable which represents the animal model struct
	var animal models.Animal

	//bind the animal model together and check if an error occured while binding
	if err := c.BindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//user the validator package to validate the user inputes to make sure some
	//some required fields are not left blank and throw an error
	validateErr := validate.Struct(&animal)
	if validateErr != nil {
		// log.Panic(validateErr)
		msg := "Some fields are left blank or the description is too small"
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}

	//count throw al the database in the database to to check for an exception
	//if the name of the animal privided is already availavle in the database or not
	count, err := animalCollection.CountDocuments(ctx, bson.M{"name": animal.Name})
	if err != nil {
		log.Panic(err)
		msg := "Sorry, an error occured"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	//if the privided animal name exists, throw an error
	//notifying the user that the animal name exists already
	if count > 0 {
		msg := "The name already exists"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	//add a created at time to the animal data to be created
	animal.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	// animal.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//Init a mongodb default id for the Id fields in the models
	animal.ID = primitive.NewObjectID()
	//assign that ID to the animal_id so that the are the same
	animal.Animal_id = animal.ID.Hex()

	//tru to save the prvided animal data and if an error occured
	//return an error message
	insert, err := animalCollection.InsertOne(ctx, animal)
	if err != nil {
		msg := "An error occured while trying the save the animal information"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	//if there has been no errors this far
	//then save the animal data
	c.JSON(http.StatusOK, gin.H{
		"name":      animal.Name,
		"desc":      animal.Description,
		"image":     animal.Image,
		"animal_id": insert,
	})
}

func GetAnimalsHandler(c *gin.Context) {
	//open up a dataase connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection on the insertion is done
	defer cancel()

	//declear an aniamls variable ro hold an array of the animal model
	var animals []models.Animal

	//declear a variable to find all the datas in the database
	//check if there is an error and handle that error appropriately
	cusor, err := animalCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Panic(err)
		msg := "Sorry, something went wrong. Please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	//user the initially create varivale to finc all the aninmal data and handle the
	// error if any
	if err = cusor.All(ctx, &animals); err != nil {
		log.Panic(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//close the connection
	defer cusor.Close(ctx)

	//check if there is an error in retrieving the animal datas
	if err := cusor.Err(); err != nil {
		log.Panic(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//list out the animal datas found
	c.JSON(http.StatusOK, gin.H{
		"animal": animals,
	})
}

func GetAnimalHandler(c *gin.Context) {
	//open up a dataase connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection on the insertion is done
	defer cancel()

	//declear a variable which holds the animal model
	var animal models.Animal

	//get the url parameter and throw an error if its empty
	animal_id := c.Param("animal_id")
	if animal_id == "" {
		msg := "Invalid parameter"
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
		return
	}

	// match animal_id gotter from the url with the animal_id in the database
	//and chek if there is an error
	animals_id, err := primitive.ObjectIDFromHex(animal_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//query through the database to get a match for the animal_if from the url
	//and check for an error
	animalResult := animalCollection.FindOne(ctx, bson.M{"_id": animals_id})
	if err := animalResult.Err(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//if there is a match, then decode it to the decleared animal variavle
	//and if an error was ecountered while decoding, handle the error
	if err := animalResult.Decode(&animal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//if everything goes well, display the individual animal with matches the animal_id
	c.JSON(http.StatusOK, gin.H{
		"animal": animal,
	})

}

func UpdateAnimalHandler(c *gin.Context) {
	//open up a dataase connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection on the insertion is done
	defer cancel()
	var animal models.Animal

	//get the url parameter and throw an error if its empty
	animal_id := c.Param("animal_id")
	if animal_id == "" {
		msg := "Invalid parameter"
		c.JSON(http.StatusBadRequest, gin.H{
			"error": msg,
		})
	}

	// match animal_id gotter from the url with the animal_id in the database
	//and chek if there is an error
	animals_id, err := primitive.ObjectIDFromHex(animal_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//bind all of the provided datas togeter and handler errors if any
	if err := c.BindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//set up a mongodb query which queries the data- animal_id
	filter := bson.D{primitive.E{Key: "_id", Value: animals_id}}

	//speifying the fields available for update
	update := bson.D{
		{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: animal.Name},
			{Key: "desc", Value: animal.Description},
			{Key: "image", Value: animal.Image},
			{Key: "habitat", Value: animal.Habitat},
			{Key: "domain", Value: animal.Domain},
			{Key: "kingdom", Value: animal.Kingdom},
			{Key: "phylum", Value: animal.Phylum},
			{Key: "class", Value: animal.Class},
			{Key: "order", Value: animal.Order},
			{Key: "family", Value: animal.Family},
			{Key: "genus", Value: animal.Genus},
			{Key: "specie", Value: animal.Specie},
			{Key: "color", Value: animal.Color},
			{Key: "predator", Value: animal.Predator},
			{Key: "food_type", Value: animal.FoodType},
			{Key: "ovulation_period", Value: animal.OvuationPeriod},
			{Key: "gestation_period", Value: animal.GestationPeriod},
			{Key: "extimated_population", Value: animal.EstimatedPopulation},
			{Key: "extinction_status", Value: animal.ExtinctionStatus},
			{Key: "reproduction", Value: animal.Reproduction},
			{Key: "motility", Value: animal.Motility},
			{Key: "mating_season", Value: animal.MatingSeason},
			{Key: "mode_of_birth", Value: animal.ModeOfBirth},
		}}}

	//only the error that is needed
	//try to update a specific animal data and if an error occured, handler the error
	_, err = animalCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Done()

	//if everything goes fine output some information to the user at to what has been updated
	c.JSON(http.StatusOK, gin.H{
		"name":    animal.Name,
		"desc":    animal.Description,
		"message": "successfully updated the animal detail",
	})
}

func DeleteAnimalHandler(c *gin.Context) {
	//open up a dataase connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection on the insertion is done
	defer cancel()

	//get the url parameter and throw an error if its empty
	animal_id := c.Param("animal_id")
	if animal_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}

	// match animal_id gotter from the url with the animal_id in the database
	//and chek if there is an error
	animals_id, err := primitive.ObjectIDFromHex(animal_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//set up a mongodb query which queries the data- animal_id
	filter := bson.D{primitive.E{Key: "_id", Value: animals_id}}

	//only the error that is needed
	//try to delete a specific animal data and if an error occured, handler the error
	_, err = animalCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong while deleting the data",
		})
		return
	}
	ctx.Done()

	//if everything goes fine
	//let the user know that the animal data has been successfully deleted with the animal_id
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted animal with the provided ID",
		"animal":  animals_id,
	})
}

func SearchAnimalHandler(c *gin.Context) {
	//open up a dataase connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection on the insertion is done
	defer cancel()

	//declear a variable which holds the animal models
	var animal []models.Animal

	//get the providel query parameter from the url
	//and if the query paramater is empty throw an error
	query := c.Query("search")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}

	//searh the database to get a match for the  query paramater the user provided
	//and if there was an error, handle it
	queryAnimal, err := animalCollection.Find(ctx, bson.M{"name": bson.M{"$regex": query}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	//Get all the possible match of the query parameter the user entered
	//and if there is an error, handle it
	if err := queryAnimal.All(ctx, &animal); err != nil {
		msg := "Something went wrong while trying to find all the names"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
	}
	defer queryAnimal.Close(ctx)

	//if there is an while trying to get all the related data throw an error
	if err := queryAnimal.Err(); err != nil {
		msg := "Invalid request"
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": msg,
		})
		return
	}

	//if everything goes fine show the list of all the match from the database
	c.JSON(http.StatusOK, gin.H{
		"animal": animal,
	})

}
