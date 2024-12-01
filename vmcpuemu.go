package vm

import (
	"strings"
	"unsafe"
)

var supportedInstructions []string = []string{
	"PUSH",
	"POP",
	"MOV",
	"LEA",
	"XCHG",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"INC",
	"DEC",
	"AND",
	"OR",
	"XOR",
	"CMP",
	"TEST",
	"JE",
	"JNE",
	"JG",
	"JL",
	"JGE",
	"JLE",
	"JMP",
	"LOCK",
	"SYSCALL",
}

/**
 * This is the most unsafe shit but fuck it its a cpu
 * so stfu Ethan
 *
 * Seriously though, this shit is a sure way to blow your
 * leg off
 */

func (VM *vm) ExecInstruction() error {

	if VM.rip >= len(VM.code) {
		return VM.vmSegfault("rip at invalid memory address")
	}

	switch strings.ToUpper(VM.opcodes[VM.code[VM.rip]]) {

	case "PUSH":

		if VM.rip+1 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 {
			return VM.vmSegfault("invalid operand")
		}

		VM.stack = append(VM.stack, VM.registers[operandOneIndex])

		VM.rip += 2
		break

	case "POP":

		if VM.rip+1 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 {
			return VM.vmSegfault("invalid operand")
		}

		if length := len(VM.stack); length == 0 {
			return VM.vmSegfault("stack underflow")
		} else {
			VM.registers[operandOneIndex] = VM.stack[length-1]
			VM.stack = VM.stack[:length-1]
		}

		VM.rip += 2
		break

	case "MOV":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		VM.registers[operandOneIndex] = VM.registers[operandTwoIndex]

		VM.rip += 3
		break

	case "LEA":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]
		switch operandTwo.(type) {

		case unsafe.Pointer:
			VM.registers[operandOneIndex] = operandTwo
			break

		default:
			return VM.vmSegfault("invalid data type for operand two")
		}

		VM.rip += 3
		break

	case "XCHG":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var tmp interface{} = VM.registers[operandOneIndex]
		VM.registers[operandOneIndex] = VM.registers[operandTwoIndex]
		VM.registers[operandTwoIndex] = tmp

		VM.rip += 3
		break

	case "ADD":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

			case int:

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) + operandTwo.(byte)
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) + byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) + int(operandTwo.(byte))
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) + operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "SUB":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

			case int:

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) - operandTwo.(byte)
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) - byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) - int(operandTwo.(byte))
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) - operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "MUL":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

			case int:

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) * operandTwo.(byte)
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) * byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) * int(operandTwo.(byte))
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) * operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "DIV":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				if int(operandTwo.(byte)) == 0 {
					return VM.vmSegfault("division by zero")
				}

			case int:

				if operandTwo.(int) == 0 {
					return VM.vmSegfault("division by zero")
				}

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				if int(operandTwo.(byte)) == 0 {
					return VM.vmSegfault("division by zero")
				}

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) / operandTwo.(byte)
				break

			case int:

				if operandTwo.(int) == 0 {
					return VM.vmSegfault("division by zero")
				}

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) / byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				if int(operandTwo.(byte)) == 0 {
					return VM.vmSegfault("division by zero")
				}

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) / int(operandTwo.(byte))
				break

			case int:

				if operandTwo.(int) == 0 {
					return VM.vmSegfault("division by zero")
				}

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) / operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "INC":

		if VM.rip+1 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 {
			return VM.vmSegfault("invalid operand")
		}

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

		case byte:

			VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) + 1
			break

		case int:

			VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) + 1
			break

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 2
		break

	case "DEC":

		if VM.rip+1 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 {
			return VM.vmSegfault("invalid operand")
		}

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

		case byte:

			VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) - 1
			break

		case int:

			VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) - 1
			break

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 2
		break

	case "AND":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

			case int:

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) & operandTwo.(byte)
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) & byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) & int(operandTwo.(byte))
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) & operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "OR":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

			case int:

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) | operandTwo.(byte)
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) | byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) | int(operandTwo.(byte))
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) | operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "XOR":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		var operandTwo interface{} = VM.registers[operandTwoIndex]

		switch VM.registers[operandOneIndex].(type) {

		case unsafe.Pointer:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

			case int:

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case byte:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) ^ operandTwo.(byte)
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(byte) ^ byte(operandTwo.(int))
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		case int:

			switch operandTwo.(type) {

			case unsafe.Pointer:

			case byte:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) ^ int(operandTwo.(byte))
				break

			case int:

				VM.registers[operandOneIndex] = VM.registers[operandOneIndex].(int) ^ operandTwo.(int)
				break

			default:
				return VM.vmSegfault("invalid data type for operand two")
			}

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		VM.rip += 3
		break

	case "CMP":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		break

	case "TEST":

		if VM.rip+2 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])
		var operandTwoIndex int = int(VM.code[VM.rip+2])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 ||
			operandTwoIndex >= len(VM.registers) || operandTwoIndex < 0 {
			return VM.vmSegfault("invalid operands")
		}

		break

	case "JE":

		break

	case "JNE":

		break

	case "JG":

		break

	case "JL":

		break

	case "JGE":

		break

	case "JLE":

		break

	case "JMP":

		if VM.rip+1 >= len(VM.code) {
			return VM.vmSegfault("insufficient number of operands for opcode")
		}

		var operandOneIndex int = int(VM.code[VM.rip+1])

		if operandOneIndex >= len(VM.registers) || operandOneIndex < 0 {
			return VM.vmSegfault("invalid operand")
		}

		var operandOne interface{} = VM.registers[operandOneIndex]
		switch operandOne.(type) {

		case unsafe.Pointer:

			var absoluteAddr int = *(*int)(operandOne.(unsafe.Pointer))
			VM.rip = absoluteAddr
			break

		case byte:

			VM.rip += int(operandOne.(byte))
			break

		case int:

			VM.rip += operandOne.(int)
			break

		default:
			return VM.vmSegfault("invalid data type for operand one")
		}

		break

	case "LOCK":

		break

	case "SYSCALL":

		break

	default:
		return VM.vmSegfault("invalid opcode")
	}

	return nil
}
