package tests

import (
	"context"
	"testing"
	"url-shortener/models/url"
)

func allStrinsEqual(a []string) bool {
	for i := 1; i < len(a); i++ {
        if a[i] != a[0] {
            return false
        }
    }
    return true
}


func TestShortAllCodesUnique(t * testing.T) {

	codes := []string{}

	plainUrl := &url.PlainUrl{
		Url: "https://www.google.com",
	}

	iterCount := 1 << 16

	ctx := context.Background()
	for i := 0; i < iterCount; i++ {
		value, _ := url.Shorten(ctx, i, url.DefaultCutSize, plainUrl)
		codes = append(codes, value.ShortCode)
	}

	if allStrinsEqual(codes) {
		t.Errorf("allStringsEqual(): all codes must be unique, but aren't")
	}

}
