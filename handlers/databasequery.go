package handlers

import (
	"context"
	"encoding/json"
	"entdemo/ent"
	"entdemo/queries"
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

	if _, err = queries.CreateBookEntry(ctx, client, userBook, userAuthor, userPublishDate); err != nil {
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

	var u, errr = queries.QueryforAllBook(ctx, client)
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

	var u, errr = queries.QueryforBookName(ctx, client, userBook)
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

	var errr = queries.DeleteByBookName(ctx, client, userBook)
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

	var u, errr = queries.QueryforAuthorName(ctx, client, userAuthor)
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

	var u, errr = queries.QueryforPublishDate(ctx, client, userPublishDate)
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Your Book Details are , %s", u)))
}
