package FileOp

import (
    "encoding/binary"
    "fmt"
    "os"
)

type ThreeIntegers struct {
    First  uint8
}

func ReadBinary(path string, readsize int)([]uint8) {

    file, err := os.Open(path)
    if err != nil {
        fmt.Println("error occured 'os.Open()'")
        panic(err)
	}
	
	memory := make([]uint8,readsize)

	var threeIntegers ThreeIntegers

	for i :=0 ;i<readsize;i++{
		errb := binary.Read(file, binary.BigEndian, &threeIntegers)
		if errb != nil {
			fmt.Println("error occured 'binary.Read()'")
			panic(errb)
		}
		//fmt.Printf("%d  : %02x\n",i, threeIntegers.First)
		memory[i] = threeIntegers.First
	}
	
	return memory

}