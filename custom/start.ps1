<#---
title: Start
tag: start
api: post
---
#>
if ((Split-Path -Leaf (Split-Path  -Parent -Path $PSScriptRoot)) -eq "sessions"){
    Set-Location "$PSScriptRoot/../../caddy/"
    ./caddy run --watch
}
else{
    Set-Location "$PSScriptRoot/../.koksmat/caddy/"
    ./caddy run --watch
}
