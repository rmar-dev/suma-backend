# Test API endpoints

Write-Host "Testing Finance App API Endpoints" -ForegroundColor Green
Write-Host "=================================" -ForegroundColor Green

# Test health endpoint
Write-Host "`nTesting /health endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/health" -Method GET
    Write-Host "✓ Health check passed:" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "✗ Health check failed: $_" -ForegroundColor Red
}

# Test API ping endpoint
Write-Host "`nTesting /api/v1/ping endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/ping" -Method GET
    Write-Host "✓ API ping passed:" -ForegroundColor Green
    $response | ConvertTo-Json
} catch {
    Write-Host "✗ API ping failed: $_" -ForegroundColor Red
}

Write-Host "`nAPI test complete!" -ForegroundColor Green