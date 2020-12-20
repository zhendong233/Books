package mockableuuid

import (
	"sync"
	"testing"

	"github.com/google/uuid"
)

type UUID struct{}

var (
	backend string
	mutex   sync.RWMutex
)

func New() UUID {
	return UUID{}
}

func (u UUID) String() string {
	if b := getBackend(); b != "" {
		return b
	}
	return uuid.New().String()
}

func getBackend() string {
	mutex.Lock()
	defer mutex.Unlock()
	return backend
}

func Mock(t *testing.T, fake string) func() {
	t.Helper()
	mutex.Lock()
	defer mutex.Unlock()
	origin := backend
	backend = fake
	return func() { Mock(t, origin) }
}
