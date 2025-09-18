package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
)

func main() {
	objs := verificationfailObjects{}
	if err := loadVerificationfailObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()


	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.VerifierFail, nil)
	if err != nil {
		log.Fatalf("[VERIFIER ERROR] Failed to attach eBPF program: %v", err)
	}
	defer tp.Close()

	log.Println("eBPF program attached (unexpected: should have failed verification). Press Ctrl+C to exit.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
