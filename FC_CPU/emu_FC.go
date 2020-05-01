package main

import (
	"fmt"
)

var register = []string{"A", "X", "Y", "S", "P"}

const (
	mem_cap = 65535
	//register = []string{"A", "X", "Y", "S", "P"}
)

type emu struct {
	// A X Y S P
	regi map[string]uint8
	// PC 
	reg_pc uint16
	// memory
	memory [mem_cap]uint8
}

func Debug(degemu *emu){
	for _, value := range register {
		fmt.Println("regi",value,"=", degemu.regi[value])
	}
	fmt.Printf("PC = 0x%x\n",degemu.reg_pc)
}

func Init_reg(fc_emu *emu){
	fc_emu.regi = make(map[string]uint8)
	fc_emu.regi["A"] = 4
	fc_emu.regi["X"] = 4
	fc_emu.regi["Y"] = 4
	fc_emu.regi["S"] = 4
	fc_emu.regi["P"] = 4
}

func Init_pc(fc_emu *emu){
	fc_emu.reg_pc = 0x8000
}

func Init_mem(fc_emu *emu){
	fc_emu.memory[fc_emu.reg_pc] = 0xa9
	fc_emu.memory[fc_emu.reg_pc+1] = 0x01
}

func Init_emu(fc_emu *emu){
	Init_reg(fc_emu)
	Init_pc(fc_emu)
	Init_mem(fc_emu)
}

func main(){
	var fc_emu = emu{}
	Init_emu(&fc_emu)
	//fc_emu.reg_pc = fc_emu.reg_pc + 1
	fc_emu.Execute()
	Debug(&fc_emu)
}