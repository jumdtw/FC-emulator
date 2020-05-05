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
	var register = []string{"A", "X", "Y", "S", "P"}
	for _, value := range register {
		fmt.Print("regi", value, " = ")
		fmt.Printf("0x%x\n", degemu.regi[value])
	}
	fmt.Printf("PC = 0x%x\n", degemu.regPc)
}

func initReg(fcEmu *emu) {
	fcEmu.regi = make(map[string]uint8)
	fcEmu.regi["A"] = 4
	fcEmu.regi["X"] = 4
	fcEmu.regi["Y"] = 6
	fcEmu.regi["S"] = 0xff
	fcEmu.regi["P"] = 0b00000010
}

func initPc(fcEmu *emu) {
	fcEmu.regPc = 0x8000
}

func checkNes(checkbuf []uint8) uint8 {
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
	return progSize
}

func readFile(fcEmu *emu) uint8 {
	file, err := os.Open(`C:\Users\ttnmr\HOME\emulator\FC\amegure.nes`)
	if err != nil {
		panic("file error")
	}
	defer file.Close()

	checkbuf := make([]uint8, 16)
	_, checkerr := file.Read(checkbuf)
	if checkerr != nil {
		// Readエラー処理
		panic("error read : checkbuf")

	}
	// progSize 1 == 16384
	progSize := checkNes(checkbuf)
	progcounter := uint32(progSize) * 16384
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
		progcounter--
		// nes file の pro 部分を 0x8000 から 16384*progSize 分読み込む
		fcEmu.memory[fcEmu.regPc] = writebuf[0]
		fcEmu.regPc++
	}
	initPc(fcEmu)
	return progSize
}

func initMem(fcEmu *emu) uint8 {
	progSize := readFile(fcEmu)
	/*
		fcEmu.memory[0x18] = 0x13
		fcEmu.memory[0x0206] = 0x19
		fcEmu.memory[fcEmu.regPc] = 0xf0
		fcEmu.memory[fcEmu.regPc+1] = 0x04
		fcEmu.memory[fcEmu.regPc+2] = 0x02
	*/
	return progSize
}

func initEmu(fcEmu *emu) uint8 {
	initReg(fcEmu)
	initPc(fcEmu)
	progSize := initMem(fcEmu)
	return progSize
}

func main() {
	var fcEmu = emu{}
	//_ = initEmu(&fcEmu)

	progSize := initEmu(&fcEmu)
	progcounter := uint16(progSize) * 16384
	exesize := progcounter + 0x8000
	for fcEmu.regPc < exesize {
		fcEmu.Execute()
		fmt.Printf("main : 0x%x\n", fcEmu.regPc)
	}

	//fcEmu.Execute()
	debug(&fcEmu)
	//fmt.Printf("0x%x\n", fcEmu.memory[0])
}
