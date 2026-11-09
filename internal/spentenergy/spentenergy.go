// Package spentenergy calculates distance, average speed and calories used.
package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// Distance returns the distance in kilometres.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

// MeanSpeed returns the average speed in kilometres per hour.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories returns the calories spent while running.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры бега")
	}
	calories := weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH
	return calories, nil
}

// WalkingSpentCalories returns the calories spent while walking.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные параметры ходьбы")
	}
	calories := weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH
	return calories * walkingCaloriesCoefficient, nil
}
