#!/bin/bash

PYLON_PATH=/opt/pylon5

# HW libs available: /opt/pylon5/lib64/libpylon_TL_{bcon,camemu,gige,usb}.so
libs=(
	-lpylonc 
	-lpylonbase 
	-lpylon_TL_usb    # FIX here for other HW
	-lpylonutility 
	-lGenApi_gcc_v3_1_Basler_pylon 
	-lGCBase_gcc_v3_1_Basler_pylon
)

CPATH=${PYLON_PATH}/include/
CGO_CPPFLAGS=-I${CPATH}
CGO_LDFLAGS="-L${PYLON_PATH}/lib64 -L. -Wl,--start-group ${libs[@]} -Wl,--end-group"

cmd=test                    # defaults to test run
if [ -n "$1" ]; then
    case "$1" in
    -*) ;;                  # pass flags directly
    *)  cmd="$1"; shift;;   # get first non-flag word as command
    esac
fi

set -xe
cd pylon
exec env \
    LD_LIBRARY_PATH="${PYLON_PATH}/lib64" \
    PYLON_PATH="${PYLON_PATH}" \
    CGO_CPPFLAGS="${CGO_CPPFLAGS}" \
    CGO_LDFLAGS="${CGO_LDFLAGS}" \
    go "$cmd" "$@"

# EOF #
