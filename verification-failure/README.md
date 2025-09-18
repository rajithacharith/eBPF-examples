# Verification Failure eBPF Example

This example demonstrates a simple eBPF program that will fail kernel verification. It is intended for educational purposes to show what happens when the eBPF verifier rejects a program.

## What This Example Does

The eBPF program attempts to write to an invalid memory address, which is not allowed and will be rejected by the kernel verifier.

## How to Build and Run

1. Build the Docker image:

   ```sh
   docker build -t ebpf-verification-failure .
   ```

2. Run the container (with required mounts):

   ```sh
   docker run --rm -it --privileged \
     -v /sys/kernel/debug:/sys/kernel/debug \
     -v /sys/kernel/tracing:/sys/kernel/tracing \
     ebpf-verification-failure
   ```

You should see a verifier error message in the output, indicating that the eBPF program was rejected.

---

**Note:** This example is for learning and demonstration only. The program will not be loaded into the kernel.
