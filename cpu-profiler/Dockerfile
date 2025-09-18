FROM golang:1.23-bullseye

# Install dependencies for eBPF development (multi-arch)
RUN set -e; \
  arch=$(dpkg --print-architecture); \
  if [ "$arch" = "arm64" ]; then \
    headers=linux-headers-arm64; \
    inc="-I/usr/include -I/usr/include/aarch64-linux-gnu"; \
  else \
    headers=linux-headers-amd64; \
    inc="-I/usr/include -I/usr/include/x86_64-linux-gnu"; \
  fi; \
  apt-get update && apt-get install -y \
    clang \
    llvm \
    libbpf-dev \
    libelf-dev \
    gcc \
    make \
    pkg-config \
    bpftool \
    git \
    linux-libc-dev \
    $headers \
    && rm -rf /var/lib/apt/lists/* && \
  echo "$inc" > /bpf_includes

WORKDIR /app

RUN go install github.com/cilium/ebpf/cmd/bpf2go@latest

# Copy eBPF C code
COPY epbf/cpu_profiler.c epbf/

# Copy Go loader and go.mod
COPY loader/main.go .
COPY loader/go.mod .

# Generate Go bindings for eBPF code (auto-detect arch)
RUN inc=$(cat /bpf_includes) && bpf2go -go-package main CpuProfiler epbf/cpu_profiler.c -- $inc

# Build Go loader
RUN go mod tidy && go build -o cpu-profiler .

CMD ["./cpu-profiler"]
