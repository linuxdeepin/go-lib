// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package notify

const (
	dbusDest = "org.freedesktop.Notifications"
	dbusPath = "/org/freedesktop/Notifications"
)

const (
	ExpiresDefault     = -1
	ExpiresNever       = 0
	ExpiresMillisecond = 1
	ExpiresSecond      = 1000
)

// hints
const (
	// When set, a server that has the "action-icons" capability will attempt to interpret
	// any action identifier as a named icon. The localized display name will be used to
	// annotate the icon for accessibility purposes.
	// The icon name should be compliant with the Freedesktop.org Icon Naming Specification.
	// type: bool
	HintActionIcons = "action-icons"

	// The type of notification this is.
	// type: string
	HintCategory = "category"

	// This specifies the name of the desktop filename representing the calling program.
	// This should be the same as the prefix used for the application's .desktop file.
	// An example would be "rhythmbox" from "rhythmbox.desktop".
	// This can be used by the daemon to retrieve the correct icon for the application,
	// for logging purposes, etc.
	// type: string
	HintDesktopEntry = "desktop-entry"

	// 	This is a raw data image format which describes the width, height, rowstride,
	// has alpha, bits per sample, channels and image data respectively.
	// type: (iiibiiay)
	HintImageData = "image-data"

	// Alternative way to define the notification image.
	// type: string
	HintImagePath = "image-path"

	// When set the server will not automatically remove the notification when an action
	// has been invoked. The notification will remain resident in the server until it is
	// explicitly removed by the user or by the sender. This hint is likely only useful
	// when the server has the "persistence" capability.
	// type: bool
	HintResident = "resident"

	// The path to a sound file to play when the notification pops up.
	// type: string
	HintSoundFile = "sound-file"

	// A themeable named sound from the freedesktop.org sound naming specification to play
	// when the notification pops up. Similar to icon-name, only for sounds.
	// An example would be "message-new-instant".
	// type: string
	HintSoundName = "sound-name"

	// Causes the server to suppress playing any sounds, if it has that ability.
	// This is usually set when the client itself is going to play its own sound.
	// type: bool
	HintSuppressSound = "suppress-sound"

	// When set the server will treat the notification as transient and by-pass the server's
	// persistence capability, if it should exist.
	// type: bool
	HintTransient = "transient"

	// Specifies the X location on the screen that the notification should point to.
	// The "y" hint must also be specified.
	// type: int32
	HintX = "x"

	// Specifies the Y location on the screen that the notification should point to.
	// The "x" hint must also be specified.
	// type: int32
	HintY = "y"

	// The urgency level.
	// type: byte
	HintUrgency = "urgency"
)

type ClosedReason int

const (
	ClosedReasonInvalid               = -1
	ClosedReasonExpired               = 1
	ClosedReasonDismissedByUser       = 2
	ClosedReasonCallCloseNotification = 3
	ClosedReasonUnknown               = 4
)

func (i ClosedReason) String() string {
	switch i {
	case ClosedReasonInvalid:
		return "Invalid reason"
	case ClosedReasonExpired:
		return "The notification expired"
	case ClosedReasonDismissedByUser:
		return "The notification was dismissed by the user"
	case ClosedReasonCallCloseNotification:
		return "The notification was closed by a call to CloseNotification"
	default:
		return "Unknown reason"
	}
}

const (
	UrgencyLow      = 0
	UrgencyNormal   = 1
	UrgencyCritical = 2
)
