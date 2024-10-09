package sdkserver

import (
	"context"
	"sync"
	"time"

	"github.com/gptscript-ai/broadcaster"
	"github.com/sanjay920/gptscript/pkg/engine"
	"github.com/sanjay920/gptscript/pkg/runner"
	gserver "github.com/sanjay920/gptscript/pkg/server"
	"github.com/sanjay920/gptscript/pkg/types"
)

type SessionFactory struct {
	events *broadcaster.Broadcaster[event]
}

func NewSessionFactory(events *broadcaster.Broadcaster[event]) *SessionFactory {
	return &SessionFactory{
		events: events,
	}
}

func (s SessionFactory) Start(ctx context.Context, prg *types.Program, env []string, input string) (runner.Monitor, error) {
	id := gserver.RunIDFromContext(ctx)
	category := engine.ToolCategoryFromContext(ctx)

	if category == engine.NoCategory {
		s.events.C <- event{
			Event: gserver.Event{
				Event: runner.Event{
					Time: time.Now(),
					Type: runner.EventTypeRunStart,
				},
				Input:   input,
				RunID:   id,
				Program: prg,
			},
		}
	}

	return &Session{
		id:     id,
		prj:    prg,
		env:    env,
		events: s.events,
	}, nil
}

func (s SessionFactory) Pause() func() {
	return func() {}
}

type Session struct {
	id      string
	prj     *types.Program
	env     []string
	events  *broadcaster.Broadcaster[event]
	runLock sync.Mutex
}

func (s *Session) Event(e runner.Event) {
	s.runLock.Lock()
	defer s.runLock.Unlock()
	s.events.C <- event{
		Event: gserver.Event{
			Event: e,
			RunID: s.id,
		},
	}
}

func (s *Session) Stop(ctx context.Context, output string, err error) {
	category := engine.ToolCategoryFromContext(ctx)

	if category != engine.NoCategory {
		return
	}

	e := event{
		Event: gserver.Event{
			Event: runner.Event{
				Time: time.Now(),
				Type: runner.EventTypeRunFinish,
			},
			RunID:  s.id,
			Output: output,
		},
	}
	if err != nil {
		e.Err = err.Error()
	}

	s.runLock.Lock()
	defer s.runLock.Unlock()
	s.events.C <- e
}

func (s *Session) Pause() func() {
	s.runLock.Lock()
	return func() {
		s.runLock.Unlock()
	}
}
