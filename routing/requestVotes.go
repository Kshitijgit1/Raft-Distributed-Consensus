package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// RequestVotes is the signal that triggers the raft
// behaviour in server clusters
func RequestVotes(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request from client in requestVotes.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq ReqVotesRequest
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in requestVotes.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := types.ServerData[r.Host]
	vote := voting(state, r.Host, newReq.ServerID)
	outJSON, err := json.Marshal(vote)
	if err != nil {
		log.Printf("Can't Marshall to JSON in requestVotes.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}

func voting(state *types.State, IP string, RemoteID int) (vote int) {
	state.Lock.Lock()
	if state.VotedFor == -2 {
		log.Printf("I %s can vote\n", IP)
		vote = 1
		state.VotedFor = RemoteID
	} else {
		log.Printf("I %s cant vote\n", IP)
		vote = 0
	}
	state.Lock.Unlock()
	return
}
