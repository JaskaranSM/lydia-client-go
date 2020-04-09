package types

type payload struct{
	SessionId string `json:"session_id"`
	Language string `json:"language"`
	Available bool `json:"available"`
	Expires int `json:"expires"`
}

type SessionResponse struct{
	Success string `json:"success"`
	ResponseCode int `json:"response_code"`
	Payload payload `json:"payload"`
}

type ThoughtPayload struct{
	Output string `json:"output"`
}

type ThoughtResponse struct{
	Success string `json:"success"`
	ResponseCode int `json:"response_code"`
	Payload ThoughtPayload `json:"payload"`
}

