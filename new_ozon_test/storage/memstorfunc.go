package storage

import (
	"errors"
	"new_ozon_test/createlink"
)

type Data struct {
	FullLink  string `json:"full_link" db:"full_link"`
	ShortLink string `json:"short_link" db:"short_link"`
}

var longLinks = make(map[string]string)
var shortLinks = make(map[string]string)

func (s *Data) SaveLongLink() (string, error) {
	if shortLink, ok := longLinks[s.FullLink]; !ok {
		link := createlink.CreateLink()
		longLinks[s.FullLink] = link
		shortLinks[link] = s.FullLink
		return link, nil
	} else {
		return shortLink, nil
	}
}

func (s *Data) GetLongLink() (string, error) {
	if fullLink, ok := shortLinks[s.ShortLink]; !ok {
		return "", errors.New("full URL does not exist")
	} else {
		return fullLink, nil
	}
}
