[1.10.2] 2019-04-08
*   feat(desktopappinfo): do not process X-Deepin-Vendor field

[1.10.1] 2019-04-03
*   fix: build failed
*   chore(desktopappinfo): remove redundant println

[1.10.0] 2019-03-15
*   fix(asound): SelemHasCaptureSwitchJoined wrong
*   change(api): asound add more functions
*   fix(dbusutil): proxy.Object removeHandler not work if ruleAuto is false

[1.9.2] 2019-03-01
*   chore(dbusutil): RemoveHandler check signal ext
*   chore: use github.com/linuxdeepin/go-gir

[1.9.1] 2019-02-25
*   fix(proxy): abnormally clear env vars when modifying config

[1.9.0] 2019-02-22
*   refactor: fix a typo
*   change(api): appinfo AppLaunchContext add method SetCmdSuffixes
*   fix(desktpappinfo): failed to parse Exec field of wine programs desktop files

## [1.8.0] 2019-01-03
*   chore: use dh golang build system on `sw_64`
*   feat(dbusutil): gsprop.Enum add methods GetString and SetString
*   fix(dbus): failed to find session bus address in systemd 240+

## [1.7.0] 2018-12-29
*   chore(`sound_effect`): player add method Finder

## [1.6.0] 2018-12-10
*   fix(dbus1): defaultSignalHandler.DeliverSignal

## [1.5.0] 2018-12-07
*   chore(dbus1): skip test transport nonce tcp
*   chore(dbus1): update to the latest upstream code
*   feat(imgutil): add method CanDecodeConfig
*   feat: add lib imgutil

## [1.4.0] 2018-11-23
*   feat(desktopappinfo): support key X-Deepin-Vendor

## [1.3.0] 2018-10-25
*   fix(dbusutil): gsprop nil pointer panic
*   feat: add lib shell
*   chore(gettext): do not run test
*   fix: make install lost a file
*   chore: add makefile for `sw_64`

## [1.2.16] 2018-08-07
*   feat(sound-effect): check theme and event validity

## [1.2.15] 2018-07-31
*   fix(desktopappinfo): expandFieldCode

## [1.2.14] 2018-07-20
*   chore(proxy): defer proxySettings init

## [1.2.13] 2018-07-19
*   fix: test failed

## [1.2.12] 2018-07-19
*   chore(dbusutil): gsprop add mutex prevent data race
*   chore(dbus1): update to the latest upstream code
*   chore(dbusutil): request name error include name
*   chore: remove polkit
*   chore(gdkpixbuf): use go-x11-client
*   fix: play wav file with pulseaudio get error
*   chore(debian): update debian control
*   chore(appinfo): use go-x11-client
*   feat(audio): improve event handling

## [1.2.11] 2018-06-07
*   fix(pulse): no subscribe server event
*   fix(pulse): Context data race
*   fix(dbus1): exportedObj data race
*   refactor(pulse): Context add some methods

## [1.2.10] 2018-05-23
*   chore: update dbus1

## [1.2.9] - 2018-05-15
*   fix(pulse): event error should handle in  go space
*   perf(pulse): avoid wasting CPU time to poll connect state

## [1.2.8] - 2018-05-14
*   feat(dbusutil): dbusutil-gen add do not edit header
*   fix(pam): cbPAMConv return two values
*   fix(dbus1): test failed
*   fix(dbusutil): failed to set property of object
*   refactor(pulse): ignore pulseaudio's strange behavior
*   fix(pulse): memory leak about pa_operation
*   perf(pulse): remove useless memcpy
*   fix(pulse): unity all mainloop_lock to safeDo
*   fix(pulse): deadlock on pendingCallback
*   fix(pulse): block on ck.Feed
*   refactor(pulse): print info about failed pulse operation
*   fix(pulse): too early free C.String
*   refactor(pulse): erase magic number about PA init State
*   refactor(pulse): move all pa_threaded_mainloop_lock to one file
*   fix(pulse): move all callback out of mainloop thread

## [1.2.7] - 2018-04-17
*   fix(desktopappinfo): incorrect use of os.Chdir to set working directory of cmd
*   fix(sound_effect): fix alsa backend play failed
*   chore(debian): add dependency on libpam0g-dev
*   feat(pam): add lib pam
*   Merge "feat(dbusutil): proxy.Object can choose not to use auto rule"
*   feat(dbusutil): proxy.Object can choose not to use auto rule
*   fix(pulse): protect pa_stream
*   fix(pulse): protect sink_suspend and meter
*   fix(pulse): protect pa_context_unref
*   feat(dbusutil): add support for auto generate code

## [1.2.6] - 2018-03-19
*   perf: unref the `pa_operation` in macro "DEFINE"
*   feat: dbusutil/gsprop add Uint64 type
*   fix(pulse): data race on pulse.Context.free
*   refactor(pulse): hold lock when starting mainloop
*   feat(dbusutil): new api
*   refactor(pulse): remove global variable of `pa_threaded_mainloop`
*   refactor(pulse): hide internal functions

## [1.2.5] - 2018-03-07
*   feat(dbusutil): hide log output
*   feat(dbusutil): allow field type pointer implements the Property interface
*   feat(dbusutil): delay emit property changed
*   doc: add document about dbusutil
*   feat(dbusutil): improves object introspection
*   fix(dbusutil): PropsMaster Begin and End
*   fix(dbusutil): no check impl is nil
*   feat(dbusutil): add PropsMaster
*   fix(dbusutil): method RequestName panic
*   feat(dbusutil) add method NameHasOwenr and GetNameOwner
*   fix(dbusutil/gsprop): use path property to connect changed
*   fix(dbusutil): auto quit
*   fix(dbusutil): ToError panic if param err is nil
*   feat: add lib dbusutil
*   fix(encoding/kv): constants value wrong
*   refactor: reduce global `success_cb` variable
*   fix: protect global variable "sourceMeterCBs"
*   chore: update license
*   feat: lib dbus/property and proxy use new lib gsettings
*   feat: add lib gsettings
*   fix(dbus1): test SystemBus failed in build env
*   add lib dbus1

## [1.2.4] - 2018-01-24
*   fix Adapt lintian
*   asound: add more types and functions
*   `sound_effect`: fix alsa play backend HWParams wrong
*   notify: fix notification Update and Show
*   `sound_effect`: fix can not compile with go 1.4
*   dbus: fix SetAutoDestroyHandler
*   add lib cgroup
*   pulse/simple: fix wrong error handling
*   appinfo: do not call GetStartupNotifyId if timestamp is 0

## [1.2.3] - 2017-12-13
*   add some audio libs
*   fix package golang-dlib-dev depends wrong
*   fix: race condition on HasNewMessage
*   asound: add method SetRateNear and GetDeviceNameHints


## [1.2.2] - 2017-11-28
+ fix concurrent access dbus PropertyProxy
+ add StartCommand method for DestkopAppInfo and DesktopAction


## [1.2.1] - 2017-11-16
+ add field Section for DesktopAction
+ add SetDataDirs in desktopappinfo


## [1.2.0] - 2017-10-12
### Added
+ add pulse init timeout

### Changed
+ update license
+ replace syscall 'statfs' with 'statvfs'
+ make transport endian aware

### Fixed
+ fix dbus introspection map concurrency
