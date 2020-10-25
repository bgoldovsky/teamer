package scheduler

import (
	"testing"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
)

var persons = []models.Person{
	{ID: 1, TeamID: 1, FirstName: "FirstA", LastName: "LastA", Slack: "QWERTY1", DutyOrder: 15, IsActive: true},
	{ID: 2, TeamID: 1, FirstName: "FirstB", LastName: "LastB", Slack: "QWERTY2", DutyOrder: 10, IsActive: true},
	{ID: 3, TeamID: 1, FirstName: "FirstC", LastName: "LastC", Slack: "QWERTY3", DutyOrder: 20, IsActive: true},
	{ID: 4, TeamID: 1, FirstName: "FirstD", LastName: "LastD", Slack: "QWERTY4", DutyOrder: 5, IsActive: true},
	{ID: 5, TeamID: 1, FirstName: "FirstE", LastName: "LastE", Slack: "QWERTY5", DutyOrder: 14, IsActive: true},
}

var dataTestGetCurrentIdx = []struct {
	CurrentOrder int64
	ExpectedIdx  int64
}{
	// current duty
	{15, 0},
	{10, 1},
	{20, 2},
	{5, 3},
	{14, 4},

	// next duty
	{-100, 3},
	{0, 3},
	{6, 1},
	{11, 4},
	{16, 2},

	// first duty
	{21, 3},
	{100, 3},
	{100000, 3},
}

func TestGetCurrentIdx(t *testing.T) {
	tempPersons := getTempPersons(persons)
	for _, val := range dataTestGetCurrentIdx {
		if act := getCurrentIdx(val.CurrentOrder, tempPersons); act != val.ExpectedIdx {
			t.Errorf("expected %d but found %d", val.ExpectedIdx, act)
		}
	}
}

func TestGetCurrentIdxError(t *testing.T) {
	if act := getCurrentIdx(10, []*models.Person{}); act != -1 {
		t.Errorf("expected %d but found %d", -1, act)
	}
}

var dataTestScheduleCount = []struct {
	CurrentOrder int64
	DaysCount    int64
}{
	{15, 0},
	{10, 5},
	{20, 15},
}

func TestScheduleCount(t *testing.T) {
	for _, val := range dataTestScheduleCount {
		if act, err := Schedule(val.CurrentOrder, persons, time.Now(), val.DaysCount); int64(len(act)) != val.DaysCount {
			t.Errorf("expected %d but found %d (err: %v)", val.DaysCount, len(act), err)
		}
	}
}

func TestScheduleDetails(t *testing.T) {
	date := time.Date(2020, time.March, 23, 0, 0, 0, 0, time.UTC)
	expected := []models.Duty{
		{
			TeamID:    persons[1].TeamID,
			PersonID:  persons[1].ID,
			FirstName: persons[1].FirstName,
			LastName:  persons[1].LastName,
			Slack:     persons[1].Slack,
			Order:     persons[1].DutyOrder,
			Month:     time.March,
			Day:       23,
		},
		{
			TeamID:    persons[4].TeamID,
			PersonID:  persons[4].ID,
			FirstName: persons[4].FirstName,
			LastName:  persons[4].LastName,
			Slack:     persons[4].Slack,
			Order:     persons[4].DutyOrder,
			Month:     time.March,
			Day:       24,
		},
		{
			TeamID:    persons[0].TeamID,
			PersonID:  persons[0].ID,
			FirstName: persons[0].FirstName,
			LastName:  persons[0].LastName,
			Slack:     persons[0].Slack,
			Order:     persons[0].DutyOrder,
			Month:     time.March,
			Day:       25,
		},
		{
			TeamID:    persons[2].TeamID,
			PersonID:  persons[2].ID,
			FirstName: persons[2].FirstName,
			LastName:  persons[2].LastName,
			Slack:     persons[2].Slack,
			Order:     persons[2].DutyOrder,
			Month:     time.March,
			Day:       26,
		},
		{
			TeamID:    persons[3].TeamID,
			PersonID:  persons[3].ID,
			FirstName: persons[3].FirstName,
			LastName:  persons[3].LastName,
			Slack:     persons[3].Slack,
			Order:     persons[3].DutyOrder,
			Month:     time.March,
			Day:       27,
		},
		{
			TeamID:    persons[1].TeamID,
			PersonID:  persons[1].ID,
			FirstName: persons[1].FirstName,
			LastName:  persons[1].LastName,
			Slack:     persons[1].Slack,
			Order:     persons[1].DutyOrder,
			Month:     time.March,
			Day:       30,
		},
		{
			TeamID:    persons[4].TeamID,
			PersonID:  persons[4].ID,
			FirstName: persons[4].FirstName,
			LastName:  persons[4].LastName,
			Slack:     persons[4].Slack,
			Order:     persons[4].DutyOrder,
			Month:     time.March,
			Day:       31,
		},
		{
			TeamID:    persons[0].TeamID,
			PersonID:  persons[0].ID,
			FirstName: persons[0].FirstName,
			LastName:  persons[0].LastName,
			Slack:     persons[0].Slack,
			Order:     persons[0].DutyOrder,
			Month:     time.April,
			Day:       1,
		},
		{
			TeamID:    persons[2].TeamID,
			PersonID:  persons[2].ID,
			FirstName: persons[2].FirstName,
			LastName:  persons[2].LastName,
			Slack:     persons[2].Slack,
			Order:     persons[2].DutyOrder,
			Month:     time.April,
			Day:       2,
		},
		{
			TeamID:    persons[3].TeamID,
			PersonID:  persons[3].ID,
			FirstName: persons[3].FirstName,
			LastName:  persons[3].LastName,
			Slack:     persons[3].Slack,
			Order:     persons[3].DutyOrder,
			Month:     time.April,
			Day:       3,
		},
	}

	act, err := Schedule(7, persons, date, 10)
	if err != nil {
		t.Fatal(err)
	}

	for idx := range act {
		if act[idx] != expected[idx] {
			t.Errorf("expected %v but found %v", expected[idx], act[idx])
		}
	}
}

var dataTestCurrent = []struct {
	CurrentOrder int64
	ID           int64
}{
	// current duty
	{15, 1},
	{10, 2},
	{20, 3},
	{5, 4},
	{14, 5},

	// next duty
	{-100, 4},
	{0, 4},
	{6, 2},
	{11, 5},
	{16, 3},

	// first duty
	{21, 4},
	{100, 4},
	{100000, 4},
}

func TestCurrent(t *testing.T) {
	date := time.Date(2020, time.March, 23, 0, 0, 0, 0, time.UTC)

	for _, expected := range dataTestCurrent {
		act, err := Current(expected.CurrentOrder, persons, date)
		if err != nil {
			t.Fatal(err)
		}

		if act.PersonID != expected.ID {
			t.Errorf("expected %d but found %d", expected.ID, act.PersonID)
		}
	}
}
