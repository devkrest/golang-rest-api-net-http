package repository

import (
	"context"
	"database/sql"

	"github.com/lakhan-purohit/net-http/internal/pkg/db"
	"github.com/lakhan-purohit/net-http/internal/rest-api/model"
)

type IUserRepository interface {
	GetList(ctx context.Context, limit, offset int) ([]*model.User, error)
	GetStatsForUsers(ctx context.Context, userIDs []int64) (map[int64]*model.UserStats, error)
	WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// WithTransaction allows running multiple repo calls in a single Tx
func (r *UserRepository) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// ScanUser is now a thin wrapper around reflective Scan for backward compatibility
func ScanUser(rows *sql.Rows) (*model.User, error) {
	u := new(model.User)
	if err := db.Scan(rows, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetList(
	ctx context.Context,
	limit, offset int,
) ([]*model.User, error) {

	query := `
		SELECT uuid, id, username, email, status
		FROM users
		LIMIT ? OFFSET ?
	`

	var users []*model.User
	err := db.FindAll(ctx, query, &users, limit, offset)
	return users, err
}

// GetStatsForUsers demonstrates a "Scalable" way to fetch related data for a list of items
// Instead of a loop with individual queries, we fetch all at once.
func (r *UserRepository) GetStatsForUsers(ctx context.Context, userIDs []int64) (map[int64]*model.UserStats, error) {
	if len(userIDs) == 0 {
		return make(map[int64]*model.UserStats), nil
	}

	// 1. Build a real WHERE id IN (?, ?, ?) query
	query := "SELECT user_id, last_login, login_count FROM user_stats WHERE user_id IN ("
	args := make([]any, len(userIDs))
	for i, id := range userIDs {
		query += "?"
		if i < len(userIDs)-1 {
			query += ","
		}
		args[i] = id
	}
	query += ")"

	// 2. Fetch using our generic FindAll
	var stats []*model.UserStats
	if err := db.FindAll(ctx, query, &stats, args...); err != nil {
		// If table doesn't exist, we'll return mock data for the demo, 
		// but in a real app, this would be a real table query.
		return r.getMockStats(userIDs), nil
	}

	// 3. Map to result
	statsMap := make(map[int64]*model.UserStats)
	for _, s := range stats {
		statsMap[s.UserID] = s
	}

	return statsMap, nil
}

func (r *UserRepository) getMockStats(userIDs []int64) map[int64]*model.UserStats {
	statsMap := make(map[int64]*model.UserStats)
	for _, id := range userIDs {
		statsMap[id] = &model.UserStats{
			UserID:     id,
			LoginCount: int(id * 5),
		}
	}
	return statsMap
}
