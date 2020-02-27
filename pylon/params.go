package pylon

type ParamGroup int

func (p ParamGroup) Description() string {
	return groupParamDescription[p]
}

func (p ParamGroup) String() string {
	return groupParamName[p]
}

type Param struct {
	Name             string
	Group            ParamGroup
	Description      string
	OriginalType     string
	OriginalEnumType string
}

type ParamList []Param

func (pl ParamList) ParamGroup(pg ParamGroup) ParamList {
	r := ParamList{}
	for _, p := range pl {
		if p.Group == pg {
			r = append(r, p)
		}
	}
	return r
}

func (pl ParamList) ParamByName(n string) *Param {
	for _, p := range pl {
		if p.Name == n {
			return &p
		}
	}
	return nil
}

const (
	SequenceControl ParamGroup = iota
	SequenceControlConfiguration
	AnalogControls
	ImageFormat
	DeviceInformation
	ColorImprovementsControl
	AOI
	StackedZoneImaging
	AcquisitionTrigger
	DigitalIO
	VirtualInput
	ShaftEncoderModule
	FrequencyConverter
	TimerControls
	TimerSequence
	LUTControls
	TransportLayer
	ActionControl
	DeviceControl
	UserSets
	AutoFunctions
	AutoFunctionAOIs
	ColorOverexposureCompensation
	Shading
	UserDefinedValues
	FeatureSets
	RemoveParamLimits
	ExpertFeatureAccess
	ChunkDataStreams
	ChunkData
	EventsGeneration
	ExposureEndEventData
	LineStartOvertriggerEventData
	FrameStartOvertriggerEventData
	FrameStartEventData
	AcquisitionStartEventData
	AcquisitionStartOvertriggerEventData
	FrameTimeoutEventData
	EventOverrunEventData
	CriticalTemperatureEventData
	OverTemperatureEventData
	Line1RisingEdgeEventData
	Line2RisingEdgeEventData
	Line3RisingEdgeEventData
	Line4RisingEdgeEventData
	VirtualLine1RisingEdgeEventData
	VirtualLine2RisingEdgeEventData
	VirtualLine3RisingEdgeEventData
	VirtualLine4RisingEdgeEventData
	FileAccessControl

	SequencerControl	// USB specific
	PGIControl
	AutoFunctionROIControl
	SoftwareSignalControl
	EventExposureEndData
	EventFrameStartData
	EventFrameBurstStartData
	EventFrameStartOvertriggerData
	EventFrameBurstStartOvertriggerData
	EventTestData
	EventCriticalTemperatureData
	EventOverTemperatureData
	EventFrameStartWaitData
	EventFrameBurstStartWaitData
	RemoveParameterLimitControl

	ImageFormatControl	= DeviceInformation // WTF?
	AnalogControl		= AnalogControls
	ImageQualityControl	= ColorImprovementsControl
	AcquisitionControl	= AcquisitionTrigger
	AutoFunctionControl	= AutoFunctions
	AutoFunctionAOIControl	= AutoFunctionAOIs
	DigitalIOControl	= DigitalIO
	CounterAndTimerControl	= TimerControls
	UserSetControl		= UserSets
	ChunkDataControl	= ChunkDataStreams
	EventControl		= EventsGeneration
	TransportLayerControl	= TransportLayer

	OriginalTypeGenApiIBoolean      = "GenApi::IBoolean"
	OriginalTypeGenApiICommand      = "GenApi::ICommand"
	OriginalTypeGenApiIInteger      = "GenApi::IInteger"
	OriginalTypeGenApiIEnumerationT = "GenApi::IEnumerationT"
	OriginalTypeGenApiIFloat        = "GenApi::IFloat"
	OriginalTypeGenApiIString       = "GenApi::IString"
	OriginalTypeGenApiIRegister     = "GenApi::IRegister"
)

var (
	groupParamDescription = [...]string{
		"This category includes items that control the sequencer feature",
		"This category includes items that control the sequence set advance",
		"This category includes items that control the analog characteristics of the video signal",
		"This category includes items that control the size of the acquired image and the format of the transferred pixel data",
		"This category includes items that describe the device and its sensor",
		"This category includes items that control color improvement",
		"This category includes items used to set the size and position of the area of interest",
		"",
		"This category includes items used to set the image acquisition parameters and to start and stop acquisition",
		"This category includes items used to control the operation of the camera's digital I/O lines",
		"This category includes items used to control the operation of the camera’s virtual input I/O lines",
		"This category provides controls for operating the camera's shaft encoder module.",
		"This category includes items used to control the operation of the camera's frequency converter module",
		"This category includes items used to control the operation of the camera's timers",
		"",
		"This category includes items used to control the operation of the camera's lookup table (LUT)",
		"This category includes items related to the GigE Vision transport layer",
		"This category includes items that control the action control feature",
		"",
		"This category includes items that control the configuration sets feature that is used to save sets of parameters in the camera",
		"This category includes items that parameterize the Auto Functions",
		"Portion of the sensor array used for auto function control",
		"Compensates for deviations of hue resulting from overexposure",
		"Includes items used to control the operation of the camera's shading correction.",
		"",
		"",
		"This category includes items that allow removing the limits of camera parameters",
		"",
		"This category includes items that control the chunk features available on the camera.",
		"This category includes items related to the chunk data that can be appended to the image data",
		"This category includes items that control event generation by the camera.",
		"This category includes items available for an exposure end event",
		"This category includes items available for an line start overtrigger event",
		"This category includes items available for an frame start overtrigger event",
		"This category includes items available for an frame start event",
		"This category includes items available for an acquisition start event",
		"This category includes items available for an acquisition start overtrigger event",
		"This category includes items available for an frame timeout event",
		"This category includes items available for an event overrun event",
		"This category includes items available for a critical temperature event",
		"This category includes items available for an over temperature event",
		"This category includes items available for an io line 1 rising edge event",
		"This category includes items available for an io line 2 rising edge event",
		"This category includes items available for an io line 3 rising edge event",
		"This category includes items available for an io line 4 rising edge event",
		"This category includes items available for an io virtual line 1 rising edge event",
		"This category includes items available for an io virtual line 2 rising edge event",
		"This category includes items available for an io virtual line 3 rising edge event",
		"This category includes items available for an io virtual line 4 rising edge event",
		"This category includes items used to conduct file operations",
		"USB SequencerControl group",	// USB specific
		"USB PGIControl group",
		"USB AutoFunctionROIControl group",
		"USB SoftwareSignalControl group",
		"USB EventExposureEndData group",
		"USB EventFrameStartData group",
		"USB EventFrameBurstStartData group",
		"USB EventFrameStartOvertriggerData group",
		"USB EventFrameBurstStartOvertriggerData group",
		"USB EventTestData group",
		"USB EventCriticalTemperatureData group",
		"USB EventOverTemperatureData group",
		"USB EventFrameStartWaitData group",
		"USB EventFrameBurstStartWaitData group",
		"USB RemoveParameterLimitControl group",
	}

	groupParamName = [...]string{
		"SequenceControl",
		"SequenceControlConfiguration",
		"AnalogControls",
		"ImageFormat",
		"DeviceInformation",
		"ColorImprovementsControl",
		"AOI",
		"StackedZoneImaging",
		"AcquisitionTrigger",
		"DigitalIO",
		"VirtualInput",
		"ShaftEncoderModule",
		"FrequencyConverter",
		"TimerControls",
		"TimerSequence",
		"LUTControls",
		"TransportLayer",
		"ActionControl",
		"DeviceControl",
		"UserSets",
		"AutoFunctions",
		"AutoFunctionAOIs",
		"ColorOverexposureCompensation",
		"Shading",
		"UserDefinedValues",
		"FeatureSets",
		"RemoveParamLimits",
		"ExpertFeatureAccess",
		"ChunkDataStreams",
		"ChunkData",
		"EventsGeneration",
		"ExposureEndEventData",
		"LineStartOvertriggerEventData",
		"FrameStartOvertriggerEventData",
		"FrameStartEventData",
		"AcquisitionStartEventData",
		"AcquisitionStartOvertriggerEventData",
		"FrameTimeoutEventData",
		"EventOverrunEventData",
		"CriticalTemperatureEventData",
		"OverTemperatureEventData",
		"Line1RisingEdgeEventData",
		"Line2RisingEdgeEventData",
		"Line3RisingEdgeEventData",
		"Line4RisingEdgeEventData",
		"VirtualLine1RisingEdgeEventData",
		"VirtualLine2RisingEdgeEventData",
		"VirtualLine3RisingEdgeEventData",
		"VirtualLine4RisingEdgeEventData",
		"FileAccessControl",
		"SequencerControl",	// USB specific
		"PGIControl",
		"AutoFunctionROIControl",
		"SoftwareSignalControl",
		"EventExposureEndData",
		"EventFrameStartData",
		"EventFrameBurstStartData",
		"EventFrameStartOvertriggerData",
		"EventFrameBurstStartOvertriggerData",
		"EventTestData",
		"EventCriticalTemperatureData",
		"EventOverTemperatureData",
		"EventFrameStartWaitData",
		"EventFrameBurstStartWaitData",
		"RemoveParameterLimitControl",
	}

	// SequenceControl - This category includes items that control the sequencer feature
	SequenceEnable         = Param{Name: "SequenceEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the sequencer.", Group: SequenceControl}
	SequenceAsyncRestart   = Param{Name: "SequenceAsyncRestart", OriginalType: OriginalTypeGenApiICommand, Description: "Allows asynchronous restart of the sequence of sequence sets.", Group: SequenceControl}
	SequenceAsyncAdvance   = Param{Name: "SequenceAsyncAdvance", OriginalType: OriginalTypeGenApiICommand, Description: "Allows asynchronous advance from one sequence set to the next.", Group: SequenceControl}
	SequenceCurrentSet     = Param{Name: "SequenceCurrentSet", OriginalType: OriginalTypeGenApiIInteger, Description: "Current sequence set.", Group: SequenceControl}
	SequenceSetTotalNumber = Param{Name: "SequenceSetTotalNumber", OriginalType: OriginalTypeGenApiIInteger, Description: "Total number of sequence sets.", Group: SequenceControl}
	SequenceSetIndex       = Param{Name: "SequenceSetIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Selects the index number of a sequence set.", Group: SequenceControl}
	SequenceSetLoad        = Param{Name: "SequenceSetLoad", OriginalType: OriginalTypeGenApiICommand, Description: "Loads a sequence set.", Group: SequenceControl}
	SequenceSetStore       = Param{Name: "SequenceSetStore", OriginalType: OriginalTypeGenApiICommand, Description: "Stores the current sequence set.", Group: SequenceControl}
	SequenceSetExecutions  = Param{Name: "SequenceSetExecutions", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the number of sequence set executions.", Group: SequenceControl}

	// SequenceControlConfiguration - This category includes items that control the sequence set advance
	SequenceAdvanceMode        = Param{Name: "SequenceAdvanceMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the sequence set advance mode.", Group: SequenceControlConfiguration, OriginalEnumType: "SequenceAdvanceModeEnums"}
	SequenceControlSelector    = Param{Name: "SequenceControlSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects between sequence restart or sequence set advance.", Group: SequenceControlConfiguration, OriginalEnumType: "SequenceControlSelectorEnums"}
	SequenceControlSource      = Param{Name: "SequenceControlSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the source for sequence control.", Group: SequenceControlConfiguration, OriginalEnumType: "SequenceControlSourceEnums"}
	SequenceAddressBitSelector = Param{Name: "SequenceAddressBitSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects a bit of the sequence set address.", Group: SequenceControlConfiguration, OriginalEnumType: "SequenceAddressBitSelectorEnums"}
	SequenceAddressBitSource   = Param{Name: "SequenceAddressBitSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the source for the selected bit of the sequence set address.", Group: SequenceControlConfiguration, OriginalEnumType: "SequenceAddressBitSourceEnums"}

	// AnalogControls - This category includes items that control the analog characteristics of the video signal
	GainAuto           = Param{Name: "GainAuto", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Gain Auto is the 'automatic' counterpart of the manual gain feature.", Group: AnalogControls, OriginalEnumType: "GainAutoEnums"}
	GainRaw            = Param{Name: "GainRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "This is an integer value that sets the selected gain control in device specific units.", Group: AnalogControls}
	GainAbs            = Param{Name: "GainAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "This is a float value that sets the selected gain control in dB.", Group: AnalogControls}
	GainSelector       = Param{Name: "GainSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the gain control to configure.", Group: AnalogControls, OriginalEnumType: "GainSelectorEnums"}
	BlackLevelRaw      = Param{Name: "BlackLevelRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the value of the selected black level control as an integer.", Group: AnalogControls}
	BlackLevelAbs      = Param{Name: "BlackLevelAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the value of the selected black level control as a float.", Group: AnalogControls}
	BlackLevelSelector = Param{Name: "BlackLevelSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selcts a black level control to configure.", Group: AnalogControls, OriginalEnumType: "BlackLevelSelectorEnums"}
	GammaEnable        = Param{Name: "GammaEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the gamma correction.", Group: AnalogControls}
	GammaSelector      = Param{Name: "GammaSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "This enumeration selects the type of gamma to apply.", Group: AnalogControls, OriginalEnumType: "GammaSelectorEnums"}
	Gamma              = Param{Name: "Gamma", OriginalType: OriginalTypeGenApiIFloat, Description: "This feature is used to perform gamma correction of pixel intensity.", Group: AnalogControls}
	DigitalShift       = Param{Name: "DigitalShift", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the value of the selected digital shift control.", Group: AnalogControls}
	SubstrateVoltage   = Param{Name: "SubstrateVoltage", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the substrate voltage.", Group: AnalogControls}

	// ImageFormat - This category includes items that control the size of the acquired image and the format of the transferred pixel data
	PixelFormat                   = Param{Name: "PixelFormat", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the format of the pixel data transmitted for acquired images.", Group: ImageFormat, OriginalEnumType: "PixelFormatEnums"}
	PixelCoding                   = Param{Name: "PixelCoding", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the color coding of the pixels in the acquired images.", Group: ImageFormat, OriginalEnumType: "PixelCodingEnums"}
	PixelSize                     = Param{Name: "PixelSize", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Indicates the depth of the pixel values in the image in bits per pixel.", Group: ImageFormat, OriginalEnumType: "PixelSizeEnums"}
	PixelColorFilter              = Param{Name: "PixelColorFilter", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Indicates the alignment of the camera's Bayer filter to the pixels in the acquired images.", Group: ImageFormat, OriginalEnumType: "PixelColorFilterEnums"}
	ProcessedRawEnable            = Param{Name: "ProcessedRawEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables color improved RGB raw output.", Group: ImageFormat}
	SpatialCorrection             = Param{Name: "SpatialCorrection", OriginalType: OriginalTypeGenApiIInteger, Description: "Specifies the direction of imaging and the separation (consecutive numbers) of related line captures.", Group: ImageFormat}
	SpatialCorrectionAmount       = Param{Name: "SpatialCorrectionAmount", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Invisible", Group: ImageFormat}
	SpatialCorrectionStartingLine = Param{Name: "SpatialCorrectionStartingLine", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Invisible", Group: ImageFormat, OriginalEnumType: "SpatialCorrectionStartingLineEnums"}
	SensorBitDepth                = Param{Name: "SensorBitDepth", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "This feature selects the amount of data bits the sensor produces for one sample.", Group: ImageFormat, OriginalEnumType: "SensorBitDepthEnums"}
	ReverseX                      = Param{Name: "ReverseX", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the horizontal flipping of the image.", Group: ImageFormat}
	ReverseY                      = Param{Name: "ReverseY", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the vertical flipping of the image.", Group: ImageFormat}
	PixelDynamicRangeMin          = Param{Name: "PixelDynamicRangeMin", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the minimum possible pixel value that could be transferred from the camera.", Group: ImageFormat}
	PixelDynamicRangeMax          = Param{Name: "PixelDynamicRangeMax", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the maximum possible pixel value that could be transferred from the camera.", Group: ImageFormat}
	FieldOutputMode               = Param{Name: "FieldOutputMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the mode to output the fields.", Group: ImageFormat, OriginalEnumType: "FieldOutputModeEnums"}
	TestImageSelector             = Param{Name: "TestImageSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selecting a test image from the list will enable the test image.", Group: ImageFormat, OriginalEnumType: "TestImageSelectorEnums"}
	SensorDigitizationTaps        = Param{Name: "SensorDigitizationTaps", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "This feature represents the number of digitized samples outputted simultaneously by the camera A/D conversion stage.", Group: ImageFormat, OriginalEnumType: "SensorDigitizationTapsEnums"}

	// DeviceInformation - This category includes items that describe the device and its sensor
	SensorWidth            = Param{Name: "SensorWidth", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the width of the camera's sensor in pixels.", Group: DeviceInformation}
	SensorHeight           = Param{Name: "SensorHeight", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the height of the camera's sensor in pixels.", Group: DeviceInformation}
	WidthMax               = Param{Name: "WidthMax", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the maximum allowed width of the image in pixels.", Group: DeviceInformation}
	HeightMax              = Param{Name: "HeightMax", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the maximum allowed height of the image in pixels.", Group: DeviceInformation}
	DeviceVendorName       = Param{Name: "DeviceVendorName", OriginalType: OriginalTypeGenApiIString, Description: "Indicates the name of the device's vendor.", Group: DeviceInformation}
	DeviceModelName        = Param{Name: "DeviceModelName", OriginalType: OriginalTypeGenApiIString, Description: "Indicates the model name of the device.", Group: DeviceInformation}
	DeviceManufacturerInfo = Param{Name: "DeviceManufacturerInfo", OriginalType: OriginalTypeGenApiIString, Description: "Provides additional information from the vendor about the device.", Group: DeviceInformation}
	DeviceVersion          = Param{Name: "DeviceVersion", OriginalType: OriginalTypeGenApiIString, Description: "Indicates the version of the device.", Group: DeviceInformation}
	DeviceFirmwareVersion  = Param{Name: "DeviceFirmwareVersion", OriginalType: OriginalTypeGenApiIString, Description: "Indicates the version of the device's firmware and software.", Group: DeviceInformation}
	DeviceID               = Param{Name: "DeviceID", OriginalType: OriginalTypeGenApiIString, Description: "A unique identifier for the device such as a serial number or a GUID.", Group: DeviceInformation}
	DeviceUserID           = Param{Name: "DeviceUserID", OriginalType: OriginalTypeGenApiIString, Description: "A device ID that is user programmable.", Group: DeviceInformation}
	DeviceReset            = Param{Name: "DeviceReset", OriginalType: OriginalTypeGenApiICommand, Description: "Immediately resets and reboots the device.", Group: DeviceInformation}
	DeviceScanType         = Param{Name: "DeviceScanType", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Indicates the scan type of the device's sensor.", Group: DeviceInformation, OriginalEnumType: "DeviceScanTypeEnums"}
	LastError              = Param{Name: "LastError", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Indicates the error that was detected last.", Group: DeviceInformation, OriginalEnumType: "LastErrorEnums"}
	ClearLastError         = Param{Name: "ClearLastError", OriginalType: OriginalTypeGenApiICommand, Description: "Erases the last error and possibly reveals a previous error.", Group: DeviceInformation}
	TemperatureSelector    = Param{Name: "TemperatureSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Lists the temperature sources available for readout.", Group: DeviceInformation, OriginalEnumType: "TemperatureSelectorEnums"}
	TemperatureAbs         = Param{Name: "TemperatureAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Shows the current temperature of the selected target in degrees centigrade.", Group: DeviceInformation}
	TemperatureState       = Param{Name: "TemperatureState", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Temperature State.", Group: DeviceInformation, OriginalEnumType: "TemperatureStateEnums"}
	CriticalTemperature    = Param{Name: "CriticalTemperature", OriginalType: OriginalTypeGenApiIBoolean, Description: "Shows the over temperature state of the selected target.", Group: DeviceInformation}
	OverTemperature        = Param{Name: "OverTemperature", OriginalType: OriginalTypeGenApiIBoolean, Description: "Shows the over temperature state of the selected target.", Group: DeviceInformation}

	// ColorImprovementsControl - This category includes items that control color improvement
	ColorTransformationSelector        = Param{Name: "ColorTransformationSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the matrix color transformation between color spaces.", Group: ColorImprovementsControl, OriginalEnumType: "ColorTransformationSelectorEnums"}
	LightSourceSelector                = Param{Name: "LightSourceSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the type of light source to be considered for matrix color transformation.", Group: ColorImprovementsControl, OriginalEnumType: "LightSourceSelectorEnums"}
	ColorTransformationMatrixFactor    = Param{Name: "ColorTransformationMatrixFactor", OriginalType: OriginalTypeGenApiIFloat, Description: "Defines the extent to which the selected light source will be considered (float)", Group: ColorImprovementsControl}
	ColorTransformationMatrixFactorRaw = Param{Name: "ColorTransformationMatrixFactorRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Defines the extent to which the selected light source will be considered (integer)", Group: ColorImprovementsControl}
	ColorTransformationValueSelector   = Param{Name: "ColorTransformationValueSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the element to be entered in the color transformation matrix.", Group: ColorImprovementsControl, OriginalEnumType: "ColorTransformationValueSelectorEnums"}
	ColorTransformationValue           = Param{Name: "ColorTransformationValue", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets a floating point value for the selected element in the color transformation matrix.", Group: ColorImprovementsControl}
	ColorTransformationValueRaw        = Param{Name: "ColorTransformationValueRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets an integer value for the selected element in the color transformation matrix.", Group: ColorImprovementsControl}
	ColorAdjustmentEnable              = Param{Name: "ColorAdjustmentEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables color adjustment.", Group: ColorImprovementsControl}
	ColorAdjustmentReset               = Param{Name: "ColorAdjustmentReset", OriginalType: OriginalTypeGenApiICommand, Description: "Allows returning to previous settings.", Group: ColorImprovementsControl}
	ColorAdjustmentSelector            = Param{Name: "ColorAdjustmentSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the color for color adjustment.", Group: ColorImprovementsControl, OriginalEnumType: "ColorAdjustmentSelectorEnums"}
	ColorAdjustmentHue                 = Param{Name: "ColorAdjustmentHue", OriginalType: OriginalTypeGenApiIFloat, Description: "Adjustment of hue of the selected color (float)", Group: ColorImprovementsControl}
	ColorAdjustmentHueRaw              = Param{Name: "ColorAdjustmentHueRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Adjustment of hue of the selected color (integer)", Group: ColorImprovementsControl}
	ColorAdjustmentSaturation          = Param{Name: "ColorAdjustmentSaturation", OriginalType: OriginalTypeGenApiIFloat, Description: "Adjustment of saturation of the selected color (float)", Group: ColorImprovementsControl}
	ColorAdjustmentSaturationRaw       = Param{Name: "ColorAdjustmentSaturationRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Adjustment of saturation of the selected color (integer)", Group: ColorImprovementsControl}
	BalanceWhiteReset                  = Param{Name: "BalanceWhiteReset", OriginalType: OriginalTypeGenApiICommand, Description: "Allows returning to previous settings.", Group: ColorImprovementsControl}
	BalanceWhiteAuto                   = Param{Name: "BalanceWhiteAuto", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Balance White Auto is the 'automatic' counterpart of the manual white balance feature.", Group: ColorImprovementsControl, OriginalEnumType: "BalanceWhiteAutoEnums"}
	BalanceRatioSelector               = Param{Name: "BalanceRatioSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects a balance ratio to configure.", Group: ColorImprovementsControl, OriginalEnumType: "BalanceRatioSelectorEnums"}
	BalanceRatioAbs                    = Param{Name: "BalanceRatioAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the value of the selected balance ratio control as a float.", Group: ColorImprovementsControl}
	BalanceRatioRaw                    = Param{Name: "BalanceRatioRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the value of the selected balance ratio control as an integer.", Group: ColorImprovementsControl}

	// AOI - This category includes items used to set the size and position of the area of interest
	Width                 = Param{Name: "Width", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the width of the area of interest in pixels.", Group: AOI}
	Height                = Param{Name: "Height", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the height of the area of interest in pixels.", Group: AOI}
	OffsetX               = Param{Name: "OffsetX", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the X offset (left offset) of the area of interest in pixels.", Group: AOI}
	OffsetY               = Param{Name: "OffsetY", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the Y offset (top offset) for the area of interest in pixels.", Group: AOI}
	CenterX               = Param{Name: "CenterX", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the horizontal centering of the image.", Group: AOI}
	CenterY               = Param{Name: "CenterY", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the vertical centering of the image.", Group: AOI}
	BinningModeVertical   = Param{Name: "BinningModeVertical", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the vertical binning mode.", Group: AOI, OriginalEnumType: "BinningModeVerticalEnums"}
	BinningModeHorizontal = Param{Name: "BinningModeHorizontal", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the horizontal binning mode.", Group: AOI, OriginalEnumType: "BinningModeHorizontalEnums"}
	BinningVertical       = Param{Name: "BinningVertical", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the number of adjacent vertical pixes to be summed.", Group: AOI}
	BinningHorizontal     = Param{Name: "BinningHorizontal", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the number of adjacent horizontal pixes to be summed.", Group: AOI}
	LegacyBinningVertical = Param{Name: "LegacyBinningVertical", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the vertical binning feature.", Group: AOI, OriginalEnumType: "LegacyBinningVerticalEnums"}
	DecimationVertical    = Param{Name: "DecimationVertical", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets vertical sub-sampling.", Group: AOI}
	DecimationHorizontal  = Param{Name: "DecimationHorizontal", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets horizontal sub-sampling.", Group: AOI}

	// StackedZoneImaging -
	StackedZoneImagingEnable      = Param{Name: "StackedZoneImagingEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the stacked zone imaging feature.", Group: StackedZoneImaging}
	StackedZoneImagingIndex       = Param{Name: "StackedZoneImagingIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "This value sets the zone to access.", Group: StackedZoneImaging}
	StackedZoneImagingZoneEnable  = Param{Name: "StackedZoneImagingZoneEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the selected zone.", Group: StackedZoneImaging}
	StackedZoneImagingZoneOffsetY = Param{Name: "StackedZoneImagingZoneOffsetY", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the Y offset (top offset) for the selected zone.", Group: StackedZoneImaging}
	StackedZoneImagingZoneHeight  = Param{Name: "StackedZoneImagingZoneHeight", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the height for the selected zone.", Group: StackedZoneImaging}

	// AcquisitionTrigger - This category includes items used to set the image acquisition parameters and to start and stop acquisition
	EnableBurstAcquisition       = Param{Name: "EnableBurstAcquisition", OriginalType: OriginalTypeGenApiIBoolean, Description: "When enabled, the maximum frame rate does not depend on the image transfer rate out of the camera.", Group: AcquisitionTrigger}
	AcquisitionMode              = Param{Name: "AcquisitionMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the image acquisition mode.", Group: AcquisitionTrigger, OriginalEnumType: "AcquisitionModeEnums"}
	AcquisitionStart             = Param{Name: "AcquisitionStart", OriginalType: OriginalTypeGenApiICommand, Description: "Starts the acquisition of images.", Group: AcquisitionTrigger}
	AcquisitionStop              = Param{Name: "AcquisitionStop", OriginalType: OriginalTypeGenApiICommand, Description: "Stops the acquisition of images.", Group: AcquisitionTrigger}
	AcquisitionAbort             = Param{Name: "AcquisitionAbort", OriginalType: OriginalTypeGenApiICommand, Description: "Immediately aborts the acquisition of images.", Group: AcquisitionTrigger}
	AcquisitionFrameCount        = Param{Name: "AcquisitionFrameCount", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the number of frames acquired in the multiframe acquisition mode.", Group: AcquisitionTrigger}
	TriggerControlImplementation = Param{Name: "TriggerControlImplementation", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Expert", Group: AcquisitionTrigger, OriginalEnumType: "TriggerControlImplementationEnums"}
	TriggerSelector              = Param{Name: "TriggerSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the trigger type to configure.", Group: AcquisitionTrigger, OriginalEnumType: "TriggerSelectorEnums"}
	TriggerMode                  = Param{Name: "TriggerMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the mode for the selected trigger.", Group: AcquisitionTrigger, OriginalEnumType: "TriggerModeEnums"}
	TriggerSoftware              = Param{Name: "TriggerSoftware", OriginalType: OriginalTypeGenApiICommand, Description: "Generates a software trigger signal that is used when the trigger source is set to 'software'.", Group: AcquisitionTrigger}
	TriggerSource                = Param{Name: "TriggerSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the signal source for the selected trigger.", Group: AcquisitionTrigger, OriginalEnumType: "TriggerSourceEnums"}
	TriggerActivation            = Param{Name: "TriggerActivation", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the signal transition needed to activate the selected trigger.", Group: AcquisitionTrigger, OriginalEnumType: "TriggerActivationEnums"}
	TriggerPartialClosingFrame   = Param{Name: "TriggerPartialClosingFrame", OriginalType: OriginalTypeGenApiIBoolean, Description: "Determines whether a partial or complete frame is transmitted when the frame start trigger prematurely transitions.", Group: AcquisitionTrigger}
	TriggerDelayAbs              = Param{Name: "TriggerDelayAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the trigger delay time in microseconds.", Group: AcquisitionTrigger}
	TriggerDelaySource           = Param{Name: "TriggerDelaySource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the kind of trigger delay.", Group: AcquisitionTrigger, OriginalEnumType: "TriggerDelaySourceEnums"}
	TriggerDelayLineTriggerCount = Param{Name: "TriggerDelayLineTriggerCount", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the trigger delay expressed as number of line triggers.", Group: AcquisitionTrigger}
	ExposureStartDelayRaw        = Param{Name: "ExposureStartDelayRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: AcquisitionTrigger}
	ExposureStartDelayAbs        = Param{Name: "ExposureStartDelayAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "<b>Visibility</b> = Beginner", Group: AcquisitionTrigger}
	ExposureAuto                 = Param{Name: "ExposureAuto", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Exposure Auto is the 'automatic' counterpart to manually setting an 'absolute' exposure time.", Group: AcquisitionTrigger, OriginalEnumType: "ExposureAutoEnums"}
	ExposureTimeRaw              = Param{Name: "ExposureTimeRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the 'raw' exposure time.", Group: AcquisitionTrigger}
	ExposureTimeAbs              = Param{Name: "ExposureTimeAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Directly sets the camera's exposure time in microseconds.", Group: AcquisitionTrigger}
	ExposureTimeBaseAbs          = Param{Name: "ExposureTimeBaseAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the time base (in microseconds) that is used when the exposure time is set with the 'exposure time raw' setting.", Group: AcquisitionTrigger}
	ExposureTimeBaseAbsEnable    = Param{Name: "ExposureTimeBaseAbsEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the use of the exposure time base.", Group: AcquisitionTrigger}
	AcquisitionLineRateAbs       = Param{Name: "AcquisitionLineRateAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the camera's acquisition line rate in lines per second.", Group: AcquisitionTrigger}
	ResultingLinePeriodAbs       = Param{Name: "ResultingLinePeriodAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Indicates the minimum allowed line acquisition period (in microseconds) given the current settings for the area of interest, exposure time, and bandwidth.", Group: AcquisitionTrigger}
	ResultingLineRateAbs         = Param{Name: "ResultingLineRateAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Indicates the maximum allowed line acquisition rate (in lines per second) given the current settings for the area of interest, exposure time, and bandwidth.", Group: AcquisitionTrigger}
	AcquisitionFrameRateAbs      = Param{Name: "AcquisitionFrameRateAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "If the acquisition frame rate feature is enabled, this value sets the camera's acquisition frame rate in frames per second.", Group: AcquisitionTrigger}
	AcquisitionFrameRateEnable   = Param{Name: "AcquisitionFrameRateEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables setting the camera's acquisition frame rate to a specified value.", Group: AcquisitionTrigger}
	ResultingFramePeriodAbs      = Param{Name: "ResultingFramePeriodAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Indicates the minimum allowed frame acquisition period (in microseconds) given the current settings for the area of interest, exposure time, and bandwidth.", Group: AcquisitionTrigger}
	ResultingFrameRateAbs        = Param{Name: "ResultingFrameRateAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Indicates the maximum allowed frame acquisition rate (in frames per second) given the current settings for the area of interest, exposure time, and bandwidth.", Group: AcquisitionTrigger}
	ExposureMode                 = Param{Name: "ExposureMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the exposure mode.", Group: AcquisitionTrigger, OriginalEnumType: "ExposureModeEnums"}
	GlobalResetReleaseModeEnable = Param{Name: "GlobalResetReleaseModeEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enable the Global Reset Release Mode.", Group: AcquisitionTrigger}
	ShutterMode                  = Param{Name: "ShutterMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the shutter mode.", Group: AcquisitionTrigger, OriginalEnumType: "ShutterModeEnums"}
	ExposureOverlapTimeMaxRaw    = Param{Name: "ExposureOverlapTimeMaxRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the maximum overlap of the sensor exposure with the sensor readout in TriggerWidth exposure mode in raw units.", Group: AcquisitionTrigger}
	InterlacedIntegrationMode    = Param{Name: "InterlacedIntegrationMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the Interlaced Integration Mode.", Group: AcquisitionTrigger, OriginalEnumType: "InterlacedIntegrationModeEnums"}
	ExposureOverlapTimeMaxAbs    = Param{Name: "ExposureOverlapTimeMaxAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the maximum overlap of the sensor exposure with sensor readout in TriggerWidth exposure mode in microseconds.", Group: AcquisitionTrigger}
	ReadoutTimeAbs               = Param{Name: "ReadoutTimeAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Indicates the sensor readout time given the current settings.", Group: AcquisitionTrigger}
	AcquisitionStatusSelector    = Param{Name: "AcquisitionStatusSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "This enumeration is used to select which internal acquisition signal to read using AcquisitionStatus.", Group: AcquisitionTrigger, OriginalEnumType: "AcquisitionStatusSelectorEnums"}
	AcquisitionStatus            = Param{Name: "AcquisitionStatus", OriginalType: OriginalTypeGenApiIBoolean, Description: "Reads the selected acquisition status.", Group: AcquisitionTrigger}
	FrameTimeoutEnable           = Param{Name: "FrameTimeoutEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the frame timeout.", Group: AcquisitionTrigger}
	FrameTimeoutAbs              = Param{Name: "FrameTimeoutAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the frame timeout in microseconds.", Group: AcquisitionTrigger}

	// DigitalIO - This category includes items used to control the operation of the camera's digital I/O lines
	LineSelector           = Param{Name: "LineSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the I/O line to configure.", Group: DigitalIO, OriginalEnumType: "LineSelectorEnums"}
	LineInverter           = Param{Name: "LineInverter", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the signal inverter function for the selected input or output line.", Group: DigitalIO}
	LineTermination        = Param{Name: "LineTermination", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the termination resistor for the selected input line.", Group: DigitalIO}
	LineDebouncerTimeRaw   = Param{Name: "LineDebouncerTimeRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the raw value of the selected line debouncer time.", Group: DigitalIO}
	LineDebouncerTimeAbs   = Param{Name: "LineDebouncerTimeAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the absolute value of the selected line debouncer time in microseconds.", Group: DigitalIO}
	MinOutPulseWidthRaw    = Param{Name: "MinOutPulseWidthRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the raw value for the minimum signal width of an output signal.", Group: DigitalIO}
	MinOutPulseWidthAbs    = Param{Name: "MinOutPulseWidthAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the absolute value (in microseconds) for the minimum signal width of an output signal.", Group: DigitalIO}
	LineStatus             = Param{Name: "LineStatus", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates the current logical state for the selected line.", Group: DigitalIO}
	LineStatusAll          = Param{Name: "LineStatusAll", OriginalType: OriginalTypeGenApiIInteger, Description: "A single bitfield indicating the current logical state of all available line signals at time of polling.", Group: DigitalIO}
	UserOutputValue        = Param{Name: "UserOutputValue", OriginalType: OriginalTypeGenApiIBoolean, Description: "Sets the state of the selected user settable output signal.", Group: DigitalIO}
	UserOutputValueAll     = Param{Name: "UserOutputValueAll", OriginalType: OriginalTypeGenApiIInteger, Description: "A single bitfield that sets the state of all user settable output signals in one access.", Group: DigitalIO}
	UserOutputValueAllMask = Param{Name: "UserOutputValueAllMask", OriginalType: OriginalTypeGenApiIInteger, Description: "Defines a mask that is used when the User Output Value All setting is used to set all of the user settable output signals in one access.", Group: DigitalIO}
	SyncUserOutputValue    = Param{Name: "SyncUserOutputValue", OriginalType: OriginalTypeGenApiIBoolean, Description: "Sets the state of the selected user settable synchronous output signal.", Group: DigitalIO}
	SyncUserOutputValueAll = Param{Name: "SyncUserOutputValueAll", OriginalType: OriginalTypeGenApiIInteger, Description: "A single bitfield that sets the state of all user settable synchronous output signals in one access.", Group: DigitalIO}
	LineMode               = Param{Name: "LineMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the mode for the selected line.", Group: DigitalIO, OriginalEnumType: "LineModeEnums"}
	LineSource             = Param{Name: "LineSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the source signal for the selected line (if the selected line is an output)", Group: DigitalIO, OriginalEnumType: "LineSourceEnums"}
	LineLogic              = Param{Name: "LineLogic", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Beginner", Group: DigitalIO, OriginalEnumType: "LineLogicEnums"}
	LineFormat             = Param{Name: "LineFormat", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the electrical configuration of the selected line.", Group: DigitalIO, OriginalEnumType: "LineFormatEnums"}
	UserOutputSelector     = Param{Name: "UserOutputSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the user settable output signal to configure.", Group: DigitalIO, OriginalEnumType: "UserOutputSelectorEnums"}
	SyncUserOutputSelector = Param{Name: "SyncUserOutputSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Beginner", Group: DigitalIO, OriginalEnumType: "SyncUserOutputSelectorEnums"}

	// VirtualInput - This category includes items used to control the operation of the camera’s virtual input I/O lines
	VInpSignalSource            = Param{Name: "VInpSignalSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the I/O line on which the camera receives the virtual input signal.", Group: VirtualInput, OriginalEnumType: "VInpSignalSourceEnums"}
	VInpBitLength               = Param{Name: "VInpBitLength", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the length of the input bit.", Group: VirtualInput}
	VInpSamplingPoint           = Param{Name: "VInpSamplingPoint", OriginalType: OriginalTypeGenApiIInteger, Description: "Time span between the beginning of the input bit and the time when the high/low status is evaluated.", Group: VirtualInput}
	VInpSignalReadoutActivation = Param{Name: "VInpSignalReadoutActivation", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects when to start the signal evaluation.", Group: VirtualInput, OriginalEnumType: "VInpSignalReadoutActivationEnums"}

	// ShaftEncoderModule - This category provides controls for operating the camera's shaft encoder module.
	ShaftEncoderModuleLineSelector        = Param{Name: "ShaftEncoderModuleLineSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the phase of the shaft encoder.", Group: ShaftEncoderModule, OriginalEnumType: "ShaftEncoderModuleLineSelectorEnums"}
	ShaftEncoderModuleLineSource          = Param{Name: "ShaftEncoderModuleLineSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the input line as signal source for the shaft encoder module.", Group: ShaftEncoderModule, OriginalEnumType: "ShaftEncoderModuleLineSourceEnums"}
	ShaftEncoderModuleMode                = Param{Name: "ShaftEncoderModuleMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the circumstances for the shaft encoder module to output trigger signals.", Group: ShaftEncoderModule, OriginalEnumType: "ShaftEncoderModuleModeEnums"}
	ShaftEncoderModuleCounterMode         = Param{Name: "ShaftEncoderModuleCounterMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the counting mode of the tick counter.", Group: ShaftEncoderModule, OriginalEnumType: "ShaftEncoderModuleCounterModeEnums"}
	ShaftEncoderModuleCounter             = Param{Name: "ShaftEncoderModuleCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the current value of the tick counter.", Group: ShaftEncoderModule}
	ShaftEncoderModuleCounterMax          = Param{Name: "ShaftEncoderModuleCounterMax", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the maximum value for the tick counter.", Group: ShaftEncoderModule}
	ShaftEncoderModuleCounterReset        = Param{Name: "ShaftEncoderModuleCounterReset", OriginalType: OriginalTypeGenApiICommand, Description: "Resets the tick counter to 0.", Group: ShaftEncoderModule}
	ShaftEncoderModuleReverseCounterMax   = Param{Name: "ShaftEncoderModuleReverseCounterMax", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the maximum value for the reverse counter.", Group: ShaftEncoderModule}
	ShaftEncoderModuleReverseCounterReset = Param{Name: "ShaftEncoderModuleReverseCounterReset", OriginalType: OriginalTypeGenApiICommand, Description: "Resets the reverse counter to 0.", Group: ShaftEncoderModule}

	// FrequencyConverter - This category includes items used to control the operation of the camera's frequency converter module
	FrequencyConverterInputSource        = Param{Name: "FrequencyConverterInputSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the input source.", Group: FrequencyConverter, OriginalEnumType: "FrequencyConverterInputSourceEnums"}
	FrequencyConverterSignalAlignment    = Param{Name: "FrequencyConverterSignalAlignment", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the signal transition relationships between received and generated signals.", Group: FrequencyConverter, OriginalEnumType: "FrequencyConverterSignalAlignmentEnums"}
	FrequencyConverterPreDivider         = Param{Name: "FrequencyConverterPreDivider", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the pre-divider value for the pre-divider sub-module.", Group: FrequencyConverter}
	FrequencyConverterMultiplier         = Param{Name: "FrequencyConverterMultiplier", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the multiplier value for the multiplier sub-module.", Group: FrequencyConverter}
	FrequencyConverterPostDivider        = Param{Name: "FrequencyConverterPostDivider", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the post-divider value for the post-divider sub-module.", Group: FrequencyConverter}
	FrequencyConverterPreventOvertrigger = Param{Name: "FrequencyConverterPreventOvertrigger", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables overtriggering protection.", Group: FrequencyConverter}

	// TimerControls - This category includes items used to control the operation of the camera's timers
	TimerDurationTimebaseAbs = Param{Name: "TimerDurationTimebaseAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the time base (in microseconds) that is used when a timer duration is set with the 'timer duration raw' setting.", Group: TimerControls}
	TimerDelayTimebaseAbs    = Param{Name: "TimerDelayTimebaseAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the time base (in microseconds) that is used when a timer delay is set with the 'timer delay raw' setting.", Group: TimerControls}
	TimerSelector            = Param{Name: "TimerSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the timer to configure.", Group: TimerControls, OriginalEnumType: "TimerSelectorEnums"}
	TimerDurationAbs         = Param{Name: "TimerDurationAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Directly sets the duration for the selected timer in microseconds.", Group: TimerControls}
	TimerDurationRaw         = Param{Name: "TimerDurationRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the 'raw' duration for the selected timer.", Group: TimerControls}
	TimerDelayAbs            = Param{Name: "TimerDelayAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Directly sets the delay for the selected timer in microseconds.", Group: TimerControls}
	TimerDelayRaw            = Param{Name: "TimerDelayRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the 'raw' delay for the selected timer.", Group: TimerControls}
	TimerTriggerSource       = Param{Name: "TimerTriggerSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the internal camera signal used to trigger the selected timer.", Group: TimerControls, OriginalEnumType: "TimerTriggerSourceEnums"}
	TimerTriggerActivation   = Param{Name: "TimerTriggerActivation", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the type of signal transistion that will start the timer.", Group: TimerControls, OriginalEnumType: "TimerTriggerActivationEnums"}
	CounterSelector          = Param{Name: "CounterSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the counter to configure.", Group: TimerControls, OriginalEnumType: "CounterSelectorEnums"}
	CounterEventSource       = Param{Name: "CounterEventSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the event that will be the source to increment the counter.", Group: TimerControls, OriginalEnumType: "CounterEventSourceEnums"}
	CounterResetSource       = Param{Name: "CounterResetSource", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the source of the reset for the selected counter.", Group: TimerControls, OriginalEnumType: "CounterResetSourceEnums"}
	CounterReset             = Param{Name: "CounterReset", OriginalType: OriginalTypeGenApiICommand, Description: "Immediately resets the selected counter.", Group: TimerControls}

	// TimerSequence -
	TimerSequenceEnable            = Param{Name: "TimerSequenceEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}
	TimerSequenceLastEntryIndex    = Param{Name: "TimerSequenceLastEntryIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}
	TimerSequenceCurrentEntryIndex = Param{Name: "TimerSequenceCurrentEntryIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}
	TimerSequenceEntrySelector     = Param{Name: "TimerSequenceEntrySelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Guru", Group: TimerSequence, OriginalEnumType: "TimerSequenceEntrySelectorEnums"}
	TimerSequenceTimerSelector     = Param{Name: "TimerSequenceTimerSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Guru", Group: TimerSequence, OriginalEnumType: "TimerSequenceTimerSelectorEnums"}
	TimerSequenceTimerEnable       = Param{Name: "TimerSequenceTimerEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}
	TimerSequenceTimerInverter     = Param{Name: "TimerSequenceTimerInverter", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}
	TimerSequenceTimerDelayRaw     = Param{Name: "TimerSequenceTimerDelayRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}
	TimerSequenceTimerDurationRaw  = Param{Name: "TimerSequenceTimerDurationRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Guru", Group: TimerSequence}

	// LUTControls - This category includes items used to control the operation of the camera's lookup table (LUT)
	LUTEnable   = Param{Name: "LUTEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the selected LUT.", Group: LUTControls}
	LUTIndex    = Param{Name: "LUTIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the LUT element to access.", Group: LUTControls}
	LUTValue    = Param{Name: "LUTValue", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the value of the LUT element at the LUT index.", Group: LUTControls}
	LUTValueAll = Param{Name: "LUTValueAll", OriginalType: OriginalTypeGenApiIRegister, Description: "Accesses the entire content of the selected LUT in one chunk access.", Group: LUTControls}
	LUTSelector = Param{Name: "LUTSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the lookup table (LUT) to configure.", Group: LUTControls, OriginalEnumType: "LUTSelectorEnums"}

	// TransportLayer - This category includes items related to the GigE Vision transport layer
	PixelFormatLegacy                         = Param{Name: "PixelFormatLegacy", OriginalType: OriginalTypeGenApiIBoolean, Description: "Select legacy pixel format encoding.", Group: TransportLayer}
	PayloadSize                               = Param{Name: "PayloadSize", OriginalType: OriginalTypeGenApiIInteger, Description: "Size of the payload in bytes.", Group: TransportLayer}
	GevInterfaceSelector                      = Param{Name: "GevInterfaceSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the physical network interface to configure.", Group: TransportLayer, OriginalEnumType: "GevInterfaceSelectorEnums"}
	GevVersionMajor                           = Param{Name: "GevVersionMajor", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the major version number of the GigE Vision specification supported by this device.", Group: TransportLayer}
	GevVersionMinor                           = Param{Name: "GevVersionMinor", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the minor version number of the GigE Vision specification supported by this device.", Group: TransportLayer}
	GevDeviceModeIsBigEndian                  = Param{Name: "GevDeviceModeIsBigEndian", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates the endianess of the bootstrap registers.", Group: TransportLayer}
	GevDeviceModeCharacterSet                 = Param{Name: "GevDeviceModeCharacterSet", OriginalType: OriginalTypeGenApiIInteger, Description: "Indictes the character set.", Group: TransportLayer}
	GevMACAddress                             = Param{Name: "GevMACAddress", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the MAC address for the selected network interface.", Group: TransportLayer}
	GevSupportedIPConfigurationLLA            = Param{Name: "GevSupportedIPConfigurationLLA", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether the selected network interface supports auto IP addressing (also known as LLA)", Group: TransportLayer}
	GevSupportedIPConfigurationDHCP           = Param{Name: "GevSupportedIPConfigurationDHCP", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether the selected network interface supports DHCP IP addressing.", Group: TransportLayer}
	GevSupportedIPConfigurationPersistentIP   = Param{Name: "GevSupportedIPConfigurationPersistentIP", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether the selected network interface supports fixed IP addressing (also known as persistent IP addressing)", Group: TransportLayer}
	GevCurrentIPConfiguration                 = Param{Name: "GevCurrentIPConfiguration", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the current IP configuration of the selected network interface.", Group: TransportLayer}
	GevCurrentIPAddress                       = Param{Name: "GevCurrentIPAddress", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the current IP address for the selected network interface.", Group: TransportLayer}
	GevCurrentSubnetMask                      = Param{Name: "GevCurrentSubnetMask", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the current subnet mask for the selected network interface.", Group: TransportLayer}
	GevCurrentDefaultGateway                  = Param{Name: "GevCurrentDefaultGateway", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the current default gateway for the selected network interface.", Group: TransportLayer}
	GevPersistentIPAddress                    = Param{Name: "GevPersistentIPAddress", OriginalType: OriginalTypeGenApiIInteger, Description: "If fixed (persistent) IP addressing is supported by the device and enabled, sets the fixed IP address for the selected network interface.", Group: TransportLayer}
	GevPersistentSubnetMask                   = Param{Name: "GevPersistentSubnetMask", OriginalType: OriginalTypeGenApiIInteger, Description: "If fixed (persistent) IP addressing is supported by the device and enabled, sets the fixed subnet mask for the selected network interface.", Group: TransportLayer}
	GevPersistentDefaultGateway               = Param{Name: "GevPersistentDefaultGateway", OriginalType: OriginalTypeGenApiIInteger, Description: "If fixed (persistent) IP addressing is supported by the device and enabled, sets the fixed default gateway for the selected network interface.", Group: TransportLayer}
	GevLinkSpeed                              = Param{Name: "GevLinkSpeed", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the connection speed in Mbps for the selected network interface.", Group: TransportLayer}
	GevLinkMaster                             = Param{Name: "GevLinkMaster", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether the selected network interface is the clock master.", Group: TransportLayer}
	GevLinkFullDuplex                         = Param{Name: "GevLinkFullDuplex", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether the selected network interface operates in full-duplex mode.", Group: TransportLayer}
	GevLinkCrossover                          = Param{Name: "GevLinkCrossover", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates the state of medium-dependent interface crossover (MDIX) for the selected network interface.", Group: TransportLayer}
	GevFirstURL                               = Param{Name: "GevFirstURL", OriginalType: OriginalTypeGenApiIString, Description: "Indicates the first URL to the XML device description file.", Group: TransportLayer}
	GevSecondURL                              = Param{Name: "GevSecondURL", OriginalType: OriginalTypeGenApiIString, Description: "Indicates the second URL to the XML device description file.", Group: TransportLayer}
	GevNumberOfInterfaces                     = Param{Name: "GevNumberOfInterfaces", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the number of network interfaces on the device.", Group: TransportLayer}
	GevMessageChannelCount                    = Param{Name: "GevMessageChannelCount", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the number of message channels supported by the device.", Group: TransportLayer}
	GevStreamChannelCount                     = Param{Name: "GevStreamChannelCount", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the number of stream channels supported by the device.", Group: TransportLayer}
	GevSupportedOptionalCommandsEVENTDATA     = Param{Name: "GevSupportedOptionalCommandsEVENTDATA", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether EVENTDATA_CMD and EVENTDATA_ACK are supported.", Group: TransportLayer}
	GevSupportedOptionalCommandsEVENT         = Param{Name: "GevSupportedOptionalCommandsEVENT", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether EVENT_CMD and EVENT_ACK are supported.", Group: TransportLayer}
	GevSupportedOptionalCommandsPACKETRESEND  = Param{Name: "GevSupportedOptionalCommandsPACKETRESEND", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether PACKETRESEND_CMD is supported.", Group: TransportLayer}
	GevSupportedOptionalCommandsWRITEMEM      = Param{Name: "GevSupportedOptionalCommandsWRITEMEM", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether WRITEMEM_CMD and WRITEMEM_ACK are supported.", Group: TransportLayer}
	GevSupportedOptionalCommandsConcatenation = Param{Name: "GevSupportedOptionalCommandsConcatenation", OriginalType: OriginalTypeGenApiIBoolean, Description: "Indicates whether multiple operations in a single message are supported.", Group: TransportLayer}
	GevHeartbeatTimeout                       = Param{Name: "GevHeartbeatTimeout", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the heartbeat timeout in milliseconds.", Group: TransportLayer}
	GevTimestampTickFrequency                 = Param{Name: "GevTimestampTickFrequency", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the number of timestamp clock ticks in 1 second.", Group: TransportLayer}
	GevTimestampControlReset                  = Param{Name: "GevTimestampControlReset", OriginalType: OriginalTypeGenApiICommand, Description: "Resets the timestamp value for the device.", Group: TransportLayer}
	GevTimestampControlLatch                  = Param{Name: "GevTimestampControlLatch", OriginalType: OriginalTypeGenApiICommand, Description: "Latches the current timestamp value of the device.", Group: TransportLayer}
	GevTimestampControlLatchReset             = Param{Name: "GevTimestampControlLatchReset", OriginalType: OriginalTypeGenApiICommand, Description: "Resets the timestamp control latch.", Group: TransportLayer}
	GevTimestampValue                         = Param{Name: "GevTimestampValue", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the latched value of the timestamp.", Group: TransportLayer}
	GevCCP                                    = Param{Name: "GevCCP", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the control channel privilege feature.", Group: TransportLayer, OriginalEnumType: "GevCCPEnums"}
	GevStreamChannelSelector                  = Param{Name: "GevStreamChannelSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the stream channel to configure.", Group: TransportLayer, OriginalEnumType: "GevStreamChannelSelectorEnums"}
	GevSCPInterfaceIndex                      = Param{Name: "GevSCPInterfaceIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the index of the network interface to use.", Group: TransportLayer}
	GevSCPHostPort                            = Param{Name: "GevSCPHostPort", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the port to which the device must send data streams.", Group: TransportLayer}
	GevSCPSFireTestPacket                     = Param{Name: "GevSCPSFireTestPacket", OriginalType: OriginalTypeGenApiICommand, Description: "<b>Visibility</b> = Guru", Group: TransportLayer}
	GevSCPSDoNotFragment                      = Param{Name: "GevSCPSDoNotFragment", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Guru", Group: TransportLayer}
	GevSCPSBigEndian                          = Param{Name: "GevSCPSBigEndian", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Guru", Group: TransportLayer}
	GevSCPSPacketSize                         = Param{Name: "GevSCPSPacketSize", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the packet size in bytes for the selected stream channel.", Group: TransportLayer}
	GevSCDA                                   = Param{Name: "GevSCDA", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the stream channel destination IPv4 address for the selected stream channel.", Group: TransportLayer}
	GevSCPD                                   = Param{Name: "GevSCPD", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the inter-packet delay (in ticks) for the selected stream channel.", Group: TransportLayer}
	GevSCFTD                                  = Param{Name: "GevSCFTD", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the frame transfer start delay (in ticks) for the selected stream channel.", Group: TransportLayer}
	GevSCBWR                                  = Param{Name: "GevSCBWR", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets a percentage of the Ethernet bandwidth assigned to the camera to be held in reserve.", Group: TransportLayer}
	GevSCBWRA                                 = Param{Name: "GevSCBWRA", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets a multiplier for the Bandwidth Reserve parameter.", Group: TransportLayer}
	GevSCBWA                                  = Param{Name: "GevSCBWA", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the bandwidth (in bytes per second) that will be used by the camera to transmit image and chunk feature data and to handle resends and control data transmissions.", Group: TransportLayer}
	GevSCDMT                                  = Param{Name: "GevSCDMT", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the maximum amount of data (in bytes per second) that the camera could generate given its current settings and ideal conditions, i.e., unlimited bandwidth and no packet resends.", Group: TransportLayer}
	GevSCDCT                                  = Param{Name: "GevSCDCT", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the actual bandwidth (in bytes per second) that the camera will use to transmit image data and chunk data given the current AOI settings, chunk feature settings, and the pixel format setting.", Group: TransportLayer}
	GevSCFJM                                  = Param{Name: "GevSCFJM", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the maximum time (in ticks) that the next frame transmission could be delayed due to a burst of resends.", Group: TransportLayer}
	TLParamsLocked                            = Param{Name: "TLParamsLocked", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates whether a live grab is under way.", Group: TransportLayer}

	// ActionControl - This category includes items that control the action control feature
	ActionSelector     = Param{Name: "ActionSelector", OriginalType: OriginalTypeGenApiIInteger, Description: "Selects the action command to configure.", Group: ActionControl}
	ActionDeviceKey    = Param{Name: "ActionDeviceKey", OriginalType: OriginalTypeGenApiIInteger, Description: "Authorization key.", Group: ActionControl}
	ActionCommandCount = Param{Name: "ActionCommandCount", OriginalType: OriginalTypeGenApiIInteger, Description: "Number of action command interfaces.", Group: ActionControl}
	ActionGroupKey     = Param{Name: "ActionGroupKey", OriginalType: OriginalTypeGenApiIInteger, Description: "Defines a group of devices.", Group: ActionControl}
	ActionGroupMask    = Param{Name: "ActionGroupMask", OriginalType: OriginalTypeGenApiIInteger, Description: "Filters out particular devices from its group.", Group: ActionControl}

	// DeviceControl -
	DeviceRegistersStreamingStart = Param{Name: "DeviceRegistersStreamingStart", OriginalType: OriginalTypeGenApiICommand, Description: "Prepare the device for registers streaming.", Group: DeviceControl}
	DeviceRegistersStreamingEnd   = Param{Name: "DeviceRegistersStreamingEnd", OriginalType: OriginalTypeGenApiICommand, Description: "Announce the end of registers streaming.", Group: DeviceControl}

	// UserSets - This category includes items that control the configuration sets feature that is used to save sets of parameters in the camera
	UserSetLoad            = Param{Name: "UserSetLoad", OriginalType: OriginalTypeGenApiICommand, Description: "Loads the selected configuration into the camera's volatile memory and makes it the active configuration set.", Group: UserSets}
	UserSetSave            = Param{Name: "UserSetSave", OriginalType: OriginalTypeGenApiICommand, Description: "Saves the current active configuration set into the selected user set.", Group: UserSets}
	UserSetSelector        = Param{Name: "UserSetSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the configuration set to load, save, or configure.", Group: UserSets, OriginalEnumType: "UserSetSelectorEnums"}
	UserSetDefaultSelector = Param{Name: "UserSetDefaultSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the configuration set to be used as the default startup set.", Group: UserSets, OriginalEnumType: "UserSetDefaultSelectorEnums"}
	DefaultSetSelector     = Param{Name: "DefaultSetSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the which factory setting will be used as default set.", Group: UserSets, OriginalEnumType: "DefaultSetSelectorEnums"}

	// AutoFunctions - This category includes items that parameterize the Auto Functions
	AutoTargetValue                  = Param{Name: "AutoTargetValue", OriginalType: OriginalTypeGenApiIInteger, Description: "Target average grey value for Gain Auto and Exposure Auto.", Group: AutoFunctions}
	GrayValueAdjustmentDampingAbs    = Param{Name: "GrayValueAdjustmentDampingAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Gray value adjustment damping for Gain Auto and Exposure Auto.", Group: AutoFunctions}
	GrayValueAdjustmentDampingRaw    = Param{Name: "GrayValueAdjustmentDampingRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Gray value adjustment damping for Gain Auto and Exposure Auto.", Group: AutoFunctions}
	BalanceWhiteAdjustmentDampingAbs = Param{Name: "BalanceWhiteAdjustmentDampingAbs", OriginalType: OriginalTypeGenApiIFloat, Description: "Balance White adjustment damping for Balance White Auto.", Group: AutoFunctions}
	BalanceWhiteAdjustmentDampingRaw = Param{Name: "BalanceWhiteAdjustmentDampingRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Balance White adjustment damping for Balance White Auto.", Group: AutoFunctions}
	AutoGainRawLowerLimit            = Param{Name: "AutoGainRawLowerLimit", OriginalType: OriginalTypeGenApiIInteger, Description: "Lower limit of the Auto Gain (Raw) parameter.", Group: AutoFunctions}
	AutoGainRawUpperLimit            = Param{Name: "AutoGainRawUpperLimit", OriginalType: OriginalTypeGenApiIInteger, Description: "Upper limit of the Auto Gain (Raw) parameter.", Group: AutoFunctions}
	AutoExposureTimeAbsLowerLimit    = Param{Name: "AutoExposureTimeAbsLowerLimit", OriginalType: OriginalTypeGenApiIFloat, Description: "Lower limit of the Auto Exposure Time (Abs) [us] parameter.", Group: AutoFunctions}
	AutoExposureTimeAbsUpperLimit    = Param{Name: "AutoExposureTimeAbsUpperLimit", OriginalType: OriginalTypeGenApiIFloat, Description: "Upper limit of the Auto Exposure Time (Abs) [us] parameter.", Group: AutoFunctions}
	AutoFunctionProfile              = Param{Name: "AutoFunctionProfile", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the strategy for controlling gain and shutter simultaneously.", Group: AutoFunctions, OriginalEnumType: "AutoFunctionProfileEnums"}

	// AutoFunctionAOIs - Portion of the sensor array used for auto function control
	AutoFunctionAOIWidth                   = Param{Name: "AutoFunctionAOIWidth", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the width of the auto function area of interest in pixels.", Group: AutoFunctionAOIs}
	AutoFunctionAOIHeight                  = Param{Name: "AutoFunctionAOIHeight", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the height of the auto function area of interest in pixels.", Group: AutoFunctionAOIs}
	AutoFunctionAOIOffsetX                 = Param{Name: "AutoFunctionAOIOffsetX", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the starting column of the auto function area of interest in pixels.", Group: AutoFunctionAOIs}
	AutoFunctionAOIOffsetY                 = Param{Name: "AutoFunctionAOIOffsetY", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the starting line of the auto function area of interest in pixels.", Group: AutoFunctionAOIs}
	AutoFunctionAOIUsageIntensity          = Param{Name: "AutoFunctionAOIUsageIntensity", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Beginner", Group: AutoFunctionAOIs}
	AutoFunctionAOIUsageWhiteBalance       = Param{Name: "AutoFunctionAOIUsageWhiteBalance", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Beginner", Group: AutoFunctionAOIs}
	AutoFunctionAOIUsageRedLightCorrection = Param{Name: "AutoFunctionAOIUsageRedLightCorrection", OriginalType: OriginalTypeGenApiIBoolean, Description: "<b>Visibility</b> = Beginner", Group: AutoFunctionAOIs}
	AutoFunctionAOISelector                = Param{Name: "AutoFunctionAOISelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the Auto Function AOI.", Group: AutoFunctionAOIs, OriginalEnumType: "AutoFunctionAOISelectorEnums"}

	// ColorOverexposureCompensation - Compensates for deviations of hue resulting from overexposure
	ColorOverexposureCompensationAOISelector  = Param{Name: "ColorOverexposureCompensationAOISelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selcts the AOI for color overexposure compensation.", Group: ColorOverexposureCompensation, OriginalEnumType: "ColorOverexposureCompensationAOISelectorEnums"}
	ColorOverexposureCompensationAOIEnable    = Param{Name: "ColorOverexposureCompensationAOIEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables color overexposure compensation.", Group: ColorOverexposureCompensation}
	ColorOverexposureCompensationAOIFactor    = Param{Name: "ColorOverexposureCompensationAOIFactor", OriginalType: OriginalTypeGenApiIFloat, Description: "Sets the color overexposure compensation factor for the selected C.O.C.", Group: ColorOverexposureCompensation}
	ColorOverexposureCompensationAOIFactorRaw = Param{Name: "ColorOverexposureCompensationAOIFactorRaw", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the raw value for the color overexposure compensation factor.", Group: ColorOverexposureCompensation}
	ColorOverexposureCompensationAOIWidth     = Param{Name: "ColorOverexposureCompensationAOIWidth", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the width for the selected C.O.C.", Group: ColorOverexposureCompensation}
	ColorOverexposureCompensationAOIHeight    = Param{Name: "ColorOverexposureCompensationAOIHeight", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the height for the selected C.O.C.", Group: ColorOverexposureCompensation}
	ColorOverexposureCompensationAOIOffsetX   = Param{Name: "ColorOverexposureCompensationAOIOffsetX", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the X offset for the selected C.O.C.", Group: ColorOverexposureCompensation}
	ColorOverexposureCompensationAOIOffsetY   = Param{Name: "ColorOverexposureCompensationAOIOffsetY", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the Y offset for the selected C.O.C.", Group: ColorOverexposureCompensation}

	// Shading - Includes items used to control the operation of the camera's shading correction.
	ShadingSelector           = Param{Name: "ShadingSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the kind of shading correction.", Group: Shading, OriginalEnumType: "ShadingSelectorEnums"}
	ShadingEnable             = Param{Name: "ShadingEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the selected kind of shading correction.", Group: Shading}
	ShadingStatus             = Param{Name: "ShadingStatus", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Indicates error statuses related to shading correction.", Group: Shading, OriginalEnumType: "ShadingStatusEnums"}
	ShadingSetDefaultSelector = Param{Name: "ShadingSetDefaultSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the bootup shading set.", Group: Shading, OriginalEnumType: "ShadingSetDefaultSelectorEnums"}
	ShadingSetSelector        = Param{Name: "ShadingSetSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the shading set to which the activate command will be applied.", Group: Shading, OriginalEnumType: "ShadingSetSelectorEnums"}
	ShadingSetActivate        = Param{Name: "ShadingSetActivate", OriginalType: OriginalTypeGenApiICommand, Description: "Activates the selected shading set.", Group: Shading}
	ShadingSetCreate          = Param{Name: "ShadingSetCreate", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Creates a shading set.", Group: Shading, OriginalEnumType: "ShadingSetCreateEnums"}

	// UserDefinedValues -
	UserDefinedValueSelector = Param{Name: "UserDefinedValueSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "<b>Visibility</b> = Guru", Group: UserDefinedValues, OriginalEnumType: "UserDefinedValueSelectorEnums"}
	UserDefinedValue         = Param{Name: "UserDefinedValue", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Guru", Group: UserDefinedValues}

	// FeatureSets -
	GenicamXmlFileDefault = Param{Name: "GenicamXmlFileDefault", OriginalType: OriginalTypeGenApiIInteger, Description: "Select default genicam XML file.", Group: FeatureSets}
	FeatureSet            = Param{Name: "FeatureSet", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Select a camera description file.", Group: FeatureSets, OriginalEnumType: "FeatureSetEnums"}

	// RemoveParamLimits - This category includes items that allow removing the limits of camera parameters
	RemoveLimits      = Param{Name: "RemoveLimits", OriginalType: OriginalTypeGenApiIBoolean, Description: "Removes the factory-set limits of the selected parameter.", Group: RemoveParamLimits}
	ParameterSelector = Param{Name: "ParameterSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the parameter to configure.", Group: RemoveParamLimits, OriginalEnumType: "ParameterSelectorEnums"}
	Prelines          = Param{Name: "Prelines", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the number of prelines.", Group: RemoveParamLimits}

	// ExpertFeatureAccess -
	ExpertFeatureAccessSelector = Param{Name: "ExpertFeatureAccessSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the feature to configure.", Group: ExpertFeatureAccess, OriginalEnumType: "ExpertFeatureAccessSelectorEnums"}
	ExpertFeatureAccessKey      = Param{Name: "ExpertFeatureAccessKey", OriginalType: OriginalTypeGenApiIInteger, Description: "Sets the key to access the selected feature.", Group: ExpertFeatureAccess}
	ExpertFeatureEnable         = Param{Name: "ExpertFeatureEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enable the selected Feature.", Group: ExpertFeatureAccess}

	// ChunkDataStreams - This category includes items that control the chunk features available on the camera.
	ChunkModeActive = Param{Name: "ChunkModeActive", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the chunk mode.", Group: ChunkDataStreams}
	ChunkEnable     = Param{Name: "ChunkEnable", OriginalType: OriginalTypeGenApiIBoolean, Description: "Enables the inclusion of the selected chunk in the payload data.", Group: ChunkDataStreams}
	ChunkSelector   = Param{Name: "ChunkSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects chunks for enabling.", Group: ChunkDataStreams, OriginalEnumType: "ChunkSelectorEnums"}

	// ChunkData - This category includes items related to the chunk data that can be appended to the image data
	ChunkStride                              = Param{Name: "ChunkStride", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the number of bytes of data between the beginning of one line in the acquired image and the beginning of the next line in the acquired image.", Group: ChunkData}
	ChunkSequenceSetIndex                    = Param{Name: "ChunkSequenceSetIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the sequence set index number related to the acquired image.", Group: ChunkData}
	ChunkOffsetX                             = Param{Name: "ChunkOffsetX", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the X offset of the area of interest represented in the acquired image.", Group: ChunkData}
	ChunkOffsetY                             = Param{Name: "ChunkOffsetY", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the Y offset of the area of interest represented in the acquired image.", Group: ChunkData}
	ChunkWidth                               = Param{Name: "ChunkWidth", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the widtth of the area of interest represented in the acquired image.", Group: ChunkData}
	ChunkHeight                              = Param{Name: "ChunkHeight", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the height of the area of interest represented in the acquired image.", Group: ChunkData}
	ChunkDynamicRangeMin                     = Param{Name: "ChunkDynamicRangeMin", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the minimum possible pixel value in the acquired image.", Group: ChunkData}
	ChunkDynamicRangeMax                     = Param{Name: "ChunkDynamicRangeMax", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the maximum possible pixel value in the acquired image.", Group: ChunkData}
	ChunkPixelFormat                         = Param{Name: "ChunkPixelFormat", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Indicates the format of the pixel data in the acquired image.", Group: ChunkData, OriginalEnumType: "ChunkPixelFormatEnums"}
	ChunkTimestamp                           = Param{Name: "ChunkTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the value of the timestamp when the image was acquired.", Group: ChunkData}
	ChunkFramecounter                        = Param{Name: "ChunkFramecounter", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the value of the frame counter when the image was acquired.", Group: ChunkData}
	ChunkLineStatusAll                       = Param{Name: "ChunkLineStatusAll", OriginalType: OriginalTypeGenApiIInteger, Description: "A bit field that indicates the status of all of the camera's input and output lines when the image was acquired.", Group: ChunkData}
	ChunkVirtLineStatusAll                   = Param{Name: "ChunkVirtLineStatusAll", OriginalType: OriginalTypeGenApiIInteger, Description: "A bit field that indicates the status of all of the camera's virtual input and output lines when the image was acquired.", Group: ChunkData}
	ChunkTriggerinputcounter                 = Param{Name: "ChunkTriggerinputcounter", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the value of the trigger input counter when the image was acquired.", Group: ChunkData}
	ChunkLineTriggerIgnoredCounter           = Param{Name: "ChunkLineTriggerIgnoredCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}
	ChunkFrameTriggerIgnoredCounter          = Param{Name: "ChunkFrameTriggerIgnoredCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}
	ChunkLineTriggerEndToEndCounter          = Param{Name: "ChunkLineTriggerEndToEndCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}
	ChunkFrameTriggerCounter                 = Param{Name: "ChunkFrameTriggerCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}
	ChunkFramesPerTriggerCounter             = Param{Name: "ChunkFramesPerTriggerCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}
	ChunkInputStatusAtLineTriggerBitsPerLine = Param{Name: "ChunkInputStatusAtLineTriggerBitsPerLine", OriginalType: OriginalTypeGenApiIInteger, Description: "Number of bits per status.", Group: ChunkData}
	ChunkInputStatusAtLineTriggerIndex       = Param{Name: "ChunkInputStatusAtLineTriggerIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Used to select a certain status.", Group: ChunkData}
	ChunkInputStatusAtLineTriggerValue       = Param{Name: "ChunkInputStatusAtLineTriggerValue", OriginalType: OriginalTypeGenApiIInteger, Description: "Value of the status selected by 'Index'.", Group: ChunkData}
	ChunkShaftEncoderCounter                 = Param{Name: "ChunkShaftEncoderCounter", OriginalType: OriginalTypeGenApiIInteger, Description: "Shaft encoder counter at frame trigger.", Group: ChunkData}
	ChunkPayloadCRC16                        = Param{Name: "ChunkPayloadCRC16", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the value of CRC checksum.", Group: ChunkData}
	ChunkExposureTime                        = Param{Name: "ChunkExposureTime", OriginalType: OriginalTypeGenApiIFloat, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}
	ChunkGainAll                             = Param{Name: "ChunkGainAll", OriginalType: OriginalTypeGenApiIInteger, Description: "<b>Visibility</b> = Beginner", Group: ChunkData}

	// EventsGeneration - This category includes items that control event generation by the camera.
	EventSelector     = Param{Name: "EventSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the type of event for enabling.", Group: EventsGeneration, OriginalEnumType: "EventSelectorEnums"}
	EventNotification = Param{Name: "EventNotification", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Sets the notification type that will be sent to the host application for the selected event.", Group: EventsGeneration, OriginalEnumType: "EventNotificationEnums"}

	// ExposureEndEventData - This category includes items available for an exposure end event
	ExposureEndEventStreamChannelIndex = Param{Name: "ExposureEndEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an exposure end event.", Group: ExposureEndEventData}
	ExposureEndEventFrameID            = Param{Name: "ExposureEndEventFrameID", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the frame ID for an exposure end event.", Group: ExposureEndEventData}
	ExposureEndEventTimestamp          = Param{Name: "ExposureEndEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an exposure end event.", Group: ExposureEndEventData}

	// LineStartOvertriggerEventData - This category includes items available for an line start overtrigger event
	LineStartOvertriggerEventStreamChannelIndex = Param{Name: "LineStartOvertriggerEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an line start overtrigger event.", Group: LineStartOvertriggerEventData}
	LineStartOvertriggerEventTimestamp          = Param{Name: "LineStartOvertriggerEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an line start overtrigger event.", Group: LineStartOvertriggerEventData}

	// FrameStartOvertriggerEventData - This category includes items available for an frame start overtrigger event
	FrameStartOvertriggerEventStreamChannelIndex = Param{Name: "FrameStartOvertriggerEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an frame start overtrigger event.", Group: FrameStartOvertriggerEventData}
	FrameStartOvertriggerEventTimestamp          = Param{Name: "FrameStartOvertriggerEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an frame start overtrigger event.", Group: FrameStartOvertriggerEventData}

	// FrameStartEventData - This category includes items available for an frame start event
	FrameStartEventStreamChannelIndex = Param{Name: "FrameStartEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an frame start event.", Group: FrameStartEventData}
	FrameStartEventTimestamp          = Param{Name: "FrameStartEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an frame start event.", Group: FrameStartEventData}

	// AcquisitionStartEventData - This category includes items available for an acquisition start event
	AcquisitionStartEventStreamChannelIndex = Param{Name: "AcquisitionStartEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an acquisition start event.", Group: AcquisitionStartEventData}
	AcquisitionStartEventTimestamp          = Param{Name: "AcquisitionStartEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an acquisition start event.", Group: AcquisitionStartEventData}

	// AcquisitionStartOvertriggerEventData - This category includes items available for an acquisition start overtrigger event
	AcquisitionStartOvertriggerEventStreamChannelIndex = Param{Name: "AcquisitionStartOvertriggerEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an acquisition start overtrigger event.", Group: AcquisitionStartOvertriggerEventData}
	AcquisitionStartOvertriggerEventTimestamp          = Param{Name: "AcquisitionStartOvertriggerEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an Acquisition start overtrigger event.", Group: AcquisitionStartOvertriggerEventData}

	// FrameTimeoutEventData - This category includes items available for an frame timeout event
	FrameTimeoutEventStreamChannelIndex = Param{Name: "FrameTimeoutEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an frame timeout event.", Group: FrameTimeoutEventData}
	FrameTimeoutEventTimestamp          = Param{Name: "FrameTimeoutEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an frame timeout event.", Group: FrameTimeoutEventData}

	// EventOverrunEventData - This category includes items available for an event overrun event
	EventOverrunEventStreamChannelIndex = Param{Name: "EventOverrunEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an event overrun event.", Group: EventOverrunEventData}
	EventOverrunEventFrameID            = Param{Name: "EventOverrunEventFrameID", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the frame ID for an event overrun event.", Group: EventOverrunEventData}
	EventOverrunEventTimestamp          = Param{Name: "EventOverrunEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an event overrun event.", Group: EventOverrunEventData}

	// CriticalTemperatureEventData - This category includes items available for a critical temperature event
	CriticalTemperatureEventStreamChannelIndex = Param{Name: "CriticalTemperatureEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for a critical temperature event.", Group: CriticalTemperatureEventData}
	CriticalTemperatureEventTimestamp          = Param{Name: "CriticalTemperatureEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a critical temperature event.", Group: CriticalTemperatureEventData}

	// OverTemperatureEventData - This category includes items available for an over temperature event
	OverTemperatureEventStreamChannelIndex = Param{Name: "OverTemperatureEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an over temperature event.", Group: OverTemperatureEventData}
	OverTemperatureEventTimestamp          = Param{Name: "OverTemperatureEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for an over temperature event.", Group: OverTemperatureEventData}

	// Line1RisingEdgeEventData - This category includes items available for an io line 1 rising edge event
	Line1RisingEdgeEventStreamChannelIndex = Param{Name: "Line1RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io line 1 rising edge event.", Group: Line1RisingEdgeEventData}
	Line1RisingEdgeEventTimestamp          = Param{Name: "Line1RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a line 1 rising edge event.", Group: Line1RisingEdgeEventData}

	// Line2RisingEdgeEventData - This category includes items available for an io line 2 rising edge event
	Line2RisingEdgeEventStreamChannelIndex = Param{Name: "Line2RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io line 2 rising edge event.", Group: Line2RisingEdgeEventData}
	Line2RisingEdgeEventTimestamp          = Param{Name: "Line2RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a line 2 rising edge event.", Group: Line2RisingEdgeEventData}

	// Line3RisingEdgeEventData - This category includes items available for an io line 3 rising edge event
	Line3RisingEdgeEventStreamChannelIndex = Param{Name: "Line3RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io line 3 rising edge event.", Group: Line3RisingEdgeEventData}
	Line3RisingEdgeEventTimestamp          = Param{Name: "Line3RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a line 3 rising edge event.", Group: Line3RisingEdgeEventData}

	// Line4RisingEdgeEventData - This category includes items available for an io line 4 rising edge event
	Line4RisingEdgeEventStreamChannelIndex = Param{Name: "Line4RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io line 4 rising edge event.", Group: Line4RisingEdgeEventData}
	Line4RisingEdgeEventTimestamp          = Param{Name: "Line4RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a line 4 rising edge event.", Group: Line4RisingEdgeEventData}

	// VirtualLine1RisingEdgeEventData - This category includes items available for an io virtual line 1 rising edge event
	VirtualLine1RisingEdgeEventStreamChannelIndex = Param{Name: "VirtualLine1RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io virtual line 1 rising edge event.", Group: VirtualLine1RisingEdgeEventData}
	VirtualLine1RisingEdgeEventTimestamp          = Param{Name: "VirtualLine1RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a virtual line 1 rising edge event.", Group: VirtualLine1RisingEdgeEventData}

	// VirtualLine2RisingEdgeEventData - This category includes items available for an io virtual line 2 rising edge event
	VirtualLine2RisingEdgeEventStreamChannelIndex = Param{Name: "VirtualLine2RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io virtual line 2 rising edge event.", Group: VirtualLine2RisingEdgeEventData}
	VirtualLine2RisingEdgeEventTimestamp          = Param{Name: "VirtualLine2RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a virtual line 2 rising edge event.", Group: VirtualLine2RisingEdgeEventData}

	// VirtualLine3RisingEdgeEventData - This category includes items available for an io virtual line 3 rising edge event
	VirtualLine3RisingEdgeEventStreamChannelIndex = Param{Name: "VirtualLine3RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io virtual line 3 rising edge event.", Group: VirtualLine3RisingEdgeEventData}
	VirtualLine3RisingEdgeEventTimestamp          = Param{Name: "VirtualLine3RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a virtual line 3 rising edge event.", Group: VirtualLine3RisingEdgeEventData}

	// VirtualLine4RisingEdgeEventData - This category includes items available for an io virtual line 4 rising edge event
	VirtualLine4RisingEdgeEventStreamChannelIndex = Param{Name: "VirtualLine4RisingEdgeEventStreamChannelIndex", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the stream channel index for an io virtual line 4 rising edge event.", Group: VirtualLine4RisingEdgeEventData}
	VirtualLine4RisingEdgeEventTimestamp          = Param{Name: "VirtualLine4RisingEdgeEventTimestamp", OriginalType: OriginalTypeGenApiIInteger, Description: "Indicates the time stamp for a virtual line 4 rising edge event.", Group: VirtualLine4RisingEdgeEventData}

	// FileAccessControl - This category includes items used to conduct file operations
	FileSelector          = Param{Name: "FileSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "This feature selects the target file in the device.", Group: FileAccessControl, OriginalEnumType: "FileSelectorEnums"}
	FileOperationSelector = Param{Name: "FileOperationSelector", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the target operation for the selected file.", Group: FileAccessControl, OriginalEnumType: "FileOperationSelectorEnums"}
	FileOpenMode          = Param{Name: "FileOpenMode", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Selects the access mode in which a file is opened.", Group: FileAccessControl, OriginalEnumType: "FileOpenModeEnums"}
	FileAccessBuffer      = Param{Name: "FileAccessBuffer", OriginalType: OriginalTypeGenApiIRegister, Description: "Defines the intermediate access buffer.", Group: FileAccessControl}
	FileAccessOffset      = Param{Name: "FileAccessOffset", OriginalType: OriginalTypeGenApiIInteger, Description: "Controls the mapping between the device file storage and the FileAccessBuffer.", Group: FileAccessControl}
	FileAccessLength      = Param{Name: "FileAccessLength", OriginalType: OriginalTypeGenApiIInteger, Description: "Controls the mapping between the device file storage and the FileAccessBuffer.", Group: FileAccessControl}
	FileOperationStatus   = Param{Name: "FileOperationStatus", OriginalType: OriginalTypeGenApiIEnumerationT, Description: "Represents the file operation execution status.", Group: FileAccessControl, OriginalEnumType: "FileOperationStatusEnums"}
	FileOperationResult   = Param{Name: "FileOperationResult", OriginalType: OriginalTypeGenApiIInteger, Description: "Represents the file operation result.", Group: FileAccessControl}
	FileSize              = Param{Name: "FileSize", OriginalType: OriginalTypeGenApiIInteger, Description: "Represents the size of the selected file.", Group: FileAccessControl}
	FileOperationExecute  = Param{Name: "FileOperationExecute", OriginalType: OriginalTypeGenApiICommand, Description: "Executes the selected operation.", Group: FileAccessControl}
)

var (
	GigEParams = ParamList{
		SequenceEnable,
		SequenceAsyncRestart,
		SequenceAsyncAdvance,
		SequenceCurrentSet,
		SequenceSetTotalNumber,
		SequenceSetIndex,
		SequenceSetLoad,
		SequenceSetStore,
		SequenceSetExecutions,
		SequenceAdvanceMode,
		SequenceControlSelector,
		SequenceControlSource,
		SequenceAddressBitSelector,
		SequenceAddressBitSource,
		GainAuto,
		GainRaw,
		GainAbs,
		GainSelector,
		BlackLevelRaw,
		BlackLevelAbs,
		BlackLevelSelector,
		GammaEnable,
		GammaSelector,
		Gamma,
		DigitalShift,
		SubstrateVoltage,
		PixelFormat,
		PixelCoding,
		PixelSize,
		PixelColorFilter,
		ProcessedRawEnable,
		SpatialCorrection,
		SpatialCorrectionAmount,
		SpatialCorrectionStartingLine,
		SensorBitDepth,
		ReverseX,
		ReverseY,
		PixelDynamicRangeMin,
		PixelDynamicRangeMax,
		FieldOutputMode,
		TestImageSelector,
		SensorDigitizationTaps,
		SensorWidth,
		SensorHeight,
		WidthMax,
		HeightMax,
		DeviceVendorName,
		DeviceModelName,
		DeviceManufacturerInfo,
		DeviceVersion,
		DeviceFirmwareVersion,
		DeviceID,
		DeviceUserID,
		DeviceReset,
		DeviceScanType,
		LastError,
		ClearLastError,
		TemperatureSelector,
		TemperatureAbs,
		TemperatureState,
		CriticalTemperature,
		OverTemperature,
		ColorTransformationSelector,
		LightSourceSelector,
		ColorTransformationMatrixFactor,
		ColorTransformationMatrixFactorRaw,
		ColorTransformationValueSelector,
		ColorTransformationValue,
		ColorTransformationValueRaw,
		ColorAdjustmentEnable,
		ColorAdjustmentReset,
		ColorAdjustmentSelector,
		ColorAdjustmentHue,
		ColorAdjustmentHueRaw,
		ColorAdjustmentSaturation,
		ColorAdjustmentSaturationRaw,
		BalanceWhiteReset,
		BalanceWhiteAuto,
		BalanceRatioSelector,
		BalanceRatioAbs,
		BalanceRatioRaw,
		Width,
		Height,
		OffsetX,
		OffsetY,
		CenterX,
		CenterY,
		BinningModeVertical,
		BinningModeHorizontal,
		BinningVertical,
		BinningHorizontal,
		LegacyBinningVertical,
		DecimationVertical,
		DecimationHorizontal,
		StackedZoneImagingEnable,
		StackedZoneImagingIndex,
		StackedZoneImagingZoneEnable,
		StackedZoneImagingZoneOffsetY,
		StackedZoneImagingZoneHeight,
		EnableBurstAcquisition,
		AcquisitionMode,
		AcquisitionStart,
		AcquisitionStop,
		AcquisitionAbort,
		AcquisitionFrameCount,
		TriggerControlImplementation,
		TriggerSelector,
		TriggerMode,
		TriggerSoftware,
		TriggerSource,
		TriggerActivation,
		TriggerPartialClosingFrame,
		TriggerDelayAbs,
		TriggerDelaySource,
		TriggerDelayLineTriggerCount,
		ExposureStartDelayRaw,
		ExposureStartDelayAbs,
		ExposureAuto,
		ExposureTimeRaw,
		ExposureTimeAbs,
		ExposureTimeBaseAbs,
		ExposureTimeBaseAbsEnable,
		AcquisitionLineRateAbs,
		ResultingLinePeriodAbs,
		ResultingLineRateAbs,
		AcquisitionFrameRateAbs,
		AcquisitionFrameRateEnable,
		ResultingFramePeriodAbs,
		ResultingFrameRateAbs,
		ExposureMode,
		GlobalResetReleaseModeEnable,
		ShutterMode,
		ExposureOverlapTimeMaxRaw,
		InterlacedIntegrationMode,
		ExposureOverlapTimeMaxAbs,
		ReadoutTimeAbs,
		AcquisitionStatusSelector,
		AcquisitionStatus,
		FrameTimeoutEnable,
		FrameTimeoutAbs,
		LineSelector,
		LineInverter,
		LineTermination,
		LineDebouncerTimeRaw,
		LineDebouncerTimeAbs,
		MinOutPulseWidthRaw,
		MinOutPulseWidthAbs,
		LineStatus,
		LineStatusAll,
		UserOutputValue,
		UserOutputValueAll,
		UserOutputValueAllMask,
		SyncUserOutputValue,
		SyncUserOutputValueAll,
		LineMode,
		LineSource,
		LineLogic,
		LineFormat,
		UserOutputSelector,
		SyncUserOutputSelector,
		VInpSignalSource,
		VInpBitLength,
		VInpSamplingPoint,
		VInpSignalReadoutActivation,
		ShaftEncoderModuleLineSelector,
		ShaftEncoderModuleLineSource,
		ShaftEncoderModuleMode,
		ShaftEncoderModuleCounterMode,
		ShaftEncoderModuleCounter,
		ShaftEncoderModuleCounterMax,
		ShaftEncoderModuleCounterReset,
		ShaftEncoderModuleReverseCounterMax,
		ShaftEncoderModuleReverseCounterReset,
		FrequencyConverterInputSource,
		FrequencyConverterSignalAlignment,
		FrequencyConverterPreDivider,
		FrequencyConverterMultiplier,
		FrequencyConverterPostDivider,
		FrequencyConverterPreventOvertrigger,
		TimerDurationTimebaseAbs,
		TimerDelayTimebaseAbs,
		TimerSelector,
		TimerDurationAbs,
		TimerDurationRaw,
		TimerDelayAbs,
		TimerDelayRaw,
		TimerTriggerSource,
		TimerTriggerActivation,
		CounterSelector,
		CounterEventSource,
		CounterResetSource,
		CounterReset,
		TimerSequenceEnable,
		TimerSequenceLastEntryIndex,
		TimerSequenceCurrentEntryIndex,
		TimerSequenceEntrySelector,
		TimerSequenceTimerSelector,
		TimerSequenceTimerEnable,
		TimerSequenceTimerInverter,
		TimerSequenceTimerDelayRaw,
		TimerSequenceTimerDurationRaw,
		LUTEnable,
		LUTIndex,
		LUTValue,
		LUTValueAll,
		LUTSelector,
		PixelFormatLegacy,
		PayloadSize,
		GevInterfaceSelector,
		GevVersionMajor,
		GevVersionMinor,
		GevDeviceModeIsBigEndian,
		GevDeviceModeCharacterSet,
		GevMACAddress,
		GevSupportedIPConfigurationLLA,
		GevSupportedIPConfigurationDHCP,
		GevSupportedIPConfigurationPersistentIP,
		GevCurrentIPConfiguration,
		GevCurrentIPAddress,
		GevCurrentSubnetMask,
		GevCurrentDefaultGateway,
		GevPersistentIPAddress,
		GevPersistentSubnetMask,
		GevPersistentDefaultGateway,
		GevLinkSpeed,
		GevLinkMaster,
		GevLinkFullDuplex,
		GevLinkCrossover,
		GevFirstURL,
		GevSecondURL,
		GevNumberOfInterfaces,
		GevMessageChannelCount,
		GevStreamChannelCount,
		GevSupportedOptionalCommandsEVENTDATA,
		GevSupportedOptionalCommandsEVENT,
		GevSupportedOptionalCommandsPACKETRESEND,
		GevSupportedOptionalCommandsWRITEMEM,
		GevSupportedOptionalCommandsConcatenation,
		GevHeartbeatTimeout,
		GevTimestampTickFrequency,
		GevTimestampControlReset,
		GevTimestampControlLatch,
		GevTimestampControlLatchReset,
		GevTimestampValue,
		GevCCP,
		GevStreamChannelSelector,
		GevSCPInterfaceIndex,
		GevSCPHostPort,
		GevSCPSFireTestPacket,
		GevSCPSDoNotFragment,
		GevSCPSBigEndian,
		GevSCPSPacketSize,
		GevSCDA,
		GevSCPD,
		GevSCFTD,
		GevSCBWR,
		GevSCBWRA,
		GevSCBWA,
		GevSCDMT,
		GevSCDCT,
		GevSCFJM,
		TLParamsLocked,
		ActionSelector,
		ActionDeviceKey,
		ActionCommandCount,
		ActionGroupKey,
		ActionGroupMask,
		DeviceRegistersStreamingStart,
		DeviceRegistersStreamingEnd,
		UserSetLoad,
		UserSetSave,
		UserSetSelector,
		UserSetDefaultSelector,
		DefaultSetSelector,
		AutoTargetValue,
		GrayValueAdjustmentDampingAbs,
		GrayValueAdjustmentDampingRaw,
		BalanceWhiteAdjustmentDampingAbs,
		BalanceWhiteAdjustmentDampingRaw,
		AutoGainRawLowerLimit,
		AutoGainRawUpperLimit,
		AutoExposureTimeAbsLowerLimit,
		AutoExposureTimeAbsUpperLimit,
		AutoFunctionProfile,
		AutoFunctionAOIWidth,
		AutoFunctionAOIHeight,
		AutoFunctionAOIOffsetX,
		AutoFunctionAOIOffsetY,
		AutoFunctionAOIUsageIntensity,
		AutoFunctionAOIUsageWhiteBalance,
		AutoFunctionAOIUsageRedLightCorrection,
		AutoFunctionAOISelector,
		ColorOverexposureCompensationAOISelector,
		ColorOverexposureCompensationAOIEnable,
		ColorOverexposureCompensationAOIFactor,
		ColorOverexposureCompensationAOIFactorRaw,
		ColorOverexposureCompensationAOIWidth,
		ColorOverexposureCompensationAOIHeight,
		ColorOverexposureCompensationAOIOffsetX,
		ColorOverexposureCompensationAOIOffsetY,
		ShadingSelector,
		ShadingEnable,
		ShadingStatus,
		ShadingSetDefaultSelector,
		ShadingSetSelector,
		ShadingSetActivate,
		ShadingSetCreate,
		UserDefinedValueSelector,
		UserDefinedValue,
		GenicamXmlFileDefault,
		FeatureSet,
		RemoveLimits,
		ParameterSelector,
		Prelines,
		ExpertFeatureAccessSelector,
		ExpertFeatureAccessKey,
		ExpertFeatureEnable,
		ChunkModeActive,
		ChunkEnable,
		ChunkSelector,
		ChunkStride,
		ChunkSequenceSetIndex,
		ChunkOffsetX,
		ChunkOffsetY,
		ChunkWidth,
		ChunkHeight,
		ChunkDynamicRangeMin,
		ChunkDynamicRangeMax,
		ChunkPixelFormat,
		ChunkTimestamp,
		ChunkFramecounter,
		ChunkLineStatusAll,
		ChunkVirtLineStatusAll,
		ChunkTriggerinputcounter,
		ChunkLineTriggerIgnoredCounter,
		ChunkFrameTriggerIgnoredCounter,
		ChunkLineTriggerEndToEndCounter,
		ChunkFrameTriggerCounter,
		ChunkFramesPerTriggerCounter,
		ChunkInputStatusAtLineTriggerBitsPerLine,
		ChunkInputStatusAtLineTriggerIndex,
		ChunkInputStatusAtLineTriggerValue,
		ChunkShaftEncoderCounter,
		ChunkPayloadCRC16,
		ChunkExposureTime,
		ChunkGainAll,
		EventSelector,
		EventNotification,
		ExposureEndEventStreamChannelIndex,
		ExposureEndEventFrameID,
		ExposureEndEventTimestamp,
		LineStartOvertriggerEventStreamChannelIndex,
		LineStartOvertriggerEventTimestamp,
		FrameStartOvertriggerEventStreamChannelIndex,
		FrameStartOvertriggerEventTimestamp,
		FrameStartEventStreamChannelIndex,
		FrameStartEventTimestamp,
		AcquisitionStartEventStreamChannelIndex,
		AcquisitionStartEventTimestamp,
		AcquisitionStartOvertriggerEventStreamChannelIndex,
		AcquisitionStartOvertriggerEventTimestamp,
		FrameTimeoutEventStreamChannelIndex,
		FrameTimeoutEventTimestamp,
		EventOverrunEventStreamChannelIndex,
		EventOverrunEventFrameID,
		EventOverrunEventTimestamp,
		CriticalTemperatureEventStreamChannelIndex,
		CriticalTemperatureEventTimestamp,
		OverTemperatureEventStreamChannelIndex,
		OverTemperatureEventTimestamp,
		Line1RisingEdgeEventStreamChannelIndex,
		Line1RisingEdgeEventTimestamp,
		Line2RisingEdgeEventStreamChannelIndex,
		Line2RisingEdgeEventTimestamp,
		Line3RisingEdgeEventStreamChannelIndex,
		Line3RisingEdgeEventTimestamp,
		Line4RisingEdgeEventStreamChannelIndex,
		Line4RisingEdgeEventTimestamp,
		VirtualLine1RisingEdgeEventStreamChannelIndex,
		VirtualLine1RisingEdgeEventTimestamp,
		VirtualLine2RisingEdgeEventStreamChannelIndex,
		VirtualLine2RisingEdgeEventTimestamp,
		VirtualLine3RisingEdgeEventStreamChannelIndex,
		VirtualLine3RisingEdgeEventTimestamp,
		VirtualLine4RisingEdgeEventStreamChannelIndex,
		VirtualLine4RisingEdgeEventTimestamp,
		FileSelector,
		FileOperationSelector,
		FileOpenMode,
		FileAccessBuffer,
		FileAccessOffset,
		FileAccessLength,
		FileOperationStatus,
		FileOperationResult,
		FileSize,
		FileOperationExecute,
	}
)
