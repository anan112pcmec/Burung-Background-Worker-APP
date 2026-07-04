Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
   
    'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\auth\auth_handle.go'

)

foreach ($file in $files) {
    if (Test-Path -Path $file) {
        $content = Get-Content -Raw -Path $file
        
        # Regex untuk mendeteksi '5 * time.Second' atau 'time.Second * 5' (fleksibel spasi)
        # \s* artinya spasi boleh ada atau tidak (misal: 5*time atau 5 * time)
        # $pattern = '(([1-9]|10)\s*\*\s*time\.Second|time\.Second\s*\*\s*([1-9]|10))'
        
        # # Lakukan replace menggunakan regex pattern di atas
        # $content = $content -replace $pattern, 'settings.TimeoutContext'

        #  Menggunakan \bse\b agar hanya mencari kata "se" yang ber
        $content = $content -replace 'environment', 'cache'
        # $pattern = '\*Data mb_cud_serializer\.Parsed\*DataMessage'
        # $replacement = 'Data *mb_cud_serializer.ParsedDataMessage'
        
        # $content = $content -replace $pattern, $replacement
        
        Set-Content -Path $file -Value $content -Encoding utf8
        Write-Host "Updated $file"
    } else {
        Write-Warning "File tidak ditemukan: $file"
    }
}

Write-Host "Proses selesai!"