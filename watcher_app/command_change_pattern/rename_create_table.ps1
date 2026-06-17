Set-Location -Path 'c:\Burung_App\Project_Source\Backend-2'
$paths = Get-ChildItem -Path .\watcher_app\database\cassandra\models -Filter *.go
foreach ($file in $paths) {
    $content = Get-Content -Raw -Path $file.FullName
    $updated = $content -replace 'func \([^\)]*\) CreateTable\(', 'func $1 CreateHistoricalTable('
    if ($updated -ne $content) {
        Set-Content -Path $file.FullName -Value $updated -Encoding utf8
        Write-Host "Updated $($file.Name)"
    }
}
$files = @(
    'c:\Burung_App\Project_Source\Backend-2\watcher_app\database\cassandra\hystorical_db\migrations\up_relation.go',
    'c:\Burung_App\Project_Source\Backend-2\watcher_app\database\cassandra\models\contract.go'
)
foreach ($file in $files) {
    $content = Get-Content -Raw -Path $file
    $updated = $content -replace 'CreateTable\(', 'CreateHistoricalTable('
    if ($updated -ne $content) {
        Set-Content -Path $file -Value $updated -Encoding utf8
        Write-Host "Updated $file"
    }
}
