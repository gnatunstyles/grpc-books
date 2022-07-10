package pkg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gnatunstyles/grpc-books/api"
)

//gRPC server structure
type GRPCServer struct {
	DBConn *sql.DB
	api.UnimplementedGetServer
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

//Getting author of the book handler
func (s *GRPCServer) Authors(ctx context.Context, req *api.GetAuthorRequest) (*api.GetAuthorResponse, error) {
	results := []string{}
	rows, err := s.DBConn.Query("select * from books where name = ?", req.GetBookName())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(rows)

	for rows.Next() {
		book := api.Book{}
		err = rows.Scan(&book.Id, &book.Bookname, &book.Author)
		if err != nil {
			return nil, err
		}
		results = append(results, book.Author)
	}
	return &api.GetAuthorResponse{Authors: results}, nil
}

//Getting author's books handler
func (s *GRPCServer) Books(ctx context.Context, req *api.GetBooksRequest) (*api.GetBooksResponse, error) {
	results := []string{}

	rows, err := s.DBConn.Query("select * from books where author = ?", req.GetAuthorName())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(rows)

	for rows.Next() {
		book := api.Book{}
		err = rows.Scan(&book.Id, &book.Bookname, &book.Author)
		if err != nil {
			return nil, err
		}
		results = append(results, book.Bookname)
	}

	return &api.GetBooksResponse{Books: results}, nil
}
