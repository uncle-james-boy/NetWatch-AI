# ADR-001: AWS Account Closure and Local-First Pivot

## Status
Accepted (2026-10-09)

## Context
On 2026-09-25, AWS notified that the project account was closed because
the 6-month free plan period ended. Reopening requires upgrading to a
paid plan by 2026-12-24, after which the account is permanently deleted.
The project currently has no payment method available for account
verification.

## Decision
Continue NetWatch development against a local AWS emulator (LocalStack)
for all serverless services (Lambda, API Gateway, DynamoDB, S3, SNS,
EventBridge). All code, AWS CLI usage, and Terraform will remain fully
AWS-compatible so deployment to real AWS can occur later without rewrite.

## Alternatives Considered
1. Stop the project until a payment method is available. (Rejected: loses momentum and learning time.)
2. Create a new AWS account to bypass closure. (Rejected: violates AWS Terms of Service.)
3. Switch to a different cloud provider. (Rejected: AWS remains the target platform for career goals.)

## Consequences
- Zero cost and no card required during development.
- Same AWS APIs, CLI commands, and Terraform provider (via custom endpoint).
- Minor behavioral differences from real AWS must be documented when found.
- Real AWS deployment deferred until the reopen window (before 2026-12-24)
  or until student/educational credits are obtained.
