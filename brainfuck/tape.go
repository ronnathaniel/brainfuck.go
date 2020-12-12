
package brainfuck

import "fmt"

/* tape allows for a mapping/understanding of the current landscape */
type tape struct {
	cells 	[TAPE_SIZE_DEFAULT]uint8
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
		cells:		[TAPE_SIZE_DEFAULT]uint8{},
		loops:		[]uint32{},
		pointer: 	0,
		windup:		0,
	}
}


func (t *tape) do(op func() error) {
	err := op()
	Check(err)
	t.advanceWindup()
}

func (t *tape) shiftRight() error {
	if t.pointer >= TAPE_SIZE_DEFAULT - 1 {
		return fmt.Errorf("inaccessible cell @ i=%d - Shift Right failed", t.pointer + 1)
	}

	Log("shifting right")
	t.plusplusPtr()
	return nil
}

func (t *tape) shiftLeft() error {
	if t.pointer <= 0 {
		return fmt.Errorf("innacessible cell @ i=%d minus1 - Shift Left failed", t.pointer)
	}

	Log("shifting left")
	t.minusminusPtr()
	return nil
}

func (t *tape) increment() error {
	if t.getCell() >= 255 {
		return fmt.Errorf("unable to increment cell @ i=%d, cell_v=%d", t.pointer, t.getCell())
	}

	Log("incrementing")
	t.plusplusCell()
	return nil
}

func (t *tape) decrement() error {
	if t.getCell() <= 0 {
		return fmt.Errorf("unable to decrement cell @ i=%d, cell_v=%d", t.pointer, t.getCell())
	}

	Log("decrementing")
	t.minusminusCell()
	return nil
}

func (t *tape) inputByte() error {
	/* TODO - readline one char */
	return fmt.Errorf("not yet implemented")
}

func (t *tape) outputByte() error {
	Log("outputting", t.getCell())
	t.stdout += fmt.Sprintf("%c", t.getCell())
	return nil
}

func (t *tape) openLoop() error {
	t.loops = append(t.loops, t.windup)
	Log("loops:", t.loops)
	return nil
}

func (t *tape) closeLoop() error {
	if t.getCell() == 0 {
		t.popLoop()
		return nil
	}

	t.windup = t.peakLoop()

	Log("moving windup back to cell @ index=", t.windup)
	return nil
}

func (t *tape) noOp() error {
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
