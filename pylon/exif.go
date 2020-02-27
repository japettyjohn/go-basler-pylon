// Copyright(C) jno, 2019; License: MIT

// Provide a mean to inject EXIF metadata into fresh image *before* writing
package pylon

import (
	"fmt"
	"bytes"
	"time"
	"encoding/binary"
	"github.com/dsoprea/go-exif"
	"github.com/dsoprea/go-jpeg-image-structure"
	"os"
)

var ByteOrder = binary.LittleEndian // binary.BigEndian // == exif.TestDefaultByteOrder
var refTagIndex = exif.NewTagIndex()
var standardIfdPathes = []string{
	exif.IfdPathStandard,
	exif.IfdPathStandardExif,
	exif.IfdPathStandardExifIop,
	exif.IfdPathStandardGps,
}

// attempts to find a suitable `path` for `name`
func name2path(name string) (string, error) {
	for _, path := range standardIfdPathes {
		_, err := refTagIndex.GetWithName(path, name)
		if err == nil {
			return path, nil
		}
	}
	return "", newError("No default path for %+q", name)
}

/*******************************************************************************
 * types to carry lists of EXIF tags to be injected
 */
type Exif map[string]interface{}

type ExifTag struct {
	Path string
	Name string
	Value interface{}
}

type ExifTagList []*ExifTag

// empty ("") `path` means "search for path yourself"
func NewExifTag(path, name string, value interface{}) *ExifTag {
	var e error
	tag := new(ExifTag)
	if path == "" {
		path, e = name2path(name)
		if e != nil { panic(e); }
	}
	tag.Path = path
	tag.Name = name
	tag.Value = value
	return tag
}

func TimestampExifTagValue(value time.Time) string {
	return exif.ExifFullTimestampString(value)
}
func StringExifTagValue(value string) string {
	return value
}
func StringExifTag(path, name string, value string) *ExifTag {
	return NewExifTag(path, name, StringExifTagValue(value))
}

func UserCommentExifTagValue(value string) exif.TagUnknownType_9298_UserComment {
	v := exif.TagUnknownType_9298_UserComment{
		EncodingType: exif.TagUnknownType_9298_UserComment_Encoding_ASCII,
		EncodingBytes: []byte(value),
	}
	return v
}
func UserCommentExifTag(path, name string, value string) *ExifTag {
	return NewExifTag(path, name, UserCommentExifTagValue(value))
}

func ShortExifTagValue(value uint16) []uint16 {
	return []uint16{uint16(value)}
}
func ShortExifTag(path, name string, value uint16) *ExifTag {
	return NewExifTag(path, name, ShortExifTagValue(value))
}

func RationalExifTagValue(num uint32, den uint32) []exif.Rational {
	v := []exif.Rational{exif.Rational{Numerator: num, Denominator: den}}
	return v
}
func RationalExifTag(path, name string, num uint32, den uint32) *ExifTag {
	return NewExifTag(path, name, RationalExifTagValue(num, den))
}

func SRationalExifTagValue(num int32, den int32) []exif.SignedRational {
	v := exif.SignedRational{Numerator: num, Denominator: den}
	return []exif.SignedRational{v}
}
func SRationalExifTag(path, name string, num int32, den int32) *ExifTag {
	return NewExifTag(path, name, SRationalExifTagValue(num, den))
}
/******************************************************************************/

/* "class" to handle EXIF data and insertion of tags */
type ExifInjector struct {
	Jpeg []byte
	Tags ExifTagList

	jmp *jpegstructure.JpegMediaParser
	sl *jpegstructure.SegmentList
	rootIb *exif.IfdBuilder
	dqt map[int]*jpegstructure.Segment
}
func (self *ExifInjector) blankRootIfdBuilder() error {
	im := exif.NewIfdMapping()
	if e := exif.LoadStandardIfds(im); e != nil {
		return newError("exif.LoadStandardIfds: %v", e)
	}
	self.rootIb = exif.NewIfdBuilder(
		im,
		exif.NewTagIndex(),
		exif.IfdPathStandard,
		ByteOrder,
	)
	return nil
}
func (self *ExifInjector) makeRootIfdBuilder() error {
	if ib, e := self.sl.ConstructExifBuilder(); e != nil {
		if e.Error() == "no exif data" { // our case: new image has no EXIF!
			if e := self.blankRootIfdBuilder(); e != nil {
				return newError("blankIfdBuilder: %v", e)
			}
		} else {
			return newError("ConstructExifBuilder: %v", e)
		}
	} else {
		self.rootIb = ib
	}
	return nil
}
// sometimes the library just loose DQT...
func (self *ExifInjector) saveDQT() {
	for i, seg := range self.sl.Segments() {
		if seg.MarkerId == jpegstructure.MARKER_DQT {
			if self.dqt == nil {
				self.dqt = make(map[int]*jpegstructure.Segment)
			}
			self.dqt[i] = seg
		}
	}
}
func (self *ExifInjector) hasSavedDQT() bool {
	return self.dqt != nil && len(self.dqt) > 0
}
// do we still have a DQT?
func (self *ExifInjector) HasDQT() bool {
	for _, seg := range self.sl.Segments() {
		if seg.MarkerId == jpegstructure.MARKER_DQT {
			return true
		}
	}
	return false
}
func (self *ExifInjector) restoreDQT() {
	if !self.HasDQT() {
		fmt.Fprintln(os.Stderr, "WARNING: Restoring lost DQT...")
		if !self.hasSavedDQT() {
			panic("I have no saved DQT to restore from!")
		}
		xsl := jpegstructure.NewSegmentList([]*jpegstructure.Segment{})
		for i, seg := range self.sl.Segments() {
			if self.hasSavedDQT() {
				xs, found := self.dqt[i]
				if found  {
					xsl.Add(xs)
					delete(self.dqt, i)
				}
			}
			xsl.Add(seg)
		}
		self.sl = xsl
	}
}
// Normally you don't need to call LoadBytes() as it is called from NewExifInjector()
// The call sets .Jpeg and initializes interim data structures for .Inject()ion.
func (self *ExifInjector) LoadBytes(data []byte) error {
	var e error

	self.Jpeg = data
	self.jmp = jpegstructure.NewJpegMediaParser()

	self.sl, e = self.jmp.ParseBytes(self.Jpeg)
	if e != nil {
		return newError("ParseBytes: %v", e)
	}
	// self.saveDQT()
	e = self.makeRootIfdBuilder()
	if e != nil {
		return newError("makeRootIfdBuilder: %v", e)
	}
	return nil
}
// AddTag: adds an EXIF tag to be .Inject()ed
func (self *ExifInjector) AddTag(path, name string, value interface{}) {
	tag := NewExifTag(path, name, value)
	self.Tags = append(self.Tags, tag)
}
// AddTagList: adds an entire bunch of EXIF tags to be .Inject()ed
func (self *ExifInjector) AddTagList(list *ExifTagList) {
	for _, tag := range *list {
		self.Tags = append(self.Tags, tag)
	}
}
func (self *ExifInjector) AddTags(list map[string]interface{}) {
	for name, value := range list {
		tag := NewExifTag("", name, value)
		self.Tags = append(self.Tags, tag)
	}
}
func (self *ExifInjector) setTag(path, name string, value interface{}) error {
	if ifdIb, e := exif.GetOrCreateIbFromRootIb(self.rootIb, path); e != nil {
		return newError("Cannot GetOrCreateIbFromRootIb(*, %q): %v", path, e)
	} else if e = ifdIb.SetStandardWithName(name, value); e != nil {
		return newError("Cannot SetStandardWithName(%q, %+v): %v", name, value, e)
	}
	return nil
}
// Inject: will inject tags added with .AddTag() and/or .AddTagList() calls
func (self *ExifInjector) Inject() error {
	var e error

	for _, tag := range self.Tags {
		if e = self.setTag(tag.Path, tag.Name, tag.Value); e != nil {
			return newError("Cannot setTag(*, %q, %q, %+v): %v",
					tag.Path, tag.Name, tag.Value, e)
		}
	}

	if e = self.sl.SetExif(self.rootIb); e != nil {
		return newError("Cannot SetExif: %v", e)
	}
	// self.restoreDQT()

	buf := new(bytes.Buffer)
	if e = self.sl.Write(buf); e != nil {
		return newError("Cannot sl.Write: %v", e)
	}
	self.Jpeg = buf.Bytes()

	fmt.Println("EXIF added:")
	self.sl.Print()
	// self.sl.Validate(???)
	return nil
}
// NewExifInjector: constructor; I don't want to see uninitialized injectors
func NewExifInjector(data []byte) (*ExifInjector, error) {
	ei := new(ExifInjector)
	err := ei.LoadBytes(data)
	fmt.Println("Loaded:")
	ei.sl.Print()
	return ei, err
}

/* EOF */
