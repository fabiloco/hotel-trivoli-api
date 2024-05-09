package utils

import "regexp"

func ReceiptIdPatternMatch(str string) bool {
	receiptPattern := regexp.MustCompile(`^r-\d+$`)
	return receiptPattern.MatchString(str)
}

func IndividualReceiptIdPatternMatch(str string) bool {
	individualReceiptPattern := regexp.MustCompile(`^ir-\d+$`)
	return individualReceiptPattern.MatchString(str)
}
