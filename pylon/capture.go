package pylon

// #include <string.h>
// #include <stdlib.h>
// #include "capture.h"
import "C"

import (
	"sync"
	"unsafe"
)

func init() {
	C.pylonInitialize();
}

type CameraInfo struct {
	FullName, VendorName, ModelName, SerialNumber, DeviceVersion string
	ProductId, VendorId, Width, Height int
}

type Camera struct {
	startMutex, attachedMutex, openMutex sync.Mutex
}

func (cam *Camera) Pointer() unsafe.Pointer {
	return unsafe.Pointer(cam)
}

func (cam *Camera) Info() (*CameraInfo, error) {
	if !C.isAttached() {
		return nil, newError("Info: No device attached")
	}
	var i *CameraInfo = new(CameraInfo)
	i.Width = int(C.width())
	i.Height = int(C.height())
	i.FullName = C.GoString(C.fullName())
	i.VendorName = C.GoString(C.vendorName())
	i.ModelName = C.GoString(C.modelName())
	i.SerialNumber = C.GoString(C.serialNumber())
	i.DeviceVersion = C.GoString(C.deviceVersion())
	i.ProductId = int(C.productId())
	i.VendorId = int(C.vendorId())
	return i, nil
}

func (cam *Camera) OpenCamera() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		if e := C.GoString(C.openCamera()); e != "" {
			return newError("OpenCamera: %v", e)
		}
	}
	return nil
}

func (cam *Camera) CloseCamera() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if C.isOpen() {
		if e := C.GoString(C.closeCamera()); e != "" {
			return newError("CloseCamera: %v", e)
		}
	}
	return nil
}

func (cam *Camera) StartCapture(max int) error {
	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if !C.isCameraGrabbing() {
		if e := C.GoString(C.startCapture(C.int(max))); e != "" {
			return newError("StartCapture: %v", e)
		}
	}
	return nil
}

func (cam *Camera) IsGrabbing() bool {
	if C.isCameraGrabbing() {
		return true
	}
	return false
}

func (cam *Camera) StopCapture() error {
	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		if e := C.GoString(C.stopCapture()); e != "" {
			return newError("StopCapture: %v", e)
		}
	}
	return nil
}

func (cam *Camera) AttachDevice() error {
	cam.attachedMutex.Lock()
	defer cam.attachedMutex.Unlock()
	if !C.isAttached() {
		if e := C.GoString(C.attachDevice()); e != "" {
			return newError("AttachDevice: %v", e)
		}
	}
	return nil
}

func (cam *Camera) ConfigureCamera() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return newError("ConfigureCamera: Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return newError("ConfigureCamera: Camera is grabbing.")
	}

	if e := C.GoString(C.configureCamera()); e != "" {
		return newError("ConfigureCamera: %v", e)
	}
	return nil
}

func (cam *Camera) SetHardwareTriggerConfiguration() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return newError("Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return newError("Camera is grabbing.")
	}

	if e := C.GoString(C.setHardwareTriggerConfiguration()); e != "" {
		return newError("SetHardwareTriggerConfiguration: %v", e)
	}
	return nil
}

func (cam *Camera) RetrieveAndSave(batch, timeout int, outputPath string) error {
	cOutputPath := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cOutputPath))
	if e := C.GoString(C.retrieveAndSave(C.int(batch),
					     C.int(timeout),
					     cOutputPath)); e != "" {
		return newError("RetrieveAndSave: %v", e)
	}
	return nil
}

func (cam *Camera) SetParam(p Param, value interface{}) error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return newError("SetParam: Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return newError("SetParam: Camera is grabbing.")
	}

	cName := C.CString(p.Name)
	defer C.free(unsafe.Pointer(cName))

	switch v := value.(type) {
	case string:
		// Multiple types for a string
		switch p.OriginalType {
		case OriginalTypeGenApiIEnumerationT:
			cValue := C.CString(v)
			defer C.free(unsafe.Pointer(cValue))
			C.setNodeMapEnumParam(cName, cValue)
		case OriginalTypeGenApiIString, OriginalTypeGenApiICommand:
			return newError("SetParam(%+q, %v): Original type %s not implemented.",
					  p.Name, value, p.OriginalType)
		default:
			return newError("SetParam(%+q, %v): Unexpected string for type %s",
					  p.Name, value, p.OriginalType)
		}

	case int64:
		if p.OriginalType != OriginalTypeGenApiIInteger {
			return newError("SetParam(%+q, %v): Unexpected int64 for type %s",
					  p.Name, value, p.OriginalType)
		}
		cValue := C.int(v)
		C.setNodeMapIntParam(cName, cValue)

	case float64:
		if p.OriginalType != OriginalTypeGenApiIFloat {
			return newError("SetParam(%+q, %v): Unexpected float64 for type %s",
					  p.Name, value, p.OriginalType)
		}

		cValue := C.double(v)
		C.setNodeMapFloatParam(cName, cValue)

	default:
		return newError("SetParam(%+q, %v): Value type %T is not implemented.",
				  p.Name, value, value)
	}
	return nil
}

func (cam *Camera) GetParam(p Param) (value interface{}, e error) {
	cam.openMutex.Lock(); defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return nil, newError("GetParam: Camera is not open.")
	}

/*
	cam.startMutex.Lock(); defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return nil, newError("GetParam: Camera is grabbing.")
	}
*/
	cName := C.CString(p.Name); defer C.free(unsafe.Pointer(cName))

	switch p.OriginalType {
	case OriginalTypeGenApiIEnumerationT:
		cValue := C.CString(string(make([]byte, 512)))
		defer C.free(unsafe.Pointer(cValue))
		msg := C.GoString(C.getNodeMapEnumParam(cName, cValue))
		if len(msg) != 0 {
			return nil, newError("C.getNodeMapEnumParam: %s", msg)
		}
		return C.GoString(cValue), nil
	// case OriginalTypeGenApiIString:
	case OriginalTypeGenApiIInteger:
		cValue := C.int64_t(0)
		msg := C.GoString(C.getNodeMapIntParam(cName, &cValue))
		if len(msg) != 0 {
			return nil, newError("C.getNodeMapEnumParam: %s", msg)
		}
		return int64(cValue), nil
	case OriginalTypeGenApiIFloat:
		cValue := C.double(0.0)
		msg := C.GoString(C.getNodeMapFloatParam(cName, &cValue))
		if len(msg) != 0 {
			return nil, newError("C.getNodeMapEnumParam: %s", msg)
		}
		return float64(cValue), nil
	}
	return nil, newError("GetParam(%+q): unsupported type %#v (%T)",
				p.Name, p.OriginalType, p.OriginalType)
}
