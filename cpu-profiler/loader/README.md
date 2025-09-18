# eBPF CPU Profiler Example

This example demonstrates how to use eBPF to profile CPU usage per process in real time using the Linux `sched_switch` tracepoint.

## What does this show?

- **Live CPU usage per process:**
	- Every 2 seconds, the loader prints the amount of CPU time (in nanoseconds) consumed by each process during the last interval.
	- Output includes the process PID, command name, and CPU time delta.
- **How it works:**
	- The eBPF program attaches to the `sched:sched_switch` tracepoint, which fires on every context switch.
	- On each switch, it records the timestamp for the process being switched out, accumulating CPU time in a BPF map keyed by PID and command.
	- The Go loader periodically reads this map and prints the per-process CPU usage delta since the last interval.

Example output:
```
✅ CPU profiler eBPF program attached. Press Ctrl+C to exit.

CPU time (ns) per process (delta in last 2s):
PID: 1 COMM: systemd CPU Time (ns): 10000000
PID: 1234 COMM: bash CPU Time (ns): 5000000
...
```

## Structure
- `epbf/cpu_profiler.c`: eBPF program that tracks CPU time per process.
- `loader/main.go`: Go loader that attaches the eBPF program and prints real-time CPU usage per process.

## Build & Run

### 1. Build the eBPF object and Go loader

```sh
cd loader
# Install bpf2go if not already installed
GO111MODULE=on go install github.com/cilium/ebpf/cmd/bpf2go@latest
# Generate Go bindings for the eBPF program
bpf2go -cc clang -cflags "-O2 -g -Wall" -type trace_event_raw_sched_switch CpuProfiler ../epbf/cpu_profiler.c -- -I/usr/include
# Tidy and build
go mod tidy
go build -o cpu-profiler .
```

### 2. Run the loader (requires root)

```sh
sudo ./cpu-profiler
```

You will see live CPU usage deltas every 2 seconds. Press Ctrl+C to exit.

## Requirements
- Linux with eBPF and tracepoints enabled
- Go 1.18+
- clang, llvm, and kernel headers for your system
