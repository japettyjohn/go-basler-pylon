package pylon

var (
SequencerMode = Param{Name: "SequencerMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SequencerModeEnums",
	Group: SequencerControl,
	Description: "Sets whether the sequencer can be used for image acquisition. Applies to: ace"}
SequencerSetActive = Param{Name: "SequencerSetActive",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: SequencerControl,
	Description: "Index number of the currently active sequencer set. Applies to: ace"}
SequencerConfigurationMode = Param{Name: "SequencerConfigurationMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SequencerConfigurationModeEnums",
	Group: SequencerControl,
	Description: "Sets whether the sequencer can be configured. Applies to: ace"}
SequencerSetStart = Param{Name: "SequencerSetStart",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: SequencerControl,
	Description: "Sequencer set that will be used with the first frame start trigger after SequencerMode was set to On. Applies to: ace"}
SequencerSetSelector = Param{Name: "SequencerSetSelector",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: SequencerControl,
	Description: "Sets a sequencer set by its index number. Applies to: ace"}
SequencerSetLoad = Param{Name: "SequencerSetLoad",
	OriginalType: OriginalTypeGenApiICommand,
	Group: SequencerControl,
	Description: "Loads the parameter values of a sequencer set into the active set. Applies to: ace"}
SequencerSetSave = Param{Name: "SequencerSetSave",
	OriginalType: OriginalTypeGenApiICommand,
	Group: SequencerControl,
	Description: "Saves the sequencer parameter values that are currently in the active set. Applies to: ace"}
SequencerPathSelector = Param{Name: "SequencerPathSelector",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: SequencerControl,
	Description: "Sets the sequencer path. Applies to: ace"}
SequencerSetNext = Param{Name: "SequencerSetNext",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: SequencerControl,
	Description: "Next sequencer set to follow after the current one. Applies to: ace"}
SequencerTriggerSource = Param{Name: "SequencerTriggerSource",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SequencerTriggerSourceEnums",
	Group: SequencerControl,
	Description: "Sets the trigger source for sequencer set advance. Applies to: ace"}
SequencerTriggerActivation = Param{Name: "SequencerTriggerActivation",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SequencerTriggerActivationEnums",
	Group: SequencerControl,
	Description: "Sets the effective logical level for sequencer set advance. Applies to: ace"}
/*
SensorWidth = Param{Name: "SensorWidth",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Width of the camera's sensor in pixels. Applies to: ace, dart, pulse"}
SensorHeight = Param{Name: "SensorHeight",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Height of the camera's sensor in pixels. Applies to: ace, dart, pulse"}
WidthMax = Param{Name: "WidthMax",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Maximum allowed width of the region of interest in pixels. Applies to: ace, dart, pulse"}
HeightMax = Param{Name: "HeightMax",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Maximum allowed height of the region of interest in pixels. Applies to: ace, dart, pulse"}
Width = Param{Name: "Width",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Width of the camera's region of interest in pixels. Applies to: ace, dart, pulse"}
Height = Param{Name: "Height",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Height of the camera's region of interest in pixels. Applies to: ace, dart, pulse"}
OffsetX = Param{Name: "OffsetX",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Horizontal offset from the left side of the sensor to the region of interest (in pixels). Applies to: ace, dart, pulse"}
OffsetY = Param{Name: "OffsetY",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Vertical offset from the top of the sensor to the region of interest (in pixels). Applies to: ace, dart, pulse"}
*/
LinePitchEnable = Param{Name: "LinePitchEnable",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ImageFormatControl,
	Description: "Enables the line pitch feature. Applies to: ace"}
LinePitch = Param{Name: "LinePitch",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Number of bytes separating the starting pixels of two consecutive lines. Applies to: ace"}
/*
CenterX = Param{Name: "CenterX",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ImageFormatControl,
	Description: "Enables horizontal centering of the image. Applies to: ace"}
CenterY = Param{Name: "CenterY",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ImageFormatControl,
	Description: "Enables vertical centering of the image. Applies to: ace"}
*/
BinningHorizontalMode = Param{Name: "BinningHorizontalMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BinningHorizontalModeEnums",
	Group: ImageFormatControl,
	Description: "Sets the binning horizontal mode. Applies to: ace, dart, pulse"}
/*
BinningHorizontal = Param{Name: "BinningHorizontal",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Number of adjacent horizontal pixels to be summed. Applies to: ace, dart, pulse"}
*/
BinningVerticalMode = Param{Name: "BinningVerticalMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BinningVerticalModeEnums",
	Group: ImageFormatControl,
	Description: "Sets the binning vertical mode. Applies to: ace, dart, pulse"}
/*
BinningVertical = Param{Name: "BinningVertical",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Number of adjacent vertical pixels to be summed. Applies to: ace, dart, pulse"}
DecimationHorizontal = Param{Name: "DecimationHorizontal",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Horizontal decimation factor. Applies to: ace"}
DecimationVertical = Param{Name: "DecimationVertical",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Vertical decimation factor. Applies to: ace"}
*/
ScalingHorizontal = Param{Name: "ScalingHorizontal",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageFormatControl,
	Description: "Horizontal scaling factor. Applies to: ace"}
ScalingVertical = Param{Name: "ScalingVertical",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageFormatControl,
	Description: "Vertical scaling factor. Applies to: ace"}
/*
ReverseX = Param{Name: "ReverseX",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ImageFormatControl,
	Description: "Enables horizontal mirroring of the image. Applies to: ace, dart, pulse"}
ReverseY = Param{Name: "ReverseY",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ImageFormatControl,
	Description: "Enables vertical mirroring of the image. Applies to: ace, dart, pulse"}
PixelFormat = Param{Name: "PixelFormat",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "PixelFormatEnums",
	Group: ImageFormatControl,
	Description: "Sets the format of the pixel data transmitted by the camera. Applies to: ace, dart, pulse"}
PixelSize = Param{Name: "PixelSize",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "PixelSizeEnums",
	Group: ImageFormatControl,
	Description: "Returns the depth of the pixel values in the image (in bits per pixel). Applies to: ace, dart, pulse"}
PixelColorFilter = Param{Name: "PixelColorFilter",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "PixelColorFilterEnums",
	Group: ImageFormatControl,
	Description: "Returns the alignment of the camera's Bayer filter to the pixels in the acquired images. Applies to: ace, dart, pulse"}
PixelDynamicRangeMin = Param{Name: "PixelDynamicRangeMin",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Minimum possible pixel value that could be transferred from the camera. Applies to: ace, dart, pulse"}
PixelDynamicRangeMax = Param{Name: "PixelDynamicRangeMax",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Maximum possible pixel value that could be transferred from the camera. Applies to: ace, dart, pulse"}
TestImageSelector = Param{Name: "TestImageSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TestImageSelectorEnums",
	Group: ImageFormatControl,
	Description: "Sets the test image to display. Applies to: ace"}
*/
TestImageResetAndHold = Param{Name: "TestImageResetAndHold",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ImageFormatControl,
	Description: "Holds all moving test images at their starting position. Applies to: ace"}
TestPattern = Param{Name: "TestPattern",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TestPatternEnums",
	Group: ImageFormatControl,
	Description: "Selects the type of image test pattern that is generated by the device. Applies to: dart, pulse"}
ROIZoneSelector = Param{Name: "ROIZoneSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ROIZoneSelectorEnums",
	Group: ImageFormatControl,
	Description: "Sets a ROI zone"}
ROIZoneMode = Param{Name: "ROIZoneMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ROIZoneModeEnums",
	Group: ImageFormatControl,
	Description: "Provides for enabling/disabling a ROI zone."}
ROIZoneSize = Param{Name: "ROIZoneSize",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Sets a ROI zone size"}
ROIZoneOffset = Param{Name: "ROIZoneOffset",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageFormatControl,
	Description: "Sets a ROI zone offset"}
/*
GainAuto = Param{Name: "GainAuto",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "GainAutoEnums",
	Group: AnalogControl,
	Description: "Sets the operation mode of the gain auto function. Applies to: ace, dart, pulse"}
GainSelector = Param{Name: "GainSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "GainSelectorEnums",
	Group: AnalogControl,
	Description: "Sets the gain channel or tap to be adjusted. Applies to: ace, dart, pulse"}
*/
Gain = Param{Name: "Gain",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AnalogControl,
	Description: "Value of the currently selected gain control in dB. Applies to: ace, dart, pulse"}
/*
BlackLevelSelector = Param{Name: "BlackLevelSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BlackLevelSelectorEnums",
	Group: AnalogControl,
	Description: "Sets the black level channel or tap to be adjusted. Applies to: ace, dart, pulse"}
*/
BlackLevel = Param{Name: "BlackLevel",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AnalogControl,
	Description: "Value of the currently selected black level channel or tap. Applies to: ace, dart, pulse"}
/*
Gamma = Param{Name: "Gamma",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AnalogControl,
	Description: "Gamma correction value. Applies to: ace, dart, pulse"}
*/
ColorSpace = Param{Name: "ColorSpace",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ColorSpaceEnums",
	Group: AnalogControl,
	Description: "Returns the color space set for image acquisitions. Applies to: ace"}
/*
DigitalShift = Param{Name: "DigitalShift",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AnalogControl,
	Description: "Value set for digital shift. Applies to: ace"}
*/
BslColorSpaceMode = Param{Name: "BslColorSpaceMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BslColorSpaceModeEnums",
	Group: AnalogControl,
	Description: "Sets the color space for image acquisition. Applies to: dart, pulse"}
LightSourcePreset = Param{Name: "LightSourcePreset",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LightSourcePresetEnums",
	Group: ImageQualityControl,
	Description: "Sets the light source preset. Applies to: ace, dart, pulse"}
/*
BalanceWhiteAuto = Param{Name: "BalanceWhiteAuto",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BalanceWhiteAutoEnums",
	Group: ImageQualityControl,
	Description: "Sets the operation mode of the balance white auto function. Applies to: ace, dart, pulse"}
BalanceRatioSelector = Param{Name: "BalanceRatioSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BalanceRatioSelectorEnums",
	Group: ImageQualityControl,
	Description: "Sets the color channel to be adjusted for manual white balance. Applies to: ace, dart, pulse"}
*/
BalanceRatio = Param{Name: "BalanceRatio",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Value of the currently selected balance ratio channel or tap. Applies to: ace, dart, pulse"}
/*
ColorAdjustmentSelector = Param{Name: "ColorAdjustmentSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ColorAdjustmentSelectorEnums",
	Group: ImageQualityControl,
	Description: "Sets the color for color adjustment. Applies to: ace"}
ColorAdjustmentHue = Param{Name: "ColorAdjustmentHue",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Hue adjustment value for the currently selected color. Applies to: ace"}
ColorAdjustmentSaturation = Param{Name: "ColorAdjustmentSaturation",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Saturation adjustment value for the currently selected color. Applies to: ace"}
ColorTransformationSelector = Param{Name: "ColorTransformationSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ColorTransformationSelectorEnums",
	Group: ImageQualityControl,
	Description: "Sets the type of color transformation that will be performed. Applies to: ace"}
ColorTransformationValueSelector = Param{Name: "ColorTransformationValueSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ColorTransformationValueSelectorEnums",
	Group: ImageQualityControl,
	Description: "Sets the element to be entered in the color transformation matrix. Applies to: ace"}
ColorTransformationValue = Param{Name: "ColorTransformationValue",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Transformation value for the selected element in the color transformation matrix. Applies to: ace"}
*/
BslContrastMode = Param{Name: "BslContrastMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BslContrastModeEnums",
	Group: ImageQualityControl,
	Description: "Sets the contrast mode. Applies to: dart, pulse"}
BslBrightness = Param{Name: "BslBrightness",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Amount of brightness to be applied. Applies to: dart, pulse"}
BslContrast = Param{Name: "BslContrast",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Amount of contrast to be applied. Applies to: dart, pulse"}
DefectPixelCorrectionMode = Param{Name: "DefectPixelCorrectionMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "DefectPixelCorrectionModeEnums",
	Group: ImageQualityControl,
	Description: "Identifies outlier pixels and adjusts their intensity value. Applies to: dart, pulse"}
BslHue = Param{Name: "BslHue",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Hue shift to be applied."}
BslHueValue = Param{Name: "BslHueValue",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ImageQualityControl,
	Description: "Hue shift to be applied. Applies to: dart, pulse"}
BslSaturation = Param{Name: "BslSaturation",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Saturation to be applied."}
BslSaturationValue = Param{Name: "BslSaturationValue",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ImageQualityControl,
	Description: "Amount of saturation to be applied. Applies to: dart, pulse"}
DemosaicingMode = Param{Name: "DemosaicingMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "DemosaicingModeEnums",
	Group: PGIControl,
	Description: "Sets the demosaicing mode. Applies to: ace"}
PgiMode = Param{Name: "PgiMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "PgiModeEnums",
	Group: PGIControl,
	Description: "Enables Basler PGI image optimizations."}
NoiseReduction = Param{Name: "NoiseReduction",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: PGIControl,
	Description: "Amount of noise reduction to apply. Applies to: ace"}
SharpnessEnhancement = Param{Name: "SharpnessEnhancement",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: PGIControl,
	Description: "Amount of sharpening to apply. Applies to: ace, dart, pulse"}
/*
AcquisitionMode = Param{Name: "AcquisitionMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "AcquisitionModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the image acquisition mode. Applies to: ace, dart, pulse"}
AcquisitionStart = Param{Name: "AcquisitionStart",
	OriginalType: OriginalTypeGenApiICommand,
	Group: AcquisitionControl,
	Description: "Starts the acquisition of images. Applies to: ace, dart, pulse"}
AcquisitionStop = Param{Name: "AcquisitionStop",
	OriginalType: OriginalTypeGenApiICommand,
	Group: AcquisitionControl,
	Description: "Stops the acquisition of images. Applies to: ace, dart, pulse"}
ShutterMode = Param{Name: "ShutterMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ShutterModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the shutter mode. Applies to: ace"}
ExposureAuto = Param{Name: "ExposureAuto",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ExposureAutoEnums",
	Group: AcquisitionControl,
	Description: "Sets the operation mode of the exposure auto function. Applies to: ace, dart, pulse"}
ExposureMode = Param{Name: "ExposureMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ExposureModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the exposure mode. Applies to: ace, dart, pulse"}
*/
ExposureTime = Param{Name: "ExposureTime",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AcquisitionControl,
	Description: "Exposure time of the camera in microseconds. Applies to: ace, dart, pulse"}
ExposureOverlapTimeMode = Param{Name: "ExposureOverlapTimeMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ExposureOverlapTimeModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the exposure overlap time mode. Applies to: ace"}
ExposureOverlapTimeMax = Param{Name: "ExposureOverlapTimeMax",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AcquisitionControl,
	Description: "Maximum overlap of the sensor exposure with sensor readout in TriggerWidth exposure mode (in microseconds). Applies to: ace"}
SensorReadoutMode = Param{Name: "SensorReadoutMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SensorReadoutModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the sensor readout mode. Applies to: ace"}
AcquisitionBurstFrameCount = Param{Name: "AcquisitionBurstFrameCount",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AcquisitionControl,
	Description: "Number of frames to acquire for each FrameBurstStart trigger. Applies to: ace"}
/*
TriggerSelector = Param{Name: "TriggerSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TriggerSelectorEnums",
	Group: AcquisitionControl,
	Description: "Sets the trigger type to be configured. Applies to: ace, dart, pulse"}
TriggerMode = Param{Name: "TriggerMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TriggerModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the mode for the currently selected trigger. Applies to: ace, dart, pulse"}
TriggerSoftware = Param{Name: "TriggerSoftware",
	OriginalType: OriginalTypeGenApiICommand,
	Group: AcquisitionControl,
	Description: "Generates a software trigger signal. Applies to: ace, dart, pulse"}
TriggerSource = Param{Name: "TriggerSource",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TriggerSourceEnums",
	Group: AcquisitionControl,
	Description: "Sets the signal source for the selected trigger. Applies to: ace, dart, pulse"}
TriggerActivation = Param{Name: "TriggerActivation",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TriggerActivationEnums",
	Group: AcquisitionControl,
	Description: "Sets the signal transition that activates the selected trigger. Applies to: ace, dart, pulse"}
*/
TriggerDelay = Param{Name: "TriggerDelay",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AcquisitionControl,
	Description: "Trigger delay time in microseconds. Applies to: ace"}
/*
AcquisitionFrameRateEnable = Param{Name: "AcquisitionFrameRateEnable",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: AcquisitionControl,
	Description: "Enables setting the camera's acquisition frame rate to a specified value. Applies to: ace"}
*/
AcquisitionFrameRate = Param{Name: "AcquisitionFrameRate",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AcquisitionControl,
	Description: "Acquisition frame rate of the camera in frames per second. Applies to: ace, dart, pulse"}
ResultingFrameRate = Param{Name: "ResultingFrameRate",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AcquisitionControl,
	Description: "Maximum allowed frame acquisition rate. Applies to: ace, dart, pulse"}
SensorReadoutTime = Param{Name: "SensorReadoutTime",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AcquisitionControl,
	Description: "Sensor readout time given the current settings. Applies to: ace"}
/*
AcquisitionStatusSelector = Param{Name: "AcquisitionStatusSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "AcquisitionStatusSelectorEnums",
	Group: AcquisitionControl,
	Description: "Sets the acquisition status to be checked. Applies to: ace"}
AcquisitionStatus = Param{Name: "AcquisitionStatus",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: AcquisitionControl,
	Description: "Indicates the status (true or false) of the currently selected acquisition signal. Applies to: ace"}
*/
SensorShutterMode = Param{Name: "SensorShutterMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SensorShutterModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the shutter mode of the device. Applies to: dart, pulse"}
OverlapMode = Param{Name: "OverlapMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "OverlapModeEnums",
	Group: AcquisitionControl,
	Description: "Configures overlapping exposure and image readout. Applies to: dart, pulse"}
BslImmediateTriggerMode = Param{Name: "BslImmediateTriggerMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BslImmediateTriggerModeEnums",
	Group: AcquisitionControl,
	Description: "Sets the immediate trigger mode. Applies to: dart, pulse"}
AutoTargetBrightness = Param{Name: "AutoTargetBrightness",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AutoFunctionControl,
	Description: "Target average brightness for the gain auto function and the exposure auto function. Applies to: ace, dart, pulse"}
/*
AutoFunctionProfile = Param{Name: "AutoFunctionProfile",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "AutoFunctionProfileEnums",
	Group: AutoFunctionControl,
	Description: "Sets how gain and exposure time will be balanced when the device is making automatic adjustments. Applies to: ace, dart, pulse"}
*/
AutoGainLowerLimit = Param{Name: "AutoGainLowerLimit",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AutoFunctionControl,
	Description: "Lower limit for the Gain parameter when the gain auto function is active. Applies to: ace, dart, pulse"}
AutoGainUpperLimit = Param{Name: "AutoGainUpperLimit",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AutoFunctionControl,
	Description: "Upper limit for the Gain parameter when the gain auto function is active. Applies to: ace, dart, pulse"}
AutoExposureTimeLowerLimit = Param{Name: "AutoExposureTimeLowerLimit",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AutoFunctionControl,
	Description: "Lower limit for the ExposureTime parameter when the exposure auto function is active. Applies to: ace, dart, pulse"}
AutoExposureTimeUpperLimit = Param{Name: "AutoExposureTimeUpperLimit",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AutoFunctionControl,
	Description: "Upper limit for the ExposureTime parameter when the exposure auto function is active. Applies to: ace, dart, pulse"}
AutoBacklightCompensation = Param{Name: "AutoBacklightCompensation",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: AutoFunctionControl,
	Description: "Sets the backlight compensation. Applies to: dart, pulse"}
AutoFunctionROISelector = Param{Name: "AutoFunctionROISelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "AutoFunctionROISelectorEnums",
	Group: AutoFunctionROIControl,
	Description: "Sets which auto function ROI can be adjusted. Applies to: ace, dart, pulse"}
AutoFunctionROIWidth = Param{Name: "AutoFunctionROIWidth",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionROIControl,
	Description: "Width of the auto function ROI (in pixels). Applies to: ace, dart, pulse"}
AutoFunctionROIHeight = Param{Name: "AutoFunctionROIHeight",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionROIControl,
	Description: "Height of the auto function ROI (in pixels). Applies to: ace, dart, pulse"}
AutoFunctionROIOffsetX = Param{Name: "AutoFunctionROIOffsetX",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionROIControl,
	Description: "Horizontal offset from the left side of the sensor to the auto function ROI (in pixels). Applies to: ace, dart, pulse"}
AutoFunctionROIOffsetY = Param{Name: "AutoFunctionROIOffsetY",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionROIControl,
	Description: "Vertical offset from the top of the sensor to the auto function ROI (in pixels). Applies to: ace, dart, pulse"}
AutoFunctionROIUseBrightness = Param{Name: "AutoFunctionROIUseBrightness",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: AutoFunctionROIControl,
	Description: "Assigns the Gain Auto and the Exposure Auto functions to the currently selected auto function ROI. Applies to: ace, dart, pulse"}
AutoFunctionROIUseWhiteBalance = Param{Name: "AutoFunctionROIUseWhiteBalance",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: AutoFunctionROIControl,
	Description: "Assigns the Balance White auto function to the currently selected auto function ROI. Applies to: ace, dart, pulse"}
/*
AutoFunctionAOISelector = Param{Name: "AutoFunctionAOISelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "AutoFunctionAOISelectorEnums",
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Sets which auto function AOI can be adjusted."}
AutoFunctionAOIWidth = Param{Name: "AutoFunctionAOIWidth",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Width of the auto function AOI (in pixels)."}
AutoFunctionAOIHeight = Param{Name: "AutoFunctionAOIHeight",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Height of the auto function AOI (in pixels)."}
AutoFunctionAOIOffsetX = Param{Name: "AutoFunctionAOIOffsetX",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Horizontal offset from the left side of the sensor to the auto function AOI (in pixels)."}
AutoFunctionAOIOffsetY = Param{Name: "AutoFunctionAOIOffsetY",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Vertical offset from the top of the sensor to the auto function AOI (in pixels)."}
*/
AutoFunctionAOIUseBrightness = Param{Name: "AutoFunctionAOIUseBrightness",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Assigns the Gain Auto and the Exposure Auto functions to the currently selected auto function AOI."}
AutoFunctionAOIUseWhiteBalance = Param{Name: "AutoFunctionAOIUseWhiteBalance",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: AutoFunctionAOIControl,
	Description: "This feature has been \b deprecated. Assigns the Balance White auto function to the currently selected auto function AOI."}
/*
LUTSelector = Param{Name: "LUTSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LUTSelectorEnums",
	Group: LUTControl,
	Description: "Sets the lookup table (LUT) to be configured. Applies to: ace"}
LUTEnable = Param{Name: "LUTEnable",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: LUTControl,
	Description: "Enables the selected lookup table (LUT). Applies to: ace"}
LUTIndex = Param{Name: "LUTIndex",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: LUTControl,
	Description: "Index of the LUT element to access. Applies to: ace"}
LUTValue = Param{Name: "LUTValue",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: LUTControl,
	Description: "Value of the LUT element at the LUT index position. Applies to: ace"}
LUTValueAll = Param{Name: "LUTValueAll",
	OriginalType: OriginalTypeGenApiIRegister,
	Group: LUTControl,
	Description: "A single register that lets you access all LUT coefficients. Applies to: ace"}
LineSelector = Param{Name: "LineSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LineSelectorEnums",
	Group: DigitalIOControl,
	Description: "Sets the I/O line to be configured. Applies to: ace, dart"}
LineMode = Param{Name: "LineMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LineModeEnums",
	Group: DigitalIOControl,
	Description: "Sets the mode for the selected line. Applies to: ace, dart"}
LineFormat = Param{Name: "LineFormat",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LineFormatEnums",
	Group: DigitalIOControl,
	Description: "Returns the electrical configuration of the currently selected line. Applies to: ace, dart"}
LineLogic = Param{Name: "LineLogic",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LineLogicEnums",
	Group: DigitalIOControl,
	Description: "Returns the line logic of the currently selected line. Applies to: ace"}
LineSource = Param{Name: "LineSource",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "LineSourceEnums",
	Group: DigitalIOControl,
	Description: "Sets the source signal for the currently selected line. Applies to: ace, dart"}
LineInverter = Param{Name: "LineInverter",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: DigitalIOControl,
	Description: "Enables the signal inverter function for the currently selected input or output line. Applies to: ace, dart"}
*/
LineDebouncerTime = Param{Name: "LineDebouncerTime",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: DigitalIOControl,
	Description: "Value of the selected line debouncer time in microseconds. Applies to: ace, dart"}
LineMinimumOutputPulseWidth = Param{Name: "LineMinimumOutputPulseWidth",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: DigitalIOControl,
	Description: "Value for the minimum signal width of an output signal (in microseconds) . Applies to: ace"}
LineOverloadStatus = Param{Name: "LineOverloadStatus",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: DigitalIOControl,
	Description: "Indicates whether an overload condition was detected on the selected line. Applies to: ace"}
LineOverloadReset = Param{Name: "LineOverloadReset",
	OriginalType: OriginalTypeGenApiICommand,
	Group: DigitalIOControl,
	Description: "Resets the overload status of the selected line. Applies to: ace"}
/*
LineStatus = Param{Name: "LineStatus",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: DigitalIOControl,
	Description: "Indicates the current logical state of the selected line. Applies to: ace, dart"}
LineStatusAll = Param{Name: "LineStatusAll",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DigitalIOControl,
	Description: "A single bit field indicating the current logical state of all available line signals at time of polling. Applies to: ace, dart"}
UserOutputSelector = Param{Name: "UserOutputSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "UserOutputSelectorEnums",
	Group: DigitalIOControl,
	Description: "Sets the user settable output signal to be configured. Applies to: ace, dart"}
UserOutputValue = Param{Name: "UserOutputValue",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: DigitalIOControl,
	Description: "Enables the selected user settable output line. Applies to: ace, dart"}
UserOutputValueAll = Param{Name: "UserOutputValueAll",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DigitalIOControl,
	Description: "A single bit field that sets the state of all user settable output signals in one access. Applies to: ace"}
*/
SoftwareSignalSelector = Param{Name: "SoftwareSignalSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "SoftwareSignalSelectorEnums",
	Group: SoftwareSignalControl,
	Description: "Sets the software signal to control. Applies to: ace"}
SoftwareSignalPulse = Param{Name: "SoftwareSignalPulse",
	OriginalType: OriginalTypeGenApiICommand,
	Group: SoftwareSignalControl,
	Description: "Generates a signal that can be used as a software trigger. Applies to: ace"}
/*
TimerSelector = Param{Name: "TimerSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TimerSelectorEnums",
	Group: CounterAndTimerControl,
	Description: "Sets the timer to be configured. Applies to: ace"}
*/
TimerDuration = Param{Name: "TimerDuration",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: CounterAndTimerControl,
	Description: "Duration of the currently selected timer in microseconds. Applies to: ace"}
TimerDelay = Param{Name: "TimerDelay",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: CounterAndTimerControl,
	Description: "Delay of the currently selected timer in microseconds. Applies to: ace"}
/*
TimerTriggerSource = Param{Name: "TimerTriggerSource",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TimerTriggerSourceEnums",
	Group: CounterAndTimerControl,
	Description: "Sets the internal camera signal used to trigger the selected timer. Applies to: ace"}
CounterSelector = Param{Name: "CounterSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "CounterSelectorEnums",
	Group: CounterAndTimerControl,
	Description: "Sets the counter to be configured. Applies to: ace"}
CounterEventSource = Param{Name: "CounterEventSource",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "CounterEventSourceEnums",
	Group: CounterAndTimerControl,
	Description: "Sets the event that increments the currently selected counter. Applies to: ace"}
CounterResetSource = Param{Name: "CounterResetSource",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "CounterResetSourceEnums",
	Group: CounterAndTimerControl,
	Description: "Sets the source signal that can reset the currently selected counter. Applies to: ace"}
*/
CounterResetActivation = Param{Name: "CounterResetActivation",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "CounterResetActivationEnums",
	Group: CounterAndTimerControl,
	Description: "Sets the signal on which the counter will be reset. Applies to: ace"}
/*
CounterReset = Param{Name: "CounterReset",
	OriginalType: OriginalTypeGenApiICommand,
	Group: CounterAndTimerControl,
	Description: "Immediately resets the selected counter. Applies to: ace"}
*/
CounterDuration = Param{Name: "CounterDuration",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: CounterAndTimerControl,
	Description: "Duration (or number of events) before the CounterEnd event is generated. Applies to: ace"}
/*
UserSetSelector = Param{Name: "UserSetSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "UserSetSelectorEnums",
	Group: UserSetControl,
	Description: "Sets the user set or the factory set to load, save or configure. Applies to: ace, dart, pulse"}
UserSetLoad = Param{Name: "UserSetLoad",
	OriginalType: OriginalTypeGenApiICommand,
	Group: UserSetControl,
	Description: "Loads the selected set into the camera's volatile memory and makes it the active configuration set. Applies to: ace, dart, pulse"}
UserSetSave = Param{Name: "UserSetSave",
	OriginalType: OriginalTypeGenApiICommand,
	Group: UserSetControl,
	Description: "Saves the current active set into the selected user set. Applies to: ace, dart, pulse"}
*/
UserSetDefault = Param{Name: "UserSetDefault",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "UserSetDefaultEnums",
	Group: UserSetControl,
	Description: "Sets the user set or the factory set to be used as the startup set. Applies to: ace, dart, pulse"}
/*
ChunkModeActive = Param{Name: "ChunkModeActive",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ChunkDataControl,
	Description: "Enables the chunk mode. Applies to: ace"}
ChunkSelector = Param{Name: "ChunkSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ChunkSelectorEnums",
	Group: ChunkDataControl,
	Description: "Sets the chunk to be enabled. Applies to: ace"}
ChunkEnable = Param{Name: "ChunkEnable",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ChunkDataControl,
	Description: "Enables the inclusion of the currently selected chunk in the payload data. Applies to: ace"}
*/
ChunkGainSelector = Param{Name: "ChunkGainSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ChunkGainSelectorEnums",
	Group: ChunkDataControl,
	Description: "Sets which gain channel to retrieve chunk data from. Applies to: ace"}
ChunkGain = Param{Name: "ChunkGain",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ChunkDataControl,
	Description: "Gain used to acquire the image. Applies to: ace"}
/*
ChunkExposureTime = Param{Name: "ChunkExposureTime",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: ChunkDataControl,
	Description: "Exposure time used to acquire the image. Applies to: ace"}
ChunkTimestamp = Param{Name: "ChunkTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ChunkDataControl,
	Description: "Value of the timestamp when the image was acquired. Applies to: ace"}
ChunkLineStatusAll = Param{Name: "ChunkLineStatusAll",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ChunkDataControl,
	Description: "A bit field that indicates the status of all of the camera's input and output lines when the image was acquired. Applies to: ace"}
*/
ChunkCounterSelector = Param{Name: "ChunkCounterSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ChunkCounterSelectorEnums",
	Group: ChunkDataControl,
	Description: "Sets which counter to retrieve chunk data from. Applies to: ace"}
ChunkCounterValue = Param{Name: "ChunkCounterValue",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ChunkDataControl,
	Description: "Value of the selected chunk counter. Applies to: ace"}
ChunkSequencerSetActive = Param{Name: "ChunkSequencerSetActive",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ChunkDataControl,
	Description: "Index of the active sequencer set. Applies to: ace"}
/*
ChunkPayloadCRC16 = Param{Name: "ChunkPayloadCRC16",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ChunkDataControl,
	Description: "CRC checksum of the acquired image. Applies to: ace"}
EventSelector = Param{Name: "EventSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "EventSelectorEnums",
	Group: EventControl,
	Description: "Sets the event notification to be enabled. Applies to: ace"}
EventNotification = Param{Name: "EventNotification",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "EventNotificationEnums",
	Group: EventControl,
	Description: "Enables event notifications for the currently selected event. Applies to: ace"}
*/
TriggerEventTest = Param{Name: "TriggerEventTest",
	OriginalType: OriginalTypeGenApiICommand,
	Group: EventControl,
	Description: "Generates an event test signal."}
EventExposureEnd = Param{Name: "EventExposureEnd",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventExposureEndData,
	Description: "Unique identifier of the exposure end event. Applies to: ace"}
EventExposureEndTimestamp = Param{Name: "EventExposureEndTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventExposureEndData,
	Description: "Time stamp of the exposure end event. Applies to: ace"}
EventExposureEndFrameID = Param{Name: "EventExposureEndFrameID",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventExposureEndData,
	Description: "Frame ID of the exposure end event. Applies to: ace"}
EventFrameStart = Param{Name: "EventFrameStart",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartData,
	Description: "Unique identifier of the frame start event. Applies to: ace"}
EventFrameStartTimestamp = Param{Name: "EventFrameStartTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartData,
	Description: "Time stamp of the frame start event. Applies to: ace"}
EventFrameStartFrameID = Param{Name: "EventFrameStartFrameID",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartData,
	Description: "Frame ID of the frame start event. Applies to: ace"}
EventFrameBurstStart = Param{Name: "EventFrameBurstStart",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartData,
	Description: "Unique identifier of the frame burst start event. Applies to: ace"}
EventFrameBurstStartTimestamp = Param{Name: "EventFrameBurstStartTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartData,
	Description: "Time stamp of the frame burst start event. Applies to: ace"}
EventFrameBurstStartFrameID = Param{Name: "EventFrameBurstStartFrameID",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartData,
	Description: "Frame ID of the frame burst start event. Applies to: ace"}
EventFrameStartOvertrigger = Param{Name: "EventFrameStartOvertrigger",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartOvertriggerData,
	Description: "Unique identifier of the frame start overtrigger event. Applies to: ace"}
EventFrameStartOvertriggerTimestamp = Param{Name: "EventFrameStartOvertriggerTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartOvertriggerData,
	Description: "Time stamp of the frame start overtrigger event. Applies to: ace"}
EventFrameStartOvertriggerFrameID = Param{Name: "EventFrameStartOvertriggerFrameID",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartOvertriggerData,
	Description: "Frame ID of the frame start overtrigger event. Applies to: ace"}
EventFrameBurstStartOvertrigger = Param{Name: "EventFrameBurstStartOvertrigger",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartOvertriggerData,
	Description: "Unique identifier of the frame burst start overtrigger event. Applies to: ace"}
EventFrameBurstStartOvertriggerTimestamp = Param{Name: "EventFrameBurstStartOvertriggerTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartOvertriggerData,
	Description: "Time stamp of the frame burst start overtrigger event. Applies to: ace"}
EventFrameBurstStartOvertriggerFrameID = Param{Name: "EventFrameBurstStartOvertriggerFrameID",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartOvertriggerData,
	Description: "Frame ID of the frame burst start overtrigger event. Applies to: ace"}
EventTest = Param{Name: "EventTest",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventTestData,
	Description: "Unique identifier of the test event."}
EventTestTimestamp = Param{Name: "EventTestTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventTestData,
	Description: "Time stamp of the test event."}
EventCriticalTemperature = Param{Name: "EventCriticalTemperature",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventCriticalTemperatureData,
	Description: "Unique identifier of the critical temperature event. Applies to: ace"}
EventCriticalTemperatureTimestamp = Param{Name: "EventCriticalTemperatureTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventCriticalTemperatureData,
	Description: "Time stamp of the crititical temperature event. Applies to: ace"}
EventOverTemperature = Param{Name: "EventOverTemperature",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventOverTemperatureData,
	Description: "Unique identifier of the over temperature event. Applies to: ace"}
EventOverTemperatureTimestamp = Param{Name: "EventOverTemperatureTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventOverTemperatureData,
	Description: "Time stamp of the over temperature event. Applies to: ace"}
EventFrameStartWait = Param{Name: "EventFrameStartWait",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartWaitData,
	Description: "Unique identifier of the frame start wait event. Applies to: ace"}
EventFrameStartWaitTimestamp = Param{Name: "EventFrameStartWaitTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameStartWaitData,
	Description: "Time stamp of the frame start wait event. Applies to: ace"}
EventFrameBurstStartWait = Param{Name: "EventFrameBurstStartWait",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartWaitData,
	Description: "Unique identifier of the frame burst start wait event. Applies to: ace"}
EventFrameBurstStartWaitTimestamp = Param{Name: "EventFrameBurstStartWaitTimestamp",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: EventFrameBurstStartWaitData,
	Description: "Time stamp of the frame brust start wait event. Applies to: ace"}
/*
PayloadSize = Param{Name: "PayloadSize",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Size of the payload in bytes. Applies to: ace, dart, pulse"}
*/
BslUSBSpeedMode = Param{Name: "BslUSBSpeedMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "BslUSBSpeedModeEnums",
	Group: TransportLayerControl,
	Description: "Returns the speed mode of the USB port. Applies to: ace, dart, pulse"}
SIPayloadTransferSize = Param{Name: "SIPayloadTransferSize",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Expected size of a single payload transfer. Applies to: ace"}
SIPayloadTransferCount = Param{Name: "SIPayloadTransferCount",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Expected number of payload transfers. Applies to: ace"}
SIPayloadFinalTransfer1Size = Param{Name: "SIPayloadFinalTransfer1Size",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Size of the first final payload transfer. Applies to: ace"}
SIPayloadFinalTransfer2Size = Param{Name: "SIPayloadFinalTransfer2Size",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Size of the second final payload transfer. Applies to: ace"}
TestPendingAck = Param{Name: "TestPendingAck",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Test pending acknowledging time in milliseconds. Applies to: dart, pulse"}
PayloadTransferBlockDelay = Param{Name: "PayloadTransferBlockDelay",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: ""}
PayloadTransferSize = Param{Name: "PayloadTransferSize",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Payload Transfer Size. Applies to: dart, pulse"}
PayloadTransferCount = Param{Name: "PayloadTransferCount",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Payload Transfer Count. Applies to: dart, pulse"}
PayloadFinalTransfer1Size = Param{Name: "PayloadFinalTransfer1Size",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Payload Final Transfer 1 Size. Applies to: dart, pulse"}
PayloadFinalTransfer2Size = Param{Name: "PayloadFinalTransfer2Size",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: TransportLayerControl,
	Description: "Payload Final Transfer 2 Size. Applies to: dart, pulse"}
/*
DeviceVendorName = Param{Name: "DeviceVendorName",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "Name of the device's vendor. Applies to: ace, dart, pulse"}
DeviceModelName = Param{Name: "DeviceModelName",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "Model name of the device. Applies to: ace, dart, pulse"}
DeviceManufacturerInfo = Param{Name: "DeviceManufacturerInfo",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "Additional information from the vendor about the camera. Applies to: ace, dart, pulse"}
DeviceVersion = Param{Name: "DeviceVersion",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "Version of the device. Applies to: ace, dart, pulse"}
DeviceFirmwareVersion = Param{Name: "DeviceFirmwareVersion",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "Version of the device's firmware. Applies to: ace, dart, pulse"}
*/
DeviceSerialNumber = Param{Name: "DeviceSerialNumber",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "Serial number of the device. Applies to: ace, dart, pulse"}
/*
DeviceUserID = Param{Name: "DeviceUserID",
	OriginalType: OriginalTypeGenApiIString,
	Group: DeviceControl,
	Description: "User-settable ID of the device. Applies to: ace, dart, pulse"}
DeviceScanType = Param{Name: "DeviceScanType",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "DeviceScanTypeEnums",
	Group: DeviceControl,
	Description: "Returns the scan type of the device's sensor (area or line scan). Applies to: ace, dart, pulse"}
*/
TimestampLatch = Param{Name: "TimestampLatch",
	OriginalType: OriginalTypeGenApiICommand,
	Group: DeviceControl,
	Description: "Latches the current timestamp counter and stores its value in TimestampLatchValue. Applies to: ace"}
TimestampLatchValue = Param{Name: "TimestampLatchValue",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Latched value of the timestamp counter. Applies to: ace"}
DeviceLinkSelector = Param{Name: "DeviceLinkSelector",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Sets which link of the device to control. Applies to: ace, dart, pulse"}
DeviceLinkSpeed = Param{Name: "DeviceLinkSpeed",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Speed of transmission negotiated on the selected link. Applies to: ace, dart, pulse"}
DeviceLinkThroughputLimitMode = Param{Name: "DeviceLinkThroughputLimitMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "DeviceLinkThroughputLimitModeEnums",
	Group: DeviceControl,
	Description: "Controls if the device link throughput limit is active. Applies to: ace, dart, pulse"}
DeviceLinkThroughputLimit = Param{Name: "DeviceLinkThroughputLimit",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Value set to limit the maximum bandwidth of the data that will be streamed out by the device (in bytes per second). Applies to: ace, dart, pulse"}
DeviceLinkCurrentThroughput = Param{Name: "DeviceLinkCurrentThroughput",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Actual bandwidth the camera will use. Applies to: ace"}
DeviceTemperatureSelector = Param{Name: "DeviceTemperatureSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "DeviceTemperatureSelectorEnums",
	Group: DeviceControl,
	Description: "Sets the location within the device where the temperature will be measured. Applies to: ace"}
DeviceTemperature = Param{Name: "DeviceTemperature",
	OriginalType: OriginalTypeGenApiIFloat,
	Group: DeviceControl,
	Description: "Temperature of the selected location within the device (in degrees centigrade). Applies to: ace"}
/*
TemperatureState = Param{Name: "TemperatureState",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "TemperatureStateEnums",
	Group: DeviceControl,
	Description: "Returns the temperature state. Applies to: ace"}
DeviceReset = Param{Name: "DeviceReset",
	OriginalType: OriginalTypeGenApiICommand,
	Group: DeviceControl,
	Description: "Immediately resets and reboots the device. Applies to: ace, dart, pulse"}
*/
DeviceSFNCVersionMajor = Param{Name: "DeviceSFNCVersionMajor",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Major version number of the SFNC specification that the device is compatible with. Applies to: ace, dart, pulse"}
DeviceSFNCVersionMinor = Param{Name: "DeviceSFNCVersionMinor",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Minor version number of the SFNC specification that the device is compatible with. Applies to: ace, dart, pulse"}
DeviceSFNCVersionSubMinor = Param{Name: "DeviceSFNCVersionSubMinor",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: DeviceControl,
	Description: "Subminor version number of the SFNC specification that the device is compatible with. Applies to: ace, dart, pulse"}
DeviceIndicatorMode = Param{Name: "DeviceIndicatorMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "DeviceIndicatorModeEnums",
	Group: DeviceControl,
	Description: "Controls the behavior of the indicators (such as LEDs) showing the status of the device. Applies to: dart"}
/*
DeviceRegistersStreamingStart = Param{Name: "DeviceRegistersStreamingStart",
	OriginalType: OriginalTypeGenApiICommand,
	Group: DeviceControl,
	Description: "Prepare the device for registers streaming without checking for consistency. Applies to: dart, pulse"}
DeviceRegistersStreamingEnd = Param{Name: "DeviceRegistersStreamingEnd",
	OriginalType: OriginalTypeGenApiICommand,
	Group: DeviceControl,
	Description: "Announce the end of registers streaming. Applies to: dart, pulse"}
UserDefinedValueSelector = Param{Name: "UserDefinedValueSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "UserDefinedValueSelectorEnums",
	Group: UserDefinedValueControl,
	Description: "Sets the user-defined value to set or read. Applies to: ace"}
UserDefinedValue = Param{Name: "UserDefinedValue",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: UserDefinedValueControl,
	Description: "A user defined value. Applies to: ace"}
*/
RemoveParameterLimitSelector = Param{Name: "RemoveParameterLimitSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "RemoveParameterLimitSelectorEnums",
	Group: RemoveParameterLimitControl,
	Description: "Sets the parameter whose factory limits should be removed. Applies to: ace"}
RemoveParameterLimit = Param{Name: "RemoveParameterLimit",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: RemoveParameterLimitControl,
	Description: "Removes the factory limit of the selected parameter. Applies to: ace"}
/*
ExpertFeatureAccessSelector = Param{Name: "ExpertFeatureAccessSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "ExpertFeatureAccessSelectorEnums",
	Group: ExpertFeatureAccess,
	Description: "Sets the expert feature to be configured. Applies to: ace"}
ExpertFeatureAccessKey = Param{Name: "ExpertFeatureAccessKey",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: ExpertFeatureAccess,
	Description: "Key to access the selected expert feature. Applies to: ace"}
ExpertFeatureEnable = Param{Name: "ExpertFeatureEnable",
	OriginalType: OriginalTypeGenApiIBoolean,
	Group: ExpertFeatureAccess,
	Description: "Enables the currently selected expert feature. Applies to: ace"}
*/
/*
FileSelector = Param{Name: "FileSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "FileSelectorEnums",
	Group: FileAccessControl,
	Description: "Sets the target file in the device. Applies to: ace"}
FileOperationSelector = Param{Name: "FileOperationSelector",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "FileOperationSelectorEnums",
	Group: FileAccessControl,
	Description: "Sets the target operation for the currently selected file. Applies to: ace"}
FileOpenMode = Param{Name: "FileOpenMode",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "FileOpenModeEnums",
	Group: FileAccessControl,
	Description: "Sets the access mode in which a file is opened in the device. Applies to: ace"}
FileAccessBuffer = Param{Name: "FileAccessBuffer",
	OriginalType: OriginalTypeGenApiIRegister,
	Group: FileAccessControl,
	Description: "Access buffer for file operations. Applies to: ace"}
FileAccessOffset = Param{Name: "FileAccessOffset",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: FileAccessControl,
	Description: "File access offset. Applies to: ace"}
FileAccessLength = Param{Name: "FileAccessLength",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: FileAccessControl,
	Description: "File access length. Applies to: ace"}
FileOperationStatus = Param{Name: "FileOperationStatus",
	OriginalType: OriginalTypeGenApiIEnumerationT,
	OriginalEnumType: "FileOperationStatusEnums",
	Group: FileAccessControl,
	Description: "Returns the file operation execution status. Applies to: ace"}
FileOperationResult = Param{Name: "FileOperationResult",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: FileAccessControl,
	Description: "File operation result. Applies to: ace"}
FileSize = Param{Name: "FileSize",
	OriginalType: OriginalTypeGenApiIInteger,
	Group: FileAccessControl,
	Description: "Size of the currently selected file in bytes. Applies to: ace"}
FileOperationExecute = Param{Name: "FileOperationExecute",
	OriginalType: OriginalTypeGenApiICommand,
	Group: FileAccessControl,
	Description: "Executes the operation selected by FileOperationSelector. Applies to: ace"}
*/
)

var UsbParams = ParamList{
		AcquisitionBurstFrameCount,
		AcquisitionFrameRate,
		AcquisitionFrameRateEnable,
		AcquisitionMode,
		AcquisitionStart,
		AcquisitionStatus,
		AcquisitionStatusSelector,
		AcquisitionStop,
		AutoBacklightCompensation,
		AutoExposureTimeLowerLimit,
		AutoExposureTimeUpperLimit,
		AutoFunctionAOIHeight,
		AutoFunctionAOIOffsetX,
		AutoFunctionAOIOffsetY,
		AutoFunctionAOISelector,
		AutoFunctionAOIUseBrightness,
		AutoFunctionAOIUseWhiteBalance,
		AutoFunctionAOIWidth,
		AutoFunctionProfile,
		AutoFunctionROIHeight,
		AutoFunctionROIOffsetX,
		AutoFunctionROIOffsetY,
		AutoFunctionROISelector,
		AutoFunctionROIUseBrightness,
		AutoFunctionROIUseWhiteBalance,
		AutoFunctionROIWidth,
		AutoGainLowerLimit,
		AutoGainUpperLimit,
		AutoTargetBrightness,
		BalanceRatio,
		BalanceRatioSelector,
		BalanceWhiteAuto,
		BinningHorizontal,
		BinningHorizontalMode,
		BinningVertical,
		BinningVerticalMode,
		BlackLevel,
		BlackLevelSelector,
		BslBrightness,
		BslColorSpaceMode,
		BslContrast,
		BslContrastMode,
		BslHue,
		BslHueValue,
		BslImmediateTriggerMode,
		BslSaturation,
		BslSaturationValue,
		BslUSBSpeedMode,
		CenterX,
		CenterY,
		ChunkCounterSelector,
		ChunkCounterValue,
		ChunkEnable,
		ChunkExposureTime,
		ChunkGain,
		ChunkGainSelector,
		ChunkLineStatusAll,
		ChunkModeActive,
		ChunkPayloadCRC16,
		ChunkSelector,
		ChunkSequencerSetActive,
		ChunkTimestamp,
		ColorAdjustmentHue,
		ColorAdjustmentSaturation,
		ColorAdjustmentSelector,
		ColorSpace,
		ColorTransformationSelector,
		ColorTransformationValue,
		ColorTransformationValueSelector,
		CounterDuration,
		CounterEventSource,
		CounterReset,
		CounterResetActivation,
		CounterResetSource,
		CounterSelector,
		DecimationHorizontal,
		DecimationVertical,
		DefectPixelCorrectionMode,
		DemosaicingMode,
		DeviceFirmwareVersion,
		DeviceIndicatorMode,
		DeviceLinkCurrentThroughput,
		DeviceLinkSelector,
		DeviceLinkSpeed,
		DeviceLinkThroughputLimit,
		DeviceLinkThroughputLimitMode,
		DeviceManufacturerInfo,
		DeviceModelName,
		DeviceRegistersStreamingEnd,
		DeviceRegistersStreamingStart,
		DeviceReset,
		DeviceSFNCVersionMajor,
		DeviceSFNCVersionMinor,
		DeviceSFNCVersionSubMinor,
		DeviceScanType,
		DeviceSerialNumber,
		DeviceTemperature,
		DeviceTemperatureSelector,
		DeviceUserID,
		DeviceVendorName,
		DeviceVersion,
		DigitalShift,
		EventCriticalTemperature,
		EventCriticalTemperatureTimestamp,
		EventExposureEnd,
		EventExposureEndFrameID,
		EventExposureEndTimestamp,
		EventFrameBurstStart,
		EventFrameBurstStartFrameID,
		EventFrameBurstStartOvertrigger,
		EventFrameBurstStartOvertriggerFrameID,
		EventFrameBurstStartOvertriggerTimestamp,
		EventFrameBurstStartTimestamp,
		EventFrameBurstStartWait,
		EventFrameBurstStartWaitTimestamp,
		EventFrameStart,
		EventFrameStartFrameID,
		EventFrameStartOvertrigger,
		EventFrameStartOvertriggerFrameID,
		EventFrameStartOvertriggerTimestamp,
		EventFrameStartTimestamp,
		EventFrameStartWait,
		EventFrameStartWaitTimestamp,
		EventNotification,
		EventOverTemperature,
		EventOverTemperatureTimestamp,
		EventSelector,
		EventTest,
		EventTestTimestamp,
		ExpertFeatureAccessKey,
		ExpertFeatureAccessSelector,
		ExpertFeatureEnable,
		ExposureAuto,
		ExposureMode,
		ExposureOverlapTimeMax,
		ExposureOverlapTimeMode,
		ExposureTime,
		FileAccessBuffer,
		FileAccessLength,
		FileAccessOffset,
		FileOpenMode,
		FileOperationExecute,
		FileOperationResult,
		FileOperationSelector,
		FileOperationStatus,
		FileSelector,
		FileSize,
		Gain,
		GainAuto,
		GainSelector,
		Gamma,
		Height,
		HeightMax,
		LUTEnable,
		LUTIndex,
		LUTSelector,
		LUTValue,
		LUTValueAll,
		LightSourcePreset,
		LineDebouncerTime,
		LineFormat,
		LineInverter,
		LineLogic,
		LineMinimumOutputPulseWidth,
		LineMode,
		LineOverloadReset,
		LineOverloadStatus,
		LinePitch,
		LinePitchEnable,
		LineSelector,
		LineSource,
		LineStatus,
		LineStatusAll,
		NoiseReduction,
		OffsetX,
		OffsetY,
		OverlapMode,
		PayloadFinalTransfer1Size,
		PayloadFinalTransfer2Size,
		PayloadSize,
		PayloadTransferBlockDelay,
		PayloadTransferCount,
		PayloadTransferSize,
		PgiMode,
		PixelColorFilter,
		PixelDynamicRangeMax,
		PixelDynamicRangeMin,
		PixelFormat,
		PixelSize,
		ROIZoneMode,
		ROIZoneOffset,
		ROIZoneSelector,
		ROIZoneSize,
		RemoveParameterLimit,
		RemoveParameterLimitSelector,
		ResultingFrameRate,
		ReverseX,
		ReverseY,
		SIPayloadFinalTransfer1Size,
		SIPayloadFinalTransfer2Size,
		SIPayloadTransferCount,
		SIPayloadTransferSize,
		ScalingHorizontal,
		ScalingVertical,
		SensorHeight,
		SensorReadoutMode,
		SensorReadoutTime,
		SensorShutterMode,
		SensorWidth,
		SequencerConfigurationMode,
		SequencerMode,
		SequencerPathSelector,
		SequencerSetActive,
		SequencerSetLoad,
		SequencerSetNext,
		SequencerSetSave,
		SequencerSetSelector,
		SequencerSetStart,
		SequencerTriggerActivation,
		SequencerTriggerSource,
		SharpnessEnhancement,
		ShutterMode,
		SoftwareSignalPulse,
		SoftwareSignalSelector,
		TemperatureState,
		TestImageResetAndHold,
		TestImageSelector,
		TestPattern,
		TestPendingAck,
		TimerDelay,
		TimerDuration,
		TimerSelector,
		TimerTriggerSource,
		TimestampLatch,
		TimestampLatchValue,
		TriggerActivation,
		TriggerDelay,
		TriggerEventTest,
		TriggerMode,
		TriggerSelector,
		TriggerSoftware,
		TriggerSource,
		UserDefinedValue,
		UserDefinedValueSelector,
		UserOutputSelector,
		UserOutputValue,
		UserOutputValueAll,
		UserSetDefault,
		UserSetLoad,
		UserSetSave,
		UserSetSelector,
		Width,
		WidthMax,
}

/* EOF */
