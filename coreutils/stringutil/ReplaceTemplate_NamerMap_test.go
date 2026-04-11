package stringutil

import (
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage_internal — replaceTemplate.UsingNamerMapOptions
//
// Target: replaceTemplate.go:316-341 (both branches)
// The namer interface is unexported, so this must be an internal test.
// ══════════════════════════════════════════════════════════════════════════════

type testNamerInternal struct {
	name string
}

func (n testNamerInternal) Name() string {
	return n.name
}

func Test_UsingNamerMapOptions_CurlyKeys(t *testing.T) {
	// Arrange
	rt := replaceTemplate{}
	namerMap := map[namer]string{
		testNamerInternal{name: "host"}: "example.com",
		testNamerInternal{name: "port"}: "8080",
	}

	// Act
	result := rt.UsingNamerMapOptions(
		true,
		"https://{host}:{port}/api",
		namerMap,
	)

	// Assert
	if result != "https://example.com:8080/api" {
		t.Fatalf("expected https://example.com:8080/api, got %s", result)
	}
}

func Test_UsingNamerMapOptions_DirectKeys(t *testing.T) {
	// Arrange
	rt := replaceTemplate{}
	namerMap := map[namer]string{
		testNamerInternal{name: "HOST"}: "example.com",
	}

	// Act
	result := rt.UsingNamerMapOptions(
		false,
		"https://HOST/api",
		namerMap,
	)

	// Assert
	if result != "https://example.com/api" {
		t.Fatalf("expected https://example.com/api, got %s", result)
	}
}
