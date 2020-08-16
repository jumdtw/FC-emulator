package FC_CPU

import (
	"fmt"
	"os"
)

const (
	memCap = 65536
)

type CpuEmu struct {

	// mirror
	// 0 == horizontal
	// 1 == vertical
	Mirror int

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

	// oamへの書き込みフラッグ
	Oamnum int
	OamWriteFlag bool
	OamWritecount int
	OamWriteValue int
	DAMflag bool
	DAMvalue uint8

	// pad 
	// 初期化用の変数
	Padvalue uint16
	BotmReadcount int
	Abotmflag bool
	Bbotmflag bool
	Selectbotmflag bool
	Startbotmflag bool
	Downbotmflag bool
	Leftbotmflag bool
	Rightbotmflag bool
	Upbotmflag bool

	// draw info
	// write x pos when flag is false, write y pos when flag is true 
	Displaywriteflag bool
	DisplayX int
	DisplayY int

	// A X Y S P
	Regi map[string]uint8
	// PC
	RegPc uint16
	// Memory
	Memory [memCap]uint8

	// 内部情報
	UpdateX int
	UpdateY int
	Bufvram [256*240*4]uint8

	// debug
	Exeopcdlist []uint8
	Saveflag bool
	
}

func (degemu *CpuEmu)Debug() {
	fmt.Println("---------------------------------")
	var Register = []string{"A", "X", "Y", "S"}
	for _, value := range Register {
		fmt.Print("Regi", value, " = ")
		fmt.Printf("0x%x\n", degemu.Regi[value])
	}
	/*
	fmt.Println("--stack--")
	// 現在のSPは空っぽのやつ
	bufsp := degemu.Regi["S"]
	ddd := 0x100+uint16(bufsp)
	aaa := degemu.Memory[ddd]
	for bufsp!=0xff {
		fmt.Printf("0x%x : value 0x%x \n",ddd,aaa)
		bufsp++
		ddd = 0x100+uint16(bufsp)
		aaa = degemu.Memory[ddd]
	}
	fmt.Printf("0x%x : value 0x%x \n",0x1ff,degemu.Memory[0x1ff])
	fmt.Println("---------")
	*/
	fmt.Printf("N. V. R. B. D. I. Z. C\n")
	fmt.Printf("RegiP = 0b%08b, 0x%x\n",degemu.Regi["P"],degemu.Regi["P"])
	fmt.Printf("PC = 0x%x\n", degemu.RegPc)
	fmt.Printf("opcd  = 0x%x\n",degemu.Memory[degemu.RegPc])
	//fmt.Printf("Memory [0] = %d\n",degemu.Memory[0])
	//fmt.Printf("Memory [0x0] = 0x%x\n",degemu.Memory[0x0])
	//fmt.Printf("Memory [0x1] = 0x%x\n",degemu.Memory[0x1])
}

func initReg(fcEmu *CpuEmu) {
	fcEmu.Regi = make(map[string]uint8)
	fcEmu.Regi["A"] = 0
	fcEmu.Regi["X"] = 0
	fcEmu.Regi["Y"] = 0
	fcEmu.Regi["S"] = 0xfd
	fcEmu.Regi["P"] = 0b00100100
	fcEmu.InterruptFlag = false
}

func initPc(fcEmu *CpuEmu,resetaddr uint16) {
	fcEmu.RegPc = resetaddr
	fcEmu.RegPc = 0xc000
}

func checkNes(checkbuf []uint8) (uint8, uint8, int) {
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
	mirrorflag := checkbuf[6] & 0b00000001
	mirror := 0
	if mirrorflag == 0b00000001 {
		mirror = 1
	}
	return progSize, chrSize, mirror
}

func readFile(fcEmu *CpuEmu) ([]uint8) {
	var irqaddr, nmiaddr, resetaddr uint16 = 0, 0, 0
	var progSize, chrSize uint8 = 0, 0
	bufRegpc := 0x8000
	//file, err := os.Open(`C:\Users\ttnmr\go\src\github.com\jumdtw\FC-emulator\chickenrace2.nes`)
	//file, err := os.Open(`C:\Users\ttnmr\OneDrive\デスクトップ\software\mario.nes`)
	//file, err := os.Open(`C:\Users\ttnmr\go\src\github.com\jumdtw\FC-emulator\amegure.nes`)
	file, err := os.Open(`C:\Users\ttnmr\OneDrive\デスクトップ\software\nestest.nes`)
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
	progSize, chrSize, fcEmu.Mirror = checkNes(checkbuf)
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
		// nes file の pro 部分を 0x8000 から 16384*progSize 分読み込む
		fcEmu.Memory[bufRegpc] = writebuf[0]
		bufRegpc++
		progcounter--
	}

	if progSize == 1 {
		for i:=0 ; i < 0x4000 ; i++ {
			fcEmu.Memory[0xc000+i] = fcEmu.Memory[0x8000+i]
		}
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
	fcEmu.BotmReadcount = 0
	fcEmu.DisplayX = 0
	fcEmu.DisplayY = 0
	fcEmu.Displaywriteflag = false
	fcEmu.DAMflag = false
	fcEmu.Saveflag = false
	chrrombuf := initMem(fcEmu)
	return chrrombuf
}


