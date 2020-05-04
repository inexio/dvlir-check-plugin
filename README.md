# DvLIR-Check-Plugin

[![Go Report Card](https://goreportcard.com/badge/github.com/inexio/dvlir-check-plugin)](https://goreportcard.com/report/github.com/inexio/dvlir-check-plugin)
[![GitHub license](https://img.shields.io/badge/license-BSD-blue.svg)](https://github.com/inexio/dvlir-check-plugin/blob/master/LICENSE)
[![GitHub code style](https://img.shields.io/badge/code%20style-uber--go-brightgreen)](https://github.com/uber-go/guide/blob/master/style.md)

## Description
Monitoring check plugin for the [dvlir-restapi-go-client](https://github.com/inexio/dvlir-restapi-go-client). The plugin complies with the [Monitoring Plugins Development Guidelines](https://www.monitoring-plugins.org/doc/guidelines.html) and should therefore be compatible with [nagios](https://www.nagios.org/), [icinga2](https://icinga.com/), [zabbix](https://www.zabbix.com/), [checkmk](https://checkmk.com/), etc.

## Usage
```
  Usage
   main [OPTIONS]
   
   Application Options:
    -i, --ipAddress= The IP address of the dvlir adapter
    -p, --password= The password of the dvlir adapter
    -P, --port=     The port of the dvlir adapter, default is 80
```

## Installation

To install, use `go get` or `git clone`:

      go get https://github.com/inexio/DvLIR-Check-Plugin
      
or 

      git clone https://github.com/inexio/DvLIR-Check-Plugin.git

## Getting Help

If there are any problems or something does not work as intended, open an issue on GitHub.

## Contribution

Contributions to the project are welcome.

We are looking forward to your bug reports, suggestions and fixes.

If you want to make any contributions make sure your go report does match up with our projects score of **A+**.

When you contribute make sure your code is conform to the **uber-go** coding style.

Happy Coding!

