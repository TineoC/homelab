helm repo add external-dns https://kubernetes-sigs.github.io/external-dns/
helm upgrade --install external-dns bitnami/external-dns \
  --namespace kube-system \
  --create-namespace \
  --set provider=cloudflare \
  --set domainFilters[0]=tineochristopher.com \
  --set cloudflare.apiToken=76Xw6_8JnKYCQ-V7Tggi41Ht9mwLzXVRtrbrlKDO \
  --set policy=sync \
  --set txtOwnerId=external-dns
