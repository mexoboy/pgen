# Release Instructions

## Automatic Release

This repository is configured to automatically create releases when semantic version tags are pushed.

### How to Create a Release

1. **Ensure all your changes are committed and pushed to the `main` branch**

2. **Create and push a semantic version tag:**
   ```bash
   # For a regular release
   git tag v1.0.0
   git push origin v1.0.0

   # For a pre-release (alpha, beta, rc)
   git tag v1.0.0-alpha.1
   git push origin v1.0.0-alpha.1
   ```

3. **GitHub Actions will automatically:**
   - Run tests
   - Build binaries for all supported platforms
   - Generate checksums
   - Create a changelog based on commits
   - Create a GitHub Release with artifacts

### Tag Formats

Supported semantic version formats:
- `v1.0.0` – stable release
- `v1.0.0-alpha.1` – alpha release
- `v1.0.0-beta.1` – beta release
- `v1.0.0-rc.1` – release candidate

### Supported Platforms

Binaries will be built for:
- Linux AMD64
- Linux ARM64
- macOS AMD64 (Intel)
- macOS ARM64 (Apple Silicon)
- Windows AMD64

### Continuous Integration (CI)

The CI pipeline runs on each push and pull request to `main`, `master`, or `develop` branches, and includes:
- Test execution with race detection
- Code coverage report generation
- Code formatting check (`go fmt`)
- Static analysis (`go vet`)
- Project build

### Troubleshooting

1. **Release not created:**
   - Ensure the tag matches the format `v*.*.*`
   - Make sure all tests pass
   - Check GitHub Actions logs

2. **Build failed:**
   - Ensure the code compiles locally
   - Verify all dependencies are declared in `go.mod`

3. **Test failures:**
   - Run tests locally: `go test -v ./...`
   - Fix failing tests before tagging

### Local Checks Before Release

It's recommended to run the following commands before tagging:

```bash
# Formatting check
go fmt ./...

# Static analysis
go vet ./...

# Run tests
go test -v ./...

# Build check
go build .
```
