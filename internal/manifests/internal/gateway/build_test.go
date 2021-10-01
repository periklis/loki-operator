package gateway

import (
	"testing"

	lokiv1beta1 "github.com/ViaQ/loki-operator/api/v1beta1"
	"github.com/stretchr/testify/require"
)

func TestBuild_StaticMode(t *testing.T) {
	expTntCfg := `
tenants:
- name: test-a
  id: test
  oidc:
    clientID: test
    clientSecret: test123
    issuerCAPath: /tmp/ca/path
    issuerURL: https://127.0.0.1:5556/dex
    redirectURL: https://localhost:8443/oidc/test-a/callback
    usernameClaim: test
    groupClaim: test
  opa:
    query: data.lokistack.allow
    paths:
    - /etc/lokistack-gateway/rbac.yaml
    - /etc/lokistack-gateway/lokistack-gateway.rego
`
	expRbacCfg := `
roleBindings:
- name: test-a
  roles:
  - read-write
  subjects:
  - kind: user
    name: test@example.com
roles:
- name: some-name
  permissions:
  - read
  resources:
  - metrics
  tenants:
  - test-a
`
	opts := Options{
		Stack: lokiv1beta1.LokiStackSpec{
			Tenants: &lokiv1beta1.TenantsSpec{
				Mode: lokiv1beta1.Static,
				Authentication: []lokiv1beta1.AuthenticationSpec{
					{
						TenantName: "test-a",
						TenantID:   "test",
						OIDC: &lokiv1beta1.OIDCSpec{
							Secret: &lokiv1beta1.TenantSecretSpec{
								Name: "test",
							},
							IssuerURL:     "https://127.0.0.1:5556/dex",
							RedirectURL:   "https://localhost:8443/oidc/test-a/callback",
							GroupClaim:    "test",
							UsernameClaim: "test",
						},
					},
				},
				Authorization: &lokiv1beta1.AuthorizationSpec{
					Roles: []lokiv1beta1.RoleSpec{
						{
							Name:        "some-name",
							Resources:   []string{"metrics"},
							Tenants:     []string{"test-a"},
							Permissions: []lokiv1beta1.PermissionType{"read"},
						},
					},
					RoleBindings: []lokiv1beta1.RoleBindingsSpec{
						{
							Name: "test-a",
							Subjects: []lokiv1beta1.Subject{
								{
									Name: "test@example.com",
									Kind: "user",
								},
							},
							Roles: []string{"read-write"},
						},
					},
				},
			},
		},
		Namespace: "test-ns",
		Name:      "test",
		TenantSecrets: []*Secret{
			{
				TenantName:   "test-a",
				ClientID:     "test",
				ClientSecret: "test123",
				IssuerCAPath: "/tmp/ca/path",
			},
		},
	}
	rbacConfig, tenantsConfig, regoCfg, err := Build(opts)
	require.NoError(t, err)
	require.YAMLEq(t, expTntCfg, string(tenantsConfig))
	require.YAMLEq(t, expRbacCfg, string(rbacConfig))
	require.NotEmpty(t, regoCfg)
}

func TestBuild_DynamicMode(t *testing.T) {
	expTntCfg := `
tenants:
- name: test-a
  id: test
  oidc:
    clientID: test
    clientSecret: test123
    issuerCAPath: /tmp/ca/path
    issuerURL: https://127.0.0.1:5556/dex
    redirectURL: https://localhost:8443/oidc/test-a/callback
    usernameClaim: test
    groupClaim: test
  opa:
    url: http://127.0.0.1:8181/v1/data/observatorium/allow
`
	opts := Options{
		Stack: lokiv1beta1.LokiStackSpec{
			Tenants: &lokiv1beta1.TenantsSpec{
				Mode: lokiv1beta1.Dynamic,
				Authentication: []lokiv1beta1.AuthenticationSpec{
					{
						TenantName: "test-a",
						TenantID:   "test",
						OIDC: &lokiv1beta1.OIDCSpec{
							Secret: &lokiv1beta1.TenantSecretSpec{
								Name: "test",
							},
							IssuerURL:     "https://127.0.0.1:5556/dex",
							RedirectURL:   "https://localhost:8443/oidc/test-a/callback",
							GroupClaim:    "test",
							UsernameClaim: "test",
						},
					},
				},
				Authorization: &lokiv1beta1.AuthorizationSpec{
					OPA: &lokiv1beta1.OPASpec{
						URL: "http://127.0.0.1:8181/v1/data/observatorium/allow",
					},
				},
			},
		},
		Namespace: "test-ns",
		Name:      "test",
		TenantSecrets: []*Secret{
			{
				TenantName:   "test-a",
				ClientID:     "test",
				ClientSecret: "test123",
				IssuerCAPath: "/tmp/ca/path",
			},
		},
	}
	rbacConfig, tenantsConfig, regoCfg, err := Build(opts)
	require.NoError(t, err)
	require.YAMLEq(t, expTntCfg, string(tenantsConfig))
	require.Empty(t, rbacConfig)
	require.Empty(t, regoCfg)
}

func TestBuild_OpenshiftLoggingMode(t *testing.T) {
	expTntCfg := `
tenants:
- name: application
  id: 32e45e3e-b760-43a2-a7e1-02c5631e56e9
  oidc:
    clientID: test
    clientSecret: ZXhhbXBsZS1hcHAtc2VjcmV0
    issuerCAPath: ./tmp/certs/ca.pem
    issuerURL: https://127.0.0.1:5556/dex
    redirectURL: https://localhost:8443/oidc/application/callback
    usernameClaim: name
  opa:
    url: http://127.0.0.1:8080/v1/data/lokistack/allow
- name: infrastructure
  id: 40de0532-10a2-430c-9a00-62c46455c118
  oidc:
    clientID: test
    clientSecret: ZXhhbXBsZS1hcHAtc2VjcmV0
    issuerCAPath: ./tmp/certs/ca.pem
    issuerURL: https://127.0.0.1:5556/dex
    redirectURL: https://localhost:8443/oidc/infrastructure/callback
    usernameClaim: name
  opa:
    url: http://127.0.0.1:8080/v1/data/lokistack/allow
- name: audit
  id: 26d7c49d-182e-4d93-bade-510c6cc3243d
  oidc:
    clientID: test
    clientSecret: ZXhhbXBsZS1hcHAtc2VjcmV0
    issuerCAPath: ./tmp/certs/ca.pem
    issuerURL: https://127.0.0.1:5556/dex
    redirectURL: https://localhost:8443/oidc/audit/callback
    usernameClaim: name
  opa:
    url: http://127.0.0.1:8080/v1/data/lokistack/allow
`
	opts := Options{
		Stack: lokiv1beta1.LokiStackSpec{
			Tenants: &lokiv1beta1.TenantsSpec{
				Mode: lokiv1beta1.OpenshiftLogging,
			},
		},
		Namespace: "test-ns",
		Name:      "test",
		TenantSecrets: []*Secret{
			{
				TenantName:   "test-a",
				ClientID:     "test",
				ClientSecret: "test123",
				IssuerCAPath: "/tmp/ca/path",
			},
		},
	}
	rbacConfig, tenantsConfig, regoCfg, err := Build(opts)
	require.NoError(t, err)
	require.YAMLEq(t, expTntCfg, string(tenantsConfig))
	require.Empty(t, rbacConfig)
	require.Empty(t, regoCfg)
}