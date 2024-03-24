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

$ErrorActionPreference = "SilentlyContinue"
$ProgressPreference = "SilentlyContinue"

$magicapppath = join-path $env:KITCHENROOT ".koksmat" "magicapps.json"
$magicapps = Get-Content $magicapppath | ConvertFrom-Json

foreach ($mapicapp in $magicapps.apps) {
  write-host "Adding $($mapicapp.servicename) to Caddyfile"
  $appName = $mapicapp.servicename
  $xx = . $appName "info" "ping"
 
    write-host $xx -ForegroundColor Yellow
}


