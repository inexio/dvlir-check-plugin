package main

import (
	"github.com/inexio/dvlir-restapi-go-client"
	"github.com/inexio/go-monitoringplugin"
	"github.com/jessevdk/go-flags"
	"os"
	"strconv"
	"strings"
)

var opts struct {
	IPAddress string `short:"i" long:"ipAddress" description:"The ip-address of the dvlir" required:"true"`
	Password  string `short:"p" long:"password" description:"The password of the dvlir" required:"true"`
	Port      string `short:"P" long:"port" description:"The port of the dvlir, must be 80" required:"false"`
}

func main() {
	var err error
	_, err = flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(3)
	}

	response := monitoringplugin.NewResponse("checked")
	defer response.OutputAndExit()

	client, err := dvlirclient.NewDvLIRClient(opts.IPAddress, opts.Password)
	if err != nil {
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not create new client: "+err.Error())
		return
	}

	err = client.Login()
	if err != nil {
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not log into DvLIR: "+err.Error())
		return
	}
	defer func() {
		err = client.Logout()
		if err != nil {
			response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not log out of DvLIR: "+err.Error())
			return
		}
	}()

	file, err := client.GetDataFile(10)
	if err != nil {
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not retrieve data file: "+err.Error())
		return
	}

	value, err := client.GetMomentaryValues()
	if err != nil {
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not retrieve momentary values: "+err.Error())
		return
	}

	counter := strings.Replace(file[1].OneEightZero, ",", ".", 1)
	FlCounter, err := strconv.ParseFloat(counter, 64)
	if err != nil {
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not convert meter reading from string to float64: "+err.Error())
		return
	}

	err = response.AddPerformanceDataPoint(monitoringplugin.NewPerformanceDataPoint("Meter_Reading", FlCounter, ""))
	if err != nil {
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Could not add new performance data point: "+err.Error())
		return
	}

	switch value.Status {
	case "OK":
		response.UpdateStatus(monitoringplugin.OK, "Check completed. \nSerial number electric meter: "+file[1].MeterNumber+",\nSerial number network adapter: "+file[1].DvLIRSn+",\nMeter reading: "+file[1].OneEightZero)
	case "NTP Fehler":
		response.UpdateStatus(monitoringplugin.CRITICAL, "Check failed: "+value.Status)
	case "IR Fehler":
		response.UpdateStatus(monitoringplugin.CRITICAL, "Check failed: "+value.Status)
	case "DHCP Fehler":
		response.UpdateStatus(monitoringplugin.CRITICAL, "Check failed: "+value.Status)
	default:
		response.UpdateStatus(monitoringplugin.UNKNOWN, "Unknown status: "+value.Status)
	}

}
