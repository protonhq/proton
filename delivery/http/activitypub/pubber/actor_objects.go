package pubber

import (
	"sync"

	"github.com/go-fed/activity/vocab"
)

type personActorType struct {
	actor   vocab.PersonType
	actorMu sync.RWMutex
}

type groupActorType struct {
	actor   vocab.GroupType
	actorMu sync.RWMutex
}

type applicationActorType struct {
	actor   vocab.ApplicationType
	actorMu sync.RWMutex
}

var (
	personActors map[string]*personActorType
	groupActors  map[string]*groupActorType

	applicationActor *applicationActorType
)
