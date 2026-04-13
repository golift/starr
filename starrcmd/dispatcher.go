package starrcmd

import (
	"fmt"
	"sync"

	"golift.io/starr"
)

type hookKey struct {
	app starr.App
	typ Event
}

// Dispatcher registers callbacks for Custom Script invocations. Call Run to parse the
// environment with New and invoke matching handlers, or call Dispatch with a *CmdEvent
// from tests or custom wiring.
type Dispatcher struct {
	mu sync.Mutex
	// hooks maps (app, event) to one or more callbacks; all matching callbacks run in order.
	hooks map[hookKey][]func(*CmdEvent) error
	// OnUnknown is invoked when no handlers are registered for cmd.App and cmd.Type.
	OnUnknown func(*CmdEvent) error
}

// NewDispatcher returns an empty Dispatcher ready for Register or the typed On{App}{Event} helpers.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		hooks: make(map[hookKey][]func(*CmdEvent) error),
	}
}

// Register adds a callback for the given app and event. Multiple registrations for the same
// pair are all run, in registration order, when Dispatch matches. Nil callback is ignored.
// A nil *Dispatcher is a no-op.
func (d *Dispatcher) Register(app starr.App, typ Event, callback func(*CmdEvent) error) {
	if d == nil || callback == nil {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	if d.hooks == nil {
		d.hooks = make(map[hookKey][]func(*CmdEvent) error)
	}

	k := hookKey{app: app, typ: typ}
	d.hooks[k] = append(d.hooks[k], callback)
}

// Run calls New, then Dispatch with the resulting *CmdEvent. It returns ErrNoEventFound
// (or any error from New) before callbacks run.
func (d *Dispatcher) Run() error {
	if d == nil {
		return ErrNilDispatcher
	}

	cmd, err := New()
	if err != nil {
		return err
	}

	return d.Dispatch(cmd)
}

// Dispatch runs all handlers registered for cmd.App and cmd.Type. If none match and
// OnUnknown is set, OnUnknown(cmd) is returned. If none match and OnUnknown is nil, it
// returns nil. The first callback error stops execution and is returned.
// A nil *Dispatcher or nil cmd returns ErrNilDispatcher / ErrNilCmdEvent respectively.
func (d *Dispatcher) Dispatch(cmd *CmdEvent) error {
	if d == nil {
		return ErrNilDispatcher
	}

	if cmd == nil {
		return ErrNilCmdEvent
	}

	d.mu.Lock()
	k := hookKey{app: cmd.App, typ: cmd.Type}

	var fns []func(*CmdEvent) error

	if d.hooks != nil {
		fns = append([]func(*CmdEvent) error(nil), d.hooks[k]...)
	}

	onUnknown := d.OnUnknown
	d.mu.Unlock()

	if len(fns) == 0 {
		if onUnknown != nil {
			return onUnknown(cmd)
		}

		return nil
	}

	for _, callback := range fns {
		if err := callback(cmd); err != nil {
			return err
		}
	}

	return nil
}

// executeGet parses the payload with getter and passes it to handler.
func executeGet[T any](cmd *CmdEvent, getter func(*CmdEvent) (T, error), handler func(T) error) error {
	val, err := getter(cmd)
	if err != nil {
		return fmt.Errorf("parse env payload: %w", err)
	}

	return handler(val)
}
