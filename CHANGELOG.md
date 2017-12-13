## [Unreleased]

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
