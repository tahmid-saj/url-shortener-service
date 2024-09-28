package url

import (
	"context"
	"testing"
)

func TestShortenAndRetrieve(t *testing.T) {
	testURL := "https://github.com/tahmid-saj/url-shortener-service"
	sp := ShortenParams{
		URL: testURL,
	}
	
	res, err := Shorten(context.Background(), &sp)
	if err != nil {
		t.Fatal(err)
	}

	expectURL := testURL
	if res.URL != expectURL {
		t.Errorf("got %q, expected %q", res.URL, expectURL)
	}

	firstURL := res
	gotURL, err := Get(context.Background(), res.ID)
	if err != nil {
		t.Fatal(err)
	}

	if *gotURL != *firstURL {
		t.Errorf("got %q, expected %q", *gotURL, *firstURL)
	}
}