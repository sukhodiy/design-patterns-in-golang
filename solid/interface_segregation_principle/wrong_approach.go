// This module demonstrates the wrong approach that breaks the ISP.
package main

// Main interface
type WrongMachine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// Implementation of main interface
type MultiFunctionPrinter struct {
}

// Print
func (m *MultiFunctionPrinter) Print(d Document) {
}

// Fax
func (m *MultiFunctionPrinter) Fax(d Document) {
}

// Scan
func (m *MultiFunctionPrinter) Scan(d Document) {
}

// Another implementation of main interface
type OldFashionedPriner struct {
}

// Print
func (o *OldFashionedPriner) Print(d Document) {
	// OK
}

// XXX: Deprecated
func (o *OldFashionedPriner) Fax(d Document) {
	panic("operation not supported")
}

// XXX: Deprecated
func (o *OldFashionedPriner) Scan(d Document) {
	panic("operation not supported")
}
