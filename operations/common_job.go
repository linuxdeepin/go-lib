package operations

import (
	"container/list"
	"fmt"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/timer"
	"strings"
	"time"
)

const (
	_DesktopMIMEType string = "application/x-desktop"
)

// AmountUnit indicates which unit is used for amount.
// Bytes, Files, Directories
type AmountUnit uint16

// The AmountUnit
const (
	AmountUnitBytes AmountUnit = iota
	AmountUnitFiles
	AmountUnitDirectories
	AmountUnitSumOfFilesAndDirs = AmountUnitBytes // using Bytes as the sum of files and directories when it's useless.
)

// OpKind is the flag for which kind of operations is.
type OpKind int32

// the value of OpKind
const (
	OpKindCopy OpKind = iota
	OpKindMove
	OpKindDelete
	OpKindTrash
	OpKindList
)

const (
	_SpeedTimeoutDuration = time.Second * 5 // borrow it from kio.
)

// CommonJob is the base data field of real job.
type CommonJob struct {
	*SignalManager

	// TODO: remove these two field
	op                   OpKind
	numFilesSineProgress int

	cancellable *gio.Cancellable
	uiDelegate  IUIDelegate

	// reporter timer. in order to get reliable
	// transfer rate, there must be some duration.
	progressReporterTimer *timer.Timer
	speedTimer            *time.Timer

	// the time.Timer.C channel won't be closed and cannot be close manually.
	// using closeSpeedTimer to indicate timer is stopped and break from for loop.
	closeSpeedTimer chan struct{}
	percentage      int64

	progressUnit AmountUnit

	skippedFiles    map[string]*gio.File
	skippedDirs     map[string]*gio.File
	skipAllError    bool
	skipAllConflict bool
	mergeAll        bool
	replaceAll      bool
	deleteAll       bool

	totalAmount     map[AmountUnit]int64
	processedAmount map[AmountUnit]int64
	skippedAmount   map[AmountUnit]int64

	// TODO:
	// 1. using reportTimer to emit signals like copying for better performance.
	//    200ms is good enough according to kio.
	// 2. add Done signal for signal synchronization.
	// reportTimer *time.Timer

	err    error
	result interface{}
}

const (
	_JobSignalProcessedAmount = "processed-amount"
	_JobSignalSpeed           = "speed"
	_JobSignalPercent         = "percent"
	_JobSignalTotalAmount     = "total-amount"
	_JobSignalDone            = "job-done"
)

func (job *CommonJob) setError(err error) {
	job.err = err
}

func (job *CommonJob) GetError() error {
	return job.err
}

func (job *CommonJob) HasError() bool {
	return job.err != nil
}

func (job *CommonJob) emitDone() error {
	return job.Emit(_JobSignalDone, job.err)
}

func (job *CommonJob) ListenDone(fn func(error)) (func(), error) {
	return job.ListenSignal(_JobSignalDone, fn)
}

func (job *CommonJob) setResult(r interface{}) {
	job.result = r
}

func (job *CommonJob) Result() interface{} {
	return job.result
}

// ListenProcessedAmount adds observers to processed-amount signal.
func (job *CommonJob) ListenProcessedAmount(fn func(amount int64, unit AmountUnit)) (func(), error) {
	return job.SignalManager.ListenSignal(_JobSignalProcessedAmount, fn)
}

func (job *CommonJob) emitProcessedAmount(amount int64, unit AmountUnit) {
	err := job.Emit(_JobSignalProcessedAmount, amount, unit)
	if err != nil {
		fmt.Println("emit ProcessedAmount signal failed:", err)
	}
}

func (job *CommonJob) setProcessedAmount(size int64, unit AmountUnit) {
	shouldEmit := size != job.processedAmount[unit]
	job.processedAmount[unit] = size
	if shouldEmit {
		job.emitProcessedAmount(size, unit)
		if unit == job.progressUnit { // the AmountUnitBytes might be used to represent the sum of files and directories.
			// TODO: need processedSize signal?
			// job.emitProgressSize(size)
			job.emitPercent(job.processedAmount[unit], job.totalAmount[unit])
		}
	}
}

func (job *CommonJob) doEmitSpeed(speed uint64) {
	err := job.Emit(_JobSignalSpeed, speed)
	if err != nil {
		fmt.Println("emit Speed signal failed:", err)
	}
}

func (job *CommonJob) emitSpeed(speed uint64) {
	if job.speedTimer == nil {
		job.speedTimer = time.NewTimer(_SpeedTimeoutDuration)
		job.closeSpeedTimer = make(chan struct{})
		go func() {
			for {
				select {
				case <-job.speedTimer.C:
					job.doEmitSpeed(0)
					job.speedTimer.Stop()
				case <-job.closeSpeedTimer:
					break
				}
			}
		}()
	}

	job.doEmitSpeed(speed)
	job.speedTimer.Reset(_SpeedTimeoutDuration)
}

func (job *CommonJob) ListenSpeed(fn func(uint64)) (func(), error) {
	return job.ListenSignal(_JobSignalSpeed, fn)
}

func (job *CommonJob) doEmitPercent(percent int64) {
	err := job.Emit(_JobSignalPercent, percent)
	if err != nil {
		fmt.Println("emit Percent signal failed", err)
	}
}

func (job *CommonJob) emitPercent(processedAmount int64, totalAmount int64) {
	if totalAmount != 0 {
		oldPercentage := job.percentage
		job.percentage = int64(float64(processedAmount) / float64(totalAmount) * 100.0)
		if oldPercentage != job.percentage {
			job.doEmitPercent(job.percentage)
		}
	}
}

func (job *CommonJob) ListenPercent(fn func(int64)) (func(), error) {
	return job.ListenSignal(_JobSignalPercent, fn)
}

func (job *CommonJob) emitTotalAmount(totalAmount int64, unit AmountUnit) {
	err := job.Emit(_JobSignalTotalAmount, totalAmount, unit)
	if err != nil {
		fmt.Println("emit TotalAmount signal failed:", err)
	}
}

func (job *CommonJob) ListenTotalAmount(fn func(int64, AmountUnit)) (func(), error) {
	return job.ListenSignal(_JobSignalTotalAmount, fn)
}

func (job *CommonJob) execute() {
}

// UI returns the uiDelegate.
func (job *CommonJob) UI() IUIDelegate {
	return job.uiDelegate
}

// release resources on job.
func (job *CommonJob) finalize() {
	if job.speedTimer != nil {
		job.speedTimer.Stop()
		close(job.closeSpeedTimer)
	}

	job.cancellable.Unref()
}

// Abort the job, Finalize will not be called.
func (job *CommonJob) Abort() {
	job.cancellable.PushCurrent()
	job.cancellable.Cancel()
	job.cancellable.PopCurrent()
}

func (job *CommonJob) isAborted() bool {
	return job.cancellable.IsCancelled()
}

func (job *CommonJob) shouldSkipFile(file *gio.File) bool {
	_, ok := job.skippedFiles[file.GetUri()]
	return ok
}

func (job *CommonJob) shouldSkipDir(file *gio.File) bool {
	_, ok := job.skippedDirs[file.GetUri()]
	return ok
}

func (job *CommonJob) verifyDestination(dst *gio.File, requiredSize uint64) (destFsID string) {
	destIsSymlink := false

retry:
	queryFlags := gio.FileQueryInfoFlagsNone
	if !destIsSymlink {
		queryFlags = gio.FileQueryInfoFlagsNofollowSymlinks
	}
	fInfo, err := dst.QueryInfo(strings.Join(
		[]string{
			gio.FileAttributeStandardType,
			gio.FileAttributeIdFilesystem,
		}, ","),
		queryFlags, job.cancellable)

	if fInfo == nil {
		gerr := err.(gio.GError)
		errCode := gio.IOErrorEnum(gerr.Code)
		if errCode == gio.IOErrorEnumCancelled {
			return
		}

		primaryText := ""
		secondaryText := ""
		detailText := ""

		if errCode == gio.IOErrorEnumPermissionDenied {
			secondaryText = Tr("You do not have permissions to access the destination folder.")
		} else {
			secondaryText = Tr("There was an error getting information about the destination.")
			detailText = gerr.Message
		}

		response := job.uiDelegate.AskRetry(primaryText, secondaryText, detailText)
		switch response.Code() {
		case ResponseCancel:
			fmt.Println("AskRetry aborted")
			job.Abort()
		case ResponseRetry:
			goto retry
		}

		return
	}

	fileType := fInfo.GetFileType()
	if !destIsSymlink && fileType == gio.FileTypeSymbolicLink {
		destIsSymlink = true
		fInfo.Unref()
		goto retry
	}

	destFsID = fInfo.GetAttributeString(gio.FileAttributeIdFilesystem)
	fInfo.Unref()

	if fileType != gio.FileTypeDirectory {
		// TODO: emit error
		// primaryText = Tr("Error while copying to “%B”.") //, dest
		// secondaryText = Tr("The destination is not a folder.")
		fmt.Println("aa Abort")
		job.Abort()
	}

	if destIsSymlink {
		return
	}

	fsInfo, _ := dst.QueryFilesystemInfo(gio.FileAttributeFilesystemFree+","+gio.FileAttributeFilesystemReadonly, job.cancellable)
	if fsInfo == nil {
		return
	}
	defer fsInfo.Unref()

	if requiredSize > 0 && fsInfo.HasAttribute(gio.FileAttributeFilesystemFree) {
		freeSize := fsInfo.GetAttributeUint64(gio.FileAttributeFilesystemFree)

		if freeSize < requiredSize {
			fmt.Println("freeSize:", freeSize, "requiredSize:", requiredSize)
			sizeDifference := requiredSize - freeSize
			//TODO
			primaryText := Tr("Error while copying to “%B”.") //, dest
			secondaryText := Tr("There is not enough space on the destination. Try to remove files to make space.")

			detailText := fmt.Sprintf(Tr("%S more space is required to copy to the destination."), sizeDifference)

			response := job.uiDelegate.AskRetry(primaryText, secondaryText, detailText)

			switch response.Code() {
			case ResponseCancel:
				fmt.Println("bbb abort")
				job.Abort()
			case ResponseRetry:
				goto retry
			}
		}
	}

	if !job.isAborted() && fsInfo.GetAttributeBoolean(gio.FileAttributeFilesystemReadonly) {
		// TODO: emit error
		// primaryText = Tr("Error while copying to “%B”.") //, dest)
		// secondaryText = Tr("The destination is read-only.")
		fmt.Println("ccc abort")
		job.Abort()
	}

	return
}

func (job *CommonJob) reportCountProgress() {
	// TODO:
	// job.emitCounting(job.processedAmount[AmountUnitFiles] + job.processedAmount[AmountUnitDirectories])
}

func (job *CommonJob) scanSources(files []*gio.File) {
	// TODO:
	// job.reportCountProgress()

	for _, file := range files {
		if job.isAborted() {
			break
		}

		job.scanFile(file)
	}

	// job.reportCountProgress()

	// TODO:
	job.emitTotalAmount(job.totalAmount[AmountUnitBytes], AmountUnitBytes)
	job.emitTotalAmount(job.totalAmount[AmountUnitFiles], AmountUnitFiles)
	job.emitTotalAmount(job.totalAmount[AmountUnitDirectories], AmountUnitDirectories)
}

func (job *CommonJob) scanFile(file *gio.File) {
	dirs := list.New()

retry:
	info, err := file.QueryInfo(
		gio.FileAttributeStandardType+","+gio.FileAttributeStandardSize,
		gio.FileQueryInfoFlagsNofollowSymlinks,
		job.cancellable)

	if err == nil {
		job.countFile(info)

		if info.GetFileType() == gio.FileTypeDirectory {
			dirs.PushBack(file)
		}
		info.Unref()
	} else if job.skipAllError {
		job.skipFile(file)
	} else {
		errCode := gio.IOErrorEnum(err.(gio.GError).Code)
		if errCode != gio.IOErrorEnumCancelled {
			primaryText := job.getScanPrimary()
			secondaryText := ""
			detailText := ""

			if errCode == gio.IOErrorEnumPermissionDenied {
				secondaryText = Tr("permission denied")
			} else {
				secondaryText = Tr("error")
				detailText = err.Error()
			}

			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti|UIFlagsRetry)
			switch response.Code() {
			case ResponseCancel:
			case ResponseSkip:
				if response.ApplyToAll() {
					job.skipAllError = true
				}
				job.skipFile(file)
			case ResponseRetry:
				goto retry
			}
		}
	}

	for {
		if job.isAborted() || dirs.Len() == 0 {
			break
		}
		dir := dirs.Front()
		dirs.Remove(dir)

		job.scanDir(dir.Value.(*gio.File), dirs)
	}
}

func (job *CommonJob) scanDir(dir *gio.File, dirs *list.List) {
	savedInfo := map[AmountUnit]int64{
		AmountUnitBytes:       job.processedAmount[AmountUnitBytes],
		AmountUnitFiles:       job.processedAmount[AmountUnitFiles],
		AmountUnitDirectories: job.processedAmount[AmountUnitDirectories],
		// AmountUnitSumOfFilesAndDirs: job.processedAmount[AmountUnitSumOfFilesAndDirs],
	}
retry:
	enumerator, enumerateErr := dir.EnumerateChildren(
		gio.FileAttributeStandardName+
			","+gio.FileAttributeStandardType+
			","+gio.FileAttributeStandardSize,
		gio.FileQueryInfoFlagsNofollowSymlinks,
		job.cancellable,
	)

	if enumerateErr == nil {
		var err error
		for {
			info, nextErr := enumerator.NextFile(job.cancellable)
			err = nextErr
			if info == nil || err != nil {
				break
			}

			job.countFile(info)

			if info.GetFileType() == gio.FileTypeDirectory {
				subdir := dir.GetChild(info.GetName())

				// prepend to dirs for Deepin First Search
				dirs.PushFront(subdir)
			}

			info.Unref()
		}

		enumerator.Close(job.cancellable)
		enumerator.Unref()

		if err != nil {
			errCode := gio.IOErrorEnum(err.(gio.GError).Code)
			if errCode != gio.IOErrorEnumCancelled {
				primaryText := job.getScanPrimary()
				secondaryText := ""
				detailText := ""

				if errCode == gio.IOErrorEnumPermissionDenied {
					secondaryText = Tr("permission denied")
				} else {
					secondaryText = Tr("error")
					detailText = err.Error()
				}

				response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti)
				switch response.Code() {
				case ResponseCancel:
					fmt.Println("ddd abort")
					job.Abort()
				case ResponseRetry:
					job.processedAmount[AmountUnitBytes] = savedInfo[AmountUnitBytes]
					job.processedAmount[AmountUnitFiles] = savedInfo[AmountUnitFiles]
					job.processedAmount[AmountUnitDirectories] = savedInfo[AmountUnitDirectories]
					// job.processedAmount[AmountUnitSumOfFilesAndDirs] = savedInfo[AmountUnitSumOfFilesAndDirs]
					goto retry
				case ResponseSkip:
					job.skipReadDir(dir)
				}
			}
		}
	} else if job.skipAllError {
		job.skipFile(dir)
	} else {
		errCode := gio.IOErrorEnum(enumerateErr.(gio.GError).Code)
		if errCode != gio.IOErrorEnumCancelled {
			primaryText := job.getScanPrimary()
			secondaryText := ""
			detailText := ""

			if errCode == gio.IOErrorEnumPermissionDenied {
				secondaryText = Tr("permission denied")
			} else {
				secondaryText = Tr("error")
				detailText = enumerateErr.Error()
			}

			response := job.uiDelegate.AskSkip(primaryText, secondaryText, detailText, UIFlagsMulti|UIFlagsRetry)
			switch response.Code() {
			case ResponseCancel:
				fmt.Println("ask skip abort")
				job.Abort()
			case ResponseSkip:
				if response.ApplyToAll() {
					job.skipAllError = true
				}
				job.skipFile(dir)
			case ResponseRetry:
				goto retry
			}
		}
	}
}

func (job *CommonJob) countFile(fileInfo *gio.FileInfo) {
	// sourceInfo.NumFiles++
	// sourceInfo.NumBytes += fileInfo.GetSize()
	if fileInfo.GetFileType() == gio.FileTypeDirectory {
		job.totalAmount[AmountUnitDirectories]++
	} else {
		job.totalAmount[AmountUnitFiles]++
	}
	job.totalAmount[AmountUnitBytes] += fileInfo.GetSize()

	// if job.numFilesSineProgress > 100 {
	// 	// TODO: is count signal needed?
	// 	// job.reportCountProgress()
	// 	job.numFilesSineProgress = 0
	// } else {
	// 	job.numFilesSineProgress++
	// }
}

func (job *CommonJob) getScanPrimary() string {
	switch job.op {
	default:
		fallthrough
	case OpKindCopy:
		return Tr("Error while copying.")
	case OpKindMove:
		return Tr("Error while moving.")
	case OpKindDelete:
		return Tr("Error while deleting.")
	case OpKindTrash:
		return Tr("Error while moving files to trash.")
	case OpKindList:
		return Tr("Error while listing files.")
	}
}

func (job *CommonJob) skipFile(file *gio.File) {
	job.skippedFiles[file.GetUri()] = file
	job.skippedAmount[AmountUnitFiles]++
}

func (job *CommonJob) skipReadDir(dir *gio.File) {
	job.skippedDirs[dir.GetUri()] = dir
	job.skippedAmount[AmountUnitDirectories]++
}

func newCommon(uiDelegate IUIDelegate) *CommonJob {
	if uiDelegate == nil {
		uiDelegate = _deafultUIDelegate
	}

	cancellable := gio.NewCancellable()
	return &CommonJob{
		SignalManager:         NewSignalManager(cancellable),
		cancellable:           cancellable,
		uiDelegate:            uiDelegate,
		skippedDirs:           map[string]*gio.File{},
		skippedFiles:          map[string]*gio.File{},
		progressReporterTimer: timer.NewTimer(),
		progressUnit:          AmountUnitBytes,
		totalAmount: map[AmountUnit]int64{
			AmountUnitBytes:       0,
			AmountUnitFiles:       0,
			AmountUnitDirectories: 0,
			// AmountUnitSumOfFilesAndDirs: 0,
		},
		processedAmount: map[AmountUnit]int64{
			AmountUnitBytes:       0,
			AmountUnitFiles:       0,
			AmountUnitDirectories: 0,
			// AmountUnitSumOfFilesAndDirs: 0,
		},
		skippedAmount: map[AmountUnit]int64{
			AmountUnitBytes:       0,
			AmountUnitFiles:       0,
			AmountUnitDirectories: 0,
			// AmountUnitSumOfFilesAndDirs: 0,
		},
	}
}
