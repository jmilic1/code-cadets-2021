package validators

import "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/models"

const paymentFromInclusive = 2.0
const paymentToInclusive = 100.0
const coefficientToInclusive = 10.0

// BetRequestValidator validates event update requests.
type BetRequestValidator struct{}

// NewBetRequestValidator creates a new instance of BetRequestValidator.
func NewBetRequestValidator() *BetRequestValidator {
	return &BetRequestValidator{}
}

// isWithinRangeInclusive checks if given value is within given range [from, to]
func (b *BetRequestValidator) isWithinRangeInclusive(value, from, to float64) bool {
	return value >= from && value <= to
}

// isLessThanInclusive checks if given value is less than or equal to given upper bound
func (b *BetRequestValidator) isLessThanInclusive(value, to float64) bool {
	return value <= to
}

// isAnyFieldEmpty returns true if any field has default value, false otherwise
func (b *BetRequestValidator) isAnyFieldEmpty(dto models.BetRequestDto) bool {
	return dto.SelectionCoefficient == 0 || dto.Payment == 0 || dto.SelectionId == "" || dto.CustomerId == ""
}

// BetRequestIsValid checks if event update is valid.
// Fields are not empty
// SelectionCoefficient is <= 10.0
// Payment is in range [2.0, 100.0]
func (b *BetRequestValidator) BetRequestIsValid(betRequestDto models.BetRequestDto) bool {
	return !b.isAnyFieldEmpty(betRequestDto) && b.isLessThanInclusive(betRequestDto.SelectionCoefficient, coefficientToInclusive) && b.isWithinRangeInclusive(betRequestDto.Payment, paymentFromInclusive, paymentToInclusive)
}
