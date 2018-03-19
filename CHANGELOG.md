## [Unreleased]

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
