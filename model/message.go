package model

// Message is a struct that represents a message with a severity level and content.
type Message struct {
	Id           string `json:"messageId,omitempty"`
	Text         string `json:"messageText"`
	IsPalindrome bool   `json:"isPalindrome,omitempty"`
}
