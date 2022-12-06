package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Golang-GraphQL-API/database"
	"Golang-GraphQL-API/graph/generated"
	"Golang-GraphQL-API/graph/model"
	"context"
)

var db = database.Connect()

// CreateCharacter is the resolver for the createCharacter field.
func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.CreateCharacterInput) (*model.Character, error) {
	return db.CreateCharacter(input), nil
}

// UpdateCharacter is the resolver for the updateCharacter field.
func (r *mutationResolver) UpdateCharacter(ctx context.Context, id string, input model.UpdateCharacterInput) (*model.Character, error) {
	return db.UpdateCharacter(id, input), nil
}

// DeleteCharacter is the resolver for the deleteCharacter field.
func (r *mutationResolver) DeleteCharacter(ctx context.Context, id string) (*model.DeleteCharacterResponse, error) {
	return db.DeleteCharacter(id), nil
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.CreateAuthorInput) (*model.Author, error) {
	return db.CreateAuthor(input), nil
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, input model.UpdateAuthorInput) (*model.Author, error) {
	return db.UpdateAuthor(id, input), nil
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.DeleteAuthorResponse, error) {
	return db.DeleteAuthor(id), nil
}

// CreateComicBook is the resolver for the createComicBook field.
func (r *mutationResolver) CreateComicBook(ctx context.Context, input model.CreateComicBookInput) (*model.ComicBook, error) {
	return db.CreateComicBook(input), nil
}

// UpdateComicBook is the resolver for the updateComicBook field.
func (r *mutationResolver) UpdateComicBook(ctx context.Context, id string, input model.UpdateComicBookInput) (*model.ComicBook, error) {
	return db.UpdateComicBook(id, input), nil
}

// DeleteComicBook is the resolver for the deleteComicBook field.
func (r *mutationResolver) DeleteComicBook(ctx context.Context, id string) (*model.DeleteComicBookResponse, error) {
	return db.DeleteComicBook(id), nil
}

// CreateComicBookStore is the resolver for the createComicBookStore field.
func (r *mutationResolver) CreateComicBookStore(ctx context.Context, input model.CreateComicBookStoreInput) (*model.ComicBookStore, error) {
	return db.CreateComicBookStore(input), nil
}

// UpdateComicBookStore is the resolver for the updateComicBookStore field.
func (r *mutationResolver) UpdateComicBookStore(ctx context.Context, id string, input model.UpdateComicBookStoreInput) (*model.ComicBookStore, error) {
	return db.UpdateComicBookStore(id, input), nil
}

// DeleteComicBookStore is the resolver for the deleteComicBookStore field.
func (r *mutationResolver) DeleteComicBookStore(ctx context.Context, id string) (*model.DeleteComicBookStoreResponse, error) {
	return db.DeleteComicBookStore(id), nil
}

// CharacterDetails is the resolver for the character_details field.
func (r *queryResolver) CharacterDetails(ctx context.Context, id string) (*model.Character, error) {
	return db.GetCharacter(id), nil
}

// AuthorDetails is the resolver for the author_details field.
func (r *queryResolver) AuthorDetails(ctx context.Context, id string) (*model.Author, error) {
	return db.GetAuthor(id), nil
}

// BookDetails is the resolver for the book_details field.
func (r *queryResolver) BookDetails(ctx context.Context, id string) (*model.ComicBook, error) {
	return db.GetComicBook(id), nil
}

// StoreDetails is the resolver for the store_details field.
func (r *queryResolver) StoreDetails(ctx context.Context, id string) (*model.ComicBookStore, error) {
	return db.GetComicBookStore(id), nil
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return db.GetCharacters(), nil
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return db.GetAuthors(), nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.ComicBook, error) {
	return db.GetComicBooks(), nil
}

// Stores is the resolver for the stores field.
func (r *queryResolver) Stores(ctx context.Context) ([]*model.ComicBookStore, error) {
	return db.GetComicBookStores(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
