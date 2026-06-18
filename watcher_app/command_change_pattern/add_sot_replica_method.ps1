Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
    'watcher_app/database/cassandra/models/engagement_entity.go',
    'watcher_app/database/cassandra/models/table_pengiriman.go',
    'watcher_app/database/cassandra/models/table_media.go',
    'watcher_app/database/cassandra/models/table_payout.go',
    'watcher_app/database/cassandra/models/table_rekening.go',
    'watcher_app/database/cassandra/models/table_barang.go',
    'watcher_app/database/cassandra/models/table_transaksi.go'
)

foreach ($file in $files) {
    if (Test-Path -Path $file) {
        $resolvedPath = (Resolve-Path $file).Path
        $content = [System.IO.File]::ReadAllText($resolvedPath) -replace "`r`n", "`n"
        
        $newMethodsGenerated = @()
        
        # -------------------------------------------------------------------------
        # FASE 1: SCAN SEMUA METHOD "TableNameHistorical" (Multi-Match)
        # -------------------------------------------------------------------------
        # Regex mencari semua fungsi TableNameHistorical secara global di dalam file
        $tableNameRegex = [regex]'(?m)^func\s+\(\s*(\w+)\s+(\*?\w+)\s*\)\s+TableNameHistorical\(\)\s+string\s*\{\s*return\s*"([^"]+)"\s*\}'
        $tableMatches = $tableNameRegex.Matches($content)
        
        foreach ($match in $tableMatches) {
            $receiverVar = $match.Groups[1].Value
            $structName   = $match.Groups[2].Value
            $tableName    = $match.Groups[3].Value
            
            # Buat nama fungsi barunya
            $sotMethodName = "TableNameSotReplica"
            
            # Cek apakah struct ini sudah punya fungsi tersebut di dalam file asli
            # Contoh pemastian unik: "func (s Seller) TableNameSotReplica"
            $checkString = "func ($receiverVar $structName) $sotMethodName"
            if ($content -notlike "*$checkString*") {
                $newTableName = $tableName -replace '_historical$', '_sot_replica'
                $newMethod = "`nfunc ($receiverVar $structName) TableNameSotReplica() string {`n`treturn `"$newTableName`"`n}"
                $newMethodsGenerated += $newMethod
            }
        }

        # -------------------------------------------------------------------------
        # FASE 2: SCAN SEMUA METHOD "CreateHistoricalTable" (Multi-Match + Brace Tracker)
        # -------------------------------------------------------------------------
        # Kita scan semua teks "CreateHistoricalTable" di file ini
        $funcName = "CreateHistoricalTable"
        $searchIndex = 0
        
        while (($searchIndex = $content.IndexOf($funcName, $searchIndex)) -ge 0) {
            # Mundur untuk mencari kata kunci "func" pembuka method ini
            $startPos = $content.LastIndexOf("func", $searchIndex)
            
            if ($startPos -ge 0) {
                # Hitung kurung kurawal secara akurat untuk mengekstrak 1 blok method penuh
                $openBraces = 0
                $endPos = -1
                for ($i = $startPos; $i -lt $content.Length; $i++) {
                    if ($content[$i] -eq '{') { $openBraces++ }
                    if ($content[$i] -eq '}') { 
                        $openBraces-- 
                        if ($openBraces -eq 0) {
                            $endPos = $i
                            break
                        }
                    }
                }
                
                if ($endPos -gt $startPos) {
                    # Berhasil dapet 1 blok penuh milik struct saat ini
                    $historicalBlock = $content.Substring($startPos, ($endPos - $startPos) + 1)
                    
                    # Ekstrak tipe struct-nya dari block ini untuk validasi duplikat
                    # contoh: func (s *Seller) CreateHistoricalTable
                    if ($historicalBlock -match 'func\s+\(\s*\w+\s+(\*?\w+)\s*\)\s+CreateHistoricalTable') {
                        $structType = $Matches[1]
                        $checkSotFunc = "CreateSotReplicaTable"
                        
                        # Pastikan struct ini belum dikasih method SotReplicaTable
                        if ($content -notlike "*$structType) $checkSotFunc*") {
                            
                            # Lakukan konversi teks
                            $sotBlock = $historicalBlock
                            $sotBlock = $sotBlock -replace 'CreateHistoricalTable', 'CreateSotReplicaTable'
                            $sotBlock = $sotBlock -replace 'TableNameHistorical\(\)', 'TableNameSotReplica()'
                            $sotBlock = $sotBlock -replace 'membuat tabel %s', 'membuat tabel sot_replica'
                            
                            # Hapus field historical (aman dari variasi spasi/tab)
                            $sotBlock = $sotBlock -replace '\n\s*tahun_update\s+int,?', ''
                            $sotBlock = $sotBlock -replace '\n\s*bulan_update\s+int,?', ''
                            $sotBlock = $sotBlock -replace '\n\s*event_time\s+timestamp,?', ''
                            
                            # Ubah komposisi Primary Key
                            $sotBlock = $sotBlock -replace '(?s)PRIMARY KEY\s*\(\(\s*id\s*,\s*tahun_update\s*,\s*bulan_update\s*\)\s*,\s*event_time\s*\)', 'PRIMARY KEY (id)'
                            
                            $newMethodsGenerated += "`n" + $sotBlock
                        }
                    }
                }
            }
            # Lanjut cari ke method CreateHistoricalTable berikutnya di file yang sama
            $searchIndex += $funcName.Length
        }

        # -------------------------------------------------------------------------
        # FASE 3: PENULISAN KEMBALI KE FILE
        # -------------------------------------------------------------------------
        if ($newMethodsGenerated.Count -gt 0) {
            # Gabungkan semua method baru yang berhasil digenerate ke akhir file
            $content += "`n" + ($newMethodsGenerated -join "`n")
            
            [System.IO.File]::WriteAllText($resolvedPath, $content, [System.Text.Encoding]::UTF8)
            Write-Host "✅ Berhasil menambahkan $($newMethodsGenerated.Count) method baru di: $file" -ForegroundColor Green
        } else {
            Write-Host "🟡 Dilewati (Semua struct sudah memiliki method SotReplica): $file" -ForegroundColor Yellow
        }
    } else {
        Write-Warning "❌ File tidak ditemukan: $file"
    }
}

Write-Host "Selesai memproses seluruh file dan ratusan struct kamu secara global!" -ForegroundColor Cyan