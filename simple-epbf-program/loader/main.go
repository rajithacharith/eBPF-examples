// main.go
// Go loader for a simple eBPF program that counts syscalls (e.g., execve)
// Uses cilium/ebpf for loading and attaching the eBPF program
//
// Build: go build -o loader .
// Run: sudo ./loader

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
)


func main() {
	// Load compiled eBPF objects (bpf2go bindings)
	objs := traceexecObjects{}
	if err := loadTraceexecObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	// Attach program to tracepoint
	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.TraceExec, nil)
	if err != nil {
		log.Fatalf("attaching tracepoint: %v", err)
	}
	defer tp.Close()

	log.Println("✅ eBPF program attached. Run some commands in another shell and check trace_pipe:")

	// Wait for SIGINT or SIGTERM
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
