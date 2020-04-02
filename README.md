# DvLIR-Check-Plugin

## Description
Monitoring check plugin for the [dvlir-restapi-go-client](https://github.com/inexio/dvlir-restapi-go-client). The plugin complies with the [Monitoring Plugins Development Guidelines](https://www.monitoring-plugins.org/doc/guidelines.html) and should therefore be compatible with [nagios](https://www.nagios.org/), [icinga2](https://icinga.com/), [zabbix](https://www.zabbix.com/), [checkmk](https://checkmk.com/), etc.

## Usage
```
  Usage
   main [OPTIONS]
   
   Application Options:
    -i, --ipAddress= The username of the dvlir adapter
    -p, --password= The password of the dvlir adapter
    -P, --port=     The port of the dvlir adapter, default is 80
```

## Installation

To install, use `go get` or `git clone`:

      go get https://github.com/inexio/DvLIR-Check-Plugin
      
or 

      git clone https://github.com/inexio/DvLIR-Check-Plugin.git
