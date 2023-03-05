package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

type Provider struct {
	ctx context.Context

	guard sync.Mutex
	conn  *pgx.Conn
}

// NewProvider creates Provider instance
func NewProvider(ctx context.Context, conn *pgx.Conn) *Provider {
	return &Provider{
		ctx:  ctx,
		conn: conn,
	}
}

// PopCode returns the first matched code and remove it from db
func (p *Provider) PopCode(productID, price int) (string, error) {
	p.guard.Lock()
	defer p.guard.Unlock()

	var code string

	query := `
		SELECT code FROM codes
		WHERE id_goods = $1 AND price = $2
		LIMIT 1
	`
	if err := p.conn.QueryRow(p.ctx, query, productID, price).Scan(&code); err != nil || code == "" {
		return "", fmt.Errorf("failed to select code by price %d, err: %s", price, err.Error())
	}

	query = `
		DELETE FROM codes
		WHERE code = $1
	`
	if _, err := p.conn.Exec(p.ctx, query, code); err != nil {
		return "", fmt.Errorf("failed to delete code %s, err: %s", code, err.Error())
	}

	return code, nil
}

// SaveIssued saves issued code to db
func (p *Provider) SaveIssued(uniqueCode string, code string, price int, email string) error {
	p.guard.Lock()
	defer p.guard.Unlock()

	query := `
		INSERT INTO issued_codes(unique_code, code, price, email)
		VALUES ($1, $2, $3, $4)
	`
	if _, err := p.conn.Exec(p.ctx, query, uniqueCode, code, price, email); err != nil {
		return fmt.Errorf("failed to insert issued code %s, err: %s", code, err.Error())
	}

	return nil
}

func (p *Provider) GetIssued(uniqueCode string) (string, bool) {
	p.guard.Lock()
	defer p.guard.Unlock()

	var code string

	query := `
		SELECT code FROM issued_codes
		WHERE unique_code = $1
		LIMIT 1
	`
	if err := p.conn.QueryRow(p.ctx, query, uniqueCode).Scan(&code); err != nil || code == "" {
		return "", false
	}

	return code, true
}

// Close release db connection
func (p *Provider) Close() {
	_ = p.conn.Close(p.ctx)
}
