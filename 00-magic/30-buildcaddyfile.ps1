<#---
title: Make Caddyfile
tag: make
icon: magic.png
api: post
---
#>

param (
  $test = $false
)
$magicapppath = join-path $env:KITCHENROOT ".koksmat" "magicapps.json"

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions") {
  $path = join-path $PSScriptRoot ".." ".."
}
else {
  $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path
$caddyFilePath = join-path $koksmatDir "caddy" "Caddyfile"

$magicapps = Get-Content $magicapppath | ConvertFrom-Json

$caddyFileData = @"
###############################################################################
#
# File generated by $($PSScriptRoot)
# Input: $magicapppath 
#
###############################################################################
"@

foreach ($mapicapp in $magicapps.apps) {
  $port = $mapicapp.port
  $testdata = ""
  if ($test) {
    $testdata = @"
  respond "I am https://localhost:$port pointing to $($mapicapp.web)"
"@
  }
  $caddyFileData += @"

# $($mapicapp.name)

localhost:$port {
  $testdata
  reverse_proxy /* $($mapicapp.web)
}  

"@ 
  <# $mapicapp is the current item #>
}


Out-File -InputObject $caddyFileData -FilePath $caddyFilePath -Encoding utf8NoBOM