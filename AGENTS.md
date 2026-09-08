# AGENTS.md — Repo rules for Codex

## Goal
Provision IAM for 4 demo services using Keycloak (operator), federated to OpenLDAP as an AD alternative. Enforce authN/authZ at the edge via Gateway API (Envoy Gateway). No Ingress for control-plane apps.

## Architecture constraints
- GitOps via Flux `Kustomization`/`HelmRelease` objects only. No kubectl patching by hand.
- Cluster entrypoints live in `clusters/<cluster>/`; reusable releases in `control-plane/releases/`;
  raw charts and kustomize bases stay in `control-plane/manifests/`.
- Ordering is expressed with `dependsOn`, not sync-waves.
- Namespaces:
  - `iam` → Keycloak operator + Keycloak
  - `directory` → OpenLDAP
  - `gateway-system` → Envoy Gateway controller
  - `apps` → Four demo services (apps a–d)
  - `control-plane` → internal tools (no Ingress; traffic via Gateway only)
- Secrets via External Secrets or plain K8s secrets placeholders (ok for demo).
- All edge exposure through **Gateway API**. No `Ingress` resources.
- Authorization done by Envoy/Gateway using **JWT from Keycloak** and **claims/roles**.

## Definition of Done
- `flux bootstrap github --owner=TineoC --repository=homelab --path=clusters/<cluster>` installs the
  Flux controllers; Flux then reconciles Envoy Gateway and the Keycloak Operator.
- `make deploy` applies:
  - OpenLDAP with seeded users/groups/departments
  - Keycloak instance + Realm + LDAP Federation + OIDC clients (one per app)
  - Envoy Gateway + Gateway + HTTPRoutes + Auth policies
  - Four apps (a–d) with Services and HTTPRoutes
- `make test` validates:
  - Keycloak realm + clients exist
  - JWT can be obtained via OIDC device code or password grant (demo)
  - Each app route enforces the intended group/department access
- README updated with quickstart and curl examples.
- No exposed Ingress in `control-plane`.

## PR template (Codex must fill)
- Summary
- What changed (components)
- Verification (commands run + outputs)
- Risks/Mitigations
