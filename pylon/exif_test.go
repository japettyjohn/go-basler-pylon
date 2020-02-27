package pylon

import "testing"
import (
	"os"
	"path"
	"time"
)

const (
	ImageWidth = 2448
	ImageHeight = 2048
)

var tags = Exif{
	"DateTime": TimestampExifTagValue(time.Now().UTC()),
	"Copyright": StringExifTagValue(
	  "Copyright, Inbase of Nekst LLC, 2019. All rights reserved."),
	"Make": StringExifTagValue("Inbase"),
	"Model": StringExifTagValue("Upyr1"),
	"Software": StringExifTagValue("U1 by Inbase"),
	"Artist": StringExifTagValue("Inbase Upyr1 Bot"),
	"HostComputer": StringExifTagValue("localhost"),

	"UserComment": UserCommentExifTagValue("TEST COMMENT"),

	"ExposureTime": RationalExifTagValue(1, 350),
	"FNumber": RationalExifTagValue(14, 10),

	"ColorSpace": ShortExifTagValue(65535),
	"Sharpness": ShortExifTagValue(0),
	"SubjectDistanceRange": ShortExifTagValue(3),
	"ISOSpeedRatings": ShortExifTagValue(2),
	"ShutterSpeedValue": SRationalExifTagValue(12287712, 1000000),
	"ApertureValue": RationalExifTagValue(1, 1),
	"FocalLength": RationalExifTagValue(25, 1),

	"GPSLatitudeRef": StringExifTagValue("N"),
	"GPSLatitude": RationalExifTagValue(5575564, 100000),
	"GPSLongitudeRef": StringExifTagValue("E"),
	"GPSLongitude": RationalExifTagValue(3756544, 100000),
}

func ListDir(t *testing.T, path string) {
	if f, e := os.Open(path); e != nil {
		t.Fatalf("os.Open(%+q) failed: %v", path, e)
	} else if n, e := f.Readdirnames(-1); e != nil {
		defer f.Close()
		t.Fatalf("Readdirnames(%+q) failed: %v", path, e)
	} else if len(n) == 0 {
		defer f.Close()
		t.Fatalf("No files found in %+q.", path)
	} else {
		defer f.Close()
		t.Logf("Found files: %d", len(n))
		for i, x := range n {
			t.Logf("%5d %s", i, x)
		}
	}
}

func SingleFile(t *testing.T, name string) {
	hostname, err := os.Hostname()
	if err != nil { t.Fatalf("No host name!"); }
	tags["HostComputer"] = StringExifTagValue(hostname)

	t.Logf("[%s]", name)
	image, e := loadRaw(path.Join(imgPath, name))
	if e != nil { t.Fatalf("Cannot load %+q: %v", name, e); }

	jpeg := path.Join(imgPath, path.Base(name)) + ".jpg"
	try(t, go_save2jpeg(jpeg, ImageWidth, ImageHeight, image, tags))
}

func TestExif(t *testing.T) {
	t.Logf("Test Exif")
	if f, e := os.Open(imgPath); e != nil {
		t.Fatalf("os.Open(%#v) failed: %v", imgPath, e)
	} else if n, e := f.Readdirnames(-1); e != nil {
		defer f.Close()
		t.Fatalf("Readdirnames failed: %v", e)
	} else if len(n) == 0 {
		defer f.Close()
		t.Fatalf("No scans taken.")
	} else {
		defer f.Close()
		ListDir(t, imgPath)
		t.Logf("Scanned %d", len(n))
		for _, x := range n {
			if path.Ext(x) != ".raw" { continue; }
			SingleFile(t, x)
		}
		ListDir(t, imgPath)
	}
}

/* EOF */
