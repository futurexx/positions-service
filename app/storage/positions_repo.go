package storage

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
