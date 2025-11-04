package tests

// import (
// 	"context"
// 	"testing"
// 	"url-shortener/models/url"
// )

// func allStrinsEqual(a []string) bool {
// 	for i := 1; i < len(a); i++ {
//         if a[i] != a[0] {
//             return false
//         }
//     }
//     return true
// }

// func Test_SequenceShortening_equal_urls(t * testing.T) {

// 	ss := url.SequenceShortener{}

// 	codes := []string{}

// 	plainUrl := &url.PlainUrl{
// 		Url: "https://www.google.com",
// 	}

// 	iterCount := 10000

// 	ctx := context.Background()
// 	for i := 0; i < iterCount; i++ {
// 		value, _ := ss.Shorten(ctx, url.DefaultCutSize, plainUrl)
// 		codes = append(codes, value.ShortCode)
// 	}

// 	if allStrinsEqual(codes) == true {
// 		t.Errorf("not all unique: %v", codes)
// 	}

// }
