# Continuous Integration (CI) Pipeline

## Overview

Our CI pipeline is designed to ensure code quality, security, and reliability across different stages of development. The pipeline adapts its validation level based on the target branch and context.

## Validation Levels

### Basic Validation (Development)
Applied to feature branches and PRs to develop.

**Checks:**
- Version check
- Executable bit verification
- Format validation
- Basic compilation test

**When Used:**
- Regular development work
- Feature branches
- PRs to develop

### Extended Validation (Staging)
Applied to PRs targeting main branch.

**Checks:**
- All basic validation checks
- `go vet` analysis
- Race condition detection
- Integration tests
- Coverage reporting
- SBOM generation
- Checksum verification
- Static analysis with golangci-lint

**When Used:**
- Pull requests to main
- Pre-release validation
- Integration testing

### Release Validation
Applied to main branch builds.

**Checks:**
- All extended validation checks
- Comprehensive static analysis
- Full test suite with race detection
- Binary size constraints (1MB - 100MB)
- Dual checksum verification (SHA-256/SHA-512)
- Binary signing (if keys available)
- Extended artifact retention (90 days)

**When Used:**
- Main branch builds
- Release preparation
- Production deployments

## Build Types and Tags

### Development Builds
```bash
BUILD_TYPE=development
BUILD_TAGS="development"
# Standard build flags
```

### Staging Builds
```bash
BUILD_TYPE=staging
BUILD_TAGS="staging"
BUILD_FLAGS="-trimpath"
LDFLAGS="-w -s"  # Strip debug info
```

### Release Builds
```bash
BUILD_TYPE=release
BUILD_TAGS="release,production"
BUILD_FLAGS="-trimpath"
LDFLAGS="-w -s"  # Strip debug info
```

## Security Features

### SBOM Generation
- SPDX format
- CycloneDX format
- Generated for extended and release builds

### Checksums
- SHA-256 for all builds
- Additional SHA-512 for release builds
- Verified during CI process

### Binary Signing
- Uses Cosign for release builds
- Requires `COSIGN_PRIVATE_KEY` secret
- Generates `.sig` file for verification

## Artifacts

### Storage Duration
- Basic builds: No artifacts stored
- Extended builds: 7 days retention
- Release builds: 90 days retention

### Artifact Contents
```
api                    # Main binary
api.sha256            # SHA-256 checksum
api.sha512            # SHA-512 checksum (release only)
api.sig               # Cosign signature (release only)
sbom.spdx.json        # SPDX format SBOM
sbom.cyclonedx.json   # CycloneDX format SBOM
```

## Environment Variables

### Required Secrets
- `COSIGN_PRIVATE_KEY`: For binary signing (release builds)
- `CODECOV_TOKEN`: For coverage reporting

### Build Environment
- `CGO_ENABLED=0`: Ensures static builds
- `GOEXPERIMENT=loopvar`: Enables latest safety features
- `TEST_MODE=true`: Prevents server startup during tests

## Coverage Requirements

- Minimum: 5% (fail if below)
- Warning: < 30%
- Target: > 30%

Coverage reports are uploaded to Codecov for tracking and visualization.

## Version Information

Binaries include embedded version info:
```
Finance App API {version} ({commit}) [{buildType}] built on {date}
```

Example:
```
Finance App API 1.24 (f28f521) [release] built on 2025-10-12
```