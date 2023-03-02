package db

import (
	"fmt"
	"sync"
	"time"
)

type Provider struct {
	guard sync.Mutex
}

// NewProvider creates Provider instance
func NewProvider() *Provider {
	return &Provider{}
}

// PopCode returns the first matched code and remove it from db
func (p *Provider) PopCode(price int) (string, error) {
	p.guard.Lock()
	defer p.guard.Unlock()

	return fmt.Sprintf("%d:%d", price, time.Now().Unix()), nil
}

// SaveIssued saves issued code to db
func (p *Provider) SaveIssued(
	uniqueCode string,
	code string,
	price int,
	email string,
	datePay string,
) error {
	p.guard.Lock()
	defer p.guard.Unlock()

	return nil
}
