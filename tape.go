
package main

import (
	"fmt"
)

/* layout allows for a mapping/understanding of the current landscape */
type layout struct {
	cells 	[LAYOUT_LEN_DEFAULT]uint8
	pointer	uint32
	loops	[]uint32
}

/* layout default initializer */
func NewLayout() *layout {
	return &layout{
		cells:		[LAYOUT_LEN_DEFAULT]uint8{},
		pointer: 	0,
		loops:		[]uint32{},
	}
}

/* unsafe inner cell increment */
func (l *layout) plusplus() {
	l.cells[l.pointer]++
}

/* unsafe inner cell decrement */
func (l *layout) minusminus() {
	l.cells[l.pointer]--
}

/* unsafe inner cell set value */
func (l *layout) set(newCellVal uint8) {
	l.cells[l.pointer] = newCellVal
}

/* return current cell value */
func (l *layout) get() uint8 {
	return l.cells[l.pointer]
}

func (l *layout) pop() {
	length := len(l.loops)
	l.loops = l.loops[:length-1]
}

func (l *layout) peak() uint32 {
	length := len(l.loops)
	return l.loops[length-1]
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
	l.plusplus()
	return nil
}

func (l *layout) decrement() error {
	if l.cells[l.pointer] <= 0 {
		return fmt.Errorf("unable to decrement cell @ i=%d, cell_v=%d", l.pointer, l.cells[l.pointer])
	}
	l.minusminus()
	return nil
}

func (l *layout) inputByte() error {
	return fmt.Errorf("not yet implemented")
}

func (l *layout) outputByte() error {
	return fmt.Errorf("not yet implemented")
}

func (l *layout) openLoop() error {
	l.loops = append(l.loops, l.pointer)

	return nil
	//return fmt.Errorf("not yet implemented")
}

func (l *layout) closeLoop() error {
	if l.get() == 0 {
		l.pop()
		return nil
	}

	l.pointer = l.peak()

	return nil
	//return fmt.Errorf("not yet implemented")
}

