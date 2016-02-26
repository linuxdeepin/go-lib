package initializer

import (
	"pkg.deepin.io/lib/dbus"
)

// DoWithSessionBus starts the initialization and install dbus object to session bus.
func (i *Initializer) DoWithSessionBus(fn func() (dbus.DBusObject, error)) *Initializer {
	return i.init(func() error {
		dbusObject, err := fn()
		if err != nil {
			return err
		}

		err = dbus.InstallOnSession(dbusObject)
		if err != nil {
			return err
		}
		return nil
	})
}

// DoWithSystemBus starts the initialization and install dbus object to system bus.
func (i *Initializer) DoWithSystemBus(fn func() (dbus.DBusObject, error)) *Initializer {
	return i.init(func() error {
		dbusObject, err := fn()
		if err != nil {
			return err
		}

		err = dbus.InstallOnSystem(dbusObject)
		if err != nil {
			return err
		}
		return nil
	})
}

// DoWithSessionBus starts the initialization and install dbus object to sesison bus.
func DoWithSessionBus(fn func() (dbus.DBusObject, error)) *Initializer {
	i := new(Initializer)
	return i.DoWithSessionBus(fn)
}

// DoWithSystemBus starts the initialization and install dbus object to system bus.
func DoWithSystemBus(fn func() (dbus.DBusObject, error)) *Initializer {
	i := new(Initializer)
	return i.DoWithSystemBus(fn)
}
