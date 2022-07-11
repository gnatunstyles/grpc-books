package main

import (
	"log"
	"net"
	"os"

	"github.com/gnatunstyles/grpc-books/api"
	"github.com/gnatunstyles/grpc-books/db"
	"github.com/gnatunstyles/grpc-books/pkg"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {

	//initzializing db
	dsn := os.Getenv("MYSQL_CONN_STRING")
	Database, err := db.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()

	//initializing grpc server
	s := grpc.NewServer()
	srv := pkg.NewGRPCServer()
	srv.DBConn = Database.DB

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
