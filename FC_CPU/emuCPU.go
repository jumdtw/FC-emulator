package FC_CPU

import (
	"fmt"
	"os"
)

const (
	memCap = 65535
)

type CpuEmu struct {

	//各割り込みを行うための情報
	Irqaddr uint16
	Nmiaddr uint16
	Resetaddr uint16
	InterruptFlag bool

	// vram　へ書き込みを行うための内部情報 
	// vram addr
	VramAddr uint16
	// vram write flag
	VramWriteFlag bool
	// vram write value
	VramWriteValue uint8


	// A X Y S P
	Regi map[string]uint8
	// PC
	RegPc uint16
	// Memory
	Memory [memCap]uint8
}

func debug(degemu *CpuEmu) {
	var Register = []string{"A", "X", "Y", "S"}
	for _, value := range Register {
		fmt.Print("Regi", value, " = ")
		fmt.Printf("0x%x\n", degemu.Regi[value])
	}
	fmt.Printf("RegiP = 0b%08b\n",degemu.Regi["P"])
	fmt.Printf("PC = 0x%x\n", degemu.RegPc)
	fmt.Printf("opcd  = 0x%x\n",degemu.Memory[degemu.RegPc])
	fmt.Printf("Memory [0] = %d\n",degemu.Memory[0])
	fmt.Printf("Memory [0x208] = %d\n",degemu.Memory[0x208])
}

func initReg(fcEmu *CpuEmu) {
	fcEmu.Regi = make(map[string]uint8)
	fcEmu.Regi["A"] = 4
	fcEmu.Regi["X"] = 4
	fcEmu.Regi["Y"] = 6
	fcEmu.Regi["S"] = 0xff
	fcEmu.Regi["P"] = 0b00000000
	fcEmu.InterruptFlag = false
}

func initPc(fcEmu *CpuEmu,resetaddr uint16) {
	fcEmu.RegPc = resetaddr
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

func readFile(fcEmu *CpuEmu) ([]uint8) {
	var irqaddr, nmiaddr, resetaddr uint16 = 0, 0, 0
	bufRegpc := 0x8000
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
		fcEmu.Memory[bufRegpc] = writebuf[0]
		//fmt.Printf("rom: 0x%x\n",fcEmu.Memory[bufRegpc])
		bufRegpc++
	}

	fcEmu.Irqaddr = irqaddr
	fcEmu.Nmiaddr = nmiaddr
	fcEmu.Resetaddr = resetaddr
	
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
	return chrrombuf
}

func initMem(fcEmu *CpuEmu) ([]uint8) {
	chrrombuf := readFile(fcEmu)	
	return chrrombuf
}

func InitEmu(fcEmu *CpuEmu) ([]uint8) {
	initReg(fcEmu)
	fcEmu.RegPc = 0x8000
	chrrombuf := initMem(fcEmu)
	return chrrombuf
}


