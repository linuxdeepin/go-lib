// Package introspect provides some utilities for dealing with the DBus
// introspection format.
package introspect

import "encoding/xml"

// Node is the root element of an introspection.
type NodeInfo struct {
	XMLName    xml.Name        `xml:"node"`
	Name       string          `xml:"name,attr,omitempty"`
	Interfaces []InterfaceInfo `xml:"interface"`
	Children   []NodeInfo      `xml:"node,omitempty"`
}

// Interface describes a DBus interface that is available on the message bus.
type InterfaceInfo struct {
	Name        string           `xml:"name,attr"`
	Methods     []MethodInfo     `xml:"method"`
	Signals     []SignalInfo     `xml:"signal"`
	Properties  []PropertyInfo   `xml:"property"`
	Annotations []AnnotationInfo `xml:"annotation"`
}

// Method describes a Method on an Interface as retured by an introspection.
type MethodInfo struct {
	Name        string           `xml:"name,attr"`
	Args        []ArgInfo        `xml:"arg"`
	Annotations []AnnotationInfo `xml:"annotation"`
}

// Signal describes a Signal emitted on an Interface.
type SignalInfo struct {
	Name        string           `xml:"name,attr"`
	Args        []ArgInfo        `xml:"arg"`
	Annotations []AnnotationInfo `xml:"annotation"`
}

// Property describes a property of an Interface.
type PropertyInfo struct {
	Name        string           `xml:"name,attr"`
	Type        string           `xml:"type,attr"`
	Access      string           `xml:"access,attr"`
	Annotations []AnnotationInfo `xml:"annotation"`
}

// Arg represents an argument of a method or a signal.
type ArgInfo struct {
	Name        string           `xml:"name,attr,omitempty"`
	Type        string           `xml:"type,attr"`
	Direction   string           `xml:"direction,attr,omitempty"`
	Annotations []AnnotationInfo `xml:"annotation"`
}

// Annotation is an annotation in the introspection format.
type AnnotationInfo struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
