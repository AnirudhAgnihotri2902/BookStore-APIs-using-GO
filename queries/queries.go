package queries

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/books"
	"fmt"
	"log"
)

func CreateBookEntry(ctx context.Context, client *ent.Client, userBook string, userAuthor string, userPublishDate string) (*ent.Books, error) {
	u, err := client.Books.
		Create().
		SetBookName(userBook).
		SetAuthor(userAuthor).
		SetPublishDate(userPublishDate).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func DeleteByBookName(ctx context.Context, client *ent.Client, userBook string) error {
	_, err := client.Books.
		Delete().
		Where(books.BookName(userBook)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("Book Deleted")
	return nil
}

func QueryforAllBook(ctx context.Context, client *ent.Client) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func QueryforBookName(ctx context.Context, client *ent.Client, userBook string) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		Where(books.BookName(userBook)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func QueryforAuthorName(ctx context.Context, client *ent.Client, userAuthor string) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		Where(books.Author(userAuthor)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func QueryforPublishDate(ctx context.Context, client *ent.Client, userPublishDate string) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		Where(books.PublishDate(userPublishDate)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

// incomplete...
func QueryforPublishDateAfter(ctx context.Context, client *ent.Client, userPublishDate string) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		Where(books.PublishDate(userPublishDate)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

// incomplete..
func QueryforPublishDateBefore(ctx context.Context, client *ent.Client, userPublishDate string) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		Where(books.PublishDate(userPublishDate)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
