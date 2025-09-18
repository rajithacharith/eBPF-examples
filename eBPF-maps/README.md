# eBPF Maps Example: Syscall Counter

This example demonstrates how to use an eBPF map to count the number of times the `execve` syscall is called on the system.

## What does this show?
- Tracks and prints the number of `execve` (process execution) syscalls in real time.
- Shows how to use a BPF map to store and retrieve data between kernel and user space.

## How it works
- The eBPF program attaches to the `syscalls:sys_enter_execve` tracepoint.
- Every time a process calls `execve`, the eBPF program increments a counter in a BPF map.
- The Go loader periodically reads this map and prints the current count.

## Structure
- `epbf/syscall_counter.c`: eBPF program that increments a map value on every `execve` syscall.
- `loader/main.go`: Go loader that attaches the eBPF program and prints the syscall count every 2 seconds.

## Build & Run

### 1. Build the Docker image
```sh
cd eBPF-maps
# Build the image
docker build -t ebpf-maps-example .
```

### 2. Run the example in Docker (requires privileges)
```sh
docker run --rm -it --privileged --name ebpf-maps-example --network host \
  --cap-add=SYS_ADMIN --cap-add=SYS_RESOURCE --cap-add=SYS_PTRACE \
  -v /sys/kernel/debug:/sys/kernel/debug:rw \
  -v /sys/fs/bpf:/sys/fs/bpf:rw \
  ebpf-maps-example
```

You will see the `execve` count printed every 2 seconds. Run commands or open new terminals to see the count increase.

## Requirements
- Linux with eBPF and tracepoints enabled
- Docker with privileged mode
- Go 1.18+, clang, llvm, and kernel headers (for building the image)
