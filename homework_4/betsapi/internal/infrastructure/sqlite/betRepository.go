package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
	}
}

// readBet tries to read a single bet from given sql rows.
// If sql row has no more rows to scan, an error is returned.
func (b *BetRepository) readBet(row *sql.Rows) (models.Bet, error) {
	var id string
	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	err := row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
	if err != nil {
		return models.Bet{}, err
	}

	var payout int
	if payoutSql.Valid {
		payout = int(payoutSql.Int64)
	}

	return models.Bet{
		Id:                   id,
		CustomerId:           customerId,
		Status:               status,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
		Payout:               payout,
	}, nil
}

// GetBetById fetches a bet from the database and returns it. The second returned value indicates
// whether the bet exists in DB. If the bet does not exist, an error will not be returned.
func (b *BetRepository) GetBetById(ctx context.Context, id string) (models.Bet, bool, error) {
	storageBet, err := b.queryGetBetById(ctx, id)
	if err == sql.ErrNoRows {
		return models.Bet{}, false, nil
	}
	if err != nil {
		return models.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	exists := storageBet != models.Bet{}
	return storageBet, exists, nil
}

func (b *BetRepository) queryGetBetById(ctx context.Context, id string) (models.Bet, error) {
	selectBetSql := "SELECT  * FROM bets WHERE id=?"

	row, err := b.dbExecutor.QueryContext(ctx, selectBetSql, id)
	if err != nil {
		return models.Bet{}, err
	}
	defer row.Close()

	row.Next()

	return b.readBet(row)
}

// GetBetsByCustomerId fetches bets from database which match given customerId.
// If bets do not exist, an error will not be returned.
func (b *BetRepository) GetBetsByCustomerId(ctx context.Context, customerId string) ([]models.Bet, error) {
	storageBets, err := b.queryGetBetsByCustomerId(ctx, customerId)
	if err == sql.ErrNoRows {
		return []models.Bet{}, nil
	}
	if err != nil {
		return []models.Bet{}, errors.Wrap(err, "bet repository failed to get a bet with customerId "+customerId)
	}

	return storageBets, nil
}

func (b *BetRepository) queryGetBetsByCustomerId(ctx context.Context, customerId string) ([]models.Bet, error) {
	selectBetSql := "SELECT  * FROM bets WHERE customer_id=?"

	row, err := b.dbExecutor.QueryContext(ctx, selectBetSql, customerId)
	if err != nil {
		return []models.Bet{}, err
	}
	defer row.Close()

	var bets []models.Bet

	found := row.Next()

	for found {
		bet, err := b.readBet(row)
		if err != nil {
			return []models.Bet{}, err
		}
		bets = append(bets, bet)
		found = row.Next()
	}

	return bets, nil
}

// GetBetsByStatus fetches bets from database which match given status.
// If bets do not exist, an error will not be returned.
func (b *BetRepository) GetBetsByStatus(ctx context.Context, status string) ([]models.Bet, error) {
	storageBets, err := b.queryGetBetsByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []models.Bet{}, nil
	}
	if err != nil {
		return []models.Bet{}, errors.Wrap(err, "bet repository failed to get a bet with id "+status)
	}

	return storageBets, nil
}

func (b *BetRepository) queryGetBetsByStatus(ctx context.Context, status string) ([]models.Bet, error) {
	selectBetSql := "SELECT  * FROM bets WHERE status=?"

	row, err := b.dbExecutor.QueryContext(ctx, selectBetSql, status)
	if err != nil {
		return []models.Bet{}, err
	}
	defer row.Close()

	var bets []models.Bet

	found := row.Next()

	for found {
		bet, err := b.readBet(row)
		if err != nil {
			return []models.Bet{}, err
		}
		bets = append(bets, bet)
		found = row.Next()
	}

	return bets, nil
}
