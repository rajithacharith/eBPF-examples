package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cilium/ebpf/link"
)

//go:generate bpf2go -cc clang -cflags "-O2 -g -Wall" -type trace_event_raw_sched_switch CpuProfiler ../epbf/cpu_profiler.c -- -I/usr/include

type KeyT struct {
	Pid  uint32
	Comm [16]byte
}

func main() {
	objs := CpuProfilerObjects{}
	if err := LoadCpuProfilerObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	tp, err := link.Tracepoint("sched", "sched_switch", objs.OnCpuSwitch, nil)
	if err != nil {
		log.Fatalf("attaching tracepoint: %v", err)
	}
	defer tp.Close()

	fmt.Println("✅ CPU profiler eBPF program attached. Press Ctrl+C to exit.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	prev := make(map[KeyT]uint64)
	ticker := make(chan bool, 1)
	go func() {
		for {
			ticker <- true
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case <-sig:
			fmt.Println("\nExiting.")
			return
		case <-ticker:
			fmt.Println("\nCPU time (ns) per process (delta in last 2s):")
			var key KeyT
			var value uint64
			it := objs.CpuTime.Iterate()
			curr := make(map[KeyT]uint64)
			for it.Next(&key, &value) {
				curr[key] = value
				delta := value - prev[key]
				fmt.Printf("PID: %d COMM: %s CPU Time (ns): %d\n", key.Pid, string(key.Comm[:]), delta)
			}
			prev = curr
		}
	}
}
