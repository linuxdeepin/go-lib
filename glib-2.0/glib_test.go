package glib

import "fmt"
import "testing"

func TestGSettings(t *testing.T) {
	go StartLoop() //gtk_main()

	s := gio.NewSettings("com.deepin.dde.dock")
	if len(s.ListKeys()) != 4 {
		t.Error("ListKeys Error")
	}
	v := s.GetBoolean("active-mini-mode")
	defer func() {
		s.SetBoolean("active-mini-mode", v)
	}()
	changed_times := 0
	ch := make(chan bool, 2)
	s.Connect("changed::active-mini-mode", func(s *Settings, name string) {
		changed_times++
		ch <- true
	})

	s.SetBoolean("active-mini-mode", true)
	<-ch
	if s.GetBoolean("active-mini-mode") != true {
		t.Error("SetBoolean failed")
	}
	s.SetBoolean("active-mini-mode", false)
	<-ch
	if s.GetBoolean("active-mini-mode") != false {
		t.Error("SetBoolean failed")
	}

	if changed_times != cap(ch) {
		t.Error(fmt.Sprintf("Should only changed %d but changed %d", cap(ch), changed_times))
	}
}

func TestGSettingsWithPath(t *testing.T) {
	go StartLoop()

	s := NewSettingsWithPath("org.compiz.grid",
		"/org/compiz/profiles/deepin/plugins/grid/")

	v := s.GetInt("top-edge-action")
	fmt.Println(v)
	defer func() {
		s.SetInt("top-edge-action", v)
	}()
	changed_times := 0
	ch := make(chan bool, 2)
	s.Connect("changed::top-edge-action", func(s *Settings, name string) {
		changed_times++
		ch <- true
	})

	s.SetInt("top-edge-action", 0)
	<-ch
	if s.GetInt("top-edge-action") != 0 {
		t.Error("SetInt Error")
	}

	s.SetInt("top-edge-action", 1)
	<-ch
	if s.GetInt("top-edge-action") != 1 {
		t.Error("SetInt Error")
	}

	if changed_times != cap(ch) {
		t.Error(fmt.Sprintf("Should only changed %d but changed %d", cap(ch), changed_times))
	}
}

func TestInt(t *testing.T) {
	s := NewSettings("org.gnome.settings-daemon.peripherals.keyboard")
	old_d := s.GetUint("delay")
	old_b := s.GetInt("bell-duration")
	defer func() {
		s.SetUint("delay", old_d)
		s.SetUint("delay", old_b)
	}()
	s.SetUint("delay", 50)
	s.SetInt("bell-duration", 50)
	if s.GetUint("delay") != 50 || s.GetInt("bell-duration") != 50 {
		t.Fail()
	}
}
