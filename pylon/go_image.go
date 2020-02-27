// Copyright (C) jno, 2019; License: MIT

// Save to .JPEG using vanilla Go library
package pylon

import (
	"os"
	"fmt"
	"bytes"
	"image"
	"image/jpeg"
)

const JpegFileSuffix = ".jpg"

// save to JPEG using Go's builtin facility
var GoJpegOptions = jpeg.Options{ Quality: jpeg.DefaultQuality } // sorta const, yeah

func writeTo(path string, data []byte) error {
	of, e := os.Create(path)
	if e != nil {
		return newError("Cannot Create(%#v): %v", path, e)
	}
	defer of.Close()
	_, e = of.Write(data)
	if e != nil {
		return newError("Cannot Write(...%d...) to %q: %v",
				len(data), path, e)
	}
	return nil
}

var saveRaw = writeTo

func loadRaw(path string) (data []byte, e error) {
	f, e := os.Open(path)
	if e != nil { return nil, newError("Cannot Open(%+q): %v", path, e); }
	defer f.Close()
	s, e := f.Stat()
	if e != nil { return nil, newError("Cannot Stat(%+q): %v", path, e); }
	data = make([]byte, s.Size())
	n, e := f.Read(data)
	if e != nil { return nil, newError("Cannot Read(%+q): %v", path, e); }
	if n != len(data) {
		return data, newError("Read(%+q) returned %d instead of %d",
					path, n, len(data))
	}
	return data, nil
}

/*
    META from MakerNote:
        stamp: 2019-10-25 13:48:02.185Z
        clock: 0.668262
        sharpness: 1
        gain: 35
        fps: 35.805077159941
        noise_reduction:
            0: 0.5
            1: 0
            2: 2
        difference:
        image_number: 1
        exposure: 0.0002
    FILE:
        FileName: 016e032c8ef5.jpg
        FileDateTime: 1572011342
        FileSize: 591140
        FileType: 2
        MimeType: image/jpeg
        SectionsFound: ANY_TAG, IFD0, EXIF, GPS
    COMPUTED:
        html: width="2448" height="2048"
        Height: 2048
        Width: 2448
        IsColor: 0
        ByteOrderMotorola: 0
        ApertureFNumber: f/1.4
        UserComment: { "stamp": "2019-10-25 13:48:02.185Z", "clock": 0.668262, "sharpness": 1.0, "gain": 35.0, "fps": 35.80507715994128, "noise_reduction": [ 0.5, 0.0, 2.0 ], "difference": null, "image_number": 1, "exposure": 0.0002 }
        UserCommentEncoding: UNDEFINED
        Copyright: Copyright, Inbase of Nekst LLC, 2019. All rights reserved.
    IFD0:
        Make: Inbase
        Model: Upyr1
        XResolution: 71/1
        YResolution: 71/1
        ResolutionUnit: 2
        Software: U1 by Inbase
        DateTime: 2019-10-25 13:48:02Z
        Artist: Inbase Upyr1 Bot (from JSON file)
        HostComputer: jet-one
        Copyright: Copyright, Inbase of Nekst LLC, 2019. All rights reserved.
        ExposureTime: 2/10000
        FNumber: 14/10
        Exif_IFD_Pointer: 468
        GPS_IFD_Pointer: 1710
        UndefinedTag:0x882A: 0
        UndefinedTag:0xC614: Inbase Upyr1
        UndefinedTag:0xC615: Inbase Upyr1
        UndefinedTag:0xC62C: 1/1
        UndefinedTag:0xC6F7: 25/100
    EXIF:
        ISOSpeedRatings: 3500
        DateTimeOriginal: 2019-10-25 13:48:02Z
        DateTimeDigitized: 2019-10-25 13:48:02Z
        ShutterSpeedValue: 12287712/1000000
        ApertureValue: 1/1
        FocalLength: 25/1
        MakerNote: ...
        UserComment: { "stamp": "2019-10-25 13:48:02.185Z", "clock": 0.668262, "sharpness": 1.0, "gain": 35.0, "fps": 35.80507715994128, "noise_reduction": [ 0.5, 0.0, 2.0 ], "difference": null, "image_number": 1, "exposure": 0.0002 }
        ColorSpace: 65535
        Sharpness: 0
        SubjectDistanceRange: 3
        UndefinedTag:0xA433: Fujian
        UndefinedTag:0xA434: CCTV 25mm f/1.4 C-Mount
        UndefinedTag:0xA435: 1234567890
    GPS:
        GPSLatitudeRef: N
        GPSLatitude: 5575564/100000
        GPSLongitudeRef: E
        GPSLongitude: 3756544/100000
*/

// makeExif(jpegData) --> (exifData, error)
func makeExif(jpg []byte, tags Exif) ([]byte, error) {
	if injector, e := NewExifInjector(jpg); e != nil {
		return nil, newError("NewExifInjector: %v", e)
	} else {
		// injector.AddTagList(&tagList)
		injector.AddTags(tags)

		if e = injector.Inject(); e != nil {
			return nil, newError("Cannot Inject: %v", e)
		}

		return injector.Jpeg, nil
	}

	return nil, nil
}

// go_save2jpeg: native Go implementation
func go_save2jpeg(path string, width, height int, data []byte, tags Exif) (e error) {
	raw := image.NewGray(image.Rect(0, 0, width, height))
	raw.Pix = data

	var tmp bytes.Buffer
	e = jpeg.Encode(&tmp, raw, &GoJpegOptions)
	if e != nil {
		return newError("Cannot Encode(): %v", e)
	}
	jpg := tmp.Bytes()

	var jfif []byte
	if tags != nil {
		jfif, e = makeExif(jpg, tags)
		if e != nil {
			return newError("Cannot makeExif(): %v", e)
		}
		fmt.Printf("jfif(%d): %v\n", len(jfif), jfif[:32])
	}

	e = writeTo(path, jpg)
	if e != nil {
		return newError("Cannot writeTo(%q): %v", path, e)
	}

	if tags != nil {
		e = writeTo(path + JpegFileSuffix, jfif)
		if e != nil {
			return newError("Cannot writeTo(%q): %v", path + JpegFileSuffix, e)
		}
	}

	fmt.Printf("Written to %#v\n", path)
	return nil
}

/* EOF */
