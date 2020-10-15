package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/teams"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	teamsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/teams"
	teamsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/teams"
	"github.com/gorilla/mux"
)

var team = &models.TeamView{
	ID:          777,
	Name:        "Dream Team",
	Description: "Best team ever",
	Slack:       "QWERTY",
	Crated:      time.Now().UTC(),
	Updated:     time.Now().UTC(),
}

func TestHandlers_CreateTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := teamsSrv.New(client, repo)
	h := New(s)

	reqBody := &models.TeamForm{
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	}
	reqByte, _ := json.Marshal(reqBody)
	body := bytes.NewReader(reqByte)

	req, _ := http.NewRequest("POST", "/teams", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateTeam)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":777,"message":"successfully created"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlers_DeleteTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := teamsSrv.New(client, repo)
	h := New(s)

	req, _ := http.NewRequest("DELETE", "/teams", nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatInt(team.ID, 10)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeleteTeam)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":777,"message":"successfully removed"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlers_UpdateTeam(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := teamsSrv.New(client, repo)
	h := New(s)

	reqBody := &models.TeamForm{
		Name:        team.Name,
		Description: team.Description,
		Slack:       team.Slack,
	}
	reqByte, _ := json.Marshal(reqBody)
	body := bytes.NewReader(reqByte)

	req, _ := http.NewRequest("PUT", "/teams", body)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatInt(team.ID, 10)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdateTeam)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":777,"message":"successfully updated"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlers_GetTeams(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	s := teamsSrv.New(client, repo)
	h := New(s)

	req, _ := http.NewRequest("GET", "/teams", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetTeams)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	view := []*models.TeamView{
		{
			ID:          team.ID,
			Name:        team.Name,
			Description: team.Description,
			Slack:       team.Slack,
			Crated:      team.Crated,
			Updated:     team.Updated,
		},
	}

	byteView, _ := json.Marshal(view)
	expected := string(byteView)

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func newClientMock() *teams.Client {
	return teams.NewMock(
		team.ID,
		team.Name,
		team.Description,
		team.Slack,
		team.Crated,
		team.Updated)
}

func newRepoMock() *teamsRepo.RepositoryMock {
	return teamsRepo.NewMock(
		team.ID,
		team.Name,
		team.Description,
		team.Slack,
		team.Crated,
		team.Updated)
}
