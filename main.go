package main

import (
	"fmt"
	"github.com/asssaf/stemma-soil-go/soil"
	"log"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func main() {
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	devAddr := 0x36
	i2cPort, err := i2creg.Open("/dev/i2c-1")
	if err != nil {
		log.Panicf("could not open i2c device: %s", err)
	}

	opts := soil.DefaultOpts
	if devAddr != 0 {
		if devAddr < 0x36 || devAddr > 0x39 {
			panic(fmt.Sprintf("given address not supported by device: %x", devAddr))
		}
		opts.Addr = uint16(devAddr)
	}

	dev, err := soil.NewI2C(i2cPort, &opts)
	if err != nil {
		panic(err)
	}
	defer dev.Halt()

	values := soil.SensorValues{}
	if err := dev.Sense(&values); err != nil {
		log.Fatalf("unable to sense moisture: %s", err)
	}
	floatCap := float32(values.Capacitance)
	log.Printf("temperature: %0.2f°C capacitance : %0.2f\n", values.Temperature.Celsius(), floatCap)
}
