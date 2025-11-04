package url

import (
	"context"
	"crypto/md5"
	"github.com/jxskiss/base62"
	"errors"
	"io"
	"fmt"
)

const (
	DefaultCutSize = 6
)

type PlainUrl struct
{
	Url string `json:"url"`
}

type ShortUrl struct
{
	Id 	int				`json:"id" bson:"uniq_number"`
	Url string 			`json:"url"`
	ShortCode string 	`json:"code"`
}


func computeHash(url string) []byte {

	hash := md5.New()
	io.WriteString(hash, url)	

	return hash.Sum(nil)
}

func encodeHash(h [] byte) [] byte {
	return base62.Encode(h)
}

func cutOffBeginning(b64Encoded [] byte, num int) [] byte {
	return b64Encoded[:num]
}


func Shorten(ctx context.Context, id int, codelength int, pu * PlainUrl) (*ShortUrl, error) {

	if codelength <= 0 {
		codelength = DefaultCutSize
	}

	if pu == nil {
		return nil, errors.New("null pointer to url")
	}
	
	uniqUrl := fmt.Sprintf("%s_%d", pu.Url, id)
	hashedUrl := computeHash(uniqUrl)
	encodedUrl := encodeHash(hashedUrl)

	shortUrl := &ShortUrl{}
	shortUrl.Id = id
	shortUrl.Url = pu.Url
	shortUrl.ShortCode = string(cutOffBeginning(encodedUrl, codelength))

	return shortUrl, nil
}