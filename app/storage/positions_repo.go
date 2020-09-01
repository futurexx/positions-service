package storage

import "time"

// PositionsRepo is repo for `positions` table
type PositionsRepo struct {
	storage *Storage
}

// Summary get number of positions for domain
func (p *PositionsRepo) Summary(domain string) (uint, error) {
	var count uint
	err := p.storage.db.QueryRow(
		"SELECT COUNT(position) FROM positions WHERE domain=?",
		domain).Scan(&count)

	return count, err
}

type DomainPosition struct {
	Keyword  string
	Position uint
	URL      string
	Volume   uint
	Results  uint
	Cpc      float32
	Updated  time.Time
}

// Positions get list of position for domain
func (p *PositionsRepo) Positions(domain string, order string, limit uint64, offset uint64) ([]DomainPosition, error) {
	var res []DomainPosition

	rows, err := p.storage.db.Query(
		"SELECT keyword, position, url, volume, results, cpc, updated FROM positions WHERE domain=? ORDER BY ? LIMIT ?, ?",
		domain, order, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := DomainPosition{}

		err := rows.Scan(&t.Keyword, &t.Position, &t.URL, &t.Volume, &t.Results, &t.Cpc, &t.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, t)
	}

	return res, nil
}
