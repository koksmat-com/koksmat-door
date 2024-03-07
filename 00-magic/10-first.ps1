<#---
title: Run Caddy
tag: caddy
icon: magic.png
---
#>

# Run Caddy

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions"){
    $path = join-path $PSScriptRoot ".." ".."
}
else{
  $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path

Set-Location (Join-Path $koksmatDir "caddy")



./caddy run --watch 