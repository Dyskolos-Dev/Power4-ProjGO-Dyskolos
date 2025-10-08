package game

type User struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type GameSession struct {
	Player1     User   `json:"player1"`
	Player2     User   `json:"player2"`
	CurrentTurn string `json:"currentTurn"`
	GameMode    string `json:"gameMode"`
}

func NewGameSession(player1Name, player2Name string) *GameSession {
	return &GameSession{
		Player1: User{
			Name:  player1Name,
			Color: "red",
		},
		Player2: User{
			Name:  player2Name,
			Color: "yellow",
		},
		CurrentTurn: player1Name,
		GameMode:    "local",
	}
}
