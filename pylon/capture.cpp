#if __GNUC__ < 6 && __cplusplus
/* Jetson has ubuntu 16.04 and gcc 5.4 (and golang 1.6, but it's unusable) */
#define JETSON 1

/* It rendered bad idea to add the option to `CGO_CPPFLAGS`. I dunno why. */
#pragma GCC diagnostic warning "-std=c++11"

/* Heh, even with c++11 it has no `nullptr`, damn. */
#ifndef nullptr
#define nullptr (0)
#endif

/* Pylon5 lacks some entries on arm64... */
/* Bus 002 Device 002: ID 2676:ba02 Basler AG ace */
#define BASLER_ACE_VENDOR_ID "2676"
#define BASLER_ACE_PRODUCT_ID "ba02"
#endif /* __GNUC__ < 6 */

#include <cstdlib>

// Include files to use the PYLON API.
#include <pylon/PylonIncludes.h>

#ifdef GIGE
#include <pylon/gige/BaslerGigEInstantCamera.h>
#define Basler_XCamera Basler_GigECamera
#define CBaslerXCamera Pylon::CBaslerGigEInstantCamera
#else
#include <pylon/usb/PylonUsbIncludes.h>
#define Basler_XCamera Basler_UsbCameraParams
#define CBaslerXCamera Pylon::CBaslerUsbInstantCamera
#endif
#include <pylon/InstantCamera.h>
#include <capture.h>

class CameraWrapper {
    public:
        static CameraWrapper& getInstance() {
            static CameraWrapper instance;
            return instance;
        }
        std::string stopCapture();
        std::string attachDevice();
        std::string openCamera();
        std::string closeCamera();
        bool isGrabbing();
        bool isAttached();
        bool isOpen();
        std::string retrieveAndSave(int, int, std::string);
        std::string startCapture(int max);
        std::string configureCamera();
        std::string setHardwareTriggerConfiguration();


        std::string setNodeMapIntParam(const GenICam::gcstring,int64_t);
        std::string setNodeMapFloatParam(const GenICam::gcstring, double);
        std::string setNodeMapEnumParam(const GenICam::gcstring,
					const GenICam::gcstring);
        std::string getNodeMapIntParam(const GenICam::gcstring, int64_t*);
        std::string getNodeMapFloatParam(const GenICam::gcstring, double*);
        std::string getNodeMapEnumParam(const GenICam::gcstring, GenICam::gcstring*);

	std::string fetch(int idx);

	static const int NoIntValue = -1;
    private:
        CameraWrapper() {}
        CameraWrapper(CameraWrapper const&);  // Don't Implement.
        void operator=(CameraWrapper const&); // Don't implement

        CBaslerXCamera camera;
        Pylon::PylonAutoInitTerm autoInitTerm;
	const std::string openState(std::string prefix) {
		return prefix + (isOpen() ? "open" : "closed");
	}
	inline GenApi::INode* getNode(const GenICam::gcstring name);

    private: // device info & getters
	Pylon::CInstantCamera::DeviceInfo_t info() { return this->camera.GetDeviceInfo(); }
    	int width, height;
	std::string fullName, vendorName, modelName, serialNumber,
		    deviceVersion, productId, vendorId;
    public:
	int Width() { return width; }
	int Height() { return height; }
	const char* FullName() { return fullName.c_str(); }
	const char* VendorName() { return vendorName.c_str(); }
	const char* ModelName() { return modelName.c_str(); }
	const char* SerialNumber() { return serialNumber.c_str(); }
	const char* DeviceVersion() { return deviceVersion.c_str(); }
	const char* ProductId() { return productId.c_str(); }
	const char* VendorId() { return vendorId.c_str(); }

    private:
	uint32_t countOfImagesToGrab = 100;
	uint32_t grabTimeout = 5000; // ms
    public:
	void setFetchTimeout(uint32_t v);
	void setFetchCount(uint32_t v);
};

static inline bool hasNoValue(std::string s) {
	return strncmp(s.c_str(), "<no ", 4) == 0;
}

static inline int hex2int(std::string s, int dflt) {
	return hasNoValue(s) ? dflt : (int)strtol(s.c_str(), nullptr, 16);
}

#define CAMERA(op, ...) CameraWrapper::getInstance().op(__VA_ARGS__)
#ifdef VERBOSE
#define TRACE(msg, ...) std::cerr << __func__ << "(" #__VA_ARGS__ "):" \
				  << ((msg) == "" ? "ok" : msg) \
				  << std::endl
#else
#define TRACE(msg, ...) ((void)0)
#endif
/******************************************************************************
 * EXTERNALLY VISIBLE BEGIN
 */

// The init/shutdown is normally performed by the Pylon::PylonAutoInitTerm
// These funcs are here for completeness only!
void pylonInitialize() {
	Pylon::PylonInitialize();
}

void pylonTerminate() {
	Pylon::PylonTerminate(true);
}

const char* stopCapture() {
    std::string msg = CAMERA(stopCapture); TRACE(msg);
    return msg.c_str();
}

const char* attachDevice() {
    std::string msg = CAMERA(attachDevice); TRACE(msg);
    return msg.c_str();
}

const char* configureCamera() {
    assert(isOpen());
    std::string msg = CAMERA(configureCamera); TRACE(msg);
    return msg.c_str();
}

const char* setHardwareTriggerConfiguration() {
    std::string msg = CAMERA(setHardwareTriggerConfiguration); TRACE(msg);
    return msg.c_str();
}

const char* fetch(int idx) {
    std::string msg = CAMERA(fetch, idx); TRACE(msg, idx);
    return msg.c_str();
}

const char* retrieveAndSave(int batch, int timeout, char* outputPath) {
    std::string op;
    op.assign(outputPath);
    std::string msg = CAMERA(retrieveAndSave, batch, timeout, op);
    TRACE(msg, batch, timeout, op);
    return msg.c_str();
}
const char* startCapture(int max) {
    std::string msg = CAMERA(startCapture, max); TRACE(msg, max);
    return msg.c_str();
}

const char* openCamera() {
    assert(isAttached());
    std::string msg = CAMERA(openCamera); TRACE(msg);
    return msg.c_str();
}

const char* closeCamera() {
    std::string msg = CAMERA(closeCamera); TRACE(msg);
    return msg.c_str();
}

bool isCameraGrabbing() { return CAMERA(isGrabbing); }
bool isAttached() { return CAMERA(isAttached); }
bool isOpen() { return CAMERA(isOpen); }

const char* setNodeMapIntParam(char* name, int value) {
    GenICam::gcstring op;
    op.assign(name);
    std::string msg = CAMERA(setNodeMapIntParam, op, (int64_t)value);
    TRACE(msg, op, (int64_t)value);
    return msg.c_str();
}
const char* setNodeMapFloatParam(char* name, double value) {
    GenICam::gcstring op;
    op.assign(name);
    std::string msg = CAMERA(setNodeMapFloatParam, op, value);
    TRACE(msg, op, value);
    return msg.c_str();
}
const char* setNodeMapEnumParam(char* name, char* value) {
    GenICam::gcstring op, v;
    op.assign(name);
    v.assign(value);
    std::string msg = CAMERA(setNodeMapEnumParam, op, v);
    TRACE(msg, op, v);
    return msg.c_str();
}

const char* getNodeMapIntParam(char *name, int64_t *value) {
    GenICam::gcstring op;
    op.assign(name);
    std::string msg = CAMERA(getNodeMapIntParam, op, (int64_t*)value);
    TRACE(msg, op, (int64_t)(*value));
    return msg.c_str();
}

const char* getNodeMapFloatParam(char *name, double *value) {
    GenICam::gcstring op;
    op.assign(name);
    std::string msg = CAMERA(getNodeMapFloatParam, op, value);
    TRACE(msg, op, *value);
    return msg.c_str();
}

const char* getNodeMapEnumParam(char *name, char *value) {
    GenICam::gcstring op, v;
    op.assign(name);
    std::string msg = CAMERA(getNodeMapEnumParam, op, &v);
    TRACE(msg, op, v);
    strcpy(value, v.c_str());
    return msg.c_str();
}

const char* fullName() { return CAMERA(FullName); }
const char* vendorName() { return CAMERA(VendorName); }
const char* modelName() { return CAMERA(ModelName); }
const char* serialNumber() { return CAMERA(SerialNumber); }
const char* deviceVersion() { return CAMERA(DeviceVersion); }

int productId() {
	return hex2int(CAMERA(ProductId), CameraWrapper::NoIntValue);
}
int vendorId() {
	return hex2int(CAMERA(VendorId), CameraWrapper::NoIntValue);
}

int width() { return CAMERA(Width); }
int height() { return CAMERA(Height); }

void setFetchTimeout(uint32_t v) { CAMERA(setFetchTimeout, v); }
void setFetchCount(uint32_t v) { CAMERA(setFetchCount, v); }
/*
 * EXTERNALLY VISIBLE END
 ******************************************************************************/

inline GenApi::INode* CameraWrapper::getNode(const GenICam::gcstring name) {
    GenApi::INode *p = this->camera.GetNodeMap().GetNode(name);
    // std::cerr << "getNode(" << name << "): " << p << ";" << std::endl;
    if (!p) std::cerr << "getNode(" << name << "): none found" << std::endl;
    return p;
}

static inline std::string ExceptionValue(const char* func,
					 GenICam::GenericException &e) {
    std::string msg = func;
    msg += ": ";
    msg += e.GetDescription();
    std::cerr << std::endl << "Exception: " << msg << std::endl;
    return msg;
}

std::string CameraWrapper::stopCapture() {
    try {
	camera.Close();
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::attachDevice() {
    try {
	this->camera.Attach(Pylon::CTlFactory::GetInstance().CreateFirstDevice());
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

#define INFOS(info, name) ((info).Is##name##Available() ? \
			   (info).Get##name().c_str() : \
			   ("<no "#name">"))

#define INFOI(info, name, defaultValue) ((info).Is##name##Available() ? \
					(info).Get##name().GetValue() : \
					(defaultValue))

#define OUTS(name) "\t"#name"=[" << this->name.c_str() << "]" << std::endl

std::string CameraWrapper::openCamera() {
    try {
	this->camera.Open();
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }

    GenApi::CIntegerPtr width(camera.GetNodeMap().GetNode("Width"));
    GenApi::CIntegerPtr height(camera.GetNodeMap().GetNode("Height"));
    this->width = width->GetValue();
    this->height = height->GetValue();

    Pylon::CInstantCamera::DeviceInfo_t info = this->info();

    this->fullName.assign(INFOS(info, FullName));
    this->vendorName.assign(INFOS(info, VendorName));
    this->modelName.assign(INFOS(info, ModelName));
    this->serialNumber.assign(INFOS(info, SerialNumber));
    this->deviceVersion.assign(INFOS(info, DeviceVersion));
#ifdef JETSON
    this->productId.assign(BASLER_ACE_VENDOR_ID);
    this->vendorId.assign(BASLER_ACE_PRODUCT_ID);
#else
    this->productId.assign(INFOS(info, ProductId));
    this->vendorId.assign(INFOS(info, VendorId));
#endif

#ifdef VERBOSE
    std::cerr << "Camera [" << this->Width() << "×" << this->Height() << "]" << std::endl
	      << OUTS(fullName)
	      << OUTS(vendorName)
	      << OUTS(modelName)
	      << OUTS(serialNumber)
	      << OUTS(deviceVersion)
	      << OUTS(productId)
	      << OUTS(vendorId)
	      ;
#endif
    return "";
}

std::string CameraWrapper::closeCamera() {
    try {
	this->camera.Close();
    } catch (GenICam::GenericException &e) {
        std::string msg = __func__; msg += ": "; msg += e.GetDescription();
	std::cerr << msg << std::endl;
	return msg;
    }
    this->width = this->height = -1;
    return "";
}

bool CameraWrapper::isGrabbing() {
    return this->camera.IsGrabbing();
}

bool CameraWrapper::isAttached() {
    return this->camera.IsPylonDeviceAttached();
}

bool CameraWrapper::isOpen() {
    return this->camera.IsOpen();
}

/******************************************************************************
 * BEGIN FETCH-RELATED STUFF
 */
static inline int C_fetch_callback(int idx, // callback index
				   int w, int h, int pxt, // image dimensions
				   int size, const void *buf) { // image data
	return Go_fetch_callback(idx, w, h, pxt, size, (char*)buf);
}

static int handleGrabResult(int idx, Pylon::CGrabResultPtr ptrGrabResult) {
    if (!ptrGrabResult->GrabSucceeded()) {
	return -1;
    }

    uint32_t w = ptrGrabResult->GetWidth();
    uint32_t h = ptrGrabResult->GetHeight();
    Pylon::EPixelType pt = ptrGrabResult->GetPixelType(); // it's enum e.i. int

    size_t ImageBufferSize = ptrGrabResult->GetImageSize();
    const uint8_t *pImageBuffer = (uint8_t *) ptrGrabResult->GetBuffer();

#ifdef VERBOSE
    std::cerr << "Taken image " << w << "x" << h
	      << std::endl
	      << "Gray value of first pixel: "
    		<< (uint32_t) pImageBuffer[0]
	      << std::endl;
#endif
    int rc = C_fetch_callback(idx,	// callback selector
			      w, h, (int)pt,	// image props
			      ImageBufferSize, pImageBuffer); // image data
#ifdef VERBOSE
    std::cerr << "Callback returned " << rc
	      << std::endl
	      << std::endl;
#endif
    return 0;
}

// CameraWrapper::fetch() receieves a callback index as the argument
std::string CameraWrapper::fetch(int idx) {
    try {
	Pylon::CGrabResultPtr ptrGrabResult;
	// this->camera.MaxNumBuffer = 10; // 10 is the default
	this->camera.StartGrabbing(this->countOfImagesToGrab);
    	while (this->camera.IsGrabbing()) {
	    this->camera.RetrieveResult(this->grabTimeout,
	    				ptrGrabResult,
					Pylon::TimeoutHandling_ThrowException);

	    if (handleGrabResult(idx, ptrGrabResult)) {
	    	std::cerr << "Error: "
			  << ptrGrabResult->GetErrorCode()
			  << " "
			  << ptrGrabResult->GetErrorDescription()
			  << std::endl;
	    }
	}
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

void CameraWrapper::setFetchTimeout(uint32_t v) {
    this->grabTimeout = v;
}

void CameraWrapper::setFetchCount(uint32_t v) {
    this->countOfImagesToGrab = v;
}

/*
 * END OF FETCH
 ******************************************************************************/

std::string CameraWrapper::retrieveAndSave(int batch,
                                           int timeout,
					   std::string outputPath) {
    Pylon::CGrabResultPtr ptrGrabResult;

    try {
        if(this->camera.RetrieveResult(timeout,
	                               ptrGrabResult,
				       Pylon::TimeoutHandling_ThrowException) &&
           ptrGrabResult->GrabSucceeded()) {
            Pylon::EPixelType pixelType = ptrGrabResult->GetPixelType();
            uint32_t width = ptrGrabResult->GetWidth();
            uint32_t height = ptrGrabResult->GetHeight();
            size_t paddingX = ptrGrabResult->GetPaddingX();
            size_t bufferSize = ptrGrabResult->GetImageSize();
            void* buffer = ptrGrabResult->GetBuffer();
            
            std::ostringstream nameStream;
            nameStream << outputPath
	               << batch
		       << "_"
		       << ptrGrabResult->GetTimeStamp()
		       << ".tiff";
            
            Pylon::CImagePersistence::Save(
                    Pylon::ImageFileFormat_Tiff,
                    nameStream.str().c_str(),
                    buffer,
                    bufferSize,
                    pixelType,
                    width,
                    height,
                    paddingX,
                    Pylon::ImageOrientation_TopDown,
		    NULL);
        } else {
            std::ostringstream result;
            result << ptrGrabResult->GetErrorCode()
	           << " "
		   << ptrGrabResult->GetErrorDescription();
	    std::cerr << "grab failed: " << result.str() << std::endl;
            return result.str();
        }
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::startCapture(int max) {
    try {
	if (max <= 0) {
            // Pylon::GrabLoop_ProvidedByInstantCamera
	    this->camera.StartGrabbing(Pylon::GrabStrategy_OneByOne,
	                               Pylon::GrabLoop_ProvidedByUser);
	} else {
	    this->camera.StartGrabbing(max,
	    				Pylon::GrabStrategy_OneByOne,
					Pylon::GrabLoop_ProvidedByInstantCamera);
	}
    } catch (GenICam::GenericException &e) {
        camera.Close();
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::configureCamera() {
    std::string fn = __func__;
    try {
        this->camera.GainAuto.SetValue(Basler_XCamera::GainAuto_Continuous);
    } catch (GenICam::GenericException &e) {
    	fn += ".GainAuto";
	return ExceptionValue(fn.c_str(), e);
    }
#ifdef GIGE
    this->camera.GainRaw.SetValue(1);
    this->camera.BlackLevelRaw.SetValue(90);
#else
    /*
    try {
        this->camera.Gain.SetValue(this->camera.Gain.GetMin());
    } catch (GenICam::GenericException &e) {
    	fn += ".Gain";
	return ExceptionValue(fn.c_str(), e);
    }
    try {
        this->camera.BlackLevel.SetValue(90.0);
    } catch (GenICam::GenericException &e) {
    	fn += ".BlackLevel";
	return ExceptionValue(fn.c_str(), e);
    }
    */
#endif
    try {
        this->camera.DigitalShift.SetValue(1);
    } catch (GenICam::GenericException &e) {
    	fn += ".DigitalShift";
	return ExceptionValue(fn.c_str(), e);
    }
    try {
        this->camera.ExposureAuto.SetValue(Basler_XCamera::ExposureAuto_Continuous);
    } catch (GenICam::GenericException &e) {
    	fn += ".ExposureAuto";
	return ExceptionValue(fn.c_str(), e);
    }
    return "";
}

#define CHECK_NODE(fn, node, name)		\
    if (!(node)) {				\
    	std::string msg = "";			\
	msg += fn;				\
	msg += ": no node '";			\
    	msg += name;				\
    	msg += "'";				\
	std::cerr << msg << std::endl;		\
	return msg;				\
    }

std::string CameraWrapper::setNodeMapIntParam(const GenICam::gcstring name,
					      int64_t value) {
    GenApi::INode* node = this->getNode(name);
    CHECK_NODE(__func__, node, name);
    try {
	GenApi::CIntegerPtr(node)->SetValue(value);
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::setNodeMapFloatParam(const GenICam::gcstring name,
						double value) {
    GenApi::INode* node = this->getNode(name);
    CHECK_NODE(__func__, node, name);
    try {
	GenApi::CFloatPtr(node)->SetValue(value);
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::setNodeMapEnumParam(const GenICam::gcstring name,
					       const GenICam::gcstring value) {
    GenApi::INode* node = this->getNode(name);
    CHECK_NODE(__func__, node, name);
    try {
	GenApi::CEnumerationPtr(node)->FromString(value);
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::getNodeMapIntParam(const GenICam::gcstring name,
					      int64_t *value) {
    GenApi::INode* node = this->getNode(name);
    CHECK_NODE(__func__, node, name);
    try {
	*value = GenApi::CIntegerPtr(node)->GetValue();
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::getNodeMapFloatParam(const GenICam::gcstring name,
						double *value) {
    GenApi::INode* node = this->getNode(name);
    CHECK_NODE(__func__, node, name);
    try {
	*value = GenApi::CFloatPtr(node)->GetValue();
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::getNodeMapEnumParam(const GenICam::gcstring name,
					       GenICam::gcstring *value) {
    GenApi::StringList_t Symbolics;
    GenApi::INode* node = this->getNode(name);
    CHECK_NODE(__func__, node, name);
    try {
	GenApi::CEnumerationPtr(node)->GetSymbolics(Symbolics);
	value->assign(Symbolics[0]); // ??? somehow it works... XXX
    } catch (GenICam::GenericException &e) {
	return ExceptionValue(__func__, e);
    }
    return "";
}

std::string CameraWrapper::setHardwareTriggerConfiguration() {
    GenApi::INodeMap& nodemap = this->camera.GetNodeMap();

    //Disable all trigger types.
    // Get required enumerations.
    GenApi::CEnumerationPtr triggerSelector( nodemap.GetNode("TriggerSelector"));

    // Check the available camera trigger mode(s) to select the appropriate one:
    // acquisition start trigger mode (used by previous cameras, i.e. for
    // cameras supporting only the legacy image acquisition control mode;
    // do not confuse with acquisition start command) or frame start trigger mode
    // (used by newer cameras, i.e. for cameras using the standard image
    // acquisition control mode, equivalent to the acquisition start trigger
    // mode in the legacy image acquisition control mode).
    Pylon::String_t triggerName( "FrameStart");
    if (!IsAvailable( triggerSelector->GetEntryByName(triggerName))) {
        triggerName = "AcquisitionStart";
        if (!IsAvailable( triggerSelector->GetEntryByName(triggerName))) {
            throw RUNTIME_EXCEPTION("Could not select trigger."
	    	" Neither FrameStart nor AcquisitionStart is available.");
        }
    }

    // Get all enumeration entries of Trigger Selector.
    GenApi::NodeList_t triggerSelectorEntries;
    triggerSelector->GetEntries( triggerSelectorEntries );

    // Turn Trigger Mode off For all Trigger Selector entries.
    for (GenApi::NodeList_t::iterator it = triggerSelectorEntries.begin();
    	 it != triggerSelectorEntries.end();
	 ++it) {
        // Set Trigger Mode to off if the trigger is available.
        GenApi::CEnumEntryPtr pEntry(*it);
        if (IsAvailable(pEntry)) {
            Pylon::String_t triggerNameOfEntry(pEntry->GetSymbolic());
            triggerSelector->FromString( triggerNameOfEntry);
            if (triggerName == triggerNameOfEntry) {
                // Activate trigger.
                this->setNodeMapEnumParam("TriggerMode", "On");

                // The trigger source must be set to the trigger input, e.g.
		// 'Line1'.
                // GenApi::CEnumerationPtr(nodemap.GetNode("TriggerSource")
		//			  )->FromString("Line1");
                this->setNodeMapEnumParam("TriggerSource", "Line1");

                // The trigger activation must be set to e.g. 'RisingEdge'.
                // GenApi::CEnumerationPtr(nodemap.GetNode("TriggerActivation")
		//			  )->FromString("RisingEdge");
                this->setNodeMapEnumParam("TriggerActivation", "RisingEdge");
            } else {
                this->setNodeMapEnumParam("TriggerMode", "Off");
            }
        }
    }

    //Set acquisition mode.
    this->setNodeMapEnumParam("AcquisitionMode", "Continuous");

    return "";
}
