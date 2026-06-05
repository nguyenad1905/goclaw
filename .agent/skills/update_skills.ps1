# Cập nhật các AI Skills (understand-anything và superpowers) từ GitHub và sao chép vào dự án
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8


$ProjectSkillsDir = "$PSScriptRoot"
$UnderstandRepo = "C:\Users\anhnq.568E-local\.understand-anything\repo\understand-anything-plugin"
$SuperpowersRepo = "C:\Users\anhnq.568E-local\.superpowers\repo"

Write-Host "==============================================" -ForegroundColor Cyan
Write-Host "Cập nhật Understand-Anything" -ForegroundColor Cyan
Write-Host "==============================================" -ForegroundColor Cyan
if (Test-Path $UnderstandRepo) {
    Set-Location $UnderstandRepo
    Write-Host "Kéo code mới nhất từ GitHub..."
    git fetch --all
    git pull
    Write-Host "Cài đặt thư viện..."
    pnpm install
    Write-Host "Build @understand-anything/core..."
    pnpm --filter @understand-anything/core build
    
    Write-Host "Cập nhật vào dự án..."
    Get-ChildItem -Directory "$UnderstandRepo\skills" | ForEach-Object {
        $dest = Join-Path $ProjectSkillsDir $_.Name
        if (Test-Path $dest) { Remove-Item -Recurse -Force $dest }
        Copy-Item -Path $_.FullName -Destination $dest -Recurse -Force
        Write-Host "  -> Cập nhật: $($_.Name)" -ForegroundColor Green
    }
} else {
    Write-Host "Không tìm thấy thư mục: $UnderstandRepo" -ForegroundColor Red
}

Write-Host ""
Write-Host "==============================================" -ForegroundColor Cyan
Write-Host "Cập nhật Superpowers" -ForegroundColor Cyan
Write-Host "==============================================" -ForegroundColor Cyan
if (Test-Path $SuperpowersRepo) {
    Set-Location $SuperpowersRepo
    Write-Host "Kéo code mới nhất từ GitHub..."
    git fetch --all
    git pull
    
    Write-Host "Cập nhật vào dự án..."
    Get-ChildItem -Directory "$SuperpowersRepo\skills" | ForEach-Object {
        $dest = Join-Path $ProjectSkillsDir $_.Name
        if (Test-Path $dest) { Remove-Item -Recurse -Force $dest }
        Copy-Item -Path $_.FullName -Destination $dest -Recurse -Force
        Write-Host "  -> Cập nhật: $($_.Name)" -ForegroundColor Green
    }
} else {
    Write-Host "Không tìm thấy thư mục: $SuperpowersRepo" -ForegroundColor Red
}

Set-Location $PSScriptRoot
Write-Host ""
Write-Host "Đã cập nhật xong tất cả các skills!" -ForegroundColor Green
