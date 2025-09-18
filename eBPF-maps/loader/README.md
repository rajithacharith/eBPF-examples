# eBPF Maps Example: Syscall Counter

This example demonstrates how to use an eBPF map to count the number of times the `execve` syscall is called on the system.

## Structure
- `epbf/syscall_counter.c`: eBPF program that increments a map value on every `execve` syscall.
- `loader/main.go`: Go loader that attaches the eBPF program and periodically prints the syscall count.

## Build & Run

### 1. Build the eBPF object and Go loader

```sh
cd loader
# Install bpf2go if not already installed
GO111MODULE=on go install github.com/cilium/ebpf/cmd/bpf2go@latest
# Generate Go bindings for the eBPF program
bpf2go -cc clang -cflags "-O2 -g -Wall" -type trace_event_raw_sys_enter SyscallCounter ../epbf/syscall_counter.c -- -I/usr/include
# Tidy and build
go mod tidy
go build -o syscall-counter .
```

### 2. Run the loader (requires root)

```sh
sudo ./syscall-counter
```

You should see output like:
```
eBPF program loaded and attached. Press Ctrl+C to exit.
execve count: 3
execve count: 7
...
```

Trigger some execve syscalls (e.g., open a new terminal, run a command) to see the count increase.

## Requirements
- Linux with eBPF and tracepoints enabled
- Go 1.18+
- clang, llvm, and kernel headers for your system
