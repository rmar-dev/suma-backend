# Contributing to SUMA Finance Backend

## Development Workflow

### 1. Branching Strategy

We follow a modified GitFlow workflow:

```
main ─────────────●───────────●─────── (Production releases)
                 │           │
develop ─────●───●───●───●───●─────── (Integration branch)
            │       │   │
feature/1   ●───●───●   │           (Feature branches)
                        │
feature/2       ●───●───●           (Feature branches)
```

#### Branch Naming
- `main`: Production releases
- `develop`: Integration branch
- `feat/*`: New features
- `fix/*`: Bug fixes
- `docs/*`: Documentation updates
- `refactor/*`: Code refactoring
- `test/*`: Test additions/modifications
- `chore/*`: Maintenance tasks

### 2. Development Process

1. **Start New Feature**
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feat/feature-name
   ```

2. **Local Development**
   - Write code
   - Add tests
   - Update documentation
   - Run local checks:
     ```bash
     go test ./...
     golangci-lint run
     gosec ./...
     ```

3. **Commit Changes**
   - Follow conventional commits:
     ```
     feat: add user authentication
     fix: correct transaction date parsing
     docs: update API documentation
     test: add user service tests
     refactor: simplify transaction processing
     chore: update dependencies
     ```

4. **Push and Create PR**
   ```bash
   git push -u origin feat/feature-name
   # Create PR through GitHub UI
   ```

### 3. Pull Request Process

1. **PR Requirements**
   - [ ] Description of changes
   - [ ] Link to related issue
   - [ ] Tests added/updated
   - [ ] Documentation updated
   - [ ] Changelog updated (if applicable)

2. **Required Checks**
   - All CI checks must pass:
     - Unit tests
     - Integration tests
     - Linting
     - Security scan
     - Code coverage (minimum 80%)

3. **Review Process**
   - At least one approval required
   - All comments must be resolved
   - No unresolved conversations
   - PR must be up to date with develop

### 4. Code Standards

#### Go Guidelines
- Follow standard Go formatting (`go fmt`)
- Use meaningful variable/function names
- Write self-documenting code
- Add comments for complex logic
- Keep functions focused and small
- Use interfaces appropriately

#### Testing Requirements
- Unit tests for all new code
- Integration tests for API endpoints
- Test coverage must not decrease
- Mock external dependencies
- Use table-driven tests where appropriate

#### Documentation
- Update API documentation
- Add godoc comments
- Update README if needed
- Document configuration changes
- Add examples for new features

### 5. Security Guidelines

- No secrets in code
- Use environment variables
- Validate all inputs
- Follow secure coding practices
- Regular dependency updates
- Security scanning in CI

## Release Process

1. **Create Release Branch**
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b release/v1.x.x
   ```

2. **Release Preparation**
   - Update version numbers
   - Update CHANGELOG.md
   - Run final tests
   - Update documentation

3. **Release Review**
   - Create PR to main
   - Complete QA process
   - Get final approvals

4. **Release Deployment**
   - Merge to main
   - Tag release
   - Deploy to production
   - Monitor deployment

## Issue Guidelines

1. **Bug Reports**
   - Clear description
   - Steps to reproduce
   - Expected vs actual behavior
   - Environment details
   - Screenshots if applicable

2. **Feature Requests**
   - Clear description
   - Use case explanation
   - Proposed solution
   - Alternative solutions
   - Impact assessment

## Getting Help

- Check existing documentation
- Search closed issues
- Ask in team chat
- Create detailed issue
- Tag appropriate reviewers

## License & Attribution
- All contributions are subject to our license
- Maintain list of contributors
- Credit external resources
- Document third-party usage