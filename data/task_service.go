package data

import (
	"context"
	"fmt"
	"time"

	"Task_manager/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func StartMongo () (*mongo.Client, error) {
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	collection = client.Database("task_manager").Collection("tasks")

	fmt.Println("connected to MongoDB")
	return client, nil
}

// To get All Tasks
func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tasks []models.Task
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// To get Task by ID
func GetTaskByID(id int) *models.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var task models.Task
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil
	}
	return &task
}

// To update Task by ID
func UpdateTaskByID(id int, updatedTask models.Task) *models.Task {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return nil
	}
	fmt.Printf("Updated task %v\n", id)
	return &updatedTask
}

// To Delete Task by ID
func RemoveTaskByID(id int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deleteResult, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return false
	}

	fmt.Printf("Deleted %v documents in the tasks collection\n", deleteResult.DeletedCount)
	return true
}

// To Add New Task
func AddNewTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
