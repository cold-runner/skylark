package event

import (
	"context"
	"github.com/cold-runner/skylark/biz/entity/event"
	"github.com/cold-runner/skylark/biz/infrastructure/searchEngine"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
)

func SyncToSearchEngine() {
	storeIns := store.GetIns()
	searchEngineIns := searchEngine.GetIns()
	c := context.Background()

	event.SyncToEngine(c, storeIns, searchEngineIns)
}
