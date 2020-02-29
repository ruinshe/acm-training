github.com/ruinshe/acm-training

[![Build Status](https://travis-ci.com/ruinshe/acm-training.svg?branch=master)](https://travis-ci.com/ruinshe/acm-training)
[![codecov](https://codecov.io/gh/ruinshe/acm-training/branch/master/graph/badge.svg)](https://codecov.io/gh/ruinshe/acm-training)
[![Go Report Card](https://goreportcard.com/badge/github.com/ruinshe/acm-training)](https://goreportcard.com/report/github.com/ruinshe/acm-training)
![](https://img.shields.io/badge/version-0.0.8-blue.svg)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Join the chat at https://gitter.im//UESTC-ACM/acm-training](https://badges.gitter.im//UESTC-ACM/acm-training.svg)](https://gitter.im//UESTC-ACM/acm-training?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[http://acm.uestc.edu.cn](http://acm.uestc.edu.cn)

A web hosted system for the ACM/ICPC group based training maintenance.

This is a web writting in Go language for ACM/ICP daily grouped based training.

# Have a try

## Install

For stable channel, you can directly using `go get` to get the latest stable version of the application.

``` shell
go get -u github.com/ruinshe/acm-training
```

## Run the application

You can directly run the application using the binary named `acm-training` in the `$GOPATH/bin` folder.

``` shell
$GOPATH/bin/acm-training
```

If you need to run a daemon for the application, you can try adding the application as a service defined for [systemd](https://en.wikipedia.org/wiki/Systemd).

# Development and Contribute

## Basic tools used for this project

### `golint`

TD;DR using the command line below to perform go source code linting.

``` shell
make get lint
```

We suggest commit a change after perform lint against the go source code files. You can follow the guide in the [offical package home page](https://github.com/golang/lint).

## I want to contribute my code change, including bug fixing and new feature enhancement.
Please fllow the `CONTRIBUTING.md` file for the development process.
