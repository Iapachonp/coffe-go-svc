package models

type Process struct {
	ProcessId        int    `json:"processid"`
	Name		 string `json:"name"`
	Description      string `json:"description"`
	TimeDuration     string `json:"timeduration"`
}
