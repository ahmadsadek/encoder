package main

import (
	"bytes"
	"fmt"
	"github.com/HewlettPackard/structex"
)


type Nas5GSUpdateType struct {
	IEI           uint8 // Octet 1
	Length        uint8 // Octet 2
	SMS_requested uint8 `bitfield:"1"`// Octet 3
	NG_RAN_RCU    uint8 `bitfield:"1"`
	F5GS_PNB_CIoT uint8 `bitfield:"2"`
	EPS_PNB_CIoT  uint8 `bitfield:"2"` 
	spare0        uint8 `bitfield:"1",reserved`
	spare1        uint8 `bitfield:"1",reserved`
}


func (ie Nas5GSUpdateType) Encode(buffer *bytes.Buffer) {

	
	return 
}



func main() {
	fmt.Println("Hello GO")
	
	ie := Nas5GSUpdateType{ IEI:1, Length:2, EPS_PNB_CIoT:0, F5GS_PNB_CIoT:0, NG_RAN_RCU:1, SMS_requested:1, spare0: 0, spare1:0}
		
				
	buff, err := structex.EncodeByteBuffer(ie)
		
	if err != nil {
		panic("a problem")
		}

	fmt.Println(buff)

}
