package testutil

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
)

type Route struct {
	Name        string
	Path        string
	Method      string
	ContentType string
	Summary     string
}

type OpenAPITester interface {
	RunRouterTest(t *testing.T, router http.Handler, pathPrefix string)
}

type openAPITester struct {
	doc *openapi3.T
}

func NewOpenAPITester(t *testing.T, path string) OpenAPITester {
	t.Helper()
	tester, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return &openAPITester{doc: tester}
}

func (o *openAPITester) getRoutes(t *testing.T, pathPrefix string) []*Route {
	t.Helper()

	routes := make([]*Route, 0)
	for path, pathItem := range o.doc.Paths.Map() {
		if !strings.HasPrefix(path, pathPrefix) {
			continue
		}
		for method, op := range pathItem.Operations() {
			var contentType string
			// request body的检查比较困难暂时先不考虑对应
			for ct := range op.RequestBody.Value.Content {
				contentType = ct
			}
			routes = append(routes, &Route{
				Name:        fmt.Sprintf("%s %s", path, method),
				Path:        path,
				Method:      method,
				ContentType: contentType,
				Summary:     op.Summary,
			})
		}
	}
	return routes
}

func (o *openAPITester) validatePathAndMethod(t *testing.T, path, method string) (*openapi3.Operation, error) {
	t.Helper()

	pi := o.doc.Paths.Find(path)
	if pi == nil {
		return nil, fmt.Errorf("path '%s' does not exist", path)
	}

	op := pi.GetOperation(method)
	if op == nil {
		return nil, fmt.Errorf("method '%s' does not exist in path '%s'", method, path)
	}

	return op, nil
}

func (o *openAPITester) convertParamsInPath(t *testing.T, op *openapi3.Operation, path string) (string, error) {
	t.Helper()

	params := make([]string, 0)
	for _, ref := range op.Parameters {
		if ref.Value.In != openapi3.ParameterInPath {
			continue
		}

		pn := ref.Value.Name
		if !strings.Contains(path, pn) {
			return "", fmt.Errorf("param {%s} does not exist in path '%s'", pn, path)
		}
		params = append(params, pn)
	}

	for _, paramStr := range params {
		pm := op.Parameters.GetByInAndName(openapi3.ParameterInPath, paramStr)
		if pm == nil {
			return "", fmt.Errorf("wrong param name: %s", paramStr)
		}

		var ex string
		if pm.Example != nil {
			ex = fmt.Sprintf("%v", pm.Example)
		}
		if pm.Schema != nil {
			rawEx, err := paramsExample(pm.Schema.Value)
			if err != nil {
				t.Fatalf("failed to get example")
			}
			ex = fmt.Sprintf("%v", rawEx)
		}
		if ex == "" {
			return "", fmt.Errorf("param %s has no example", paramStr)
		}
		path = strings.Replace(path, fmt.Sprintf("{%s}", paramStr), ex, 1)
	}
	return path, nil
}

func (o *openAPITester) ExamplePath(t *testing.T, path, method string) string {
	t.Helper()

	op, err := o.validatePathAndMethod(t, path, method)
	if err != nil {
		t.Fatal(err)
	}

	exPath, err := o.convertParamsInPath(t, op, path)
	if err != nil {
		t.Fatal(err)
	}

	return exPath
}

func (o *openAPITester) RunRouterTest(t *testing.T, router http.Handler, pathPrefix string) {
	t.Helper()

	r := chi.NewRouter()
	r.Mount(pathPrefix, router)

	server := httptest.NewServer(r)
	defer server.Close()

	routes := o.getRoutes(t, pathPrefix)
	for _, tt := range routes {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			exPath := o.ExamplePath(t, tt.Path, tt.Method)
			rec, req := NewRequestAndRecorder(tt.Method, exPath, nil)
			if tt.ContentType != "" {
				req.Header.Set("Content-Type", tt.ContentType)
			}

			req = req.WithContext(SetUpContextWithDefault())
			r.ServeHTTP(rec, req)

			result := rec.Result()
			defer result.Body.Close()
			switch result.StatusCode {
			case http.StatusOK:
			case http.StatusNotImplemented:
				t.Skip("this route is not implemented")
			default:
				t.Fatalf("seems not to be declared in the router or openapi definition mistake: status=%d", result.StatusCode)
			}
		})
	}
}

// nolint
func paramsExample(schema *openapi3.Schema) (out interface{}, err error) {
	ex := schema.Example
	if ex != nil {
		return ex, nil
	}
	switch {
	case schema.Type == "boolean":
		return true, nil
	case schema.Type == "number", schema.Type == "integer":
		value := 0.0

		if schema.Min != nil && *schema.Min > value {
			value = *schema.Min
			if schema.ExclusiveMin {
				if schema.Max != nil {
					// Make the value half way.
					value = (*schema.Min + *schema.Max) / 2.0
				} else {
					value++
				}
			}
		}

		if schema.Max != nil && *schema.Max < value {
			value = *schema.Max
			if schema.ExclusiveMax {
				if schema.Min != nil {
					// Make the value half way.
					value = (*schema.Min + *schema.Max) / 2.0
				} else {
					value--
				}
			}
		}

		if schema.MultipleOf != nil && int(value)%int(*schema.MultipleOf) != 0 {
			value += float64(int(*schema.MultipleOf) - (int(value) % int(*schema.MultipleOf)))
		}

		if schema.Type == "integer" {
			return int(value), nil
		}

		return value, nil
	case schema.Type == "string":
		if ex := stringFormatExample(schema.Format); ex != "" {
			return ex, nil
		}

		example := "string"

		for schema.MinLength > uint64(len(example)) {
			example += example
		}

		if schema.MaxLength != nil && *schema.MaxLength < uint64(len(example)) {
			example = example[:*schema.MaxLength]
		}

		return example, nil
	}
	return nil, errors.New("no params example")
}

func stringFormatExample(format string) string {
	switch format {
	case "date":
		// https://tools.ietf.org/html/rfc3339
		return "2018-07-23"
	case "date-time":
		// This is the date/time of API Sprout's first commit! :-)
		return "2018-07-23T22:58:00-07:00"
	case "time":
		return "22:58:00-07:00"
	case "email":
		return "email@example.com"
	case "hostname":
		// https://tools.ietf.org/html/rfc2606#page-2
		return "example.com"
	case "ipv4":
		// https://tools.ietf.org/html/rfc5737
		return "198.51.100.0"
	case "ipv6":
		// https://tools.ietf.org/html/rfc3849
		return "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	case "uri":
		return "https://tools.ietf.org/html/rfc3986"
	case "uri-template":
		// https://tools.ietf.org/html/rfc6570
		return "http://example.com/dictionary/{term:1}/{term}"
	case "json-pointer":
		// https://tools.ietf.org/html/rfc6901
		return "#/components/parameters/term"
	case "regex":
		// https://stackoverflow.com/q/3296050/164268
		return "/^1?$|^(11+?)\\1+$/"
	case "uuid":
		// https://www.ietf.org/rfc/rfc4122.txt
		return "f81d4fae-7dec-11d0-a765-00a0c91e6bf6"
	case "password":
		return "********"
	}

	return ""
}
