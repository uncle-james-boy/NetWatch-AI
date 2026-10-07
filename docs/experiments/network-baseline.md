# Network Baseline Experiment

**Date:** 2026-10-02  
**Goal:** Establish baseline latency for local gateway vs external internet.

## Device / Wi-Fi Context

* SSID: 5G.rain_high_NT41
* Local IP: 192.168.0.186
* Gateway: 192.168.0.1
* Wi-Fi Band: 5 GHz
* Channel: 48
* Signal: 86% (RSSI -61 dBm)
* DNS Server: fdba:d471:f80c::1

## Gateway Ping (Local Wi-Fi to Router)

* Target IP: 192.168.0.1
* Average Latency: 4 ms
* Packet Loss: 0%

## External Ping (Router to ISP to Internet)

* Target IP: 8.8.8.8
* Average Latency: 30 ms
* Packet Loss: 0%

## Observations

Local gateway path is healthy: average 4ms with 0% packet loss.  
External internet path is healthy: average 30ms with 0% packet loss.  
Wi-Fi signal is strong at 86% on 5GHz channel 48.  