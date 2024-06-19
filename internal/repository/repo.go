package repository

import (
	"gomor-interpolator/internal/model"
)

type Repository interface {
	Register(user model.User) error
	AddAction(action model.Action) error
	AddVictim(victim model.Victim) error
	AddLegend(legend model.Legend) error
	AddTag(tag model.Tags) error
}
