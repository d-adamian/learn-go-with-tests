package poker

import (
	"io"
	"time"
)

type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func (p *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := (5 + time.Duration(numberOfPlayers)) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}

func NewGame(alerter BlindAlerter, store PlayerStore) Game {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}
