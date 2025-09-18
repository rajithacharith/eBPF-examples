#!/bin/bash
# build-ebpf.sh
# Build the eBPF object file for syscall_counter.c
set -e

SRC=epbf/syscall_counter.c
OUT=epbf/syscall_counter.o

if ! command -v clang >/dev/null 2>&1; then
  echo "clang is required. Please install clang." >&2
  exit 1
fi

# Use kernel headers from the host
clang -O2 -g -target bpf -c "$SRC" -o "$OUT"
echo "Built $OUT"
