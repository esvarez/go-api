package boardgame

type Reader interface {
	FindBoardGame(id int) (*BoardGame, error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	FinByID(id int) (BoardGame, error)
}
