package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
)

func main() {
	objs := networktracerObjects{}
	if err := loadNetworktracerObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	lnk, err := link.Kprobe("tcp_v4_connect", objs.TraceConnect, nil)
	if err != nil {
		log.Fatalf("attaching kprobe: %v", err)
	}
	defer lnk.Close()

	log.Println("✅ Network tracer eBPF program attached. Run 'curl' or other network commands and check trace_pipe.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
