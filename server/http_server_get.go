package server

import (
	"encoding/json"
	"net/http"

	"github.com/chehsunliu/poker"
)

type staticTestData struct {
	Info string `json:"info"`
	Data []int  `json:"data"`
}

func (h *HttpServer) handleGetStaticTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		data := staticTestData{
			Info: "Some static test data",
			Data: []int{
				1, 2, 3, 4, 5,
			},
		}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Status of poker game in open phase.
func (h *HttpServer) handlePokerGetGameOpenStatus() http.HandlerFunc {
	type gameOpenInfo struct {
		Open               bool     `json:"open"`
		Players            []player `json:"players"`
		PlayerAmount       int      `json:"playerAmount"`
		InitialPlayerMoney int      `json:"initialPlayerMoney"`
		SmallBlindValue    int      `json:"smallBlindValue"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.Info("handlePokerGetGameOpenStatus called")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		gameOpenInfo := gameOpenInfo{
			Open:               pokerGameStart.open,
			Players:            pokerGameStart.players,
			PlayerAmount:       len(pokerGameStart.players),
			InitialPlayerMoney: pokerGameStart.initialPlayerMoney,
			SmallBlindValue:    pokerGameStart.smallBlindValue,
		}

		if err := json.NewEncoder(w).Encode(gameOpenInfo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Status of poker game in active phase.
func (h *HttpServer) handlePokerGetGameActiveStatus() http.HandlerFunc {
	type gameActiveInfo struct {
		Active          bool         `json:"active"`
		CommunityCards  []poker.Card `json:"communityCards"`
		Players         []player     `json:"players"`
		PlayerAmount    int          `json:"playerAmount"`
		CurrentRound    int          `json:"currentRound"`
		CurrentPlayer   int          `json:"currentPlayer"`
		SmallBlindValue int          `json:"smallBlindValue"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		h.logger.Info("handlePokerGetGameActiveStatus called")

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		gameActiveInfo := gameActiveInfo{
			Active:          pokerGame.active,
			CommunityCards:  pokerGame.communityCards,
			Players:         pokerGame.players,
			PlayerAmount:    len(pokerGame.players),
			CurrentRound:    pokerGame.currentRound,
			CurrentPlayer:   pokerGame.currentPlayer,
			SmallBlindValue: pokerGame.smallBlindValue,
		}

		if err := json.NewEncoder(w).Encode(gameActiveInfo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
