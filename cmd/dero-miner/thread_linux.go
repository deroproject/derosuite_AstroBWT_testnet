//+build !windows

package main

import "runtime"
import "sync/atomic"
import "golang.org/x/sys/unix"

var processor int32

// sets thread affinity to avoid cache collision and thread migration
func threadaffinity() {
	var cpuset unix.CPUSet

	lock_on_cpu := atomic.AddInt32(&processor, 1)
	if lock_on_cpu >= int32(runtime.GOMAXPROCS(0)) { // threads are more than cpu, we do not know what to do
		return
	}
	cpuset.Zero()
	cpuset.Set(int(lock_on_cpu))

	unix.SchedSetaffinity(0, &cpuset)
}
