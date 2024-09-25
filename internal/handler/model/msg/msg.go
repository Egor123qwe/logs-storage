package msg

type MSG struct {
	// Type for different events in package event
	Type string `json:"type"`
	// Content for different events in package content
	Content []Log `json:"content"`
}
