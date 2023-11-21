version
[![Test](https://github.com/issue9/version/actions/workflows/go.yml/badge.svg)](https://github.com/issue9/version/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/issue9/version/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/version)
![License](https://img.shields.io/github/license/issue9/version)
[![Go version](https://img.shields.io/github/go-mod/go-version/issue9/version)](https://golang.org)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/issue9/version)](https://pkg.go.dev/github.com/issue9/version)
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

版权
----

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
