package main

import (
	"encoding/binary"
	"os"
)

// controls
const HLT = uint16(0b1000000000000000)
const MI = uint16(0b0100000000000000)
const RI = uint16(0b0010000000000000)
const RO = uint16(0b0001000000000000)
const IO = uint16(0b0000100000000000)
const II = uint16(0b0000010000000000)
const AI = uint16(0b0000001000000000)
const AO = uint16(0b0000000100000000)
const EO = uint16(0b0000000010000000)
const SU = uint16(0b0000000001000000)
const BI = uint16(0b0000000000100000)
const OI = uint16(0b0000000000010000)
const CE = uint16(0b0000000000001000)
const CO = uint16(0b0000000000000100)
const J = uint16(0b0000000000000010)

// each instruction is composed of 5 micro instructions
// to makes things easier, appended 3 emtpty bytes to each
var instructions = []uint16{
	// 0000 - NOP
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 0001 - LDA
	MI | CO, RO | II | CE, IO | MI, RO | AI, 0, 0, 0, 0,
	// 0010 - ADD
	MI | CO, RO | II | CE, IO | MI, RO | BI, EO | AI, 0, 0, 0,
	// 0011 - SUB
	MI | CO, RO | II | CE, IO | MI, RO | BI, EO | AI | SU, 0, 0, 0,
	// 0100 - STA
	MI | CO, RO | II | CE, IO | MI, AO | RI, 0, 0, 0, 0,
	// 0101 - LDI
	MI | CO, RO | II | CE, IO | AI, 0, 0, 0, 0, 0,
	// 0110 - JMP
	MI | CO, RO | II | CE, IO | J, 0, 0, 0, 0, 0,
	// 0111
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1000
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1001
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1010
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1011
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1100
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1101
	MI | CO, RO | II | CE, 0, 0, 0, 0, 0, 0,
	// 1110 - OUT
	MI | CO, RO | II | CE, AO | OI, 0, 0, 0, 0, 0,
	// 1111 - HLT
	MI | CO, RO | II | CE, HLT, 0, 0, 0, 0, 0,
}

func main() {
	// 1. create an empty binary for each EEPROM
	bin1 := make([]byte, 256) // most significant bits
	bin2 := make([]byte, 256) // least significant bits

	// 2. add the micro instructions to the binaries
	for i := 0; i < len(instructions); i++ {
		bs := binary.BigEndian.AppendUint16([]byte{}, instructions[i])
		bin1[i] = bs[0]
		bin2[i] = bs[1]
	}

	// 3. write the binary files
	f1, err := os.Create("instruction_register_1.bin")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	if _, err := f1.Write(bin1); err != nil {
		panic(err)
	}
	f2, err := os.Create("instruction_register_2.bin")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if _, err := f2.Write(bin2); err != nil {
		panic(err)
	}
}
