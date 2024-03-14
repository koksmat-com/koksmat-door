<#---
title: Show Caddyfile
tag: show
icon: magic.png
api: post
---
#>

# Show Caddyfile

if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions"){
    $path = join-path $PSScriptRoot ".." ".."
}
else{
  $path = join-path $PSScriptRoot ".." ".koksmat/"

}

$koksmatDir = Resolve-Path $path



 

 Get-Content (Join-Path $koksmatDir "caddy" "Caddyfile")
