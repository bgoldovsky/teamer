package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/persons"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	personsRepo "github.com/bgoldovsky/dutyer/gateway-api/internal/repostiory/persons"
	personsSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/persons"
	"github.com/gorilla/mux"
)

var person = &models.PersonView{
	ID:        777,
	TeamId:    888,
	FirstName: "Boris",
	LastName:  "B",
	Slack:     "QWERTY",
	Role:      2,
	Created:   time.Now().UTC(),
	Updated:   time.Now().UTC(),
}

func TestHandlers_CreatePerson(t *testing.T) {
	client := newClientMock()
	repo := newRepoMock()
	router := mux.NewRouter()
	s := personsSrv.New(client, repo)
	h := New(router, "", s)

	reqBody := &models.PersonForm{
		TeamId:    person.TeamId,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Role:      person.Role,
		IsActive:  person.IsActive,
	}
	reqByte, _ := json.Marshal(reqBody)
	body := bytes.NewReader(reqByte)

	req, _ := http.NewRequest("POST", "/teams", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreatePerson)
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
	router := mux.NewRouter()
	s := personsSrv.New(client, repo)
	h := New(router, "", s)

	req, _ := http.NewRequest("DELETE", "/teams", nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatInt(person.ID, 10)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeletePerson)
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
	router := mux.NewRouter()
	s := personsSrv.New(client, repo)
	h := New(router, "", s)

	reqBody := &models.PersonForm{
		TeamId:    person.TeamId,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Slack:     person.Slack,
		Role:      person.Role,
		IsActive:  person.IsActive,
	}
	reqByte, _ := json.Marshal(reqBody)
	body := bytes.NewReader(reqByte)

	req, _ := http.NewRequest("PUT", "/teams", body)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.FormatInt(person.ID, 10)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdatePerson)
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
	router := mux.NewRouter()
	s := personsSrv.New(client, repo)
	h := New(router, "", s)

	req, _ := http.NewRequest("GET", "/teams", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetPersons)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	view := []*models.PersonView{
		{
			ID:        person.ID,
			TeamId:    person.TeamId,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Slack:     person.Slack,
			Role:      person.Role,
			IsActive:  person.IsActive,
			Created:   person.Created,
			Updated:   person.Updated,
		},
	}

	byteView, _ := json.Marshal(view)
	expected := string(byteView)

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func newClientMock() *persons.Client {
	return persons.NewMock(
		person.ID,
		&models.PersonForm{
			TeamId:    person.TeamId,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			Slack:     person.Slack,
			Role:      person.Role,
			IsActive:  person.IsActive,
		},
		person.Created,
		person.Updated)
}

func newRepoMock() *personsRepo.RepositoryMock {
	return personsRepo.NewMock(
		person.ID,
		person.TeamId,
		person.FirstName,
		person.LastName,
		person.Slack,
		person.Role,
		person.Created,
		person.Updated)
}
