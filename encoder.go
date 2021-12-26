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
	
	fmt.Println("Enter Encode")
	
	buff, err := structex.EncodeByteBuffer(ie)
		
	if err != nil {
		panic("a problem")
		}
	
	fmt.Println("Encoded array",buff)
	
	n,_:=buffer.Write(buff)
	fmt.Println("Number of Bytes :", n)
	
	
	fmt.Printf("Bytestrom=")
	for i := 0; i < n; i++ {
		fmt.Printf("0x%02X", buff[i])
		if i != n-1 {
			fmt.Printf(",")
		}
	}
	fmt.Println()
	
	fmt.Println("Exit Encode")
}




func main() {

	fmt.Println("Enter Main")
	
	ie := Nas5GSUpdateType{ IEI:1, Length:2, SMS_requested:1, NG_RAN_RCU:1, F5GS_PNB_CIoT:0, EPS_PNB_CIoT:0, spare0: 0, spare1:0}
	
	buf := new(bytes.Buffer)
	
	ie.Encode(buf)

	fmt.Println("Exit Main")

}
