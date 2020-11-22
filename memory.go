package automaxprocs

// Author: catlittlechen@gmail.com

import (
	iruntime "go.uber.org/automaxprocs/internal/runtime"
)

func GetMemoryLimit() (int, iruntime.CPUQuotaStatus, error) {
	return iruntime.GetMemoryLimit()
}
