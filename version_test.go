package version_test

import (
	"testing"
)

import (
	"github.com/chuckpreslar/version"
)

func TestMajorVersion(t *testing.T) {
	var (
		ver = version.Establish(1, 0, 0)
		act = ver.Major()
		exp = 1
	)

	if exp != act {
		t.Errorf("expected value of major version to be %d, got %d", exp, act)
	}
}

func TestMinorVersion(t *testing.T) {
	var (
		ver = version.Establish(0, 1, 0)
		act = ver.Minor()
		exp = 1
	)

	if exp != act {
		t.Errorf("expected value of minor version to be %d, got %d", exp, act)
	}
}

func TestPatchVersion(t *testing.T) {
	var (
		ver = version.Establish(0, 0, 1)
		act = ver.Patch()
		exp = 1
	)

	if exp != act {
		t.Errorf("expected value of patch version to be %d, got %d", exp, act)
	}
}

func TestLabelVersion(t *testing.T) {
	var (
		ver = version.Establish(0, 0, 1, "rc1")
		act = ver.Label()
		exp = "rc1"
	)

	if exp != act {
		t.Errorf("expected value of label version to be %s, got %s", exp, act)
	}
}

func TestVersionStringWithLabel(t *testing.T) {
	var (
		ver = version.Establish(1, 0, 0, "rc1")
		act = ver.String()
		exp = "1.0.0-rc1"
	)

	if exp != act {
		t.Errorf("expected string value of version to be %s, got %s", exp, act)
	}
}

func TestVersionStringWithoutLabel(t *testing.T) {
	var (
		ver = version.Establish(1, 0, 0)
		act = ver.String()
		exp = "1.0.0"
	)

	if exp != act {
		t.Errorf("expected string value of version to be %s, got %s", exp, act)
	}
}
