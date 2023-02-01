package handlers

import (
	"context"
	"encoding/json"
	"entdemo/ent"
	"entdemo/ent/books"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var userBook string
var userAuthor string
var userPublishDate string

type BookDetails struct {
	Bookname    string `json:"bookName"`
	Authorname  string `json:"author"`
	Publishdate string `json:"publishDate`
}

func CreateBookEntry(ctx context.Context, client *ent.Client) (*ent.Books, error) {
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

func DeleteByBookName(ctx context.Context, client *ent.Client) error {
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

func QueryforBookName(ctx context.Context, client *ent.Client) ([]*ent.Books, error) {
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

func QueryforAuthorName(ctx context.Context, client *ent.Client) ([]*ent.Books, error) {
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

func QueryforPublishDate(ctx context.Context, client *ent.Client) ([]*ent.Books, error) {
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
func QueryforPublishDateAfter(ctx context.Context, client *ent.Client) ([]*ent.Books, error) {
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

// incomplete..
func QueryforPublishDateBefore(ctx context.Context, client *ent.Client) ([]*ent.Books, error) {
	u, err := client.Books.
		Query().
		Where(books.PublishDate(userAuthor)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var booksdetail BookDetails
	err := json.NewDecoder(r.Body).Decode(&booksdetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userBook = booksdetail.Bookname
	userAuthor = booksdetail.Authorname
	userPublishDate = booksdetail.Publishdate

	if _, err = CreateBookEntry(ctx, client); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book is Saved at Database , %s", userBook)))
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var booksdetail BookDetails
	err := json.NewDecoder(r.Body).Decode(&booksdetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userBook = booksdetail.Bookname

	var u, errr = QueryforAllBook(ctx, client)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book Details are , %s", u)))
}

func GetBookByName(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var booksdetail BookDetails
	err := json.NewDecoder(r.Body).Decode(&booksdetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userBook = booksdetail.Bookname

	var u, errr = QueryforBookName(ctx, client)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book Details are , %s", u)))
}
func DelBookByName(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var booksdetail BookDetails
	err := json.NewDecoder(r.Body).Decode(&booksdetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userBook = booksdetail.Bookname

	var errr = DeleteByBookName(ctx, client)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book Details are deleted from Database named , %s", userBook)))
}

func GetBookByAuthor(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var booksdetail BookDetails
	err := json.NewDecoder(r.Body).Decode(&booksdetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userAuthor = booksdetail.Authorname

	var u, errr = QueryforAuthorName(ctx, client)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book Details are , %s", u)))
}

func GetBookByPublishDate(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var booksdetail BookDetails
	err := json.NewDecoder(r.Body).Decode(&booksdetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userPublishDate = booksdetail.Publishdate

	var u, errr = QueryforPublishDate(ctx, client)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book Details are , %s", u)))
}
