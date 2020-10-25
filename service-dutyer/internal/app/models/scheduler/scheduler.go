package scheduler

import (
	"fmt"
	"math"
	"time"

	"github.com/bgoldovsky/dutyer/service-dutyer/internal/app/models"
)

func Schedule(currentOrder int64, persons []models.Person, date time.Time, daysCount int64) ([]models.Duty, error) {
	maxCount := getMaxCount(persons)
	count := maxCount

	duties := make([]models.Duty, 0, count)
	tempPersons := make([]*models.Person, len(persons))

	for i := int64(0); i < daysCount; i++ {
		if count == maxCount {
			tempPersons = getTempPersons(persons)
			count = 0
		}

		if date.Weekday() == time.Saturday {
			date = date.UTC().Add(time.Hour * 48)
		}

		if date.Weekday() == time.Sunday {
			date = date.UTC().Add(time.Hour * 24)
		}

		currentIdx := getCurrentIdx(currentOrder, tempPersons)
		if currentIdx == -1 {
			return nil, fmt.Errorf("scheduler can't find duty. order: %d", currentOrder)
		}

		duty := models.Duty{
			TeamID:    tempPersons[currentIdx].TeamID,
			PersonID:  tempPersons[currentIdx].ID,
			FirstName: tempPersons[currentIdx].FirstName,
			LastName:  tempPersons[currentIdx].LastName,
			Slack:     tempPersons[currentIdx].Slack,
			Order:     tempPersons[currentIdx].DutyOrder,
			Month:     date.Month(),
			Day:       int64(date.Day()),
		}

		duties = append(duties, duty)
		tempPersons[currentIdx] = nil

		date = date.Add(time.Hour * 24)
		count++
	}

	return duties, nil
}

func Current(currentOrder int64, persons []models.Person, date time.Time) (*models.Duty, error) {
	tempPersons := getTempPersons(persons)
	currentIdx := getCurrentIdx(currentOrder, tempPersons)

	if currentIdx == -1 {
		return nil, fmt.Errorf("scheduler can't find duty. order: %d", currentOrder)
	}

	return &models.Duty{
		TeamID:    tempPersons[currentIdx].TeamID,
		PersonID:  tempPersons[currentIdx].ID,
		FirstName: tempPersons[currentIdx].FirstName,
		LastName:  tempPersons[currentIdx].LastName,
		Slack:     tempPersons[currentIdx].Slack,
		Order:     tempPersons[currentIdx].DutyOrder,
		Month:     date.Month(),
		Day:       int64(date.Day()),
	}, nil
}

func getCurrentIdx(currentOrder int64, persons []*models.Person) int64 {
	var nextIdx, minIdx int64 = -1, -1
	nextOrder, minOrder := int64(math.MaxInt64), int64(math.MaxInt64)

	for idx, person := range persons {
		if person == nil || !person.IsActive {
			continue
		}

		if person.DutyOrder == currentOrder {
			return int64(idx)
		}

		if person.DutyOrder > currentOrder && person.DutyOrder < nextOrder {
			nextOrder = person.DutyOrder
			nextIdx = int64(idx)
		}

		if person.DutyOrder < minOrder {
			minOrder = person.DutyOrder
			minIdx = int64(idx)
		}
	}

	if nextIdx != -1 {
		return nextIdx
	}

	if minIdx != -1 {
		return minIdx
	}

	return -1
}

func getMaxCount(persons []models.Person) int64 {
	var maxCount int64
	for _, val := range persons {
		if val.IsActive {
			maxCount++
		}
	}

	return maxCount
}

func getTempPersons(persons []models.Person) []*models.Person {
	tempPersons := make([]*models.Person, len(persons))
	for idx, person := range persons {
		tempPerson := person
		tempPersons[idx] = &tempPerson
	}

	return tempPersons
}
