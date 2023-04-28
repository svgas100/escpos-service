package usb

import (
	"fmt"

	"github.com/karalabe/usb"
)

func TestUsb() {
	if usb.Supported() {
		fmt.Println("supported!")
	} else {
		fmt.Println("unsupported!")
		return
	}

	devices, err := usb.Enumerate(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(devices) == 0 {
		fmt.Println("No device found!")
		return
	}

	device, err := devices[0].Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	device.Write([]byte("Test USB Connection ;)"))
}
