package main

const VM_START = 0x3000
const VMEM_MAX = 0xFFFF

const (
	VOP_BR uint16 = iota
	VOP_ADD
	VOP_LD
	VOP_ST
	VOP_JSR
	VOP_AND
	VOP_LDR
	VOP_STR
	VOP_RTI
	VOP_NOT
	VOP_LDI
	VOP_STI
	VOP_JMP
	VOP_RES
	VOP_LEA
	VOP_TRAP
)

const (
	VR_R0 uint16 = iota
	VR_R1
	VR_R2
	VR_R3
	VR_R4
	VR_R5
	VR_R6
	VR_R7
	VR_PC
	VR_COND
	VR_COUNT
)

const (
	VFL_POS uint16 = 1 << 0
	VFL_ZRO uint16 = 1 << 1
	VFL_NEG uint16 = 1 << 2
)

const (
	VMR_KBSR uint16 = 0xFE00
	VMR_KBDR uint16 = 0xFE02
)
