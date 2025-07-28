package game

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

// Domain entities -----------------------------------------------------------

type Blockchain struct {
	ID          string
	Name        string
	MarketCap   float64
	TPS         int
	LastUpdated time.Time
}

type Player struct {
	ID       string
	Username string
	Wallet   string
	Balance  float64
}

// ---------------------------------------------------------------------------

type EngineState struct {
	Blockchains map[string]*Blockchain
	Players     map[string]*Player
	Epoch       uint64 // текущий внутриигровой «тик»
}

// ----------------------------------------------------------------------------

func NewState() *EngineState {
	return &EngineState{
		Blockchains: make(map[string]*Blockchain),
		Players:     make(map[string]*Player),
		Epoch:       0,
	}
}

// Snapshot / Restore ---------------------------------------------------------

func (s *EngineState) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *EngineState) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

// ---------------------------------------------------------------------------

func (s *EngineState) UpsertPlayer(p Player) {
	s.Players[p.ID] = &p
}

func (s *EngineState) GetPlayer(id string) (*Player, error) {
	pl, ok := s.Players[id]
	if !ok {
		return nil, errors.New("player not found")
	}
	return pl, nil
}
