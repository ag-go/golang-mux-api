package main

import (
	"context"
	"log"

	"./entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

//FirestoreDao will allow us to connect our API with Firebase as the data source
type FirestoreDao struct{}

const (
	projectID  string = "pragmatic-reviews"
	collection string = "posts"
)

func (FirestoreDao) Save(post *entity.Post) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore Client: %v", err)
	}

	// Close client when done (until the surrounding function returns)
	defer client.Close()

	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"title": post.Title,
		"text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func (FirestoreDao) FindAll() []entity.Post {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore Client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	var posts []entity.Post = []entity.Post{}
	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		post := entity.Post{ID: doc.Data()["ID"].(int64), Title: doc.Data()["title"].(string), Text: doc.Data()["text"].(string)}
		posts = append(posts, post)
	}
	return posts
}
