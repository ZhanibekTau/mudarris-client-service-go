package requests

type Request []RefreshedData

type RefreshedData struct {
	EventType      string    `json:"eventType"`
	EventTime      string    `json:"eventTime"`
	OrganizationId string    `json:"organizationId"`
	EventInfo      EventInfo `json:"eventInfo"`
}

type EventInfo struct {
	Id             string          `json:"id"`
	PosId          string          `json:"posId"`
	OrganizationId string          `json:"organizationId"`
	CreationStatus string          `json:"creationStatus"`
	ErrorInfo      ErrorInfo       `json:"errorInfo"`
	Order          OrderTerminal   `json:"order"`
	Group          []TerminalGroup `json:"terminalGroupsStopListsUpdates"`
}

type TerminalGroup struct {
	Id     string `json:"id"`
	IsFull bool   `json:"isFull"`
}

type OrderTerminal struct {
	TerminalGroupId string `json:"terminalGroupId"`
	Status          string `json:"status"`
}

type ErrorInfo struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}
