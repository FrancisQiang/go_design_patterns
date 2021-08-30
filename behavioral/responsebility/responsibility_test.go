package responsebility

import (
	"testing"
)

func TestResponsibility(t *testing.T) {
	twoElementValidator := TwoElementValidator{}
	var oneCardOneAccountValidator Validator
	oneCardOneAccountValidator = &OneCardOneAccountValidator{}
	twoElementValidator.SetNext(&oneCardOneAccountValidator)
	twoElementValidator.Validate("刘老肥 3212375645245877")
}
