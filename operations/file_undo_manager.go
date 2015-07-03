package operations

import (
	"sync"
)

type CommandType int32

var _FileUndoManager *FileUndoManager = nil
var _FileUndoManagerCreator sync.Once

// TODO: compact command type
const (
	CommandCopy CommandType = iota
	CommandMove
	CommandRename
	CommandLink
	CommandMkdir
	CommandTrash
	CommandPut
)

type FileUndoManager struct {
	*SignalManager
}

const (
	_UndoManagerSignalAvailable            string = "available"
	_UndoManagerSignalUndoJobFinished             = "undo-job-finished"
	_UndoManagerSignalUndoTextChanged             = "undo-text-changed"
	_UndoManagerSignalJobRecordingStarted         = "job-recording-started"
	_UndoManagerSignalJobRecordingFinished        = "job-recording-finished"
)

func (m *FileUndoManager) emitAvailable(available bool) {
	m.Emit(_UndoManagerSignalAvailable, func(f interface{}, args ...interface{}) {
		fn := f.(func(bool))
		fn(available)
	})
}

func (m *FileUndoManager) ListenAvaiable(fn func(bool)) (func(), error) {
	return m.ListenSignal(_UndoManagerSignalAvailable, fn)
}

func (m *FileUndoManager) emitUndoJobFinished() {
	m.Emit(_UndoManagerSignalUndoJobFinished, func(f interface{}, args ...interface{}) {
		fn := f.(func())
		fn()
	})
}

func (m *FileUndoManager) ListenUndoJobFinished(fn func()) (func(), error) {
	return m.ListenSignal(_UndoManagerSignalUndoJobFinished, fn)
}

func (m *FileUndoManager) emitUndoJobTextChanged(text string) {
	m.Emit(_UndoManagerSignalUndoTextChanged, func(f interface{}, args ...interface{}) {
		fn := f.(func(string))
		fn(text)
	})
}

func (m *FileUndoManager) ListenUndoJobTextChanged(fn func(string)) (func(), error) {
	return m.ListenSignal(_UndoManagerSignalUndoTextChanged, fn)
}

func (m *FileUndoManager) emitJobRecordingStarted(op CommandType) {
	m.Emit(_UndoManagerSignalJobRecordingStarted, func(f interface{}, args ...interface{}) {
		fn := f.(func(CommandType))
		fn(op)
	})
}

func (m *FileUndoManager) ListenJobRecordingStarted(fn func(CommandType)) (func(), error) {
	return m.ListenSignal(_UndoManagerSignalJobRecordingStarted, fn)
}

func (m *FileUndoManager) emitJobRecordingFinished(op CommandType) {
	m.Emit(_UndoManagerSignalJobRecordingFinished, func(f interface{}, args ...interface{}) {
		fn := f.(func(CommandType))
		fn(op)
	})
}

func (m *FileUndoManager) ListenJobRecordingFinished(fn func(CommandType)) (func(), error) {
	return m.ListenSignal(_UndoManagerSignalJobRecordingFinished, fn)
}

func (*FileUndoManager) RecordJob(op CommandType, srcURLs []string, destURL string, job interface{}) {
	switch op {
	case CommandMove, CommandCopy:
		copyMoveJob := job.(*CopyMoveJob)
		copyMoveJob.ListenCopyingMovingDone(func(srcURL string, destURL string) {
		})
	}
}

func (*FileUndoManager) Undo() {
}

func newFileUndoManager() *FileUndoManager {
	undoManager := &FileUndoManager{}
	return undoManager
}

func FileUndoManagerInstance() *FileUndoManager {
	_FileUndoManagerCreator.Do(func() {
		_FileUndoManager = newFileUndoManager()
	})

	return _FileUndoManager
}
