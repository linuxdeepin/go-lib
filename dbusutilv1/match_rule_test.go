package dbusutil

import (
	"testing"
)

func TestMatchRuleBuilder(t *testing.T) {
	ruleStr := NewMatchRuleBuilder().
		Type("t1").
		Path("p1").
		Sender("s1").
		Interface("i1").
		Member("m1").
		PathNamespace("pns1").
		Destination("d1").
		Eavesdrop(false).
		Arg(1, "a1").
		ArgPath(2, "a2path").
		ArgNamespace(3, "a3ns").BuildStr()

	ruleExpected := "type='t1',path='p1',sender='s1',interface='i1',member='m1',path_namespace='pns1',destination='d1',eavesdrop='false',arg1='a1',arg2path='a2path',arg3namespace='a3ns'"
	if ruleStr != ruleExpected {
		t.Errorf("ruleStr expected %q got %q", ruleExpected, ruleStr)
	}

	ruleStr = NewMatchRuleBuilder().ExtSignal("/a/b/c", "a.b.c", "Sig1").BuildStr()

	ruleExpected = "type='signal',path='/a/b/c',interface='a.b.c',member='Sig1'"
	if ruleStr != ruleExpected {
		t.Errorf("ruleStr expected %q got %q", ruleExpected, ruleStr)
	}

	ruleStr = NewMatchRuleBuilder().ExtPropertiesChanged("/a/b/c",
		"a.b.c").BuildStr()

	ruleExpected = "type='signal',path='/a/b/c',interface='org.freedesktop.DBus.Properties',member='PropertiesChanged',arg0namespace='a.b.c'"
	if ruleStr != ruleExpected {
		t.Errorf("ruleStr expected %q got %q", ruleExpected, ruleStr)
	}
}
