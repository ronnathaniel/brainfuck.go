
package main

import "fmt"

/* tape allows for a mapping/understanding of the current landscape */
type tape struct {
	cells 	[LAYOUT_LEN_DEFAULT]uint8
	loops	[]uint32
	stdout	string
	pointer	uint32
	windup	uint32
}

type operational interface {
	shiftRight() error
	shiftLeft() error
	increment() error
	decrement() error
	inputByte() error
	outputByte() error
	openLoop() error
	closeLoop() error
}

/* layout default initializer */
func NewTape() *tape {
	return &tape{
		cells:		[LAYOUT_LEN_DEFAULT]uint8{},
		loops:		[]uint32{},
		pointer: 	0,
		windup:		0,
	}
}


func (t *tape) do(op func() error) {
	if err := op(); DEBUG && err != nil {
		fmt.Println(err)
	}
	t.advanceWindup()
}

func (t *tape) shiftRight() error {
	if t.pointer >= LAYOUT_LEN_DEFAULT - 1 {
		return fmt.Errorf("inaccessible cell @ i=%d - Shift Right failed", t.pointer + 1)
	}

	//fmt.Println("shifting right")
	t.plusplusPtr()
	return nil
}

func (t *tape) shiftLeft() error {
	if t.pointer <= 0 {
		return fmt.Errorf("innacessible cell @ i=%d minus1 - Shift Left failed", t.pointer)
	}

	//fmt.Println("shifting left")
	t.minusminusPtr()
	return nil
}

func (t *tape) increment() error {
	if t.getCell() >= 255 {
		return fmt.Errorf("unable to increment cell @ i=%d, cell_v=%d", t.pointer, t.getCell())
	}

	//fmt.Println("incrementing")
	t.plusplusCell()
	return nil
}

func (t *tape) decrement() error {
	if t.getCell() <= 0 {
		return fmt.Errorf("unable to decrement cell @ i=%d, cell_v=%d", t.pointer, t.getCell())
	}

	//fmt.Println("decrementing")
	t.minusminusCell()
	return nil
}

func (t *tape) inputByte() error {
	/* TODO - readline one char */
	return fmt.Errorf("not yet implemented")
}

func (t *tape) outputByte() error {
	t.stdout += fmt.Sprintf("%c", t.getCell())
	return nil
}

func (t *tape) openLoop() error {
	t.loops = append(t.loops, t.windup)
	//fmt.Println("new loops:", t.loops)
	return nil
}

func (t *tape) closeLoop() error {
	if t.getCell() == 0 {
		//fmt.Println("can safely leave this loops")
		t.popLoop()
		return nil
	}

	//t.pointer = t.peakLoop()
	//t.resetWindup()
	t.windup = t.peakLoop()

	//fmt.Println("moving windup back to cell indx", t.windup)
	return nil
}

/* unsafe inner cell increment */
func (t *tape) plusplusCell() {
	t.cells[t.pointer]++
}

/* unsafe inner cell increment */
func (t *tape) plusplusPtr() {
	t.pointer++
}

/* unsafe inner cell decrement */
func (t *tape) minusminusCell() {
	t.cells[t.pointer]--
}

/* unsafe inner cell increment */
func (t *tape) minusminusPtr() {
	t.pointer--
}

/* unsafe inner cell set value */
func (t *tape) setCell(newCellVal uint8) {
	t.cells[t.pointer] = newCellVal
}

/* return current cell value */
func (t *tape) getCell() uint8 {
	return t.cells[t.pointer]
}

func (t *tape) popLoop() {
	length := len(t.loops)
	t.loops = t.loops[:length-1]
}

func (t *tape) peakLoop() uint32 {
	length := len(t.loops)
	return t.loops[length-1]
}

func (t *tape) advanceWindup() {
	t.windup++
}

func (t *tape) resetWindup() {
	t.windup = t.pointer
}

func (t *tape) showTape() {
	fmt.Println(t.cells[:10])
}
