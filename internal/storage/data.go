package storage

type historyLine struct {
	CreatedAt int    `json:"created_at"`
	Directory string `json:"pwd"`
	Command   string `json:"command"`
}
