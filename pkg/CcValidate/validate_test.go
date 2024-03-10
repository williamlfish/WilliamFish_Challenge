package CcValidate

import "testing"

func TestCCNumberIsGood(t *testing.T) {
	var tests = []struct {
		name          string
		number        string
		expectedError string
	}{
		{
			"Should be valid cc number with dashes",
			"4234-5678-9123-4567",
			"",
		},
		{
			"Should fail because bad chars",
			"4234_5678-9123-4567",
			InvalidCharsMessage,
		},
		{
			"Should fail because bad format",
			"42345678-9123-4567",
			InvalidFormatting,
		},
		{
			"Should fail because bad format with spacing",
			"423-45678-9123-4567",
			InvalidFormatSpacing,
		},
		{
			"Should fail because missing number in dashes groupings",
			"4234-5678-9123-456",
			InvalidFormatSpacing,
		},
		{
			"Should fail because first number is invalid",
			"1234-5678-9123-4564",
			InvalidFirstNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := Validate(tt.number)
			if cc.ErrorMsg != tt.expectedError {
				t.Errorf("got %s, and expecting %s on test: %s", cc.ErrorMsg, tt.expectedError, tt.name)
			}
		})
	}
}
