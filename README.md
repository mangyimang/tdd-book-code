这是一个[单体仓库](https://trunkbaseddevelopment.com/monorepos/)，包含了 O'Reilly 出版的["测试驱动开发学习指南"](https://learning.oreilly.com/library/view/learning-test-driven-development/9781098106461/)一书的所有源代码。

# 前置要求

你需要安装 [Go](https://golang.org/)、[Node.js](https://nodejs.org/en/) 和 [Python 3](https://www.python.org/) 的运行环境来运行此仓库中的代码。

除此之外，你可能还需要其他工具 —— 例如 ["act"](https://github.com/nektos/act) 用于在本地运行 GitHub actions、[gocyclo](https://github.com/fzipp/gocyclo) 用于检查 Go 代码的圈复杂度、[jshint](https://jshint.com) 用于 JavaScript 代码的静态分析，以及 [flake8](https://flake8.pycqa.org) 用于 Python 代码的静态分析。

要理解代码的演进过程和目的，你需要配套的书籍。

# 如何运行测试
简而言之，使用以下命令来运行各语言版本的测试。

## Go
```
cd go
go test -v ./...
```

## JavaScript
```
node js/test_money.js
```

## Python
```
python3 py/test_money.py
```

# 如何获取这本书

这本书可以在 [亚马逊](https://www.amazon.com/Learning-Test-Driven-Development-Polyglot-Uncluttered/dp/1098106474/ref=sr_1_3?) 和 [O'Reilly](https://learning.oreilly.com/library/view/learning-test-driven-development/9781098106461/) 购买。