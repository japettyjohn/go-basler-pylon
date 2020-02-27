// Copyright (C) jno, 2019; License: MIT

/* Save to .JPEG using MagicWand from ImageMagick
 * Use as:
 *
 *	mw := im_setup(); defer im_cleanup(mw)
 *	...
 *	if e := im_save2jpeg(mw, ...); if e != nil { t.Fatalf("save: %v", e); }
 *	...
 *
 */
package pylon
/*
import (
	"os"
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
)

const imageJpegQuality = 75

// im_save2jpeg: using ImageMagick
func im_save2jpeg(mw *imagick.MagickWand, path string, width, height int, data []byte) error {
	defer mw.Clear()

	pw := imagick.NewPixelWand(); defer pw.Destroy()

	mw.ResetIterator()
	if e := mw.NewImage(uint(width), uint(height), pw); e != nil {
		return newError("NewImage: %v", e); }
	if e := mw.ImportImagePixels(0, 0, // start
				     uint(width), uint(height), // size
				     "I", imagick.PIXEL_CHAR, // type
				     data); // data
	  e != nil { return newError("ImportImagePixels: %v", e); }
	mw.ResetIterator()
	if e := mw.SetImageFormat("JPEG"); e != nil {
		return newError("SetImageFormat: %v", e); }
	mw.ResetIterator()
	if e := mw.SetImageCompressionQuality(imageJpegQuality); e != nil {
		return newError("SetImageCompressionQuality: %v", e)
	}
//	if e := mw.SetOrientation(imagick.ORIENTATION_UNDEFINED); e != nil {
//		return newError("SetOrientation: %v", e)
//	}
	mw.ResetIterator()
	var propNames = []string{
		"EXIF:UserComment",
		"EXIF:Photo.UserComment",
		"EXIF:37510",
		"EXIF:0x9286",
		"exif:UserComment",
		"exif:Photo.UserComment",
		"exif:37510",
		"exif:0x9286",
		"UserComment",
		"37510",
		"0x9286",
	}
	const propValue string = "this is a sample user comment"
	for _, pn := range propNames {
		if e := mw.SetImageProperty(pn, propValue);
			e != nil { fmt.Printf("SetImageProperty: %v", e); }
		if e := mw.SetImageArtifact(pn, propValue);
			e != nil { fmt.Printf("SetImageArtifact: %v", e); }
	}
	if e := mw.WriteImage(path); e != nil {
		return newError("WriteImage: %v", e); }
	return nil
}

// ImageMagick setup routine
func im_setup() *imagick.MagickWand {
	imagick.Initialize()
	imName, imVer := imagick.GetVersion()
	fmt.Printf("About to use ImageMagick %#v (%v)\n\n", imName, imVer)
	return imagick.NewMagickWand()
}
// ImageMagick cleanup routine
func im_cleanup(mw *imagick.MagickWand) {
	mw.Destroy()
	imagick.Terminate()
	fmt.Printf("Done with ImageMagick\n")
}

// im_show: utility function to display an image
func im_show(mw *imagick.MagickWand, path string) error {
	defer mw.Clear()
	fmt.Printf("Viewing %#v ...\n", path)
	if e := mw.ReadImage(path); e != nil {
		return newError("ReadImage(%#v): %v", path, e)
	}
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()
	if e := mw.ResizeImage(w/2, h/2, imagick.FILTER_LANCZOS, 1); e != nil {
		return newError("ResizeImage: %v", e)
	}
	if e := mw.SetImageCompressionQuality(imageJpegQuality); e != nil {
		return newError("SetImageCompressionQuality: %v", e)
	}
	fmt.Printf("Image identification: %s", mw.IdentifyImage())
	var n int = 0
	for i, a := range mw.GetImageArtifacts("") {
		v := mw.GetImageProperty(a)
		fmt.Printf("[%#v]=[%#v]=%#v\n", i, a, v)
		n++
	}
	if n == 0 {
		fmt.Printf("Image has no artifacts.\n")
	}
	if e := mw.DisplayImage(os.Getenv("DISPLAY")); e != nil {
		return newError("DisplayImage(%#v): %v", os.Getenv("DISPLAY"), e)
	}
	return nil
}
*/
/* EOF */
