package players_killed

import "github.com/ddd/pkg/support"

var (
	WorldKillerLabel = "world"
)

type PlayerKilled interface {
	Find()
	GetTotal() int64
	Compute(Matchable)
	IsEligibleToBeAPlayer(Matchable) bool
	GetPlayers() Players
}

type PlayerKilledEntity struct {
	Total   int64
	Player  Playable
	Players Players
}

func NewPlayerKilled(player Playable) PlayerKilled {
	return &PlayerKilledEntity{
		Player: player,
	}
}

func (p *PlayerKilledEntity) Find() {

}

func (p *PlayerKilledEntity) GetTotal() int64 {
	return p.Total
}

func (p *PlayerKilledEntity) Compute(match Matchable) {
	err := match.Validation()
	support.PanicOnError(err, "")

	if !p.IsEligibleToBeAPlayer(match) {
		p.Player.KillDown(match)
	}
	p.Player.KillUp(match)
}

func (p *PlayerKilledEntity) IsEligibleToBeAPlayer(match Matchable) bool {
	return match.GetPlayerWhoKilled() != WorldKillerLabel
}

func (p *PlayerKilledEntity) GetPlayers() Players {
	return p.Player.GetPlayers()
}
