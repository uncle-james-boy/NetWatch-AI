# INC-0002: Terraform S3 DNS Resolution Failure in LocalStack

**Date:** 2026-10-10
**Severity:** High (Blocked S3 deployment)
**Status:** Resolved

## Symptoms
Running `terraform apply` to create an S3 bucket resulted in a DNS lookup error: `dial tcp: lookup netwatch-raw-telemetry-dev.localhost: no such host`.

## Root Cause Analysis
The AWS Terraform provider defaults to "Virtual-Hosted-Style" S3 URLs (e.g., `bucket.endpoint.com`). LocalStack requires "Path-Style" URLs (e.g., `endpoint.com/bucket`). Because the provider attempted to resolve the bucket name as a subdomain of localhost, the OS DNS resolver failed.
Additionally, the LocalStack container was in an "Exited" state due to a host machine reboot, causing secondary connection timeouts.

## Resolution
1. Restarted the LocalStack container via `docker-compose up`.
2. Updated `providers.tf` to include `s3_use_path_style = true` in the AWS provider block, forcing the SDK to use path-based routing.
