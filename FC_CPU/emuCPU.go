package main

import (
	"fmt"
	"os"
)

const (
	memCap = 65535
)

type emu struct {
	// A X Y S P
	regi map[string]uint8
	// PC
	regPc uint16
	// memory
	memory [memCap]uint8
}

func debug(degemu *emu) {
	var register = []string{"A", "X", "Y", "S"}
	for _, value := range register {
		fmt.Print("regi", value, " = ")
		fmt.Printf("0x%x\n", degemu.regi[value])
	}
	fmt.Printf("regiP = 0b%08b\n",degemu.regi["P"])
	fmt.Printf("PC = 0x%x\n", degemu.regPc)
	fmt.Printf("opcd  = 0x%x\n",degemu.memory[degemu.regPc])
	fmt.Printf("memory [0] = %d\n",degemu.memory[0])
	fmt.Printf("memory [0x208] = %d\n",degemu.memory[0x208])



}

func initReg(fcEmu *emu) {
	fcEmu.regi = make(map[string]uint8)
	fcEmu.regi["A"] = 4
	fcEmu.regi["X"] = 4
	fcEmu.regi["Y"] = 6
	fcEmu.regi["S"] = 0xff
	fcEmu.regi["P"] = 0b00000000
}

func initPc(fcEmu *emu,resetaddr uint16) {
	fcEmu.regPc = resetaddr
}

func checkNes(checkbuf []uint8) (uint8, uint8) {
	if checkbuf[0] != 0x4e {
		panic("not nes file : not 0x4e")
	}
	if checkbuf[1] != 0x45 {
		panic("not nes file : not 0x45")
	}
	if checkbuf[2] != 0x53 {
		panic("not nes file : not 0x53")
	}
	if checkbuf[3] != 0x1a {
		panic("not nes file : not 0x1a")
	}
	progSize := checkbuf[4]
	chrSize := checkbuf[5]
	return progSize, chrSize
}

func readFile(fcEmu *emu) (uint8, []uint8) {
	var irqaddr, nmiaddr, resetaddr uint16 = 0, 0, 0
	bufregpc := 0x8000
	file, err := os.Open(`C:\Users\ttnmr\HOME\emulator\FC\amegure.nes`)

	if err != nil {
		panic("file error")
	}
	

	checkbuf := make([]uint8, 16)
	_, checkerr := file.Read(checkbuf)
	if checkerr != nil {
		// Readエラー処理
		panic("error read : checkbuf")

	}
	defer file.Close()
	// progSize 1 == 16384
	progSize, chrSize := checkNes(checkbuf)
	// この数で配列を操作するので一引かないとずれちゃう
	progcounter := int(progSize) * 16384
	progcounter--
	chrcounter := int(chrSize) * 8192
	chrrombuf := make([]uint8, chrcounter)
	chrcounter--
	writebuf := make([]uint8, 1)
	for progcounter >= 0 {
		n, writeerr := file.Read(writebuf)
		if n == 0 {
			break
		}
		if writeerr != nil {
			// Readエラー処理
			fmt.Println("error read : writebuf")
			break
		}
		// IRQ
		if progcounter==0||progcounter==1{
			irqaddr = irqaddr >> 8
			buf := uint16(writebuf[0])
			buf = buf << 8
			irqaddr += buf
		}
		// RESET
		if progcounter==2||progcounter==3{
			resetaddr = resetaddr >> 8
			buf := uint16(writebuf[0])
			buf = buf << 8
			resetaddr += buf
		}
		// NMI
		if progcounter==4||progcounter==5{
			nmiaddr = nmiaddr >> 8
			buf := uint16(writebuf[0])
			buf = buf << 8
			nmiaddr += buf
		}
		progcounter--
		// nes file の pro 部分を 0x8000 から 16384*progSize 分読み込む
		fcEmu.memory[bufregpc] = writebuf[0]
		//fmt.Printf("rom: 0x%x\n",fcEmu.memory[bufregpc])
		bufregpc++
	}
	
	bufcounter := 0
	for chrcounter >= 0 {
		n, writeerr := file.Read(writebuf)
		if n == 0 {
			break
		}
		if writeerr != nil {
			// Readエラー処理
			fmt.Println("error read : chrrombuf")
			break
		}
		chrcounter--
		chrrombuf[bufcounter] = writebuf[0]
		bufcounter++		
	}

	initPc(fcEmu, resetaddr)
	return progSize, chrrombuf
}

func initMem(fcEmu *emu) (uint8, []uint8) {
	progSize, chrrombuf := readFile(fcEmu)
	
	/*
	fcEmu.memory[0] = 5
	
	fcEmu.memory[fcEmu.regPc] = 0xa9
	fcEmu.memory[fcEmu.regPc+1] = 0x56

	fcEmu.memory[fcEmu.regPc+2] = 0x48

	fcEmu.memory[fcEmu.regPc+3] = 0xa9
	fcEmu.memory[fcEmu.regPc+4] = 0x07

	fcEmu.memory[fcEmu.regPc+5] = 0x48

	fcEmu.memory[fcEmu.regPc+6] = 0x18

	fcEmu.memory[fcEmu.regPc+7] = 0xa9
	fcEmu.memory[fcEmu.regPc+8] = 0x00

	fcEmu.memory[fcEmu.regPc+9] = 0x68

	fcEmu.memory[fcEmu.regPc+10] = 0x85
	fcEmu.memory[fcEmu.regPc+11] = 0x00

	fcEmu.memory[fcEmu.regPc+12] = 0x68

	fcEmu.memory[fcEmu.regPc+13] = 0x65
	fcEmu.memory[fcEmu.regPc+14] = 0x00
	

	return uint8(1), make([]uint8,0)
	*/
	
	return progSize, chrrombuf
}

func initEmu(fcEmu *emu) (uint8, []uint8) {
	initReg(fcEmu)
	fcEmu.regPc = 0x8000
	progSize, chrrombuf := initMem(fcEmu)
	return progSize, chrrombuf
}

func main() {
	var fcEmu = emu{}
	var continueflag bool = false 
	var index string
	breakpoint := make([]uint16, 0,50)
	//breakpoint = append(breakpoint,0x834f)
	progSize, _ := initEmu(&fcEmu)
	fmt.Printf("progsize: %d\n",progSize)
	
	for fcEmu.regPc < 0xffff {
		for _, value := range breakpoint{
			if value == fcEmu.regPc{
				continueflag = true
			}
		}
		if continueflag == true{
			debug(&fcEmu)
			fmt.Print(" >>")
			fmt.Scan(&index)
			if index == "c"{
				continueflag = true
			}else{
				continueflag = false
			}
		}

		fcEmu.Execute()
		fmt.Printf("main : 0x%x\n", fcEmu.regPc)
	}
	
	/*
	for i:=0 ; i<10 ;i++{
		fcEmu.Execute()
		debug(&fcEmu)
		fmt.Print("\n")
	}
	*/
	//fmt.Printf("0x%x\n", fcEmu.memory[0x8099])
}
