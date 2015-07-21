package operations

import (
	"errors"
	"fmt"
	"pkg.deepin.io/lib/gio-2.0"
	"reflect"
	"regexp"
	"strings"
)

const (
	ErrorInvalidSignalName = iota
	ErrorSignalNotFound
	ErrorNoMonitor
)

type SignalError struct {
	Code       int
	SignalName string
}

func (e SignalError) Error() string {
	switch e.Code {
	case ErrorInvalidSignalName:
		return fmt.Sprintf("Invalid signal name %q", e.SignalName)
	case ErrorSignalNotFound:
		return fmt.Sprintf("No such a signal %q", e.SignalName)
	case ErrorNoMonitor:
		return fmt.Sprintf("No such a monitor %q", e.SignalName)

	}

	panic("wrong SignalError Code")
}

var (
	_NameFirstCharRegexp   = regexp.MustCompile("^[a-z\\d]")
	_NameCharInvalidRegexp = regexp.MustCompile("[^\\w]|_")
	_NameReduceDashRegexp  = regexp.MustCompile("-{2,}")
)

type SignalManager struct {
	reactors    map[string]*SignalReactor
	cancellable *gio.Cancellable
}

func NewSignalManager(cancellable *gio.Cancellable) *SignalManager {
	return &SignalManager{
		reactors:    map[string]*SignalReactor{},
		cancellable: cancellable,
	}
}

// all signal name should be consist of lower character, digits and dash character('-'),
// starts with lower character or digits and uses '-' to separate words.
func normalizeSignalNameInternal(signalName string, replaceInvalidChar bool) string {
	if replaceInvalidChar {
		signalName = _NameCharInvalidRegexp.ReplaceAllString(signalName, "-")
	}

	if !_NameFirstCharRegexp.MatchString(signalName) {
		signalName = "s-" + signalName
	}

	signalName = strings.ToLower(signalName)
	signalName = _NameReduceDashRegexp.ReplaceAllString(signalName, "-")
	return signalName
}

func normalizeSignalName(signalName string) string {
	return normalizeSignalNameInternal(signalName, true)
}

func isValidSignalName(signalName string) bool {
	return normalizeSignalNameInternal(signalName, false) == normalizeSignalName(signalName)
}

func (m *SignalManager) getSignalReactor(signalName string) (*SignalReactor, error) {
	if !isValidSignalName(signalName) {
		return nil, SignalError{Code: ErrorInvalidSignalName, SignalName: signalName}
	}

	signalName = normalizeSignalName(signalName)
	reactor, ok := m.reactors[signalName]
	if !ok {
		return nil, SignalError{Code: ErrorNoMonitor, SignalName: signalName}
	}

	return reactor, nil
}

func (m *SignalManager) createMonitor(signalName string) *SignalReactor {
	signalName = normalizeSignalName(signalName)
	m.reactors[signalName] = NewSignalReactor(signalName, m.cancellable)
	return m.reactors[signalName]
}

func (m *SignalManager) ListenSignal(signalName string, fn interface{}) (func(), error) {
	reactor, err := m.getSignalReactor(signalName)
	if err != nil {
		switch err.(SignalError).Code {
		case ErrorNoMonitor:
			reactor = m.createMonitor(signalName)
		default:
			return func() {}, err
		}
	}

	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("function is required for listening signal.")
	}

	return reactor.Add(fn), nil
}

func (m *SignalManager) Emit(signalName string, args ...interface{}) error {
	reactor, err := m.getSignalReactor(signalName)
	if err != nil {
		return err
	}

	return m.emitReactor(reactor, args...)
}

func genArgs(fnType reflect.Type, args []interface{}) []reflect.Value {
	expectNArgs := fnType.NumIn()
	actualNArgs := len(args)
	if expectNArgs != actualNArgs {
		panic(fmt.Sprintf("the function %s expect %d arguments, acctually get %d arguments", fnType, expectNArgs, actualNArgs))
	}

	argsValues := make([]reflect.Value, expectNArgs)
	for i, arg := range args {
		argType := fnType.In(i)

		argValue := reflect.ValueOf(arg)
		if !argValue.IsValid() {
			argValue = reflect.Zero(argType)
		}

		actualType := argValue.Type()
		if argType != actualType && !actualType.Implements(argType) {
			// TODO: change %dth to %dst, %dnd, %drd, %dth.
			panic(fmt.Sprintf("the %dth argument on function %s gets wrong type: expect %v, actually get %v", i, fnType, argType, actualType))
		}
		argsValues[i] = argValue
	}

	return argsValues
}

func (m *SignalManager) emitReactor(reactor *SignalReactor, args ...interface{}) error {
	enumerator := reactor.Enumerator()
	defer enumerator.Close()
	for f := range enumerator.Next() {
		if m.cancellable != nil && m.cancellable.IsCancelled() {
			return errors.New("emit operations is cancelled.")
		}

		fn := reflect.ValueOf(f)
		if fn.Kind() != reflect.Func {
			panic("function is required for emitting signal.")
		}

		argsValues := genArgs(fn.Type(), args)
		fn.Call(argsValues)
	}

	return nil
}
