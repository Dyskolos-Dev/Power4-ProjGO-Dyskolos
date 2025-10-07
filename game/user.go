package power4

// User représente un utilisateur du jeu
type User struct {
	Name  string `json:"name"`
	Color string `json:"color"` // "red" ou "yellow"
}

// GameSession représente une session de jeu
type GameSession struct {
	Player1     User   `json:"player1"`
	Player2     User   `json:"player2"`
	CurrentTurn string `json:"currentTurn"` // nom du joueur dont c'est le tour
	GameMode    string `json:"gameMode"`    // "local" pour tour par tour local
}

// NewGameSession crée une nouvelle session de jeu
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
