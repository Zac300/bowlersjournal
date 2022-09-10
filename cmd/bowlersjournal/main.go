package main

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/zac300/bowlersjournal/db/sqlc"
	"github.com/zac300/bowlersjournal/gqlgen"
)

func main() {
	// initialize the db
	d, err := db.Open("dbname=gqlgen_sqlc_example_db user=root password=secret sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer d.Close()

	// initialize the repository
	repo := db.NewRepository(d)

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	mux.Handle("/query", gqlgen.NewHandler(repo))

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
