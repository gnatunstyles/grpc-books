package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/gnatunstyles/grpc-books/api"
	"github.com/gnatunstyles/grpc-books/pkg"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {

	//initzializing db
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/db")
	if err != nil {
		panic("Failed to connect to db")
	}
	fmt.Println("database connected successfully")

	//initializing grpc server
	s := grpc.NewServer()
	srv := pkg.NewGRPCServer()
	srv.DBConn = db

	api.RegisterGetServer(s, srv)
	//open listener at port 8080
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	//accepting incoming connections with listener
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
