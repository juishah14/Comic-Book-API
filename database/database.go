package database

import (
	"context"
	"log"
	"time"

	model "Golang-GraphQL-API/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// change this later
var connectionString string = "mongodb+srv://akhil:akhil@cluster0.td5oga4.mongodb.net/test"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetCharacter(id string) *model.Character {
	var characterCollection = db.client.Database("Cluster0").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var character model.Character
	err := characterCollection.FindOne(ctx, filter).Decode(&character)
	if err != nil {
		log.Fatal(err)
	}
	return &character
}

func (db *DB) GetCharacters() []*model.Character {
	var characterCollection = db.client.Database("Cluster0").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var characters []*model.Character
	cursor, err := characterCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &characters); err != nil {
		panic(err)
	}
	return characters
}

func (db *DB) CreateCharacter(characterInfo model.CreateCharacterInput) *model.Character {
	var characterCollection = db.client.Database("Cluster0").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	insert, err := characterCollection.InsertOne(ctx, bson.M{"name": characterInfo.Name, "description": characterInfo.Description, "comic_book_series": characterInfo.ComicBookSeries, "superpower": characterInfo.Superpower, "danger_level": characterInfo.DangerLevel})
	if err != nil {
		log.Fatal(err)
	}

	insertedID := insert.InsertedID.(primitive.ObjectID).Hex()
	character := model.Character{ID: insertedID, Name: characterInfo.Name, Description: characterInfo.Description, ComicBookSeries: characterInfo.ComicBookSeries, Superpower: characterInfo.Superpower, DangerLevel: characterInfo.DangerLevel}
	return &character
}

func (db *DB) UpdateCharacter(characterId string, characterInfo model.UpdateCharacterInput) *model.Character {
	var characterCollection = db.client.Database("Cluster0").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateCharacterInfo := bson.M{}

	if characterInfo.Name != nil {
		updateCharacterInfo["name"] = characterInfo.Name
	}
	if characterInfo.Description != nil {
		updateCharacterInfo["description"] = characterInfo.Description
	}
	if characterInfo.ComicBookSeries != nil {
		updateCharacterInfo["comic_book_series"] = characterInfo.ComicBookSeries
	}
	if characterInfo.Superpower != nil {
		updateCharacterInfo["superpower"] = characterInfo.Superpower
	}
	if characterInfo.DangerLevel != nil {
		updateCharacterInfo["danger_level"] = characterInfo.DangerLevel
	}

	_id, _ := primitive.ObjectIDFromHex(characterId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateCharacterInfo}
	results := characterCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var character model.Character
	err := results.Decode(&character)
	if err != nil {
		log.Fatal(err)
	}
	return &character
}

func (db *DB) DeleteCharacter(characterId string) *model.DeleteCharacterResponse {
	var characterCollection = db.client.Database("Cluster0").Collection("characters")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(characterId)
	filter := bson.M{"_id": _id}
	_, err := characterCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteCharacterResponse{deletedCharacterId: characterId}
}

func (db *DB) GetAuthor(id string) *model.Author {
	var authorCollection = db.client.Database("Cluster0").Collection("authors")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var author model.Author
	err := authorCollection.FindOne(ctx, filter).Decode(&author)
	if err != nil {
		log.Fatal(err)
	}
	return &author
}

func (db *DB) GetAuthors() []*model.Author {
	var authorCollection = db.client.Database("Cluster0").Collection("authors")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var authors []*model.Author
	cursor, err := authorCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &authors); err != nil {
		panic(err)
	}
	return authors
}

func (db *DB) CreateAuthor(authorInfo model.CreateAuthorInput) *model.Author {
	var authorCollection = db.client.Database("Cluster0").Collection("authors")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	insert, err := authorCollection.InsertOne(ctx, bson.M{"first_name": authorInfo.FirstName, "last_name": authorInfo.LastName, "description": authorInfo.Description, "birth_date": authorInfo.BirthDate, "comic_books_written": authorInfo.ComicBooksWritten})
	if err != nil {
		log.Fatal(err)
	}

	insertedID := insert.InsertedID.(primitive.ObjectID).Hex()
	author := model.Author{ID: insertedID, FirstName: authorInfo.FirstName, LastName: authorInfo.LastName, BirthDate: authorInfo.BirthDate, ComicBooksWritten: authorInfo.ComicBooksWritten}
	return &author
}

func (db *DB) UpdateAuthor(authorId string, authorInfo model.UpdateAuthorInput) *model.Author {
	var authorCollection = db.client.Database("Cluster0").Collection("authors")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateAuthorInfo := bson.M{}

	if authorInfo.FirstName != nil {
		updateAuthorInfo["first_name"] = authorInfo.FirstName
	}
	if authorInfo.LastName != nil {
		updateAuthorInfo["last_name"] = authorInfo.LastName
	}
	if authorInfo.Description != nil {
		updateAuthorInfo["description"] = authorInfo.Description
	}
	if authorInfo.BirthDate != nil {
		updateAuthorInfo["birth_date"] = authorInfo.BirthDate
	}
	if authorInfo.ComicBooksWritten != nil {
		updateAuthorInfo["comic_books_written"] = authorInfo.ComicBooksWritten
	}

	_id, _ := primitive.ObjectIDFromHex(authorId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateAuthorInfo}
	results := authorCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var author model.Author
	err := results.Decode(&author)
	if err != nil {
		log.Fatal(err)
	}
	return &author
}

func (db *DB) DeleteAuthor(authorId string) *model.DeleteAuthorResponse {
	var authorCollection = db.client.Database("Cluster0").Collection("authors")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(authorId)
	filter := bson.M{"_id": _id}
	_, err := authorCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteAuthorResponse{deletedAuthorId: authorId}
}

func (db *DB) GetComicBook(id string) *model.ComicBook {
	var comicCollection = db.client.Database("Cluster0").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var comic model.ComicBook
	err := comicCollection.FindOne(ctx, filter).Decode(&comic)
	if err != nil {
		log.Fatal(err)
	}
	return &comic
}

func (db *DB) GetComicBooks() []*model.ComicBook {
	var comicCollection = db.client.Database("Cluster0").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var books []*model.ComicBook
	cursor, err := comicCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &books); err != nil {
		panic(err)
	}
	return books
}

func (db *DB) CreateComicBook(bookInfo model.CreateComicBookInput) *model.ComicBook {
	var comicCollection = db.client.Database("Cluster0").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	insert, err := comicCollection.InsertOne(ctx, bson.M{"name": bookInfo.Name, "description": bookInfo.Description, "publication_date": bookInfo.PublicationDate, "author": bookInfo.Author})
	if err != nil {
		log.Fatal(err)
	}

	insertedID := insert.InsertedID.(primitive.ObjectID).Hex()
	book := model.ComicBook{ID: insertedID, Name: bookInfo.Name, Description: bookInfo.Description, PublicationDate: bookInfo.PublicationDate, Author: bookInfo.Author}
	return &book
}

func (db *DB) UpdateComicBook(bookId string, bookInfo model.UpdateComicBookInput) *model.ComicBook {
	var comicCollection = db.client.Database("Cluster0").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateBookInfo := bson.M{}

	if bookInfo.Name != nil {
		updateBookInfo["name"] = bookInfo.Name
	}
	if bookInfo.Description != nil {
		updateBookInfo["description"] = bookInfo.Description
	}
	if bookInfo.PublicationDate != nil {
		updateBookInfo["publication_date"] = bookInfo.PublicationDate
	}
	if bookInfo.Author != nil {
		updateBookInfo["author"] = bookInfo.Author
	}

	_id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateBookInfo}
	results := comicCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var book model.ComicBook
	err := results.Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	return &book
}

func (db *DB) DeleteComicBook(bookId string) *model.DeleteComicBookResponse {
	var comicCollection = db.client.Database("Cluster0").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": _id}
	_, err := comicCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteComicBookResponse{deletedComicBookId: bookId}
}

func (db *DB) GetComicBookStore(id string) *model.ComicBookStore {
	var storeCollection = db.client.Database("Cluster0").Collection("stores")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var store model.ComicBookStore
	err := storeCollection.FindOne(ctx, filter).Decode(&store)
	if err != nil {
		log.Fatal(err)
	}
	return &store
}

func (db *DB) GetComicBookStores() []*model.ComicBookStore {
	var storeCollection = db.client.Database("Cluster0").Collection("stores")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var stores []*model.ComicBookStore
	cursor, err := storeCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &stores); err != nil {
		panic(err)
	}
	return stores
}

func (db *DB) CreateComicBookStore(storeInfo model.CreateComicBookStoreInput) *model.ComicBookStore {
	var storeCollection = db.client.Database("Cluster0").Collection("stores")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	insert, err := storeCollection.InsertOne(ctx, bson.M{"name": storeInfo.Name, "location": storeInfo.Location, "postal_code": storeInfo.PostalCode, "description": storeInfo.Description, "phone_number": storeInfo.PhoneNumber, "store_owner": storeInfo.StoreOwner})
	if err != nil {
		log.Fatal(err)
	}

	insertedID := insert.InsertedID.(primitive.ObjectID).Hex()
	store := model.ComicBookStore{ID: insertedID, Name: storeInfo.Name, Location: storeInfo.Location, PostalCode: storeInfo.PostalCode, Description: storeInfo.Description}
	return &store
}

func (db *DB) UpdateComicBookStore(storeId string, storeInfo model.UpdateComicBookStoreInput) *model.ComicBookStore {
	var storeCollection = db.client.Database("Cluster0").Collection("stores")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateStoreInfo := bson.M{}

	if storeInfo.Name != nil {
		updateStoreInfo["name"] = storeInfo.Name
	}
	if storeInfo.Location != nil {
		updateStoreInfo["location"] = storeInfo.Location
	}
	if storeInfo.PostalCode != nil {
		updateStoreInfo["postal_code"] = storeInfo.PostalCode
	}
	if storeInfo.Description != nil {
		updateStoreInfo["description"] = storeInfo.Description
	}
	if storeInfo.PhoneNumber != nil {
		updateStoreInfo["phone_number"] = storeInfo.PhoneNumber
	}
	if storeInfo.StoreOwner != nil {
		updateStoreInfo["store_owner"] = storeInfo.StoreOwner
	}

	_id, _ := primitive.ObjectIDFromHex(storeId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateStoreInfo}
	results := storeCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var store model.ComicBookStore
	err := results.Decode(&store)
	if err != nil {
		log.Fatal(err)
	}
	return &store
}

func (db *DB) DeleteComicBookStore(storeId string) *model.DeleteComicBookStoreResponse {
	var storeCollection = db.client.Database("Cluster0").Collection("stores")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(storeId)
	filter := bson.M{"_id": _id}
	_, err := storeCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteComicBookStoreResponse{deletedComicBookStoreId: storeId}
}
