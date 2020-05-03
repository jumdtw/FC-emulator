package main

import (
	"fmt"
)

var register = []string{"A", "X", "Y", "S", "P"}

const (
	memCap = 65535
	//register = []string{"A", "X", "Y", "S", "P"}
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
	for _, value := range register {
		fmt.Println("regi", value, "=", degemu.regi[value])
	}
	fmt.Printf("PC = 0x%x\n", degemu.regPc)
}

func initReg(fcEmu *emu) {
	fcEmu.regi = make(map[string]uint8)
	fcEmu.regi["A"] = 4
	fcEmu.regi["X"] = 4
	fcEmu.regi["Y"] = 4
	fcEmu.regi["S"] = 4
	fcEmu.regi["P"] = 4
}

func initPc(fcEmu *emu) {
	fcEmu.regPc = 0x8000
}

func initMem(fcEmu *emu) {
	fcEmu.memory[0x200] = 20
	fcEmu.memory[fcEmu.regPc] = 0xad
	fcEmu.memory[fcEmu.regPc+1] = 0x00
	fcEmu.memory[fcEmu.regPc+2] = 0x02
}

func initEmu(fcEmu *emu) {
	initReg(fcEmu)
	initPc(fcEmu)
	initMem(fcEmu)
}

func main() {
	var fcEmu = emu{}
	initEmu(&fcEmu)
	fcEmu.Execute()
	debug(&fcEmu)
}
