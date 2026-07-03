Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
    # 'watcher_app\message_broker\consumer\method.go',
    # 'watcher_app\message_broker\dispatcher\kurir\create.go',
    # 'watcher_app\message_broker\dispatcher\kurir\delete.go',
    # 'watcher_app\message_broker\dispatcher\kurir\update.go',
    # 'watcher_app\message_broker\dispatcher\pengguna\create.go', # <-- Sudah ditambahkan koma yang kurang
    # 'watcher_app\message_broker\dispatcher\pengguna\delete.go',
    # 'watcher_app\message_broker\dispatcher\pengguna\update.go',
    # 'watcher_app\message_broker\dispatcher\seller\create.go',
    # 'watcher_app\message_broker\dispatcher\seller\delete.go',
    # 'watcher_app\message_broker\dispatcher\seller\update.go'

    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\alamat_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\credential_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\informasi_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\media_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\pengiriman_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\profiling_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\rekening_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\kurir_service\social_media_services\services_handle.go'

    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\alamat_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\barang_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\credential_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\media_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\profiling_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\social_media_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\transaction_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\pengguna_service\wishlist_services\services_handle.go'

    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\alamat_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\barang_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\credential_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\diskon_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\etalase_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\jenis_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\media_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\profiling_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\social_media_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\social_media_services\services_handle.go',
    # 'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\seller_service\transaksi_services\services_handle.go'

    'C:\Burung_App\Project_Source\Backend-2\watcher_app\service_handle\sistem_services\payout\services_handle.go'

)

foreach ($file in $files) {
    if (Test-Path -Path $file) {
        $content = Get-Content -Raw -Path $file
        
        # Menggunakan \bse\b agar hanya mencari kata "se" yang berdiri sendiri
        $content = $content -replace 'environment', 'cache'
        
        Set-Content -Path $file -Value $content -Encoding utf8
        Write-Host "Updated $file"
    } else {
        Write-Warning "File tidak ditemukan: $file"
    }
}

Write-Host "Proses selesai!"