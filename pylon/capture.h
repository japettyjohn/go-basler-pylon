#ifndef CAPTURE_H
#define CAPTURE_H

#ifdef __cplusplus 
extern "C" {
#else
#include <stdbool.h>
#endif

// prototypes
void stopCapture();
void attachDevice();
void configureCamera();
const char* retrieveAndSave(int batch, int timeout, char* outputPath);
const char* startCapture();
void openCamera();
void closeCamera();
bool isCameraGrabbing();
bool isAttached();
bool isOpen();
void setHardwareTriggerConfiguration();

#ifdef __cplusplus
}
#endif

#endif // CAPTURE_H