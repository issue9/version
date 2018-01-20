// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

import (
	"bytes"
	"strconv"
	"strings"
)

// SemVersion 是 semver 的定义，
// 具体可参考：https://semver.org/lang/zh-CN/
type SemVersion struct {
	Major      int    `version:"0,.1"`
	Minor      int    `version:"1,.2"`
	Patch      int    `version:"2,+4,-3"`
	PreRelease string `version:"3,+4"`
	Build      string `version:"4"`
}

// Compare 比较两个版本号，若相同返回 0，若 v 比较大返回正整数，否则返回负数。
func (v *SemVersion) Compare(v2 *SemVersion) int {
	switch {
	case v.Major != v2.Major:
		return v.Major - v2.Major
	case v.Minor != v2.Minor:
		return v.Minor - v2.Minor
	case v.Patch != v2.Patch:
		return v.Patch - v2.Patch
	case len(v.PreRelease) == 0:
		return len(v2.PreRelease) // v2.PreRelease 只要有内容，就比 v 的版本号小
	case len(v2.PreRelease) == 0:
		return -len(v.PreRelease) // v.PreRelease 只要有内容，就比 v2 的版本号小
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

// CompareString 将当前对象与一个版本号字符串相比较。其返回值的功能与 Compare 相同
func (v *SemVersion) CompareString(ver string) (int, error) {
	v2, err := SemVer(ver)
	if err != nil {
		return 0, err
	}

	return v.Compare(v2), nil
}

// Compatible 当前对象与 v2 是否兼容。
// semver 规定主版本号相同的，在 API 层面必须兼容。
func (v *SemVersion) Compatible(v2 *SemVersion) bool {
	return v.Major == v2.Major
}

// CompatibleString 当前对象与版本号字符串是否兼容。
// semver 规定主版本号相同的，在 API 层面必须兼容。
func (v *SemVersion) CompatibleString(ver string) (bool, error) {
	v2, err := SemVer(ver)
	if err != nil {
		return false, err
	}

	return v.Compatible(v2), nil
}

// String 转换成版本号字符串
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

// SemVer 将一个版本号字符串解析成 SemVersion 对象
func SemVer(ver string) (*SemVersion, error) {
	semver := &SemVersion{}

	if err := Parse(semver, ver); err != nil {
		return nil, err
	}
	return semver, nil
}

// SemVerCompare 比较两个 semver 版本号字符串
func SemVerCompare(ver1, ver2 string) (int, error) {
	v1, err := SemVer(ver1)
	if err != nil {
		return 0, err
	}

	return v1.CompareString(ver2)
}

// SemVerCompatible 两个 semver 版本号是否兼容
func SemVerCompatible(ver1, ver2 string) (bool, error) {
	v1, err := SemVer(ver1)
	if err != nil {
		return false, err
	}

	return v1.CompatibleString(ver2)
}

// SemVerValid 验证 semver 版本号是否符合 semver 规范。
func SemVerValid(ver string) bool {
	v := &SemVersion{
		Major: -1,
		Minor: -1,
		Patch: -1,
	}

	if err := Parse(v, ver); err != nil {
		return false
	}

	if v.Major == -1 || v.Minor == -1 || v.Patch == -1 {
		return false
	}

	return true
}
