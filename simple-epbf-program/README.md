
# Simple eBPF Program Example

## What this eBPF Program Does

This project demonstrates a simple eBPF program that traces every time a new process is executed on the system (i.e., every `execve` syscall). The eBPF program attaches to the `sys_enter_execve` tracepoint and prints the process ID and command name of each process that is started. These messages are sent to the kernel's trace pipe, which you can view in real time.

The Go loader loads and attaches the eBPF program at runtime. When running, you can observe the output by reading `/sys/kernel/debug/tracing/trace_pipe` inside the container.

## Project Structure

- `epbf/syscall_counter.c`: The eBPF program (in C) that counts `execve` syscalls.
- `loader/main.go`: The Go loader that loads, attaches, and reads the eBPF map.
- `Dockerfile`: Builds and runs the loader and eBPF program in a container.

## How to Build and Run

### Prerequisites
- Docker (with privileged mode support)
- Linux kernel with eBPF support (most modern kernels)


### 1. Build the eBPF Object File (syscall_counter.o)

You must build the eBPF object file (`epbf/syscall_counter.o`) on a Linux system with kernel headers and clang installed. If you are on macOS or Windows, use the provided Docker build script:

```sh
# From the project root
docker run --rm -v "$PWD":/src -w /src gcc:12 bash -c 'apt-get update && apt-get install -y clang linux-libc-dev && clang -O2 -g -target bpf -c epbf/syscall_counter.c -o epbf/syscall_counter.o'
```

Alternatively, on a Linux host with clang and kernel headers:

```sh
./build-ebpf.sh
```

### 2. Build the Docker Image

```sh
docker build -t simple-ebpf-example .
```

### 3. Run the Container

> **Note:** eBPF programs require privileged access to kernel features and access to debugfs/tracefs. Mount these filesystems from the host:

```sh
docker run --rm -it --privileged \
   -v /sys/kernel/debug:/sys/kernel/debug \
   -v /sys/kernel/tracing:/sys/kernel/tracing \
   simple-ebpf-example
```

## How to Experience the eBPF Program

1. **Start the container as above.**
2. The loader will print the number of `execve` syscalls every 2 seconds.
3. In another terminal, you can run commands inside the container to trigger `execve`, e.g.:

   ```sh
   docker exec -it <container_id> /bin/sh
   ls
   echo hello
   ```
   Each command (like `ls`, `echo`, etc.) will increment the counter.
4. Observe the output in the loader: the count will increase as you run more commands.

## How it Works

- The eBPF program is attached to the `execve` syscall via a kprobe.
- Each time `execve` is called, the eBPF program increments a counter in a BPF map.
- The Go loader reads this map and prints the count.

---

**Note:** This is a minimal example for educational purposes. In production, eBPF programs should be carefully audited and tested.
