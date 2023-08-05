package main

import (
	"github.com/asssaf/stemma-soil-go/soil"
	"log"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func main() {
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	devAddrs := []int{0x36, 0x37, 0x39}
	i2cPort, err := i2creg.Open("/dev/i2c-1")
	if err != nil {
		log.Panicf("could not open i2c device: %s", err)
	}

	var moistureSensors []*soil.Dev
	for _, devAddr := range devAddrs {
		opts := soil.DefaultOpts
		opts.Addr = uint16(devAddr)
		dev, err := soil.NewI2C(i2cPort, &opts)
		if err != nil {
			panic(err)
		}
		moistureSensors = append(moistureSensors, dev)
	}

	for i, dev := range moistureSensors {
		values := soil.SensorValues{} // Create the sensor value struct we will populate in the next line
		if err := dev.Sense(&values); err != nil {
			log.Printf("Sensor: %d unable to sense moisture: %s", i+1, err)
			continue
		}
		floatCap := float32(values.Capacitance) // Turn the Capacitance into a float that we can report with float formatting
		log.Printf("Sensor: %d temperature: %0.2f°C capacitance : %0.2f\n", i+1, values.Temperature.Celsius(), floatCap)
	}

}
