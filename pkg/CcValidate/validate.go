package CcValidate

import (
	"regexp"
	"strings"
)

const InvalidCharsMessage = "invalid chars in cc number"
const InvalidMissingDigits = "cc is missing numbers"
const InvalidToManyDigits = "cc has to many numbers"
const InvalidFormatting = "Incorrect formatting in cc number"
const InvalidFormatSpacing = "Incorrect spacing with dashes in cc number"
const InvalidFirstNumber = "First number should be 4, 5, or 6"

type CCNumber struct {
	number   string
	ErrorMsg string
}

// Validate should check that a credit card number being used passes all the requirements,
// should only be number chars, and potentially have dashes split in groups of 4.
func Validate(number string) CCNumber {
	cc := CCNumber{
		number: number,
	}
	cc.validate()
	return cc
}

func (c *CCNumber) validStartingChar() bool {
	first := string(c.number[0])
	switch true {
	case first == "4":
		return true
	case first == "5":
		return true
	case first == "6":
		return true
	default:
		c.ErrorMsg = InvalidFirstNumber
		return false
	}

}
func (c *CCNumber) hasValidChars() bool {
	re := regexp.MustCompile(`^[0-9-]+$`)
	if !re.MatchString(c.number) {
		c.ErrorMsg = InvalidCharsMessage
		return false
	}
	return true
}

func (c *CCNumber) isOnlyNumbers() bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(c.number) {
		return false
	}
	return true
}

func (c *CCNumber) correctNumberLength(input string) {
	if len(input) < 16 {
		c.ErrorMsg = InvalidMissingDigits
		return
	}
	if len(input) > 16 {
		c.ErrorMsg = InvalidToManyDigits
		return
	}
	return
}

func (c *CCNumber) dashCheck() {
	splitNum := strings.Split(c.number, "-")
	if len(splitNum) != 4 {
		c.ErrorMsg = InvalidFormatting
		return
	}
	for _, s := range splitNum {
		if len(s) != 4 {
			c.ErrorMsg = InvalidFormatSpacing
			break
		}
	}
}

func (c *CCNumber) validate() {
	if !c.validStartingChar() {
		return
	}
	if !c.hasValidChars() {
		return
	}
	if c.isOnlyNumbers() {
		c.correctNumberLength(c.number)
		return
	}
	c.dashCheck()
}
