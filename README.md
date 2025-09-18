# eBPF Lab Examples

This repository contains a collection of hands-on labs and example programs for learning, testing, and experimenting with eBPF (extended Berkeley Packet Filter) on Linux. Each example demonstrates a different aspect of eBPF, including tracing, monitoring, and custom kernel instrumentation, using both C and Go (with cilium/ebpf).

## What is eBPF?

eBPF is a powerful technology in the Linux kernel that allows you to run sandboxed programs in response to events such as system calls, tracepoints, network packets, and more. It is widely used for observability, security, and networking.

## Repository Structure

- `simple-epbf-program/` — Minimal eBPF loader and program example (trace execve syscalls)
- More labs and examples coming soon!

## How to Use

Each lab/example contains its own README with build and run instructions. You can use these examples to:
- Learn eBPF basics
- Experiment with tracing and monitoring
- Build your own eBPF tools

## Getting Started

1. Clone this repository
2. Follow the instructions in each example directory
3. Run the examples on a Linux system with a recent kernel

---

**Note:** Running eBPF programs usually requires root privileges and a kernel with eBPF support (Linux 4.9+ recommended).
