package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//TODO consider moving collection to the struct. lots of repeated code.

type Database interface {
	CreatePatient(ctx context.Context, p Patient) (Patient, error)
	ReadPatient(ctx context.Context, id int) (Patient, error)
	UpdatePatient(ctx context.Context, id int, p Patient) (Patient, error)
	DeletePatient(ctx context.Context, id int) error
	Search(ctx context.Context, k string, v string) ([]int, error)
}

type monngoDb struct {
	db *mongo.Database
}

func newMongoDb(cs string) Database {
	clientOptions := options.Client().ApplyURI(cs)
	fmt.Printf("connecting to mongo")
	var err error
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Printf("starting mongo: %v\n", err)
		os.Exit(1)
	}
	return &monngoDb{
		db: client.Database("PatientService"),
	}
}

func (mdb *monngoDb) CreatePatient(ctx context.Context, p Patient) (Patient, error) {
	col := mdb.db.Collection("Patients")

	_, err := col.DeleteMany(context.Background(), bson.M{"id": p.ID})
	if err != nil {
		return Patient{}, fmt.Errorf("resolving conflict: %v", err)
	}
	res, err := col.InsertOne(ctx, p)
	if err != nil {
		return Patient{}, fmt.Errorf("inserting to db: %v", err)
	}
	fmt.Println(res.InsertedID)
	patient := Patient{}
	err = col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&patient)
	if err != nil {
		return Patient{}, fmt.Errorf("finding after insert to db: %v", err)
	}
	return patient, nil
}
func (mdb *monngoDb) ReadPatient(ctx context.Context, id int) (Patient, error) {
	fmt.Println(id)
	col := mdb.db.Collection("Patients")
	patient := Patient{}

	err := col.FindOne(ctx, bson.M{"id": id}).Decode(&patient)
	if err != nil {
		return Patient{}, fmt.Errorf("finding in db: %v", err)
	}

	return patient, nil
}
func (mdb *monngoDb) UpdatePatient(ctx context.Context, id int, p Patient) (Patient, error) {
	col := mdb.db.Collection("Patients")
	patient := Patient{}

	_, err := col.ReplaceOne(ctx, bson.M{"id": id}, p)
	if err != nil {
		log.Fatal(err)
	}

	err = col.FindOne(ctx, bson.M{"id": id}).Decode(&patient)
	if err != nil {
		return Patient{}, fmt.Errorf("finding after update to db: %v", err)
	}
	return patient, nil
}
func (mdb *monngoDb) DeletePatient(ctx context.Context, id int) error {
	col := mdb.db.Collection("Patients")

	_, err := col.DeleteMany(context.Background(), bson.M{"id": id})
	if err != nil {
		return fmt.Errorf(" deleting from db: %v", err)
	}
	return nil
}
func (mdb *monngoDb) Search(ctx context.Context, k, v string) ([]int, error) {
	col := mdb.db.Collection("Patients")
	fmt.Printf("%v\n", bson.M{k: v})
	cur, err := col.Find(ctx, bson.M{k: v})
	if err != nil {
		return []int{}, fmt.Errorf("searching db: %v", err)
	}

	defer cur.Close(ctx)
	ids := []int{}
	for cur.Next(ctx) {
		fmt.Println("FOUND")
		p := Patient{}
		err := cur.Decode(&p)
		if err != nil {
			return []int{}, fmt.Errorf("paring patient search: %v", err)
		}
		ids = append(ids, p.ID)
	}

	return ids, nil
}
