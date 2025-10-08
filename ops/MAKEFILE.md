# Make targets Codex should use

- `make bootstrap` — install controllers/CRDs (Argo CD, Keycloak Operator, Envoy Gateway)
- `make deploy-ldap` — deploy OpenLDAP + LDIF seeds
- `make deploy-keycloak` — deploy Keycloak instance + realm/clients/federation
- `make deploy-gateway` — deploy Envoy Gateway + Gateway + policies
- `make deploy-apps` — deploy 4 demo apps and HTTPRoutes
- `make deploy-iam` — runs ldap + keycloak targets
- `make deploy` — runs gateway + apps + iam
- `make test` — basic validation (realm, clients, token, protected routes)
- `make destroy` — removes all components (reverse order)
