# INC-0001: LocalStack Boot Failure (Exit Code 55)

**Date:** 2026-10-10
**Severity:** High (Blocked local development)
**Status:** Resolved

## Symptoms
When running `docker-compose up` to boot the local AWS emulation (LocalStack), the container immediately crashed with `Exit Code 55`. 
The logs displayed: `License activation failed! ... LocalStack pro features can only be used with a valid license.`

## Root Cause Analysis
LocalStack offers a free Community Edition and a paid Pro Edition. The local Docker environment was attempting to boot the Pro Edition due to either a cached Pro image or lingering environment variables (`LOCALSTACK_AUTH_TOKEN`) on the host machine. Because no paid license was configured, the supervisor process terminated the container.

## Resolution
1. Destroyed the faulty container using `docker-compose down`.
2. Pinned the free community image version explicitly via `docker pull localstack/localstack:3.8.0`.
3. Updated `docker-compose.yml` to remove the deprecated `version` attribute.
4. Added environment overrides (`LOCALSTACK_AUTH_TOKEN=`) to the YAML file to force the container into Community Edition mode, ignoring any host-level Pro configurations.

## Lessons Learned
- Always pin Docker image versions (e.g., `:3.8.0` instead of `:latest`) to prevent unexpected updates or feature-gating issues.
- Environment variables on the host OS can silently leak into Docker containers. Explicitly overriding them in the `docker-compose.yml` ensures predictable Infrastructure-as-Code behavior.
