package FileOp

import (
    "encoding/binary"
    "fmt"
    "os"
)

type ThreeIntegers struct {
    First  uint8
}

func ReadBinary(path string, readsize int,memory []uint8,baseaddr uint32)([]uint8) {

    file, err := os.Open(path)
    if err != nil {
        fmt.Println("error occured 'os.Open()'")
        panic(err)
    }

	var threeIntegers ThreeIntegers

	for i :=0 ;i<40;i++{
		errb := binary.Read(file, binary.BigEndian, &threeIntegers)
		if errb != nil {
			fmt.Println("error occured 'binary.Read()'")
			panic(errb)
		}
		//fmt.Printf("%d  : %02x\n",i, threeIntegers.First)
		memory[baseaddr+uint32(i)] = threeIntegers.First
	}
	
	return memory

}