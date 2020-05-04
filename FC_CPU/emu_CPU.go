package main

import (
	"fmt"
	"os"
)

const (
	memCap  = 65535
	BUFSIZE = 1024
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
	fcEmu.regi["P"] = 4
}

func initPc(fcEmu *emu) {
	fcEmu.regPc = 0x8000
}

func checkNes(checkbuf []uint8) {
	if checkbuf[0] != 0x4e {
		panic("not nes file : not 0x4e")
	}
	if checkbuf[1] != 0x45 {
		panic("not nes file : not 0x45")
	}
	if checkbuf[2] != 0x53 {
		panic("not nes file : not 0x53")
	}
	return
}

func readFile(fcEmu *emu) {
	file, err := os.Open(`C:\Users\ttnmr\HOME\emulator\FC\amegure.nes`)
	if err != nil {
		fmt.Println("file error")
		return
	}
	defer file.Close()

	checkbuf := make([]uint8, 16)
	_, checkerr := file.Read(checkbuf)
	if checkerr != nil {
		// Readエラー処理
		panic("error read : checkbuf")

	}
	checkNes(checkbuf)
	writebuf := make([]uint8, 1)

	for {
		n, writeerr := file.Read(writebuf)
		if n == 0 {

			break
		}
		if writeerr != nil {
			// Readエラー処理
			fmt.Println("error read : writebuf")
			break
		}

	}
}

func initMem(fcEmu *emu) {
	readFile(fcEmu)
	/*
		fcEmu.memory[0x18] = 0x13
		fcEmu.memory[0x0206] = 0x19
		fcEmu.memory[fcEmu.regPc] = 0xbe
		fcEmu.memory[fcEmu.regPc+1] = 0x00
		fcEmu.memory[fcEmu.regPc+2] = 0x02
	*/
}

func initEmu(fcEmu *emu) {
	initReg(fcEmu)
	initPc(fcEmu)
	initMem(fcEmu)
}

func main() {
	var fcEmu = emu{}
	initEmu(&fcEmu)
	for fcEmu.regPc > 0xfffa {
		fcEmu.Execute()
	}
	debug(&fcEmu)
}
