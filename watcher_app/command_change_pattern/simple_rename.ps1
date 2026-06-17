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
    $content = Get-Content -Raw -Path $file
    
    # Simple replacement: func (receiver) CreateTable( -> func (receiver) CreateHistoricalTable(
    $lines = $content -split '\n'
    $output = @()
    
    foreach ($line in $lines) {
        if ($line -match ') CreateTable\(') {
            # Replace CreateTable with CreateHistoricalTable
            $line = $line -replace '\) CreateTable\(', ') CreateHistoricalTable('
        }
        $output += $line
    }
    
    Set-Content -Path $file -Value ($output -join "`n") -Encoding utf8
    Write-Host "Updated $file"
}

# Also update contract.go interface
$contractFile = 'watcher_app/database/cassandra/models/contract.go'
$contractContent = Get-Content -Raw -Path $contractFile
$contractContent = $contractContent -replace 'CreateTable\(ctx context\.Context, s \*gocql\.Session\) error', 'CreateHistoricalTable(ctx context.Context, s *gocql.Session) error'
Set-Content -Path $contractFile -Value $contractContent -Encoding utf8
Write-Host "Updated $contractFile"

# Update migration call site
$migrationFile = 'watcher_app/database/cassandra/hystorical_db/migrations/up_relation.go'
$migrationContent = Get-Content -Raw -Path $migrationFile
$migrationContent = $migrationContent -replace 'CreateTable\(ctx, session\)', 'CreateHistoricalTable(ctx, session)'
Set-Content -Path $migrationFile -Value $migrationContent -Encoding utf8
Write-Host "Updated $migrationFile"

Write-Host "All files updated successfully!"
