// Copyright (C) jno, 2019; License: MIT

// Specialized very simple Registry of call-backs to support frame grabbing.
// It follows the idea of https://github.com/golang/go/wiki/cgo
// Perhaps, more generic https://github.com/yamnikov-oleg/cgo-callback is better
// Check for `export`ed `Go_fetch_callback` function.
package pylon

import "sync"

// The type of callback function.
// It recieves the grabbed image with width of `w`, height of `h`,
// and pixel type of `pxt` in `buffer` of `size` bytes.
type FrameCallbackType func(c *Camera, w, h, pxt, size int, buffer []byte) int

// I do not want uncontrolled growth of the Registry. So, impose the limit.
const MAX_REG = 128

// It is expected to be sorta "singleton" object, but one can use more.
type Registry struct {
	slot int
	reg map[int]FrameCallbackType
	ann map[int]interface{}
	lock sync.Mutex
}

// interim things as the Go has no normal constructors
func (r *Registry) init() {
	r.lock.Lock(); defer r.lock.Unlock()
	if r.reg == nil {
		r.reg = make(map[int]FrameCallbackType)
		r.ann = make(map[int]interface{})
		r.slot = 0
	}
}

// Adds a callback to the registry within MAX_REG limit.
// Returns the index of added entry.
// Note: there is no "remove" operation.
func (r *Registry) Add(f FrameCallbackType) int {
	r.init()
	r.lock.Lock(); defer r.lock.Unlock()
	if r.slot >= MAX_REG {
		panic("Registry overflow")
	}
	r.reg[r.slot] = f
	r.slot++
	return r.slot - 1
}

func (r *Registry) Annotate(i int, a interface{}) {
	r.init()
	r.lock.Lock(); defer r.lock.Unlock()
	if r.slot >= MAX_REG {
		panic("Registry overflow")
	}
	if i >= r.slot { panic("Bad indedx"); }
	r.ann[i] = a
}

// Retruns the registered callback for given index, if any.
func (r *Registry) Get(i int) (FrameCallbackType, interface{}) {
	if r.reg == nil || r.slot == 0 {
		panic("The registry is empty")
	}
	if i < 0 || i >= r.slot {
		panic("Bad registry index")
	}
	return r.reg[i], r.ann[i]
}

// static object: global registry
var cbreg = Registry{}

/* EOF */
