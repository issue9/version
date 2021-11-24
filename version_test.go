// SPDX-License-Identifier: MIT

package version

import (
	"testing"

	"github.com/issue9/assert/v2"
)

func TestParse(t *testing.T) {
	a := assert.New(t, false)

	semver := &SemVersion{}
	a.NotError(Parse(semver, "2.3.19")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19)

	a.NotError(Parse(semver, "2.3.19+build.1")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19).
		Equal(semver.Build, "build.1")

	a.NotError(Parse(semver, "2.3.19-pre.release+build")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19).
		Equal(semver.PreRelease, "pre.release").
		Equal(semver.Build, "build")

	a.NotError(Parse(semver, "2.3.19-pre.release")).
		Equal(semver.Major, 2).
		Equal(semver.Minor, 3).
		Equal(semver.Patch, 19).
		Equal(semver.PreRelease, "pre.release")

	a.Error(Parse(semver, "2..1"))
}

func TestGetFields(t *testing.T) {
	a := assert.New(t, false)

	// 不可导出
	o1 := &struct {
		v1 int `version:"0"`
	}{}
	fields, err := getFields(o1)
	a.Error(err).Nil(fields)

	// 重复的索引值
	o2 := &struct {
		V1 int    `version:"0,.1"`
		V2 string `version:"0"`
	}{}
	fields, err = getFields(o2)
	a.Error(err).Nil(fields)

	// 路由项不存在
	o3 := &struct {
		V1 int    `version:"0,.3"`
		V2 string `version:"1,.2,+1,-0"`
		V3 string `version:"2"`
	}{}
	fields, err = getFields(o3)
	a.Error(err).Nil(fields)

	o4 := &struct {
		V2 string `version:"1,.2,+1,-0"`
		V1 int    `version:"0,.1"`
		V3 string `version:"2"`
	}{}
	fields, err = getFields(o4)
	a.NotError(err).
		Equal(fields[0].routes['.'], 1).
		Equal(fields[1].routes['+'], 1)
}
