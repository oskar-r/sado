package usecase

import (
	"log"
	"my-archive/backend"
	"time"

	"github.com/casbin/casbin"
)

//Service service interface
type usecase struct {
	repo     backend.Repository
	tz       *time.Location
	enforcer *casbin.Enforcer
}

//NewIdentifyUseCase create new idnetify use case object
func NewUseCase(r backend.Repository, timezone string, e *casbin.Enforcer) backend.UseCase {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal("Error loading TZ: ", err.Error())
	}

	return &usecase{
		repo:     r,
		tz:       loc,
		enforcer: e,
	}
}
