// SPDX-License-Identifier: MIT

package version

import (
	"testing"

	"github.com/issue9/assert/v2"
)

func TestSemVersion_Compare(t *testing.T) {
	a := assert.New(t, false)

	v1 := &SemVersion{
		Major: 1,
	}
	v2 := &SemVersion{
		Major: 1,
	}

	a.True(v1.Compare(v2) == 0)

	v1.Minor = 2
	a.True(v1.Compare(v2) > 0)
	v2.Minor = 2

	v2.Patch = 3
	a.True(v1.Compare(v2) < 0)
	v1.Patch = 3

	// build 不参与运算
	v1.Build = "111"
	v2.Build = "222"
	a.True(v1.Compare(v2) == 0)

	v1.PreRelease = "alpha"
	a.True(v1.Compare(v2) < 0)
	a.True(v2.Compare(v1) > 0)

	v2.PreRelease = "beta"
	a.True(v1.Compare(v2) < 0)

	// 相等的 preRelease
	v1.PreRelease = "1.alpha"
	v2.PreRelease = "1.alpha"
	a.True(v1.Compare(v2) == 0)

	// preRelease 数值的比较
	v1.PreRelease = "11.alpha"
	v2.PreRelease = "9.alpha"
	a.True(v1.Compare(v2) > 0)
}

func TestSemVersion_CompareString(t *testing.T) {
	a := assert.New(t, false)

	v1 := &SemVersion{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	ret, err := v1.CompareString("1.2.2+build")
	a.NotError(err).True(ret > 0)

	ret, err = v1.CompareString("1.2.3+build")
	a.NotError(err).True(ret == 0)

	ret, err = v1.CompareString("1.2.3-alpha+build")
	a.NotError(err).True(ret > 0)

	ret, err = v1.CompareString("1.2.3-alpha")
	a.NotError(err).True(ret > 0)
}

func TestSemVersion_CompatibleString(t *testing.T) {
	a := assert.New(t, false)

	v1 := &SemVersion{
		Major: 1,
		Minor: 2,
		Patch: 3,
	}

	ret, err := v1.CompatibleString("1.2.2+build")
	a.NotError(err).True(ret)

	ret, err = v1.CompatibleString("0.2.2+build")
	a.NotError(err).False(ret)
}

func TestSemVersion_String(t *testing.T) {
	a := assert.New(t, false)

	sv := &SemVersion{
		Major: 1,
	}
	a.Equal(sv.String(), "1.0.0")

	sv.Minor = 22
	a.Equal(sv.String(), "1.22.0")

	sv.Patch = 1234
	a.Equal(sv.String(), "1.22.1234")

	sv.Build = "20160615"
	a.Equal(sv.String(), "1.22.1234+20160615")

	sv.PreRelease = "alpha1.0"
	a.Equal(sv.String(), "1.22.1234-alpha1.0+20160615")

	sv.Build = ""
	a.Equal(sv.String(), "1.22.1234-alpha1.0")
}

func TestSemVerCompare(t *testing.T) {
	a := assert.New(t, false)

	v, err := SemVerCompare("1.0.0", "1.0.0")
	a.NotError(err).True(v == 0)

	v, err = SemVerCompare("1.2.0", "1.0.0")
	a.NotError(err).True(v > 0)

	v, err = SemVerCompare("1.2.0", "1.2.1")
	a.NotError(err).True(v < 0)
}

func TestSemVerCompatible(t *testing.T) {
	a := assert.New(t, false)

	ret, err := SemVerCompatible("1.0.0", "1.0.0")
	a.NotError(err).True(ret)

	ret, err = SemVerCompatible("0.0.0", "1.0.0")
	a.NotError(err).False(ret)
}

func TestSemVerValid(t *testing.T) {
	a := assert.New(t, false)

	a.True(SemVerValid("1.0.0+build"))
	a.True(SemVerValid("1.1.2-alpha.1"))
	a.True(SemVerValid("1.1.22-alpha.2.1+build"))

	a.False(SemVerValid("1.1.22.2"))
	a.False(SemVerValid("1.1.22.2-alpha.2.1+build"))
}
