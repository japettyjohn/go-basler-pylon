// Copyright(C) jno, 2019; License: MIT
package pylon

import (
	"os"
	"fmt"
	"path"
	"path/filepath"
	"time"
	"testing"
)
var imgPath = path.Join(os.TempDir(), "go-basler-pylon-test")

type Timer struct {
	Count int
	Min, Max, Taken, Average float64
	Start, Last time.Time
	Elapsed, Total time.Duration
}
func (self *Timer) Update() {
	now := time.Now()
	self.Elapsed = now.Sub(self.Last)
	self.Total = now.Sub(self.Start)
	self.Last = now
	self.Taken = self.Elapsed.Seconds()
	if self.Count == 0 {
		self.Min = self.Taken
		self.Max = self.Taken
	} else {
		if self.Min > self.Taken { self.Min = self.Taken; }
		if self.Max < self.Taken { self.Max = self.Taken; }
	}
	self.Count += 1
	self.Average = self.Total.Seconds() / float64(self.Count)
}

func NewTimer() Timer {
	var t Timer
	t.Start = time.Now()
	t.Last = t.Start
	t.Elapsed = t.Last.Sub(t.Start)
	t.Total = t.Elapsed
	t.Taken = t.Elapsed.Seconds()
	t.Min = t.Taken
	t.Max = t.Taken
	return t
}

func try(t *testing.T, e error, args ...string) {
	if len(args) > 1 { panic("try(*testing.T, error[, \"format\"])"); }
	format := "Failed in %v"
	if len(args) > 0 { format = args[0]; }
	if e != nil { t.Fatalf(format, e); }
}

func time2name(t time.Time) string {
	nsecs := t.Unix() * 1000000000 + int64(t.Nanosecond())
	return fmt.Sprintf("%013x", nsecs / 1000)
}

func TestStart(t *testing.T) {
	// mw := im_setup(); defer im_cleanup(mw)

	cam := &Camera{}
	os.RemoveAll(imgPath)

	try(t, os.MkdirAll(imgPath, 0777))

	try(t, cam.AttachDevice())
	try(t, cam.OpenCamera())
	defer cam.CloseCamera()
	try(t, cam.ConfigureCamera())

	i, e := cam.Info(); try(t, e)
	t.Logf("camera USB[%04x:%04x] is %d×%d %#v by %#v, %#v, serial %#v, version %#v",
		i.VendorId, i.ProductId,
		i.Width, i.Height, i.ModelName, i.VendorName,
		i.FullName, i.SerialNumber, i.DeviceVersion)

	v1, e := cam.GetParam(AutoGainUpperLimit)
	try(t, e)
	t.Logf("AutoGainUpperLimit=%#v (%T)", v1, v1)

	v2, e := cam.GetParam(GainAuto)
	try(t, e)
	t.Logf("GainAuto=%#v (%T)", v2, v2)

	v3, e := cam.GetParam(SensorWidth)
	try(t, e)
	t.Logf("SensorWidth=%#v (%T)", v3, v3)

	var saved []string
	total := NewTimer()

	FrameCallback := func(cam *Camera, w, h, pxt, size int, buffer []byte) int {
		pt := EPixelType(pxt)
		t.Logf("FrameCallback(w=%#v, h=%#v, pt=%08x=%s, size=%#v, buffer=%#v...)",
			w, h, pxt, pt.String(), size, buffer[0])
		tx := NewTimer()
		// DO STUFF FROM HERE

		path := filepath.Join(imgPath, time2name(tx.Last.UTC()))

		try(t, saveRaw(path + ".raw", buffer))
		// try(t, go_save2jpeg(path + ".jpg", w, h, buffer))	// vanilla Go
		// try(t, im_save2jpeg(mw, path, w, h, buffer))	// ImageMagick

		saved = append(saved, path)

		// UNTIL HERE
		tx.Update()
		total.Update()
		t.Logf("FrameCallback taken %v", tx.Elapsed)
		return 0
	}

	cam.SetFetchTimeout(5000) // ms
	cam.SetFetchCount(10)
	cb := FrameCallbackType(FrameCallback)
	try(t, cam.Fetch(cb), "Fetch failed: %v")

	t.Logf("Total esapsed %v over %d iterations at about %.3f sec each [%.3f..%.3f]",
		   total.Total, total.Count, total.Average, total.Min, total.Max)

	if f, e := os.Open(imgPath); e != nil {
		t.Fatalf("os.Open(%#v) failed: %v", imgPath, e)
	} else if n, e := f.Readdirnames(-1); e != nil {
		t.Fatalf("Readdirnames failed: %v", e)
	} else if len(n) == 0 {
		t.Fatalf("No scans taken.")
	} else {
		t.Logf("Scanned %s", n[0])
	}
}

/*
func TestHardwareTrigger(t *testing.T) {
	cam := &Camera{}
	imgPath := path.Join(os.TempDir(), "go-basler-pylon-test")
	os.RemoveAll(imgPath)

	if e := os.MkdirAll(imgPath, 0777); e != nil {
		t.Fatalf("Failed with %v", e)
	}

	cam.AttachDevice()
	cam.OpenCamera()
	defer cam.CloseCamera()
	if e := cam.SetHardwareTriggerConfiguration(); e != nil {
		t.Fatalf("Failed with %v", e)
	}
	if e := cam.ConfigureCamera(); e != nil {
		t.Fatalf("Failed with %v", e)
	}

	i := cam.Info()
	t.Logf("camera USB[%04x:%04x] is %d×%d %#v by %#v, %#v, serial %#v, version %#v",
		i.VendorId, i.ProductId,
		i.Width, i.Height, i.ModelName, i.VendorName,
		i.FullName, i.SerialNumber, i.DeviceVersion)

	if e := cam.StartCapture(500); e != nil {
		t.Fatalf("Failed with %v", e)
	}
	defer cam.StopCapture()

	if e := cam.RetrieveAndSave(1, 1000, imgPath+"/"); e == nil {
		t.Fatalf("Expected timeout.")
	}
}
*/
/*
func TestParamGroup(t *testing.T) {
	expected := "This category includes items used to conduct file operations"
	if s := FileAccessControl.Description(); s != expected {
		t.Fatalf("Expected description '%#v' and go '%#v'.", expected, s)
	}

	expected = "FileAccessControl"
	if s := FileAccessControl.String(); s != expected {
		t.Fatalf("Expected description '%#v' and go '%#v'.", expected, s)
	}

}
*/
/* EOF */
