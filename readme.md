---
title: Koksmat Door
description: Entry point to the kitchens
---
#

# Koksmat Door

Koksmat Door is based on the Caddy web server which protects all exposed services.

The services is only need in a development enviroment where the purpose is to allow for ssl connections to services running on localhost.

## .koksmat folder
```
.
├── app
└── caddy
    ├── Caddyfile
    ├── caddy
    └── ingress
```

## Magic Buttons 
To support you when developing this 2 links should be usefull

[Start Caddy locally](./koksmat-door/station/00-magic/view/10-runcaddy.ps1)
[Show Caddyfile](./koksmat-door/station/00-magic/view/20-showconfig.ps1)

## Deployment

Check [Provision](./koksmat-door/station/60-provision/view/20-install.ps1)