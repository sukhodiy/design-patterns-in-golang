// This module demonstrates the right approach that follows the ISP.
package main

// Small interface
type Printer interface {
	Print(d Document)
}

// Another small interface
type Scanner interface {
	Scan(d Document)
}

// Implementation of Printer interface
type MyPrinter struct {
}

// Print
func (m *MyPrinter) Print(d Document) {
}

// Implementation of Printer and Scanner interfaces
type Photocopier struct {
}

// Print
func (p *Photocopier) Print(d Document) {
}

// Scan
func (p *Photocopier) Scan(d Document) {
}

// Aggregation of interfaces
type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// Decorator pattern
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

// Print
func (m *MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

// Scan
func (m *MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
