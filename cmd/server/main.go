package main

import (
	"context"
	"fmt"
	"github.com/vicrosa25/comment-api/internal/db"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// application
func Run() error {
	fmt.Println("Starting application...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged to database")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
