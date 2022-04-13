package storage

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"new_ozon_test/createlink"
	"os"
	"sync"
)

var TypeStorage = os.Getenv("STORAGE")

type Storage interface {
	AddLink(full string) (Data, error)
	GetLink(short string) (Data, error)
}

type Psql struct {
	Conn *sqlx.DB
	Mu   *sync.Mutex
}

type Memory struct {
	LongLinks  map[string]string
	ShortLinks map[string]string
	Mu         *sync.Mutex
}

type Data struct {
	FullLink  string `json:"full_link" db:"full_link"`
	ShortLink string `json:"short_link" db:"short_link"`
}

func (m *Memory) AddLink(full string) (Data, error) {
	if full == "" {
		return Data{}, errors.New("empty URL")
	}
	data := Data{}
	m.Mu.Lock()
	if shortLink, ok := m.LongLinks[full]; !ok {
		m.LongLinks[full] = createlink.CreateLink()
		data.ShortLink = m.LongLinks[full]
		m.ShortLinks[data.ShortLink] = full
	} else {
		data.ShortLink = shortLink
	}
	data.FullLink = full
	m.Mu.Unlock()
	return data, nil
}

func (m *Memory) GetLink(short string) (Data, error) {
	if short == "" {
		return Data{}, errors.New("empty short link")
	}
	m.Mu.Lock()
	if fullLink, ok := m.ShortLinks[short]; !ok {
		m.Mu.Unlock()
		return Data{}, errors.New("full URL does not exist")
	} else {
		data := Data{FullLink: fullLink, ShortLink: short}
		m.Mu.Unlock()
		return data, nil
	}
}

func (p *Psql) AddLink(full string) (Data, error) {
	if full == "" {
		return Data{}, errors.New("empty URL")
	}
	data := Data{}
	p.Mu.Lock()
	err := p.Conn.Get(&data, "SELECT * FROM links WHERE full_link=$1;", full)
	p.Mu.Unlock()
	if err != nil {
		data.FullLink = full
		data.ShortLink = createlink.CreateLink()
		p.Mu.Lock()
		_, err = p.Conn.NamedQuery("INSERT INTO links VALUES (:full_link, :short_link)", data)
		p.Mu.Unlock()
		if err != nil {
			return Data{}, err
		}
	}
	return data, nil
}

func (p *Psql) GetLink(short string) (Data, error) {
	if short == "" {
		return Data{}, errors.New("empty short link")
	}
	data := Data{}
	p.Mu.Lock()
	err := p.Conn.Get(&data, "SELECT * FROM links WHERE short_link=$1;", short)
	p.Mu.Unlock()
	if err != nil {
		return data, errors.New("url does not Exist")
	}
	return data, nil
}
