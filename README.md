# realtime-go

Realtime-go is a client library for [supabase/realtime](https://github.com/supabase/realtime).

The package interface closely mirrors the [supabase/realtime-js](https://github.com/supabase/realtime-js) library.

## Usage

```go
package main

import (
	"log"
	realtimego "github.com/overseedio/realtime-go"
)

func main() {
	// phoneix realtime server endpoint
	const ENDPOINT = "https://{SUPABASE_DB_ID}.supabase.co"
	// gateway api key
	const API_KEY = "..."
	// (optional) auth token
	const RLS_TOKEN = "..."

	// create client
	c, err := realtimego.NewClient(ENDPOINT, API_KEY,
		realtimego.WithUserToken(RLS_TOKEN),
	)
	if err != nil {
		log.Fatal(err)
	}

	// connect to server
	err = c.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// create and subscribe to channel
	db := "realtime"
	schema := "public"
	table := "my_table"
	ch, err := c.Channel(realtimego.WithTable(&db, &schema, &table))
	if err != nil {
		log.Fatal(err)
	}

	// setup hooks
	ch.OnInsert = func(m realtimego.Message) {
		log.Println("***ON INSERT....", m)
	}
	ch.OnDelete = func(m realtimego.Message) {
		log.Println("***ON DELETE....", m)
	}
	ch.OnUpdate = func(m realtimego.Message) {
		log.Println("***ON UPDATE....", m)
	}

	// subscribe to channel
	err = ch.Subscribe()
	if err != nil {
		log.Fatal(err)
	}
}

```
