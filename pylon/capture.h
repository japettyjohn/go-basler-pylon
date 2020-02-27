#ifndef CAPTURE_H
#define CAPTURE_H

#undef VERBOSE /* one may #define it here ... */

#ifdef __cplusplus 
extern "C" {
#else
#include <stdbool.h>
#include <stdint.h>
#endif

// prototypes
void pylonInitialize();	// one have not to call these functions
void pylonTerminate();	// normally

const char* attachDevice();
const char* configureCamera();
const char* openCamera();
const char* startCapture(int max);
const char* retrieveAndSave(int batch, int timeout, char* outputPath);
const char* stopCapture();
const char* closeCamera();

bool isCameraGrabbing();
bool isAttached();
bool isOpen();

const char* setHardwareTriggerConfiguration();
const char* setNodeMapIntParam(char* name, int value);
const char* setNodeMapFloatParam(char* name, double value);
const char* setNodeMapEnumParam(char* name, char* value);
const char* getNodeMapIntParam(char *name, int64_t *value);
const char* getNodeMapFloatParam(char *name, double *value);
const char* getNodeMapEnumParam(char *name, char *value);

const char* fullName();
const char* vendorName();
const char* modelName();
const char* serialNumber();
const char* deviceVersion();
int productId();
int vendorId();

int width();
int height();

extern int Go_fetch_callback(int, int, int, int, int, char*);

void setFetchTimeout(uint32_t v);
void setFetchCount(uint32_t v);
const char* fetch(int idx);
#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H
