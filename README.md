<h1 align="center">gotransformer</h1>

<p align="center">
    <a href="https://travis-ci.org/snowlyg/gotransformer"><img src="https://travis-ci.org/snowlyg/gotransformer.svg?branch=master" alt="Build Status"></a>
    <a href="https://codecov.io/gh/snowlyg/gotransformer"><img src="https://codecov.io/gh/snowlyg/gotransformer/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/gotransformer"><img src="https://goreportcard.com/badge/github.com/snowlyg/gotransformer" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/gotransformer"><img src="https://godoc.org/github.com/snowlyg/gotransformer?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/gotransformer/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/gotransformer" alt="Licenses"></a>
</p>

>本项目主要是用于数据的格式化和转换。
>使用场景：当模型数据字段太多，同时很多字段数据在接口输出的时候又需要做特殊处理（比如时间，数组，json等等）。
>使用场景：不同接口或者不同角色可能需要做不同的数据处理(比如列表，详情等等)。

#### 更新日志
[更新日志](UPDATE.MD)

#### Require
- go 1.13.x

#### 支持格式化方式
- [简单格式化](_example/struct/struct.go)
- [自定义方法格式化](_example/struct/struct.go)
- [关联数据格式化](_example/struct/struct.go)
- [时间数据格式化](_example/struct/struct.go)
- [map数据格式化](_example/map/map.go)
- [excel 导入数据格式化](_example/excel/excel.go)


- 具体项目使用实例，请参考[https://github.com/snowlyg/IrisAdminApi](https://github.com/snowlyg/IrisAdminApi)

## Installation

```
go get  github.com/snowlyg/gotransformer@master
```
