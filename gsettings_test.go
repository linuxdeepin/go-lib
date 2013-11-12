package dlib

import "fmt"
import "testing"

func TestGSettings(t *testing.T) {
	go StartLoop() //gtk_main()

	s := NewSettings("com.deepin.dde.dock")
	if len(s.ListKeys()) != 4 {
		t.Error("ListKeys Error")
	}
	v := s.GetBoolean("active-mini-mode")
	defer func() {
		s.SetBoolean("active-mini-mode", v)
	}()
	changed_times := 0
	s.Connect("changed::active-mini-mode", func(s *Settings, name string) {
		changed_times++
	})

	s.SetBoolean("active-mini-mode", true)
	if s.GetBoolean("active-mini-mode") != true {
		t.Error("SetBoolean failed")
	}
	s.SetBoolean("active-mini-mode", false)
	if s.GetBoolean("active-mini-mode") != false {
		t.Error("SetBoolean failed")
	}

	if changed_times != 2 {
		fmt.Println("changed_times", changed_times)
		t.Error("ChangedError")
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
	s.Connect("changed::top-edge-action", func(s *Settings, name string) {
		changed_times++
	})

	s.SetInt("top-edge-action", v)
	if s.GetInt("top-edge-action") != v {
		t.Error("SetInt Error")
	}

	s.SetInt("top-edge-action", v)
	if s.GetInt("top-edge-action") != v {
		t.Error("SetInt Error")
	}

	if changed_times != 2 {
		fmt.Println("changed_times", changed_times)
		t.Error("ChangedError")
	}
}
