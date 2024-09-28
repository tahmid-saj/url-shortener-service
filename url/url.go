package url

import (
	"context"
	"encoding/base64"
	"math/rand"
)

var db = sqldb.NewDatabase("url", sqldb.DatabaseConfig{
	Migration: "./migrations",
})

type URL struct {
	ID  string // short form URL id
	URL string // complete URL, in long form
}

type ShortenParams struct {
	URL string // the URL to shorten
}

// Shorten shortens a URL
//
//encore:api public method=POST path=/url
func Shorten(context context.Context, p *ShortenParams) (*URL, error) {
	id, err := generateID()
	if err != nil {
		return nil, err
	} else if err := insert(context, id, p.URL); err != nil {
		return nil, err
	}

	return &URL{
		ID: id,
		URL: p.URL,
	}, nil
}

// Get retrieves the original URL for the id
// 
//encore:api public method=GET path=/url/:id
func Get(context context.Context, id string) (*URL, error) {
	url := &URL{
		ID: id,
	}
	err := db.QueryRow(context, `
		SELECT original_url FROM url
		WHERE id = $1
	`, id).Scan(&url.URL)
	
	return url, err
}

// generateID generates a random short ID
func generateID() (string, error) {
	var data [6]byte // 6 bytes of entropy
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}

// insert inserts a URL into the database
func insert(context context.Context, id, url string) error {
	_, err := db.Exec(context, `
		INSERT INTO url (id, original_url)
		VALUES ($1, $2)
	`, id, url)

	return err
}