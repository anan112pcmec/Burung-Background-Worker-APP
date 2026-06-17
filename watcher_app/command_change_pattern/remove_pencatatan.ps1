$root = Join-Path $PSScriptRoot 'watcher_app\database\cassandra\models'
Get-ChildItem -Path $root -Filter *.go | ForEach-Object {
    $text = Get-Content -Raw -Path $_.FullName
    $new = $text -replace '(?m)^[ \t]*Pencatatan[ \t]*\r?\n', ''
    $new = $new -replace '(?m)^[ \t]*"tahun_update"\s*:\s*.*\r?\n', ''
    $new = $new -replace '(?m)^[ \t]*"bulan_update"\s*:\s*.*\r?\n', ''
    $new = $new -replace '(?m)^[ \t]*"event_time"\s*:\s*.*\r?\n', ''
    if ($new -ne $text) {
        Set-Content -Path $_.FullName -Value $new -Encoding utf8
        Write-Host "Updated: $($_.Name)"
    }
}
