# Kubernetes Homelab

GitOps-managed homelab platform, reconciled by [Flux CD](https://fluxcd.io).

## Layout

| Path | Contents |
| :--- | :--- |
| `clusters/homelab/` | Entrypoint for the k3s homelab cluster. Flux `Kustomization`s wired together with `dependsOn`. |
| `clusters/local/` | Entrypoint for a throwaway local (kind) cluster. OpenBao only. |
| `control-plane/releases/` | `HelmRelease` and `HelmRepository` objects, grouped by concern. |
| `control-plane/manifests/` | Umbrella Helm charts and plain kustomize bases consumed by the releases above. |
| `ansible/` | Node provisioning: apt maintenance, NTP, UFW, DNS, k3s. |

## Quickstart

Bootstrap Flux against a cluster. `GITHUB_TOKEN` needs `repo` scope:

```sh
export GITHUB_TOKEN=$(gh auth token)

# homelab
flux bootstrap github --owner=TineoC --repository=homelab \
  --branch=main --path=clusters/homelab --personal

# local kind cluster (OpenBao only)
kind create cluster --name homelab-local
flux bootstrap github --owner=TineoC --repository=homelab \
  --branch=main --path=clusters/local --personal
```

Watch it converge:

```sh
flux get kustomizations --watch
flux get helmreleases -A
```

## OpenBao

OpenBao runs standalone with `file` storage, so it starts **sealed** and stays sealed
across pod restarts. Initialize once:

```sh
kubectl -n openbao exec -it openbao-0 -- bao operator init -key-shares=1 -key-threshold=1
```

Save the unseal key and root token somewhere safe — they are printed once and are
never stored in this repo. Then unseal (repeat after every pod restart):

```sh
kubectl -n openbao exec -it openbao-0 -- bao operator unseal <unseal-key>
```

Reach the UI at <http://127.0.0.1:8200> with:

```sh
kubectl -n openbao port-forward svc/openbao 8200:8200
```

External Secrets reads from OpenBao through the `vault-store` `ClusterSecretStore`
(`control-plane/manifests/external-secrets/`), which expects a `kubernetes` auth
mount and an `eso-controller` role inside OpenBao.
