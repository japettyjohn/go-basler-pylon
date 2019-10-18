package pylon

// #include <string.h>
// #include <stdlib.h>
// #include "capture.h"
import "C"

import (
	"fmt"
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

func (cam *Camera) Info() (*CameraInfo, error) {
	if !C.isAttached() {
		return nil, fmt.Errorf("Info: No device attached")
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
		s := C.GoString(C.openCamera())
		if s != "" {
			return fmt.Errorf("OpenCamera: %v", s)
		}
	}
	return nil
}

func (cam *Camera) CloseCamera() error {
	fmt.Println("CloseCamera() called")
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if C.isOpen() {
		s := C.GoString(C.closeCamera())
		if s != "" {
			return fmt.Errorf("CloseCamera: %v", s)
		}
	}
	return nil
}

func (cam *Camera) StartCapture(max int) error {
	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if !C.isCameraGrabbing() {
		s := C.startCapture(C.int(max))
		if errMsg := C.GoString(s); errMsg != "" {
			return fmt.Errorf("StartCapture: %v", errMsg)
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
		s := C.GoString(C.stopCapture())
		if s != "" {
			return fmt.Errorf("StopCapture: %v", s)
		}
	}
	return nil
}

func (cam *Camera) AttachDevice() error {
	cam.attachedMutex.Lock()
	defer cam.attachedMutex.Unlock()
	if !C.isAttached() {
		s := C.GoString(C.attachDevice())
		if s != "" {
			return fmt.Errorf("AttachDevice: %v", s)
		}
	}
	return nil
}

func (cam *Camera) ConfigureCamera() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return fmt.Errorf("ConfigureCamera: Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return fmt.Errorf("ConfigureCamera: Camera is grabbing.")
	}

	if msg := C.GoString(C.configureCamera()); msg != "" {
		return fmt.Errorf("ConfigureCamera: %s", msg)
	}
	return nil
}

func (cam *Camera) SetHardwareTriggerConfiguration() error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return fmt.Errorf("Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return fmt.Errorf("Camera is grabbing.")
	}

	s := C.GoString(C.setHardwareTriggerConfiguration())
	if s != "" {
		return fmt.Errorf("SetHardwareTriggerConfiguration: %v", s)
	}
	return nil
}

func (cam *Camera) RetrieveAndSave(batch, timeout int, outputPath string) error {
	cOutputPath := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cOutputPath))
	s := C.retrieveAndSave(C.int(batch), C.int(timeout), cOutputPath)
	if errMsg := C.GoString(s); errMsg != "" {
		return fmt.Errorf("RetrieveAndSave: %v", errMsg)
	}
	return nil
}

func (cam *Camera) SetParam(p Param, value interface{}) error {
	cam.openMutex.Lock()
	defer cam.openMutex.Unlock()
	if !C.isOpen() {
		return fmt.Errorf("SetParam: Camera is not open.")
	}

	cam.startMutex.Lock()
	defer cam.startMutex.Unlock()
	if C.isCameraGrabbing() {
		return fmt.Errorf("SetParam: Camera is grabbing.")
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
			return fmt.Errorf("SetParam: Original type %s not implemented.",
					  p.OriginalType)
		default:
			return fmt.Errorf("SetParam: Unexpected string for type %s",
					  p.OriginalType)
		}

	case int64:
		if p.OriginalType != OriginalTypeGenApiIInteger {
			return fmt.Errorf("SetParam: Unexpected int64 for type %s",
					  p.OriginalType)
		}
		cValue := C.int(v)
		C.setNodeMapIntParam(cName, cValue)

	case float64:
		if p.OriginalType != OriginalTypeGenApiIFloat {
			return fmt.Errorf("SetParam: Unexpected float64 for type %s",
					  p.OriginalType)
		}

		cValue := C.double(v)
		C.setNodeMapFloatParam(cName, cValue)

	default:
		return fmt.Errorf("SetParam: Value type %T of param %s not implemented.",
				  value, p.Name)
	}
	return nil
}
