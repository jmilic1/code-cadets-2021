package validators

import "github.com/superbet-group/code-cadets-2021/homework_4/betacceptance/internal/api/controllers/models"

// BetRequestValidator validates event update requests.
type BetRequestValidator struct {
	paymentFromInclusive   float64
	paymentToInclusive     float64
	coefficientToInclusive float64
}

// NewBetRequestValidator creates a new instance of BetRequestValidator.
func NewBetRequestValidator(paymentFromInclusive float64, paymentToInclusive float64, coefficientToInclusive float64) *BetRequestValidator {
	return &BetRequestValidator{
		paymentFromInclusive:   paymentFromInclusive,
		paymentToInclusive:     paymentToInclusive,
		coefficientToInclusive: coefficientToInclusive,
	}
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
	return !b.isAnyFieldEmpty(betRequestDto) &&
		betRequestDto.SelectionCoefficient <= b.coefficientToInclusive &&
		betRequestDto.Payment >= b.paymentFromInclusive && betRequestDto.Payment <= b.paymentToInclusive
}
