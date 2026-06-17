Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
    'watcher_app/database/cassandra/models/engagement_entity.go',
    'watcher_app/database/cassandra/models/table_pengiriman.go',
    'watcher_app/database/cassandra/models/table_entity.go',
    'watcher_app/database/cassandra/models/table_media.go',
    'watcher_app/database/cassandra/models/table_payout.go',
    'watcher_app/database/cassandra/models/table_rekening.go',
    'watcher_app/database/cassandra/models/table_transaksi.go'
)

foreach ($file in $files) {
    Write-Host "Processing $file..."
    
    # Get original content from git
    $original = git show HEAD:$file 2>$null
    if (-not $original) {
        Write-Host "  ERROR: Could not get file from git"
        continue
    }
    
    # Extract all receivers from original
    $receivers = @()
    $original -split '\n' | ForEach-Object {
        if ($_ -match 'func \(([^)]+)\) CreateTable\(') {
            $receivers += $matches[1]
        }
    }
    
    Write-Host "  Found $($receivers.Count) receivers"
    
    # Read current file
    $current = Get-Content -Raw -Path $file
    
    # For each receiver, replace $1 with the actual receiver
    $count = 0
    foreach ($receiver in $receivers) {
        if ($current -match 'func \(\$1\) CreateHistoricalTable\(') {
            $current = $current -replace 'func \(\$1\) CreateHistoricalTable\(', "func ($receiver) CreateHistoricalTable(" , 1
            $count++
        }
    }
    
    Write-Host "  Replaced $count receivers"
    
    # Write back
    Set-Content -Path $file -Value $current -Encoding utf8
    Write-Host "  Done"
}

Write-Host "All files processed!"
