version
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fissue9%2Fversion%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/issue9/version/goto?ref=master)
[![Build Status](https://travis-ci.org/issue9/version.svg?branch=master)](https://travis-ci.org/issue9/version)
[![codecov](https://codecov.io/gh/issue9/version/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/version)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
======

通过定义 struct tag 的相关属性，可以解析大部份版本号字符串到一个结构体中。

```go
type Version struct {
    Major int    `version:"0,number,.1"`
    Minor int    `version:"1,number,+2"`
    Build string `version:"2,string"`
}

ver := &Version{}
version.Parse(ver, "2.1+160616")
// 解析之后
// ver.Major == 2, ver.Minor == 1, ver.Build == 160616
```

同时也定义了一个 [semver](http://semver.org) 的一个内部实现。

```go
semver,err := version.SemVer("2.10.1+build")
if err != nil{
    // TODO
}

fmt.Println(semver)
// semver.Major == 2
// semver.Minor == 10
// semver.Patch == 1
// semver.Build == build
```

安装
----

```shell
go get github.com/issue9/version
```

文档
----

[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/issue9/version)
[![GoDoc](https://godoc.org/github.com/issue9/version?status.svg)](https://godoc.org/github.com/issue9/version)

版权
----

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
