# NetLens — Network Packet Inspection Engine

NetLens is a Linux network packet inspection engine written in Go.

It captures raw network frames directly from a network interface, decodes protocols manually, reconstructs flows, inspects application-layer metadata, and is designed to evolve into a concurrent DPI and policy engine.

The project is intentionally built from the bottom up:

```text
raw frame
   ↓
Ethernet
   ↓
IP
   ↓
TCP / UDP / ICMP
   ↓
5-tuple
   ↓
flow
   ↓
DNS / HTTP / TLS
   ↓
classification
   ↓
policy
```
