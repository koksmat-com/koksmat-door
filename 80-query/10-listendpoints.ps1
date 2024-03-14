<#---
title: List Endpoints
tag: endpoints
api: post
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

Get-Content "Caddyfile" 

