package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"gitlab.com/pragmaticreviews/golang-mux-api/entity"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
}

type repo struct{}

//NewPostRepository creates a new repo
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "pragmatic-reviews"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, err = client.Collection(collectionName).Doc(post.ID).Set(ctx, post)
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(string),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (*repo) FindByID(id string) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	dsnap, err := client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	post := &entity.Post{
		ID:    dsnap.Data()["ID"].(string),
		Title: dsnap.Data()["Title"].(string),
		Text:  dsnap.Data()["Text"].(string),
	}
	return post, nil
}
