package storage

import "time"

type Storage interface {
	Save(p *Page) error
	PickRandom(UserName string) (*Page, error)
	Remove(p *Page) error
	Exists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	Created  time.Time
}
