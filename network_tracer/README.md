# Network Tracer eBPF Example

This example demonstrates a simple eBPF program that traces TCP connect events (outgoing connections) on the system. It prints the PID and destination IP/port for each outgoing TCP connection.

## What This Example Does

- Attaches an eBPF program to the `tcp_v4_connect` kprobe.
- Prints the process ID, destination IP, and port for each outgoing TCP connection.

## How to Build and Run

1. Build the Docker image:

   ```sh
   docker build -t network-tracer .
   ```

2. Run the container (with required mounts):

   ```sh
   docker run --rm -it --privileged \
     -v /sys/kernel/debug:/sys/kernel/debug \
     -v /sys/kernel/tracing:/sys/kernel/tracing \
     network-tracer
   ```

You can then generate network traffic (e.g., `curl google.com`) and observe the output in `/sys/kernel/debug/tracing/trace_pipe` inside the container.

---

**Note:** This example is for learning and demonstration only. The program prints to the kernel trace pipe.
