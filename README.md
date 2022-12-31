# Leetgo

[![Go Report Card](https://goreportcard.com/badge/github.com/j178/leetgo)](https://goreportcard.com/report/github.com/j178/leetgo)
[![CI](https://github.com/j178/leetgo/actions/workflows/ci.yaml/badge.svg)](https://github.com/j178/leetgo/actions/workflows/ci.yaml)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://makeapullrequest.com)
[![Twitter Follow](https://img.shields.io/twitter/follow/niceoe)](https://twitter.com/niceoe)

*This project is still in its early development stage and many features have not yet been implemented. Everything is subject to rapid change.*

[中文](./README_zh.md) | English

`leetgo` is a command-line tool for LeetCode that provides almost all the functionality of LeetCode, 
allowing you to complete all of your LeetCode exercises without leaving the terminal. 
It can automatically generate **skeleton code and test cases**, support **local testing and debugging**, 
and you can use any IDE you like to solve problems. 

And `leetgo` also supports real-time generation of **contest questions**, so your submissions are one step ahead!

[![asciicast](https://asciinema.org/a/0sUG7psmMfgWqzy9rr57hrcnX.svg)](https://asciinema.org/a/0sUG7psmMfgWqzy9rr57hrcnX)


## Highlight of features

- Search for and view a question by its ID or slug.
- Generate skeleton code and testing code for a question.
- Run test cases on your local machine.
- Generate contest questions just in time.

## Language support

Currently, `leetgo` supports generating code and local test for the following languages:
<!-- BEGIN MATRIX -->
|  | Generate | Local Test |
| --- | --- | --- |
| Go | :white_check_mark: | :x: |
| Python | :white_check_mark: | :x: |
| C++ | :white_check_mark: | :x: |
| Rust | :white_check_mark: | :x: |
| Java | :white_check_mark: | :x: |
| JavaScript | :white_check_mark: | :x: |
| PHP | :white_check_mark: | :x: |
| C | :white_check_mark: | :x: |
| C# | :white_check_mark: | :x: |
| Ruby | :white_check_mark: | :x: |
| Swift | :white_check_mark: | :x: |
| Kotlin | :white_check_mark: | :x: |
<!-- END MATRIX -->
and many other languages are in plan. (help wanted, contributions welcome!)

## Installation

You can download the latest binary from the [release page](https://github.com/j178/leetgo/releases).

### Install via `go install`
 
```shell
go install github.com/j178/leetgo@latest
```

### Install via `brew install`

```shell
brew install j178/tap/leetgo
```

## Usage
<!-- BEGIN USAGE -->
```
Usage:
  leetgo [command]

Available Commands:
  init                    Init a leetcode workspace
  pick                    Generate a new question
  today                   Generate the question of today
  info                    Show question info
  test                    Run question test cases
  submit                  Submit solution
  contest                 Generate contest questions
  cache                   Manage local questions cache
  config                  Show leetgo config dir

Flags:
  -v, --version      version for leetgo
  -g, --gen string   language to generate: cpp, go, python ...

Use "leetgo [command] --help" for more information about a command.
```
<!-- END USAGE -->

## LeetCode Support

Leetgo uses LeetCode's GraphQL API to get questions and submit solutions. You need to provide your LeetCode session ID to authenticate.

Currently only `leetcode.cn` is supported. Support for `leetcode.com` is under development.

## Configuration

Leetgo uses two levels of configuration files, the global configuration file located at `~/.config/leetgo/config.yaml` and the local configuration file located at `leetgo.yaml` in the current directory. 
These configuration files are generated during the `leetgo init` process. 
The local configuration file in the project will override the global configuration. 
It is generally recommended to use the global configuration as the default configuration and customize it in the project by modifying the leetgo.yaml file.

<!-- BEGIN CONFIG -->
```yaml
# Your name
author: Bob
# Generate code for questions, go, python, ... (will be override by project config and flag --gen)
gen: go
# Language of the questions, zh or en
language: zh
# LeetCode configuration
leetcode:
  # LeetCode site, https://leetcode.com or https://leetcode.cn
  site: https://leetcode.cn
editor: {}
# Cache type, json or sqlite
cache: json
go:
  # Generate separate package for each question
  separate_package: true
  # Filename template for Go files
  filename_template: ""
python:
  out_dir: python
cpp:
  out_dir: cpp
java:
  out_dir: java
rust:
  out_dir: rust
c:
  out_dir: c
csharp:
  out_dir: csharp
javascript:
  out_dir: javascript
ruby:
  out_dir: ruby
swift:
  out_dir: swift
kotlin:
  out_dir: kotlin
php:
  out_dir: php
```
<!-- END CONFIG -->

## Troubleshooting

If you encounter any problems, please run your command with `DEBUG` environment variable set to `1`, copy the command output and open an issue.

## Contributions welcome!

[Good first issues](https://github.com/j178/leetgo/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) are a great place to start, 
and you can also check out some [help wanted](https://github.com/j178/leetgo/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) issues.

## Credits

- https://github.com/EndlessCheng/codeforces-go
- https://github.com/clearloop/leetcode-cli
- https://github.com/budougumi0617/leetgode
- https://github.com/skygragon/leetcode-cli
