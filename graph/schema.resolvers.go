package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Golang-GraphQL-API/database"
	"Golang-GraphQL-API/graph/generated"
	"Golang-GraphQL-API/graph/model"
	"context"
	"fmt"
)

// CreateCharacter is the resolver for the createCharacter field.
func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.CreateCharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented: CreateCharacter - createCharacter"))
}

// UpdateCharacter is the resolver for the updateCharacter field.
func (r *mutationResolver) UpdateCharacter(ctx context.Context, id string, input model.UpdateCharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented: UpdateCharacter - updateCharacter"))
}

// DeleteCharacter is the resolver for the deleteCharacter field.
func (r *mutationResolver) DeleteCharacter(ctx context.Context, id string) (*model.DeleteCharacterResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteCharacter - deleteCharacter"))
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.CreateAuthorInput) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: CreateAuthor - createAuthor"))
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, input model.UpdateAuthorInput) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: UpdateAuthor - updateAuthor"))
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.DeleteAuthorResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteAuthor - deleteAuthor"))
}

// CreateComicBook is the resolver for the createComicBook field.
func (r *mutationResolver) CreateComicBook(ctx context.Context, input model.CreateComicBookInput) (*model.ComicBook, error) {
	panic(fmt.Errorf("not implemented: CreateComicBook - createComicBook"))
}

// UpdateComicBook is the resolver for the updateComicBook field.
func (r *mutationResolver) UpdateComicBook(ctx context.Context, id string, input model.UpdateComicBookInput) (*model.ComicBook, error) {
	panic(fmt.Errorf("not implemented: UpdateComicBook - updateComicBook"))
}

// DeleteComicBook is the resolver for the deleteComicBook field.
func (r *mutationResolver) DeleteComicBook(ctx context.Context, id string) (*model.DeleteComicBookResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteComicBook - deleteComicBook"))
}

// CreateComicBookStore is the resolver for the createComicBookStore field.
func (r *mutationResolver) CreateComicBookStore(ctx context.Context, input model.CreateComicBookStoreInput) (*model.ComicBookStore, error) {
	panic(fmt.Errorf("not implemented: CreateComicBookStore - createComicBookStore"))
}

// UpdateComicBookStore is the resolver for the updateComicBookStore field.
func (r *mutationResolver) UpdateComicBookStore(ctx context.Context, id string, input model.UpdateComicBookStoreInput) (*model.ComicBookStore, error) {
	panic(fmt.Errorf("not implemented: UpdateComicBookStore - updateComicBookStore"))
}

// DeleteComicBookStore is the resolver for the deleteComicBookStore field.
func (r *mutationResolver) DeleteComicBookStore(ctx context.Context, id string) (*model.DeleteComicBookStoreResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteComicBookStore - deleteComicBookStore"))
}

// CharacterDetails is the resolver for the character_details field.
func (r *queryResolver) CharacterDetails(ctx context.Context, id string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented: CharacterDetails - character_details"))
}

// AuthorDetails is the resolver for the author_details field.
func (r *queryResolver) AuthorDetails(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: AuthorDetails - author_details"))
}

// BookDetails is the resolver for the book_details field.
func (r *queryResolver) BookDetails(ctx context.Context, id string) (*model.ComicBook, error) {
	panic(fmt.Errorf("not implemented: BookDetails - book_details"))
}

// StoreDetails is the resolver for the store_details field.
func (r *queryResolver) StoreDetails(ctx context.Context, id string) (*model.ComicBookStore, error) {
	panic(fmt.Errorf("not implemented: StoreDetails - store_details"))
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented: Characters - characters"))
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented: Authors - authors"))
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.ComicBook, error) {
	panic(fmt.Errorf("not implemented: Books - books"))
}

// Stores is the resolver for the stores field.
func (r *queryResolver) Stores(ctx context.Context) ([]*model.ComicBookStore, error) {
	panic(fmt.Errorf("not implemented: Stores - stores"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()

func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	return db.CreateJobListing(input), nil
}
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	return db.UpdateJobListing(id, input), nil
}
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	return db.DeleteJobListing(id), nil
}
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	return db.GetJobs(), nil
}
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	return db.GetJob(id), nil
}
