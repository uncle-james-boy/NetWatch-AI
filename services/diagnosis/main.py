import json

# 1. THE EVIDENCE (Mock data from our Go Telemetry)
# Imagine this is the JSON payload our Go program just sent to the cloud
telemetry_payload = """
{
  "gateway_latency_ms": 450,
  "gateway_packet_loss": 12.5,
  "external_latency_ms": 30,
  "external_packet_loss": 0.0
}
"""

# 2. PARSE THE JSON
# Convert the text string into a Python dictionary (data structure)
data = json.loads(telemetry_payload)

# 3. THE DETERMINISTIC RULES ENGINE
def diagnose_network(metrics):
    gateway_latency = metrics.get("gateway_latency_ms", 0)
    gateway_loss = metrics.get("gateway_packet_loss", 0)
    external_latency = metrics.get("external_latency_ms", 0)
    external_loss = metrics.get("external_packet_loss", 0)

    # Rule 1: Is the local Wi-Fi/Router path terrible?
    # If gateway latency is > 100ms OR packet loss is > 5%
    if gateway_latency > 100 or gateway_loss > 5.0:
        return {
            "category": "LOCAL_NETWORK_DEGRADATION",
            "severity": "HIGH",
            "confidence": 0.95,
            "reason": "Gateway latency or packet loss exceeds acceptable thresholds."
        }

    # Rule 2: Is the local path fine, but the internet path terrible?
    if external_latency > 150 or external_loss > 5.0:
        return {
            "category": "ISP_OR_EXTERNAL_DEGRADATION",
            "severity": "MEDIUM",
            "confidence": 0.90,
            "reason": "Local network is healthy, but external routing is failing."
        }

    # Rule 3: Everything is fine
    return {
        "category": "HEALTHY",
        "severity": "NONE",
        "confidence": 1.0,
        "reason": "All network metrics are within normal baselines."
    }

# 4. RUN THE DIAGNOSIS
diagnosis = diagnose_network(data)

# 5. OUTPUT THE RESULT
print("--- 🩺 NETWATCH DIAGNOSIS ENGINE ---")
print(json.dumps(diagnosis, indent=4))