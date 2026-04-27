$ErrorActionPreference = 'Stop'

# Config
$Repo = "NKS01X/Nikium"
$BinName = "nikium.exe"

Write-Host "Downloading Nikium..." -ForegroundColor Cyan

# Query GitHub API
$ApiUrl = "https://api.github.com/repos/$Repo/releases/latest"
try {
    $Release = Invoke-RestMethod -Uri $ApiUrl
} catch {
    Write-Host "Failed to query GitHub API." -ForegroundColor Red
    exit 1
}

# Find correct zip
$Asset = $Release.assets | Where-Object { $_.name -match "windows" -and $_.name -match "(x86_64|amd64)" -and $_.name -match "\.zip$" } | Select-Object -First 1

if (-not $Asset) {
    Write-Host "Cannot find Windows x86_64 release." -ForegroundColor Red
    exit 1
}

$DownloadUrl = $Asset.browser_download_url
Write-Host "Found release: $($Asset.name)" -ForegroundColor Cyan

# Download & extract
$TmpPath = Join-Path $env:TEMP "nikium.zip"
$TmpExtract = Join-Path $env:TEMP "nikium_extracted"

Invoke-WebRequest -Uri $DownloadUrl -OutFile $TmpPath
if (Test-Path $TmpExtract) { Remove-Item -Force -Recurse $TmpExtract }
Expand-Archive -Path $TmpPath -DestinationPath $TmpExtract -Force

# Pick install dir
$IsAdmin = ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if ($IsAdmin) {
    $InstallDir = Join-Path $env:ProgramFiles "Nikium"
} else {
    $InstallDir = Join-Path $env:LOCALAPPDATA "Nikium"
}

if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}

# Find extracted bin
$ExtractedBin = Get-ChildItem -Path $TmpExtract -Recurse -Filter $BinName | Select-Object -First 1
if (-not $ExtractedBin) {
    Write-Host "Cannot find $BinName in archive." -ForegroundColor Red
    Remove-Item -Force $TmpPath
    Remove-Item -Force -Recurse $TmpExtract
    exit 1
}

# Move bin
Copy-Item -Path $ExtractedBin.FullName -Destination (Join-Path $InstallDir $BinName) -Force

# Add to PATH
$PathRegKey = if ($IsAdmin) { "Machine" } else { "User" }
$CurrentPath = [Environment]::GetEnvironmentVariable("PATH", $PathRegKey)

if ($CurrentPath -notmatch [regex]::Escape($InstallDir)) {
    $NewPath = $CurrentPath + ";$InstallDir"
    [Environment]::SetEnvironmentVariable("PATH", $NewPath, $PathRegKey)
    Write-Host "Added $InstallDir to PATH." -ForegroundColor Yellow
}

# Clean
Remove-Item -Force $TmpPath
Remove-Item -Force -Recurse $TmpExtract

Write-Host "Successfully installed!" -ForegroundColor Green
Write-Host "Nikium is located at $InstallDir\$BinName"
Write-Host "Please restart your terminal to use 'nikium'." -ForegroundColor Yellow
