package test

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/gnatunstyles/grpc-books/api"
	"github.com/gnatunstyles/grpc-books/pkg"
)

//Test for GetAuthors function
func TestGRPCServer_Authors(t *testing.T) {
	database, _ := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/db")

	defer database.Close()

	type fields struct {
		DBConn                 *sql.DB
		UnimplementedGetServer api.UnimplementedGetServer
	}
	type args struct {
		ctx context.Context
		req *api.GetAuthorRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.GetAuthorResponse
		wantErr bool
	}{
		{
			name: "test_authors",
			fields: fields{
				DBConn:                 database,
				UnimplementedGetServer: api.UnimplementedGetServer{},
			},
			args: args{
				req: &api.GetAuthorRequest{
					BookName: "Moidodyr",
				},
			},
			want:    &api.GetAuthorResponse{Authors: []string{"Chukovsky"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pkg.GRPCServer{
				DBConn:                 tt.fields.DBConn,
				UnimplementedGetServer: tt.fields.UnimplementedGetServer,
			}
			got, err := s.Authors(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GRPCServer.Authors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GRPCServer.Authors() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test for GetBooks function
func TestGRPCServer_Books(t *testing.T) {
	database, _ := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/db")

	defer database.Close()
	type fields struct {
		DBConn                 *sql.DB
		UnimplementedGetServer api.UnimplementedGetServer
	}
	type args struct {
		ctx context.Context
		req *api.GetBooksRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.GetBooksResponse
		wantErr bool
	}{
		{
			name: "test_books",
			fields: fields{
				DBConn:                 database,
				UnimplementedGetServer: api.UnimplementedGetServer{},
			},
			args: args{
				req: &api.GetBooksRequest{
					AuthorName: "Rowling",
				},
			},
			want: &api.GetBooksResponse{Books: []string{
				"Harry Potter and the goblet of fire",
				"Harry Potter and the order of the phoenix"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pkg.GRPCServer{
				DBConn:                 tt.fields.DBConn,
				UnimplementedGetServer: tt.fields.UnimplementedGetServer,
			}
			got, err := s.Books(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GRPCServer.Books() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GRPCServer.Books() = %v, want %v", got, tt.want)
			}
		})
	}
}
