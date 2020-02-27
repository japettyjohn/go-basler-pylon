// Copyright (C) jno, 2019; License: MIT

// EPixelType is an `enum` (i.e. `int`).
// It is supposed to have 32 bits like 0xabCD1234
// where 0xab000000 is "pixel class" (mono/color/custom),
// 0x00CD0000 is "pixel width" (bits per pixel), and
// 0x00001234 is "pixel subtype" (just serial number).
// See $PYLON_HOME/include/pylon/PixelType.h for details
package pylon

import "fmt"

type EPixelType uint32

func NewEPixelType(class, bits, subtype int) EPixelType {
	if class & ^0x0ff != 0 { panic("invalid class"); }
	if bits & ^0x0ff != 0 { panic("invalid bit size"); }
	if subtype & ^0x0ffff != 0 { panic("invalid subtype"); }
	return EPixelType((class << 24) | (bits << 16) | subtype)
}

const ePixelType_UNDEFINED = -1 // as in the source

const ePixelType_Class_MONO = 0x01
const ePixelType_Class_COLOR = 0x02
const ePixelType_Class_CUSTOM = 0x80

const ePixelType_Valid_Class =	ePixelType_Class_MONO |
				ePixelType_Class_COLOR |
				ePixelType_Class_CUSTOM

func (pt *EPixelType) String() string {
	if !pt.IsValid() {
		return fmt.Sprintf("<EPixelType:INVALID:%08x>", *pt)
	}
	if pt.IsUndefined() {
		return "<EPixelType:UNDEFINED>"
	}
	v := PixelTypeName[*pt]
	if v == "" {
		return fmt.Sprintf("<EPixelType:UNSUPPORTED:%s(%d)%04x>",
				   pt.ClassName(), pt.Bits(), pt.Type())
	}
	return fmt.Sprintf("<EPixelType:%s>", v)
}

func (pt *EPixelType) ClassName() string {
	var class string = ""
	if pt.IsCustom() {
		class = "CUSTOM"
	}
	if pt.IsColor() {
		if class != "" {
			class += ","
		}
		class += "COLOR"
	}
	if pt.IsMono() {
		if class != "" {
			class += ","
		}
		class += "MONO"
	}
	return class
}

func (pt *EPixelType) Class() int { return int((*pt >> 24) & 0x0ff); }
func (pt *EPixelType) Bits() int { return int((*pt >> 16) & 0x0ff); }
func (pt *EPixelType) Type() int { return int(*pt & 0x0ffff); }

func (pt *EPixelType) IsValid() bool {
	return pt.IsUndefined() || pt.Class() & ^ePixelType_Valid_Class == 0
}

func (pt *EPixelType) IsMono() bool {
	return pt.Class() & ePixelType_Class_MONO == ePixelType_Class_MONO
}

func (pt *EPixelType) IsColor() bool {
	return pt.Class() & ePixelType_Class_COLOR == ePixelType_Class_COLOR
}

func (pt *EPixelType) IsCustom() bool {
	return pt.Class() & ePixelType_Class_CUSTOM == ePixelType_Class_CUSTOM
}

func (pt *EPixelType) IsUndefined() bool {
	return int(*pt) == ePixelType_UNDEFINED
}

