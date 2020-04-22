package main

type CheckKey func() bool
type GetCharacter func() uint16

type Memory struct {
	CheckKey
	GetCharacter
	Storage [VMEM_MAX]uint16
}

func (memory *Memory) WriteMemory(address uint16, value uint16) {
	memory.Storage[address] = value
}

func (memory *Memory) ReadMemory(address uint16) uint16 {
	if address == VMR_KBDR {
		memory.Storage[VMR_KBSR] = (1 << 15)
		memory.Storage[VMR_KBDR] = memory.GetCharacter()
	} else {
		memory.Storage[VMR_KBSR] = 0
	}
	return memory.Storage[address]
}
