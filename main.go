package main

import (
	"fmt"
	"github.com/asssaf/stemma-soil-go/soil"
	"log"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func main() {
	if _, err := host.Init(); err != nil { // We need to load drivers to work with i2c stuff
		panic(err)
	}

	devAddr := 0x36                           // This is the register that we found when we ran i2cscan
	i2cPort, err := i2creg.Open("/dev/i2c-1") // This is the filepath that our device is mounted to
	if err != nil {
		log.Panicf("could not open i2c device: %s", err)
	}

	opts := soil.DefaultOpts // This instantiates default settings for the soil library, the default is already 0x36, but
	// could change based on if we want to hook up multiple of these, we would need to jump the sensor,
	// and update this code
	if devAddr != 0 {
		if devAddr < 0x36 || devAddr > 0x39 {
			panic(fmt.Sprintf("given address not supported by device: %x", devAddr))
		}
		opts.Addr = uint16(devAddr)
	}

	dev, err := soil.NewI2C(i2cPort, &opts) // instantiate our soil libraries i2c connection
	if err != nil {
		panic(err)
	}
	defer dev.Halt() // Close out the connection we create when this function finishes

	values := soil.SensorValues{} // Create the sensor value struct we will populate in the next line
	if err := dev.Sense(&values); err != nil {
		log.Fatalf("unable to sense moisture: %s", err)
	}
	floatCap := float32(values.Capacitance) // Turn the Capacitance into a float that we can report with float formatting
	log.Printf("temperature: %0.2f°C capacitance : %0.2f\n", values.Temperature.Celsius(), floatCap)
}
