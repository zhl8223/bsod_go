package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"
)

type NtDLL struct {
	mod             *syscall.LazyDLL
	AdjustPrivilege *syscall.LazyProc
	RaiseHardError  *syscall.LazyProc
}

func (dll *NtDLL) init() {
	dll.mod = syscall.NewLazyDLL("ntdll.dll")
	dll.AdjustPrivilege = dll.mod.NewProc("RtlAdjustPrivilege")
	dll.RaiseHardError = dll.mod.NewProc("NtRaiseHardError")
}

const (
	FLASE                   = 0
	TRUE                    = 1
	STATUS_ACCESS_VIOLATION = 0xC0000005
)

// winnt.h

func (dll *NtDLL) bsod() {
	// NTSTATUS NtRet = NtCall(19, TRUE, FALSE, &bEnabled);
	// NtCall2(STATUS_FLOAT_MULTIPLE_FAULTS, 0, 0, 0, 6, &uResp);

	var bEnabled int8
	r1, _, _ := dll.AdjustPrivilege.Call(19, 1, 0, uintptr(unsafe.Pointer(&bEnabled)))
	fmt.Println("r1=", r1, ", bEnabled=", bEnabled)

	var uResp int32
	r1, _, _ = dll.RaiseHardError.Call(STATUS_ACCESS_VIOLATION, 0, 0, 0, 6, uintptr(unsafe.Pointer(&uResp)))
	fmt.Println("r1=", r1, ", uResp=", uResp)
}
