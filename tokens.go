
package main

type shr byte
type shl byte
type incr byte
type decr byte
type input byte
type output byte
type open byte
type close byte

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
