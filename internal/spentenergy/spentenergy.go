package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps is not valid")
	}
	if weight <= 0 {
		return 0, errors.New("weight is not valid")
	}
	if height <= 0 {
		return 0, errors.New("height is not valid")
	}
	if duration <= 0 {
		return 0, errors.New("duration is not valid")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps is not valid")
	}
	if weight <= 0 {
		return 0, errors.New("weight is not valid")
	}
	if height <= 0 {
		return 0, errors.New("height is not valid")
	}
	if duration <= 0 {
		return 0, errors.New("duration is not valid")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	if steps <= 0 {
		return 0
	}

	distanceKm := Distance(steps, height)
	speed := distanceKm / duration.Hours()

	return speed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanсeM := stepLength * float64(steps)
	distanceKm := distanсeM / mInKm
	return distanceKm
}
