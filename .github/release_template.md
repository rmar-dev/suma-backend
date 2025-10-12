# Release v{version}

## Release Information
- **Version:** {version}
- **Release Date:** {date}
- **Build Type:** {type}
- **Go Version:** {go_version}

## Security Information
- **SBOM:** Included (SPDX & CycloneDX formats)
- **Signatures:** Cosign
- **Checksums:** SHA-256, SHA-512

## What's New

### Features ✨
- Feature 1
- Feature 2

### Improvements 🚀
- Improvement 1
- Improvement 2

### Bug Fixes 🐛
- Fix 1
- Fix 2

### Security Updates 🔒
- Update 1
- Update 2

### Dependencies 📦
- Dependency 1: v1.2.3 -> v1.2.4
- Dependency 2: v2.0.0 -> v2.1.0

## Breaking Changes ⚠️
List any breaking changes here, if applicable.

## Artifacts
```
api-{version}-{os}-{arch}/
├── api              # Main binary
├── api.sha256      # SHA-256 checksum
├── api.sha512      # SHA-512 checksum
├── api.sig         # Cosign signature
└── sbom/
    ├── spdx.json   # SPDX format SBOM
    └── cyclonedx.json # CycloneDX format SBOM
```

## Verification

### Checksums
```bash
# Verify SHA-256
sha256sum -c api.sha256

# Verify SHA-512
sha512sum -c api.sha512
```

### Signature
```bash
# Verify with Cosign
cosign verify-blob --key cosign.pub --signature api.sig api
```

## Installation
```bash
# Download and verify
curl -LO https://github.com/rmar-dev/suma-backend/releases/download/v{version}/api-{version}-linux-amd64.tar.gz
tar xzf api-{version}-linux-amd64.tar.gz
cd api-{version}-linux-amd64
sha256sum -c api.sha256

# Run
./api
```

## Upgrade Instructions
Describe any special steps needed for upgrading from previous versions.

## Known Issues
List any known issues or limitations.

## Documentation
- [Full Changelog](CHANGELOG.md)
- [API Documentation](API.md)
- [Migration Guide](MIGRATION.md) (if applicable)

## Support
- Report issues: [GitHub Issues](https://github.com/rmar-dev/suma-backend/issues)
- Security concerns: security@example.com

---

## Validation Results
- [x] Extended validation passed
- [x] Release validation passed
- [x] Security scans completed
- [x] Documentation updated
- [x] Artifacts verified