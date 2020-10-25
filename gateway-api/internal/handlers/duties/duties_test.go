package duties

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/bgoldovsky/dutyer/gateway-api/internal/clients/duties"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/models"
	dutiesSrv "github.com/bgoldovsky/dutyer/gateway-api/internal/services/duties"
	"github.com/gorilla/mux"
)

var (
	duty = &models.DutyView{
		TeamId:    888,
		PersonId:  777,
		FirstName: "Boris",
		LastName:  "B",
		Slack:     "QWERTY",
		Order:     5,
	}
	secondPersonID int64 = 999
	count          int64 = 10
)

func TestHandlers_Swap(t *testing.T) {
	client := newClientMock()
	router := mux.NewRouter()
	s := dutiesSrv.New(client)
	h := New(router, "", s)

	reqBody := &models.SwapForm{
		TeamId:         duty.TeamId,
		FirstPersonID:  duty.PersonId,
		SecondPersonID: secondPersonID,
	}
	reqByte, _ := json.Marshal(reqBody)
	body := bytes.NewReader(reqByte)

	req, _ := http.NewRequest("POST", "/duties/swap", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Swap)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":888,"message":"successfully swapped"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlers_Assign(t *testing.T) {
	client := newClientMock()
	router := mux.NewRouter()
	s := dutiesSrv.New(client)
	h := New(router, "", s)

	reqBody := &models.AssignForm{
		TeamId:   duty.TeamId,
		PersonID: duty.PersonId,
	}
	reqByte, _ := json.Marshal(reqBody)
	body := bytes.NewReader(reqByte)

	req, _ := http.NewRequest("POST", "/duties/assign", body)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Assign)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id":888,"message":"successfully assigned"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlers_GetCurrentDuty(t *testing.T) {
	client := newClientMock()
	router := mux.NewRouter()
	s := dutiesSrv.New(client)
	h := New(router, "", s)

	req, _ := http.NewRequest("GET", "/duties", nil)
	req = mux.SetURLVars(req, map[string]string{"teamID": strconv.FormatInt(duty.TeamId, 10)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetCurrentDuty)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"teamId":888,"id":777,"firstName":"Boris","lastName":"B","slack":"QWERTY","order":5,"month":0,"day":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandlers_GetDuties(t *testing.T) {
	client := newClientMock()
	router := mux.NewRouter()
	s := dutiesSrv.New(client)
	h := New(router, "", s)

	req, _ := http.NewRequest("GET", "/duties", nil)
	q := req.URL.Query()
	q.Add("team-id", strconv.FormatInt(duty.TeamId, 10))
	q.Add("count", strconv.FormatInt(count, 10))
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetDuties)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	view := []*models.DutyView{
		{
			TeamId:    duty.TeamId,
			PersonId:  duty.PersonId,
			FirstName: duty.FirstName,
			LastName:  duty.LastName,
			Slack:     duty.Slack,
			Order:     duty.Order,
		},
	}

	byteView, _ := json.Marshal(view)
	expected := string(byteView)

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func newClientMock() *duties.Client {
	return duties.NewMock(
		duty.TeamId,
		duty.PersonId,
		secondPersonID,
		count,
		duty.FirstName,
		duty.LastName,
		duty.Slack,
		duty.Order)
}
