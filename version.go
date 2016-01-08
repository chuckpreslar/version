package version

import (
	"fmt"
)

// Version ...
type Version struct {
	major int
	minor int
	patch int
	label string
}

// Major ...
func (v Version) Major() int {
	return v.major
}

// Minor ...
func (v Version) Minor() int {
	return v.minor
}

// Patch ...
func (v Version) Patch() int {
	return v.patch
}

// Label ...
func (v Version) Label() string {
	return v.label
}

// String ...
func (v Version) String() string {
	var version = fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)

	if "" != v.label {
		version = fmt.Sprintf("%s-%s", version, v.label)
	}

	return version
}

// Establish ...
func Establish(major, minor, patch int, label ...string) Version {
	var version = Version{major: major, minor: minor, patch: patch}

	if 0 < len(label) {
		version.label = label[0]
	}

	return version
}
