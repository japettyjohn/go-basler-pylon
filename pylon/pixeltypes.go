// converted from $PYLON_HOME/include/pylon/PixelType.h
package pylon

const ePixelType_Class_CM = ePixelType_Class_CUSTOM | ePixelType_Class_MONO
const ePixelType_Class_CC = ePixelType_Class_CUSTOM | ePixelType_Class_COLOR

const (
    eptMn = ePixelType_Class_MONO << 24
    eptCr = ePixelType_Class_COLOR << 24
    eptCM = ePixelType_Class_CM << 24
    eptCC = ePixelType_Class_CC << 24

    ept01 =  1 << 16
    ept02 =  2 << 16
    ept04 =  4 << 16
    ept08 =  8 << 16
    ept10 = 10 << 16
    ept12 = 12 << 16
    ept16 = 16 << 16
    ept24 = 24 << 16
    ept32 = 32 << 16
    ept36 = 36 << 16
    ept48 = 48 << 16
)

const (
    PixelType_Undefined          = EPixelType(0x0ffffffff)

    PixelType_Mono8              = EPixelType(eptMn|ept08|0x0001)
    PixelType_Mono8signed        = EPixelType(eptMn|ept08|0x0002)
    PixelType_Mono10             = EPixelType(eptMn|ept16|0x0003)
    PixelType_Mono10packed       = EPixelType(eptMn|ept12|0x0004)
    PixelType_Mono10p            = EPixelType(eptMn|ept10|0x0046)
    PixelType_Mono12             = EPixelType(eptMn|ept16|0x0005)
    PixelType_Mono12packed       = EPixelType(eptMn|ept12|0x0006)
    PixelType_Mono12p            = EPixelType(eptMn|ept12|0x0047)
    PixelType_Mono16             = EPixelType(eptMn|ept16|0x0007)

    PixelType_BayerGR8           = EPixelType(eptMn|ept08|0x0008)
    PixelType_BayerRG8           = EPixelType(eptMn|ept08|0x0009)
    PixelType_BayerGB8           = EPixelType(eptMn|ept08|0x000a)
    PixelType_BayerBG8           = EPixelType(eptMn|ept08|0x000b)

    PixelType_Mono1packed        = EPixelType(eptCM|ept01|0x000c)
    PixelType_Mono2packed        = EPixelType(eptCM|ept02|0x000d)
    PixelType_Mono4packed        = EPixelType(eptCM|ept04|0x000e)

    PixelType_BayerGR10          = EPixelType(eptMn|ept16|0x000c)
    PixelType_BayerRG10          = EPixelType(eptMn|ept16|0x000d)
    PixelType_BayerGB10          = EPixelType(eptMn|ept16|0x000e)
    PixelType_BayerBG10          = EPixelType(eptMn|ept16|0x000f)

    PixelType_BayerGR12          = EPixelType(eptMn|ept16|0x0010)
    PixelType_BayerRG12          = EPixelType(eptMn|ept16|0x0011)
    PixelType_BayerGB12          = EPixelType(eptMn|ept16|0x0012)
    PixelType_BayerBG12          = EPixelType(eptMn|ept16|0x0013)

    PixelType_RGB8packed         = EPixelType(eptCr|ept24|0x0014)
    PixelType_BGR8packed         = EPixelType(eptCr|ept24|0x0015)

    PixelType_RGBA8packed        = EPixelType(eptCr|ept32|0x0016)
    PixelType_BGRA8packed        = EPixelType(eptCr|ept32|0x0017)

    PixelType_RGB10packed        = EPixelType(eptCr|ept48|0x0018)
    PixelType_BGR10packed        = EPixelType(eptCr|ept48|0x0019)

    PixelType_RGB12packed        = EPixelType(eptCr|ept48|0x001a)
    PixelType_BGR12packed        = EPixelType(eptCr|ept48|0x001b)

    PixelType_RGB16packed        = EPixelType(eptCr|ept48|0x0033)

    PixelType_BGR10V1packed      = EPixelType(eptCr|ept32|0x001c)
    PixelType_BGR10V2packed      = EPixelType(eptCr|ept32|0x001d)

    PixelType_YUV411packed       = EPixelType(eptCr|ept12|0x001e)
    PixelType_YUV422packed       = EPixelType(eptCr|ept16|0x001f)
    PixelType_YUV444packed       = EPixelType(eptCr|ept24|0x0020)

    PixelType_RGB8planar         = EPixelType(eptCr|ept24|0x0021)
    PixelType_RGB10planar        = EPixelType(eptCr|ept48|0x0022)
    PixelType_RGB12planar        = EPixelType(eptCr|ept48|0x0023)
    PixelType_RGB16planar        = EPixelType(eptCr|ept48|0x0024)

    PixelType_YUV422_YUYV_Packed = EPixelType(eptCr|ept16|0x0032)
    PixelType_YUV444planar       = EPixelType(eptCC|ept24|0x0044)
    PixelType_YUV422planar       = EPixelType(eptCC|ept16|0x0042)
    PixelType_YUV420planar       = EPixelType(eptCC|ept12|0x0040)

    PixelType_YCbCr422_8_YY_CbCr_Semiplanar= EPixelType(eptCr|ept16|0x0113)
    PixelType_YCbCr420_8_YY_CbCr_Semiplanar= EPixelType(eptCr|ept12|0x0112)

    PixelType_BayerGR12Packed    = EPixelType(eptMn|ept12|0x002A)
    PixelType_BayerRG12Packed    = EPixelType(eptMn|ept12|0x002B)
    PixelType_BayerGB12Packed    = EPixelType(eptMn|ept12|0x002C)
    PixelType_BayerBG12Packed    = EPixelType(eptMn|ept12|0x002D)

    PixelType_BayerGR10p         = EPixelType(eptMn|ept10|0x0056)
    PixelType_BayerRG10p         = EPixelType(eptMn|ept10|0x0058)
    PixelType_BayerGB10p         = EPixelType(eptMn|ept10|0x0054)
    PixelType_BayerBG10p         = EPixelType(eptMn|ept10|0x0052)
                                 
    PixelType_BayerGR12p         = EPixelType(eptMn|ept12|0x0057)
    PixelType_BayerRG12p         = EPixelType(eptMn|ept12|0x0059)
    PixelType_BayerGB12p         = EPixelType(eptMn|ept12|0x0055)
    PixelType_BayerBG12p         = EPixelType(eptMn|ept12|0x0053)

    PixelType_BayerGR16          = EPixelType(eptMn|ept16|0x002E)
    PixelType_BayerRG16          = EPixelType(eptMn|ept16|0x002F)
    PixelType_BayerGB16          = EPixelType(eptMn|ept16|0x0030)
    PixelType_BayerBG16          = EPixelType(eptMn|ept16|0x0031)

    PixelType_RGB12V1packed      = EPixelType(eptCr|ept36|0x0034)

    PixelType_Double             = EPixelType(eptCM|ept48|0x0100)
)

var PixelTypeName = map[EPixelType]string{
    PixelType_Undefined: "Undefined",

    PixelType_Mono8: "Mono8",
    PixelType_Mono8signed: "Mono8signed",
    PixelType_Mono10: "Mono10",
    PixelType_Mono10packed: "Mono10packed",
    PixelType_Mono10p: "Mono10p",
    PixelType_Mono12: "Mono12",
    PixelType_Mono12packed: "Mono12packed",
    PixelType_Mono12p: "Mono12p",
    PixelType_Mono16: "Mono16",

    PixelType_BayerGR8: "BayerGR8",
    PixelType_BayerRG8: "BayerRG8",
    PixelType_BayerGB8: "BayerGB8",
    PixelType_BayerBG8: "BayerBG8",

    PixelType_Mono1packed: "Mono1packed",
    PixelType_Mono2packed: "Mono2packed",
    PixelType_Mono4packed: "Mono4packed",

    PixelType_BayerGR10: "BayerGR10",
    PixelType_BayerRG10: "BayerRG10",
    PixelType_BayerGB10: "BayerGB10",
    PixelType_BayerBG10: "BayerBG10",

    PixelType_BayerGR12: "BayerGR12",
    PixelType_BayerRG12: "BayerRG12",
    PixelType_BayerGB12: "BayerGB12",
    PixelType_BayerBG12: "BayerBG12",

    PixelType_RGB8packed: "RGB8packed",
    PixelType_BGR8packed: "BGR8packed",

    PixelType_RGBA8packed: "RGBA8packed",
    PixelType_BGRA8packed: "BGRA8packed",

    PixelType_RGB10packed: "RGB10packed",
    PixelType_BGR10packed: "BGR10packed",

    PixelType_RGB12packed: "RGB12packed",
    PixelType_BGR12packed: "BGR12packed",

    PixelType_RGB16packed: "RGB16packed",

    PixelType_BGR10V1packed: "BGR10V1packed",
    PixelType_BGR10V2packed: "BGR10V2packed",

    PixelType_YUV411packed: "YUV411packed",
    PixelType_YUV422packed: "YUV422packed",
    PixelType_YUV444packed: "YUV444packed",

    PixelType_RGB8planar: "RGB8planar",
    PixelType_RGB10planar: "RGB10planar",
    PixelType_RGB12planar: "RGB12planar",
    PixelType_RGB16planar: "RGB16planar",

    PixelType_YUV422_YUYV_Packed: "YUV422_YUYV_Packed",
    PixelType_YUV444planar: "YUV444planar",
    PixelType_YUV422planar: "YUV422planar",
    PixelType_YUV420planar: "YUV420planar",

    PixelType_YCbCr422_8_YY_CbCr_Semiplanar: "YCbCr422_8_YY_CbCr_Semiplanar",
    PixelType_YCbCr420_8_YY_CbCr_Semiplanar: "YCbCr420_8_YY_CbCr_Semiplanar",

    PixelType_BayerGR12Packed: "BayerGR12Packed",
    PixelType_BayerRG12Packed: "BayerRG12Packed",
    PixelType_BayerGB12Packed: "BayerGB12Packed",
    PixelType_BayerBG12Packed: "BayerBG12Packed",

    PixelType_BayerGR10p: "BayerGR10p",
    PixelType_BayerRG10p: "BayerRG10p",
    PixelType_BayerGB10p: "BayerGB10p",
    PixelType_BayerBG10p: "BayerBG10p",

    PixelType_BayerGR12p: "BayerGR12p",
    PixelType_BayerRG12p: "BayerRG12p",
    PixelType_BayerGB12p: "BayerGB12p",
    PixelType_BayerBG12p: "BayerBG12p",

    PixelType_BayerGR16: "BayerGR16",
    PixelType_BayerRG16: "BayerRG16",
    PixelType_BayerGB16: "BayerGB16",
    PixelType_BayerBG16: "BayerBG16",

    PixelType_RGB12V1packed: "RGB12V1packed",

    PixelType_Double: "Double",
}

// Returns true if the pixel type is Mono and the pixel values are not byte aligned.
func IsMonoPacked(pixelType EPixelType) bool {
    switch pixelType {   	
    case PixelType_Mono1packed, PixelType_Mono2packed, PixelType_Mono4packed,
         PixelType_Mono10packed, PixelType_Mono10p,
         PixelType_Mono12packed, PixelType_Mono12p: return true
    }
    return false
}

// Returns true if the pixel type is Bayer and the pixel values are not byte aligned.
func IsBayerPacked(pixelType EPixelType) bool {
    switch pixelType {   	
    case PixelType_BayerGB12Packed, PixelType_BayerGR12Packed,
         PixelType_BayerRG12Packed, PixelType_BayerBG12Packed,
	 PixelType_BayerGB10p, PixelType_BayerGR10p,
         PixelType_BayerRG10p, PixelType_BayerBG10p,
         PixelType_BayerGB12p, PixelType_BayerGR12p,
         PixelType_BayerRG12p, PixelType_BayerBG12p: return true
    }
    return false
}

// Returns true if the pixel type is RGB and the pixel values are not byte aligned.
func IsRGBPacked(pixelType EPixelType) bool {
    switch pixelType {   	
    case PixelType_RGB12V1packed: return true
    }
    return false
}

// Returns true if the pixel type is BGR and the pixel values are not byte aligned.
func IsBGRPacked(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_BGR10V1packed, PixelType_BGR10V2packed: return true
    }
    return false
}

// Returns true if the pixels of the given pixel type are not byte aligned.
func IsPacked(pixelType EPixelType) bool {
    return IsMonoPacked(pixelType) || IsBayerPacked(pixelType) ||
    	   IsRGBPacked(pixelType) || IsBGRPacked(pixelType)
}

// Returns true if the pixel type is packed in lsb packed format.
// For lsb packed format, the data is filled lsb first in the lowest address
// byte (byte 0) starting with the first pixel and continued in the lsb of byte
// 1 (and so on).
// See the camera User's Manual or the Pixel Format Naming Convention (PFNC) of
// the GenICam standard group for more information.
func IsPackedInLsbFormat(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_Mono1packed, PixelType_Mono2packed, PixelType_Mono4packed,
	 PixelType_Mono10p   ,
         PixelType_BayerGB10p, PixelType_BayerGR10p,
	 PixelType_BayerRG10p, PixelType_BayerBG10p,
	 PixelType_BayerGB12p, PixelType_BayerGR12p,
	 PixelType_BayerRG12p, PixelType_BayerBG12p,
	 PixelType_Mono12p   : return true
    }
    return false
}

// Returns number of planes in the image composed of the pixel type.
func PlaneCount(pixelType EPixelType) uint32 {
    switch pixelType {
    case PixelType_RGB8planar, PixelType_RGB10planar, PixelType_RGB12planar,
	 PixelType_RGB16planar, PixelType_YUV444planar, PixelType_YUV422planar,
	 PixelType_YUV420planar:
	return 3
    }
    return 1
}

/// Returns the pixel type of a plane.
func GetPlanePixelType(pixelType EPixelType) EPixelType {
    switch pixelType {
    case PixelType_RGB8planar,
	 PixelType_YUV444planar,
	 PixelType_YUV422planar,
	 PixelType_YUV420planar: return PixelType_Mono8
    case PixelType_RGB10planar: return PixelType_Mono10
    case PixelType_RGB12planar: return PixelType_Mono12
    case PixelType_RGB16planar: return PixelType_Mono16
    }
    return pixelType
}

// Returns true if images of the pixel type are divided into multiple planes.
func IsPlanar(pixelType EPixelType) bool {
    return PlaneCount(pixelType) > 1
}

/// Lists the Bayer color filter types.
type EPixelColorFilter int
const (
    PCF_BayerRG = iota ///<red green
    PCF_BayerGB ///<green blue
    PCF_BayerGR ///<green red
    PCF_BayerBG ///<blue green
    PCF_Undefined ///< undefined color filter or not applicable
)

/// Returns the Bayer color filter type.
func GetPixelColorFilter(pixelType EPixelType) EPixelColorFilter {
    switch pixelType {
    case PixelType_BayerGR8: return PCF_BayerGR
    case PixelType_BayerRG8: return PCF_BayerRG
    case PixelType_BayerGB8: return PCF_BayerGB
    case PixelType_BayerBG8: return PCF_BayerBG

    case PixelType_BayerGR10: return PCF_BayerGR
    case PixelType_BayerRG10: return PCF_BayerRG
    case PixelType_BayerGB10: return PCF_BayerGB
    case PixelType_BayerBG10: return PCF_BayerBG

    case PixelType_BayerGR12: return PCF_BayerGR
    case PixelType_BayerRG12: return PCF_BayerRG
    case PixelType_BayerGB12: return PCF_BayerGB
    case PixelType_BayerBG12: return PCF_BayerBG

    case PixelType_BayerGR12Packed: return PCF_BayerGR
    case PixelType_BayerRG12Packed: return PCF_BayerRG
    case PixelType_BayerGB12Packed: return PCF_BayerGB
    case PixelType_BayerBG12Packed: return PCF_BayerBG

    case PixelType_BayerGR10p: return PCF_BayerGR
    case PixelType_BayerRG10p: return PCF_BayerRG
    case PixelType_BayerGB10p: return PCF_BayerGB
    case PixelType_BayerBG10p: return PCF_BayerBG

    case PixelType_BayerGR12p: return PCF_BayerGR
    case PixelType_BayerRG12p: return PCF_BayerRG
    case PixelType_BayerGB12p: return PCF_BayerGB
    case PixelType_BayerBG12p: return PCF_BayerBG

    case PixelType_BayerGR16: return PCF_BayerGR
    case PixelType_BayerRG16: return PCF_BayerRG
    case PixelType_BayerGB16: return PCF_BayerGB
    case PixelType_BayerBG16: return PCF_BayerBG
    }
    return PCF_Undefined
}

/// Returns true when the pixel type represents a YUV format.
func IsYUV(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_YUV411packed,
    	 PixelType_YUV422_YUYV_Packed, PixelType_YUV422packed,
	 PixelType_YUV444packed, PixelType_YUV444planar,
     	 PixelType_YUV422planar,
	 PixelType_YUV420planar: return true
    }
    return false
}

/// Returns true when the pixel type represents an RGBA format.
func IsRGBA(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_RGBA8packed: return true
    }
    return false
}

/// Returns true when the pixel type represents an RGB or RGBA format.
func IsRGB(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_RGB8packed, PixelType_RGB10packed,
    	 PixelType_RGB12packed, PixelType_RGB16packed,
	 PixelType_RGB8planar, PixelType_RGB10planar,
    	 PixelType_RGB12planar, PixelType_RGB16planar,
	 PixelType_RGB12V1packed:
	return true
    }
    return IsRGBA(pixelType)
}

/// Returns true when the pixel type represents a BGRA format.
func IsBGRA(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_BGRA8packed: return true
    }
    return false
}

/// Returns true when the pixel type represents a BGR or BGRA format.
func IsBGR(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_BGR8packed, PixelType_BGR10packed, PixelType_BGR12packed,
    	 PixelType_BGR10V1packed, PixelType_BGR10V2packed:
	return true
    }
    return IsBGRA(pixelType)
}

/// Returns true when the pixel type represents a Bayer format.
func IsBayer(pixelType EPixelType) bool {
    return GetPixelColorFilter(pixelType) != PCF_Undefined
}

/// Returns true when a given pixel is monochrome, e.g. PixelType_Mono8 or PixelType_BayerGR8.
func IsMono(pixelType EPixelType) bool {
    return pixelType.IsMono() && !pixelType.IsUndefined()
}

/// Returns true when an image using the given pixel type is monochrome, e.g. PixelType_Mono8.
func IsMonoImage(pixelType EPixelType) bool {
    return IsMono(pixelType) && !IsBayer(pixelType)
}

/// Returns true when an image using the given pixel type is a color image (RGB/BGR/RGBA/BGRA/ etc or Bayer.
func IsColorImage(pixelType EPixelType) bool {
    return pixelType != PixelType_Undefined && 
    	   (IsBayer(pixelType) || ((pixelType & ePixelType_Class_COLOR) != 0))
}

/// Returns true when an the image using the given pixel type has an alpha channel.
func HasAlpha(pixelType EPixelType) bool {
    switch pixelType {
    case PixelType_RGBA8packed, PixelType_BGRA8packed: return true
    }
    return false
}

/// Returns the minimum step size expressed in pixels for extracting an AOI.
func GetPixelIncrementX(pixelType EPixelType) uint32 {
    if IsBayer(pixelType) {
        return 2
    }

    switch pixelType {
    case PixelType_YUV422packed, PixelType_YUV422_YUYV_Packed: return 2
    case PixelType_YUV411packed: return 4
    }
    return 1
}

/// Returns the minimum step size expressed in pixels for extracting an AOI.
func GetPixelIncrementY(pixelType EPixelType) uint32 {
    if IsBayer(pixelType) {
        return 2
    }
    return 1
}

/*!
\brief Returns the pixel types needed for conversion from packed to unpacked image formats using the CImageFormatConverter class.

The following pixel types are supported:
<ul>
<li> source: PixelType_Mono1packed   target: PixelType_Mono8
<li> source: PixelType_Mono2packed   target: PixelType_Mono8
<li> source: PixelType_Mono4packed   target: PixelType_Mono8
<li> source: PixelType_Mono10packed  target: PixelType_Mono16
<li> source: PixelType_Mono10p  target: PixelType_Mono16
<li> source: PixelType_Mono12packed  target: PixelType_Mono16
<li> source: PixelType_Mono12p  target: PixelType_Mono16
<li> source: PixelType_BayerGB12Packed  imposed: PixelType_Mono12packed  target: PixelType_Mono16
<li> source: PixelType_BayerGR12Packed  imposed: PixelType_Mono12packed  target: PixelType_Mono16
<li> source: PixelType_BayerRG12Packed  imposed: PixelType_Mono12packed  target: PixelType_Mono16
<li> source: PixelType_BayerBG12Packed  imposed: PixelType_Mono12packed  target: PixelType_Mono16
<li> source: PixelType_BayerGB10p  imposed: PixelType_Mono10p  target: PixelType_Mono16
<li> source: PixelType_BayerGR10p  imposed: PixelType_Mono10p  target: PixelType_Mono16
<li> source: PixelType_BayerRG10p  imposed: PixelType_Mono10p  target: PixelType_Mono16
<li> source: PixelType_BayerBG10p  imposed: PixelType_Mono10p  target: PixelType_Mono16
<li> source: PixelType_BayerGB12p  imposed: PixelType_Mono12p  target: PixelType_Mono16
<li> source: PixelType_BayerGR12p  imposed: PixelType_Mono12p  target: PixelType_Mono16
<li> source: PixelType_BayerRG12p  imposed: PixelType_Mono12p  target: PixelType_Mono16
<li> source: PixelType_BayerBG12p  imposed: PixelType_Mono12p  target: PixelType_Mono16
</ul>

\param[in] pixelTypeSource The source pixel type.
\param[out] pixelTypeToImpose The pixel type that is used for conversion instead of the source pixel type.
                              Returns \c pixelTypeSource if changing the source pixel type is not needed.
                              Returns PixelType_Undefined if no unpacking is needed.
\param[out] pixelTypeTarget The pixel type to which the image pixel data are converted.
                            Returns PixelType_Undefined if no unpacking is needed.

\return Returns true if the source \c pixelTypeSource is in packed image format and a conversion is possible.
*/
func GetPixelTypesForUnpacking(pixelTypeSource EPixelType,
			       pixelTypeToImpose *EPixelType,
			       pixelTypeTarget *EPixelType) bool {

    *pixelTypeToImpose = pixelTypeSource
    *pixelTypeTarget = PixelType_Undefined

    switch pixelTypeSource {
    case PixelType_Mono1packed, PixelType_Mono2packed, PixelType_Mono4packed:
    	*pixelTypeTarget = PixelType_Mono8
    case PixelType_Mono10packed, PixelType_Mono10p,
    	 PixelType_Mono12packed, PixelType_Mono12p:
    	*pixelTypeTarget = PixelType_Mono16
    case PixelType_BayerGB12Packed, PixelType_BayerGR12Packed,
    	 PixelType_BayerRG12Packed, PixelType_BayerBG12Packed:
        *pixelTypeToImpose = PixelType_Mono12packed
        *pixelTypeTarget = PixelType_Mono16
    case PixelType_BayerGB10p, PixelType_BayerGR10p,
    	 PixelType_BayerRG10p, PixelType_BayerBG10p:
        *pixelTypeToImpose = PixelType_Mono10p
        *pixelTypeTarget = PixelType_Mono16
    case PixelType_BayerGB12p, PixelType_BayerGR12p,
    	 PixelType_BayerRG12p, PixelType_BayerBG12p:
        *pixelTypeToImpose = PixelType_Mono12p
        *pixelTypeTarget = PixelType_Mono16
    }

    return !pixelTypeTarget.IsUndefined()
}
/* EOF */
