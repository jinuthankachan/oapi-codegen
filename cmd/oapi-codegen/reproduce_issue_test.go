package main

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jinuthankachan/oapi-codegen/v2/pkg/codegen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEcho5ServerImports(t *testing.T) {
	// Minimal valid OpenAPI spec
	spec := `
openapi: 3.0.0
info:
  title: Test API
  version: 1.0.0
paths:
  /test:
    get:
      responses:
        '200':
          description: OK
`

	cfg := codegen.Configuration{
		PackageName: "testpackage",
		Generate: codegen.GenerateOptions{
			Echo5Server: true,
			EchoServer:  false,
			Models:      true,
		},
	}

	// Load the spec
	loader := openapi3.NewLoader()
	swagger, err := loader.LoadFromData([]byte(spec))
	require.NoError(t, err)

	// Generate code
	code, err := codegen.Generate(swagger, cfg)
	require.NoError(t, err)

	// Assertions
	// Should contain echo/v5
	assert.Contains(t, code, `"github.com/labstack/echo/v5"`)

	// Should NOT contain echo/v4
	if strings.Contains(code, `"github.com/labstack/echo/v4"`) {
		t.Errorf("Generated code should not contain github.com/labstack/echo/v4 when generating for Echo5Server")
	}
}
