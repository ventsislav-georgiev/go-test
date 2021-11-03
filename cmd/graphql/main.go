package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

type user struct {
	Name *string
}

type query struct{}

func (*query) Me(ctx context.Context) *user {
	fields := graphql.SelectedFieldsFromContext(ctx)
	fieldsJSON, _ := json.Marshal(fields)
	log.Print(string(fieldsJSON))

	name := "Ventsi"
	return &user{
		Name: &name,
	}
}

func main() {
	env := flag.String("env", "DEV", "env: DEV (for local development)")
	flag.Parse()

	srvAddress := ":8080"
	if *env == "DEV" {
		srvAddress = "localhost" + srvAddress
	}

	http.HandleFunc("/graphql", graphqlHandler)
	log.Fatal(http.ListenAndServe(srvAddress, nil))
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprint(w, err)
		}
	}()

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := `
		type Query {
			me(): User
		}

		type User {
			name: String
		}
	`
	schema := graphql.MustParseSchema(s, &query{}, graphql.UseFieldResolvers())
	response := schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
