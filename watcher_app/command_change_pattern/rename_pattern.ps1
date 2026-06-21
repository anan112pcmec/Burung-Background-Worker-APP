Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
    'watcher_app\message_broker\consumer\method.go',
    'watcher_app\message_broker\dispatcher\kurir\create.go',
    'watcher_app\message_broker\dispatcher\kurir\delete.go',
    'watcher_app\message_broker\dispatcher\kurir\update.go',
    'watcher_app\message_broker\dispatcher\pengguna\create.go', # <-- Sudah ditambahkan koma yang kurang
    'watcher_app\message_broker\dispatcher\pengguna\delete.go',
    'watcher_app\message_broker\dispatcher\pengguna\update.go',
    'watcher_app\message_broker\dispatcher\seller\create.go',
    'watcher_app\message_broker\dispatcher\seller\delete.go',
    'watcher_app\message_broker\dispatcher\seller\update.go'
)

foreach ($file in $files) {
    if (Test-Path -Path $file) {
        $content = Get-Content -Raw -Path $file
        
        # Menggunakan \bse\b agar hanya mencari kata "se" yang berdiri sendiri
        $content = $content -replace '\bse\b', 'se_index'
        
        Set-Content -Path $file -Value $content -Encoding utf8
        Write-Host "Updated $file"
    } else {
        Write-Warning "File tidak ditemukan: $file"
    }
}

Write-Host "Proses selesai!"