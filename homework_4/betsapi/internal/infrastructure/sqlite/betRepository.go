package sqlite

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/betsapi/internal/infrastructure/sqlite/models"
)

// BetRepository provides methods that operate on bets SQLite database.
type BetRepository struct {
	dbExecutor DatabaseExecutor
	betMapper  BetMapper
}

// NewBetRepository creates and returns a new BetRepository.
func NewBetRepository(dbExecutor DatabaseExecutor, betMapper BetMapper) *BetRepository {
	return &BetRepository{
		dbExecutor: dbExecutor,
		betMapper:  betMapper,
	}
}

// readBet tries to read a single bet from given sql rows.
// If sql row has no more rows to scan, an error is returned.
func (b *BetRepository) readBet(row *sql.Rows) (storagemodels.Bet, error) {
	var id string
	var customerId string
	var status string
	var selectionId string
	var selectionCoefficient int
	var payment int
	var payoutSql sql.NullInt64

	err := row.Scan(&id, &customerId, &status, &selectionId, &selectionCoefficient, &payment, &payoutSql)
	if err != nil {
		return storagemodels.Bet{}, err
	}

	var payout int
	if payoutSql.Valid {
		payout = int(payoutSql.Int64)
	}

	return storagemodels.Bet{
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
func (b *BetRepository) GetBetById(ctx context.Context, id string) (domainmodels.Bet, bool, error) {
	storageBet, err := b.queryGetBetById(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return domainmodels.Bet{}, false, errors.Wrap(err, "bet repository failed to get a bet with id "+id)
	}

	domainBet := b.betMapper.MapStorageBetToDomainBet(storageBet)
	exists := domainBet != domainmodels.Bet{}
	return domainBet, exists, nil
}

func (b *BetRepository) queryGetBetById(ctx context.Context, id string) (storagemodels.Bet, error) {
	row, err := b.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.Bet{}, err
	}
	defer row.Close()

	row.Next()

	return b.readBet(row)
}

// GetBetsByCustomerId fetches bets from database which match given customerId.
// If bets do not exist, an error will not be returned.
func (b *BetRepository) GetBetsByCustomerId(ctx context.Context, customerId string) ([]domainmodels.Bet, error) {
	storageBets, err := b.queryGetBetsByCustomerId(ctx, customerId)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, errors.Wrap(err, "bet repository failed to get a bet with customerId "+customerId)
	}

	var domainBets []domainmodels.Bet

	for _, bet := range storageBets {
		domainBet := b.betMapper.MapStorageBetToDomainBet(bet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, nil
}

func (b *BetRepository) queryGetBetsByCustomerId(ctx context.Context, customerId string) ([]storagemodels.Bet, error) {
	row, err := b.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE customer_id='"+customerId+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	found := row.Next()

	for found {
		bet, err := b.readBet(row)
		if err != nil {
			return []storagemodels.Bet{}, err
		}
		bets = append(bets, bet)
		found = row.Next()
	}

	return bets, nil
}

// GetBetsByStatus fetches bets from database which match given status.
// If bets do not exist, an error will not be returned.
func (b *BetRepository) GetBetsByStatus(ctx context.Context, status string) ([]domainmodels.Bet, error) {
	storageBets, err := b.queryGetBetsByStatus(ctx, status)
	if err == sql.ErrNoRows {
		return []domainmodels.Bet{}, nil
	}
	if err != nil {
		return []domainmodels.Bet{}, errors.Wrap(err, "bet repository failed to get a bet with id "+status)
	}

	var domainBets []domainmodels.Bet

	for _, bet := range storageBets {
		domainBet := b.betMapper.MapStorageBetToDomainBet(bet)
		domainBets = append(domainBets, domainBet)
	}

	return domainBets, nil
}

func (b *BetRepository) queryGetBetsByStatus(ctx context.Context, status string) ([]storagemodels.Bet, error) {
	row, err := b.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE status='"+status+"';")
	if err != nil {
		return []storagemodels.Bet{}, err
	}
	defer row.Close()

	var bets []storagemodels.Bet

	found := row.Next()

	for found {
		bet, err := b.readBet(row)
		if err != nil {
			return []storagemodels.Bet{}, err
		}
		bets = append(bets, bet)
		found = row.Next()
	}

	return bets, nil
}
