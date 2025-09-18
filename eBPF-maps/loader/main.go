package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/cilium/ebpf/link"
)

//go:generate bpf2go -cc clang -cflags "-O2 -g -Wall" -type trace_event_raw_sys_enter SyscallCounter ../epbf/syscall_counter.c -- -I/usr/include

func main() {
	objs := SyscallCounterObjects{}
	if err := LoadSyscallCounterObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.CountExecve, nil)
	if err != nil {
		log.Fatalf("attaching tracepoint: %v", err)
	}
	defer tp.Close()

	fmt.Println("eBPF program loaded and attached. Press Ctrl+C to exit.")

	key := uint32(0)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	for {
		select {
		case <-ticker.C:
			var count uint64
			err := objs.ExecveCount.Lookup(&key, &count)
			   if err != nil && !os.IsNotExist(err) {
				   log.Printf("map lookup failed: %v", err)
			   }
			fmt.Printf("execve count: %d\n", count)
		case <-sig:
			fmt.Println("Exiting.")
			return
		}
	}
}
