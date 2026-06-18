Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
    'watcher_app/database/cassandra/models/table_barang.go'
)

foreach ($file in $files) {
    if (Test-Path -Path $file) {
        $content = Get-Content -Raw -Path $file
        
        # Mengganti semua kata ParseToInsertType menjadi ParseToCUDType
        $content = $content -replace 'ParseToInsertType', 'ParseToCUDType'
        
        Set-Content -Path $file -Value $content -Encoding utf8
        Write-Host "Updated $file"
    } else {
        Write-Warning "File tidak ditemukan: $file"
    }
}

Write-Host "Proses selesai!"