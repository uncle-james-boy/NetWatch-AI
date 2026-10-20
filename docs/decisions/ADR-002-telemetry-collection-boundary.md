# ADR-002: Telemetry Collection Boundary (Measurement Vantage Point)

## Status
Accepted

## Context
Phase 6 moved the Go telemetry probe into a Lambda function. Testing revealed that a
Lambda-based probe measures the network path from the cloud container's location, not
from the end user's local Wi-Fi or router. Server-side probes therefore cannot observe
local degradation (weak RSSI, Wi-Fi congestion, gateway faults).

## Decision
NetWatch telemetry has two collection boundaries:
1. Client-side: the mobile app (and later a local network agent) measures gateway
   latency, RSSI, and local packet loss from inside the user's network.
2. Server-side: Lambda probes measure the cloud-side path as a stable reference
   vantage point.
The diagnosis engine correlates both vantage points to isolate fault domains.

## Consequences
- Lambda probes alone must never produce LOCAL_WIFI_* diagnoses.
- The API must accept client-submitted telemetry as primary evidence for local categories.
- Phase 13 (local network agent) is architecturally mandatory, not optional.
