# K8s manifests (GitOps layout)

## Argo CD apps
- `applications/argocd-bootstrap.yaml`         # (optional) if Argo CD not preinstalled
- `applications/envoy-gateway.yaml`
- `applications/openldap.yaml`
- `applications/keycloak-operator.yaml`
- `applications/keycloak-instance.yaml`
- `applications/keycloak-realm.yaml`
- `applications/keycloak-ldap-federation.yaml`
- `applications/apps.yaml`                     # parent app for app-a..d (or ApplicationSet)

## Manifests (referenced by the Argo apps)
- `manifests/envoy-gateway/`        # Gateway, policies
- `manifests/openldap/`             # HelmRelease/Chart + LDIF configmap
- `manifests/keycloak/operator/`    # Operator install
- `manifests/keycloak/instance/`    # Keycloak CR
- `manifests/keycloak/realm/`       # Realm + Clients + Mappers
- `manifests/keycloak/federation/`  # LDAP provider CR + sync policies
- `manifests/apps/app-a/`           # Deployment, Service, HTTPRoute
- `manifests/apps/app-b/`
- `manifests/apps/app-c/`
- `manifests/apps/app-d/`
