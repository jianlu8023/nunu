# Nunu — A CLI tool for building Go applications.

## 添加功能

nunu clone: 可以实现将仓库进行clone 并删除 .git 文件夹

### 使用场景

* 你的项目目录结构是这样

```text
app-modules
├── app    项目源码
└── tools  所使用依赖 
```

### 如何使用

```shell
# 创建module 
mkdir app-modules
# 进入目录
cd app-modules
# clone 依赖
nunu clone <repo> -b <branch>
# 创建app module
nunu create app -r <repo>
# 运行程序
cd app  && nunu run 
```

## 安装此nunu

```shell
go install gitee.com/jianlu8023/nunu@latest
```

* [Nunu中文介绍](https://github.com/go-nunu/nunu/blob/main/README_zh.md)

