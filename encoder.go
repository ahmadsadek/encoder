package main

import (
	"bytes"
	"fmt"

	"github.com/HewlettPackard/structex"
)

type Nas5GSUpdateType struct {
	IEI           uint8                 // Octet 1
	Length        uint8                 // Octet 2
	EPS_PNB_CIoT  uint8 `bitfield:"1"`  // Octet 3
	_5GS_PNB_CIoT uint8 `bitfield:"1"`
	NG_RAN_RCU    uint8 `bitfield:"2"`
	SMS_requested uint8 `bitfield:"2"`
	spare0        uint8 `bitfield:"1",spare`
	spare1        uint8 `bitfield:"1",spare`
}

func (ie Nas5GSUpdateType) Encode(buffer *bytes.Buffer) {

	return
}

func main() {
	fmt.Println("Hello GO!")

	var inquiry = new(Nas5GSUpdateType)

	// Perform IOCTL on device, returns byte reader to byte response

	if err := structex.Decode(byteReader, inquiry); err != nil {
		panic("a problem")
	}

}
