# AGENTS.md — Repo rules for Codex

## Goal
Provision IAM for 4 demo services using Keycloak (operator), federated to OpenLDAP as an AD alternative. Enforce authN/authZ at the edge via Gateway API (Envoy Gateway). No Ingress for control-plane apps.

## Architecture constraints
- GitOps via Argo CD Applications/ApplicationSets only. No kubectl patching by hand.
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
- `make bootstrap` installs CRDs/controllers (Argo CD, Envoy Gateway, Keycloak Operator).
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
