Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'

$files = @(
    'watcher_app/database/cassandra/models/table_barang.go',
    'watcher_app/database/cassandra/models/engagement_entity.go',
    'watcher_app/database/cassandra/models/table_entity.go',
    'watcher_app/database/cassandra/models/table_media.go',
    'watcher_app/database/cassandra/models/table_payout.go',
    'watcher_app/database/cassandra/models/table_pengiriman.go',
    'watcher_app/database/cassandra/models/table_rekening.go',
    'watcher_app/database/cassandra/models/table_transaksi.go'
)

foreach ($file in $files) {
    # Get original content from git HEAD
    $original = git show HEAD:$file
    
    # Extract all original function signatures with receivers
    $originalFuncs = @{}
    $original -split '\n' | Where-Object { $_ -match 'func \(([^)]+)\) CreateTable\(' } | ForEach-Object {
        if ($_ -match 'func \(([^)]+)\) CreateTable\(') {
            $receiver = $matches[1]
            $originalFuncs[$receiver] = $true
        }
    }
    
    # Read current file
    $current = Get-Content -Raw -Path $file
    
    # Replace each broken receiver
    foreach ($receiver in $originalFuncs.Keys) {
        $pattern = "func \(`$1\) CreateHistoricalTable\((.*?context\.Context.*?\) error \{)"
        $replacement = "func ($receiver) CreateHistoricalTable(`$1)"
        $current = $current -replace "func \\\(\`\$1\\\) CreateHistoricalTable\(", "func ($receiver) CreateHistoricalTable("
    }
    
    # Write back
    Set-Content -Path $file -Value $current -Encoding utf8
    Write-Host "Fixed $file"
}
