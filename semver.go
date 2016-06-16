// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

import (
	"bytes"
	"strconv"
	"strings"
)

// http://semver.org/lang/zh-CN/
type SemVersion struct {
	Major      int    `version:"0,number,.1"`
	Minor      int    `version:"1,number,.2"`
	Patch      int    `version:"2,number,+4,-3"`
	PreRelease string `version:"3,string,+4"`
	Build      string `version:"4,string"`
}

// 比较两个版本号，若相同返回 0，若 v 比较大返回正整数，否则返回负数。
func (v *SemVersion) Compare(v2 *SemVersion) int {
	switch {
	case v.Major != v2.Major:
		return v.Major - v2.Major
	case v.Minor != v2.Minor:
		return v.Minor - v2.Minor
	case v.Patch != v2.Patch:
		return v.Patch - v2.Patch
	case len(v.PreRelease) == 0:
		if len(v2.PreRelease) > 0 {
			return 1
		}
		return 0
	case len(v2.PreRelease) == 0:
		if len(v.PreRelease) > 0 {
			return -1
		}
		return 0
	}

	// pre-release 的比较，按照 semver 的规则：
	// 透过由左到右的每个被句点分隔的标识符号来比较，直到找到一个差异值后决定：
	// 只有数字的标识符号以数值高低比较，有字母或连接号时则逐字以 ASCII 的排序来比较。
	// 数字的标识符号比非数字的标识符号优先层级低。若开头的标识符号都相同时，
	// 栏位比较多的先行版本号优先层级比较高。范例：1.0.0-alpha < 1.0.0-alpha.1
	// < 1.0.0-alpha.beta < 1.0.0-beta < 1.0.0-beta.2 < 1.0.0-beta.11 < 1.0.0- rc.1 < 1.0.0。

	vReleases := strings.Split(v.PreRelease, ".")
	v2Releases := strings.Split(v2.PreRelease, ".")
	l := len(vReleases)
	if len(v2Releases) < l {
		l = len(v2Releases)
	}
	for i := 0; i < l; i++ {
		val := strings.Compare(vReleases[i], v2Releases[i])
		if val == 0 { // 当前的所有字符相同，继续比较下个字符串
			continue
		}

		// 尝试数值转换，若其中一个不能转换成数值，则直接返回 val
		v1, err := strconv.Atoi(vReleases[i])
		if err != nil {
			return val
		}
		v2, err := strconv.Atoi(v2Releases[i])
		if err != nil {
			return val
		}

		// 到这里，v1 不可能等于 v2，只要判断两者的差值即可。
		return v1 - v2
	}

	return 0
}

func (v *SemVersion) String() string {
	buf := bytes.NewBufferString(strconv.Itoa(v.Major))
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(v.Minor))
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(v.Patch))

	if len(v.PreRelease) > 0 {
		buf.WriteByte('-')
		buf.WriteString(v.PreRelease)
	}
	if len(v.Build) > 0 {
		buf.WriteByte('+')
		buf.WriteString(v.Build)
	}

	return buf.String()
}

func SemVer(ver string) (*SemVersion, error) {
	semver := &SemVersion{}

	if err := Parse(semver, ver); err != nil {
		return nil, err
	}
	return semver, nil
}

// 比较两个 semver 版本号
func SemVerCompare(ver1, ver2 string) (int, error) {
	v1, err := SemVer(ver1)
	if err != nil {
		return 0, err
	}

	v2, err := SemVer(ver2)
	if err != nil {
		return 0, err
	}

	return v1.Compare(v2), nil
}
