# Deployment workflow

## 0) Pre-reqs
- Cluster: Kubernetes 1.32 (Gateway API installed by Envoy Gateway).
- Argo CD installed (can be part of bootstrap below).
- `kubectl`, `helm`, `kustomize`, `make`.

## 1) Bootstrap operators/controllers
- Argo CD (if not already)
- Keycloak Operator
- Envoy Gateway (installs Gateway API CRDs)

Run:
```bash
make bootstrap
