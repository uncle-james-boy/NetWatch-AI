# NetWatch V1 Serverless Architecture Flow

## The Request Lifecycle

1. **Client (Mobile App)**
   - User triggers a network diagnostic.
   - Sends an HTTPS POST request to the cloud.

2. **API Gateway (The Front Door)**
   - Receives the public internet traffic.
   - Validates the request format.
   - Routes the payload to the Telemetry Lambda.

3. **Telemetry Lambda (Go)**
   - Executes concurrent Goroutines to probe Gateway, ISP, and External targets.
   - Aggregates metrics (Latency, Packet Loss) into a structured JSON payload.
   - Passes payload to the Diagnosis Lambda.

4. **Diagnosis Lambda (Python)**
   - Receives the JSON telemetry.
   - Applies deterministic threshold rules.
   - Classifies the incident (e.g., LOCAL_WIFI_DEGRADATION).

5. **DynamoDB (Storage)**
   - Saves the incident report, timestamp, and telemetry for historical tracking.

## Why Serverless?
- **Cost:** We only pay when a user actively runs a diagnostic (Free Tier friendly).
- **Scaling:** If 1,000 users run tests at once, AWS automatically spins up 1,000 Lambda instances.
- **Maintenance:** AWS manages the underlying OS and hardware patches.
