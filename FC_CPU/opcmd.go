package main

const (
	LDA_IMM = 0xa9
)

func (fc_emu emu) lda_imm(){
	fc_emu.regi["A"] = fc_emu.memory[fc_emu.reg_pc]
	fc_emu.reg_pc = fc_emu.reg_pc + 1
}

func (fc_emu emu) nop(){
}

func (fc_emu emu) Execute(){
	var opcd uint8 = fc_emu.memory[fc_emu.reg_pc]
	fc_emu.reg_pc = fc_emu.reg_pc + 1

	switch opcd{
	
	case LDA_IMM:
		fc_emu.lda_imm()
	default:
		fc_emu.nop()
	}
	
}

