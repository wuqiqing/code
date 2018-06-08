// Sample program demonstrating decoupling with interface composition.
package main

import (
	"fmt"
	"io"
)
type  Puller interface {
	Pull() error
}
type Storer interface {
	Store() error
}
type PullStorer interface {
	Puller
	Storer
}


type Xenia struct{}
func (Xenia) Pull() error {
		return nil
}
type Pillar struct{}
func (Pillar) Store() error {
	return nil
}
type System struct {
	Puller
	Storer
}
func Copy(ps PullStorer, batch int) error {
	return nil
}

func main() {
	sys := System{
		Puller: Xenia{},
		Storer: Pillar{},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
