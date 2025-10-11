# Update import paths in all Go files
$oldPath = "github.com/suma/finance-app-api"
$newPath = "github.com/rmar-dev/suma-backend"

Write-Host "Updating import paths from $oldPath to $newPath..."

Get-ChildItem -Recurse -Filter "*.go" | ForEach-Object {
    $content = Get-Content $_.FullName -Raw
    if ($content -match $oldPath) {
        Write-Host "Updating $($_.FullName)"
        $content = $content -replace $oldPath, $newPath
        Set-Content $_.FullName $content
    }
}