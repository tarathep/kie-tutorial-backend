package db

import (
	"context"
	"fmt"
	"log"

	"github.com/tarathep/tutorial-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//TutorialRepository is contain interface
type TutorialRepository interface {
	Create(tutorial model.Tutorial) error
	FindAll(title string) ([]*model.Tutorial, error)
	FindOne(id string) (model.Tutorial, error)
	Update(tutorial model.Tutorial) error
	Delete(id string) error
	DeleteAll() error
	FindAllPublished() ([]*model.Tutorial, error)
}

func (db *MongoDB) Create(tutorial model.Tutorial) error {
	collection := db.Database("tutorialdb").Collection("tutorials")
	_, err := collection.InsertOne(context.TODO(), tutorial)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) FindAll(title string) ([]*model.Tutorial, error) {

	collection := db.Database("tutorialdb").Collection("tutorials")

	findOptions := options.Find()

	var results []*model.Tutorial

	// filter := bson.D{{Key: "foo", Value: 99}}
	filter := bson.M{"title": bson.M{"$regex": title}}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Tutorial
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results, nil
}

func (db *MongoDB) FindOne(id string) (model.Tutorial, error) {
	var tutorial model.Tutorial
	collection := db.Database("tutorialdb").Collection("tutorials")

	//String hex to ObjId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("-->", err)
		return tutorial, err
	}

	if err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&tutorial); err != nil {
		fmt.Println("-->", err)
		return tutorial, err
	}
	return tutorial, nil
}

func (db *MongoDB) Update(tutorial model.Tutorial) error {
	collection := db.Database("tutorialdb").Collection("tutorials")
	//UpdateMany or UpdateOne
	result, err := collection.UpdateMany(
		context.TODO(),
		bson.M{"_id": tutorial.ID},
		bson.D{
			{"$set", bson.D{{"title", tutorial.Title}}},
			{"$set", bson.D{{"description", tutorial.Description}}},
			{"$set", bson.D{{"published", tutorial.Published}}},
			{"$set", bson.D{{"updatedat", tutorial.UpdatedAt}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return nil
}

func (db *MongoDB) Delete(id string) error {

	//String hex to ObjId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("-->", err)
		return err
	}

	collection := db.Database("tutorialdb").Collection("tutorials")
	result, err := collection.DeleteMany(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Delete(S) removed %v document(s)\n", result.DeletedCount)
	return nil
}

func (db *MongoDB) DeleteAll() error {
	collection := db.Database("tutorialdb").Collection("tutorials")
	result, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("DeleteMany removed %v document(s)\n", result.DeletedCount)
	return nil
}

//FindAllPublished func is find all published into mongodb is true
func (db *MongoDB) FindAllPublished() ([]*model.Tutorial, error) {
	var results []*model.Tutorial

	collection := db.Database("tutorialdb").Collection("tutorials")

	cur, err := collection.Find(context.TODO(), bson.M{"published": true}, options.Find())
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem model.Tutorial
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results, nil
}
