package main

import "io"

type CPU struct {
	Registers          [VR_COUNT]uint16
	Memory             *Memory
	CurrentInstruction uint16
	CurrentOperation   uint16
	IsRunning          bool
	StartAt            uint16
	Output             io.Writer
}

func CreateCPU(memory *Memory, output io.Writer) *CPU {
	return &CPU{
		StartAt: VM_START,
		Memory:  memory,
		Output:  output,
	}
}

func (cpu *CPU) ResetCPU() {
	cpu.Registers = [VR_COUNT]uint16{}
	cpu.Memory = &Memory{}
	cpu.CurrentInstruction = 0
	cpu.CurrentOperation = 0
	cpu.IsRunning = false
}
