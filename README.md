version [![Build Status](https://travis-ci.org/issue9/version.svg?branch=master)](https://travis-ci.org/issue9/version)
======

通过定义 struct tag 的相头属性，可以解析大部份版本号字符串到一个结构体中。

```go
type struct Version {
    Major int    `version:"0,number,.1"`
    Minor int    `version:"1,number,+2"`
    Build string `version:"2,string"`
}

ver := &Version{}
version.Parse(ver, "2.1+160616")
// 解析之后
// ver.Major == 2, ver.Minor == 1, ver.Build == 160616
```


### 安装

```shell
go get github.com/issue9/version
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/version)
[![GoDoc](https://godoc.org/github.com/issue9/version?status.svg)](https://godoc.org/github.com/issue9/version)


### 版权

本项目采用 [MIT](http://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
