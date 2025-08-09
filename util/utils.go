package util

import (
	"errors"
	"message-service/logger"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func ValidateMessage(msgText string) error {
	msgText = strings.TrimSpace(msgText)
	logger.Info("msgText: " + msgText)
	if len(msgText) == 0 {
		return errors.New("Message cannot be empty")
	}
	if len(msgText) > 30 {
		return errors.New("messageText cannot exceed 30 characters")
	}
	return nil
}
