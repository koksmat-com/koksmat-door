
helm template mycaddy (join-path $psscriptroot ".." ".koksmat" "caddy" "ingress" "charts" "caddy-ingress-controller") --namespace=caddy-system > (join-path $psscriptroot "mycaddy.yaml")