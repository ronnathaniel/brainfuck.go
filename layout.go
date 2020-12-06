
package main

import "fmt"

type layout struct {
	cells 	[LAYOUT_LEN_DEFAULT]uint8
	pointer	uint32
	loops	uint8
}

func NewLayout() *layout {
	return &layout{
		cells:		[LAYOUT_LEN_DEFAULT]uint8{},
		pointer: 	0,
		loops:		0,
	}
}

func (l *layout) shiftRight() error {
	if l.pointer >= LAYOUT_LEN_DEFAULT - 1 {
		return fmt.Errorf("inaccessible cell @ i=%d - Shift Right failed", l.pointer + 1)
	}

	l.pointer++
	return nil
}

func (l *layout) shiftLeft() error {
	if l.pointer <= 0 {
		return fmt.Errorf("innacessible cell @ i=%d - Shift Left failed", l.pointer - 1)
	}

	l.pointer--
	return nil
}

func (l *layout) increment() error {
	if l.cells[l.pointer] >= 255 {
		return fmt.Errorf("unable to increment cell @ i=%d, cell_v=%d", l.pointer, l.cells[l.pointer])
	}
	l.cells[l.pointer]++
	return nil
}

func (l *layout) decrement() error {
	if l.cells[l.pointer] <= 0 {
		return fmt.Errorf("unable to decrement cell @ i=%d, cell_v=%d", l.pointer, l.cells[l.pointer])
	}
	l.cells[l.pointer]--
	return nil
}

func (l *layout) inputByte() error {
	return fmt.Errorf("not yet implemented")
}

func (l *layout) outputByte() error {
	return fmt.Errorf("not yet implemented")
}

func (l *layout) openLoop() error {
	l.loops++
	return fmt.Errorf("not yet implemented")
}

func (l *layout) closeLoop() error {
	return fmt.Errorf("not yet implemented")
}

