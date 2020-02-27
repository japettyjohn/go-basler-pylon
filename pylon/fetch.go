// Copyright (C) jno, 2019; License: MIT

// As the `RetrieveAndSave` does not work for me.
// So, I follow the code from $PYLON_HOME/Samples/C++/Grab/Grab.cpp - it works.
// Plus, I prefer to handle grabbed image on "application level" and not
// bury it hardcoded somewhere in C++ underpants.
// So, the callback implementation.
package pylon

// #include <string.h>
// #include <stdlib.h>
// #include "capture.h"
import "C"

import "unsafe"
import "fmt"

//export Go_fetch_callback
func Go_fetch_callback(i C.int,				// callback index
		       w C.int, h C.int, pxt C.int,	// image props
		       size C.int, buffer *C.char,	// image data
) C.int {
	W := int(w)
	H := int(h)
	P := int(pxt)
	Z := int(size)
	B := C.GoBytes(unsafe.Pointer(buffer), size)

	if false {
		fmt.Printf("\ncallback(w=%#v, h=%#v, pxt=%#v, size=%#v, buf=%#v)\n",
			   W, H, P, Z, B[0])
	}
	f, cam := cbreg.Get(int(i))
	rc := C.int(f(cam.(*Camera), W, H, P, Z, B))
	return rc
}

func (cam *Camera) SetFetchTimeout(v uint) {
	C.setFetchTimeout(C.uint(v))
}

func (cam *Camera) SetFetchCount(v uint) {
	C.setFetchCount(C.uint(v))
}

/* Sample usage:

import (
	"fmt"
	"time"
	...
)
import "pylon"

func do() {
	cam := &pylon.Camera{}

	if e := cam.AttachDevice(); e != nil { fmt.Fatalf("cam.AttachDevice(): %v", e); }
	if e := cam.OpenCamera(); e != nil { fmt.Fatalf("cam.OpenCamera(): %v", e); }
	defer cam.CloseCamera()
	if e := cam.ConfigureCamera(); e != nil { fmt.Fatalf("cam.ConfigureCamera(): %v", e); }

	var saved []string
	const savePathPrefix = "/a/b/c/d"
	const savePathSuffix = ".jpg"

	FrameCallback := func(w, h, pxt, size int, buffer []byte) int {
		pt := pylon.EPixelType(pxt)
		fmt.Printf("FrameCallback(w=%#v, h=%#v, pt=%08x=%s, size=%#v, buffer=%#v...)\n",
			   w, h, pxt, pt.String(), size, buffer[0])
		t1 := time.Now()

		path := createFilePath(savePathPrefix, t1, savePathSuffix)

		if e := save2jpeg(path, w, h, buffer); e != nil {
			fmt.Fatalf("Cannot save to %#v: %v", path, e)
		}
		saved = append(saved, path)

		fmt.Printf("FrameCallback taken %v\n\n", time.Now().Sub(t1))
		return 0
	}

	cam.SetFetchTimeout(5000) // ms
	cam.SetFetchCount(10)

	cb := pylon.FrameCallbackType(FrameCallback)

	if e := cam.Fetch(cb); e != nil { fmt.Fatalf("Fetch failed: %v", e); }

	if len(saved) > 0 {
		// HANDLE SAVED FILES
	}
}
*/
func (cam *Camera) Fetch(cb func(c *Camera, w, h, pt, size int, buffer []byte) int) error {
	idx := cbreg.Add(cb)
	cbreg.Annotate(idx, cam)
	s := C.GoString(C.fetch(C.int(idx)));
	if s != "" {
		return newError("Fetch: %v", s)
	}
	return nil
}

/* EOF */
