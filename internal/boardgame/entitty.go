package boardgame

type BoardGame struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MinPlayers int    `json:"minPlayers"`
	MaxPlayers int    `json:"maxPlayers"`
	Duration   string `json:"duration"`
}
