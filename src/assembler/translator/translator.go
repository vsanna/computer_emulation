package translator

import (
	"computer_emulation/src/assembler/ast"
	"computer_emulation/src/assembler/tokenizer"
	"computer_emulation/src/bit"
	"computer_emulation/src/memory"
	"log"
	"os"
	"strconv"
)

type Translator struct {
	program             *ast.Program
	environment         map[string]int
	currentTextAreaLine int
}

func New(program *ast.Program) *Translator {
	return &Translator{
		program: program,
		environment: map[string]int{
			"SP":   memory.SP_WORD_ADDRESS,
			"LCL":  memory.LCL_WORD_ADDRESS,
			"ARG":  memory.ARG_WORD_ADDRESS,
			"THIS": memory.THIS_WORD_ADDRESS,
			"THAT": memory.THAT_WORD_ADDRESS,
			"R0":   memory.R0_WORD_ADDRESS,
			"R1":   memory.R1_WORD_ADDRESS,
			"R2":   memory.R2_WORD_ADDRESS,
			"R3":   memory.R3_WORD_ADDRESS,
			"R4":   memory.R4_WORD_ADDRESS,
			"R5":   memory.R5_WORD_ADDRESS,
			"R6":   memory.R6_WORD_ADDRESS,
			"R7":   memory.R7_WORD_ADDRESS,
			"R8":   memory.R8_WORD_ADDRESS,
			"R9":   memory.R9_WORD_ADDRESS,
			"R10":  memory.R10_WORD_ADDRESS,
			"R11":  memory.R11_WORD_ADDRESS,
			"R12":  memory.R12_WORD_ADDRESS,
			"R13":  memory.R13_WORD_ADDRESS,
			"R14":  memory.R14_WORD_ADDRESS,
			"R15":  memory.R15_WORD_ADDRESS,
		},
		currentTextAreaLine: 0,
	}
}

/*
sample output
	program := "0000111111111111\n" + // @4095 / SET A 4095
		"1110111111001000\n" + // M=1
		"1111110111001000\n" + // M=M+1
		"1111110111001000\n" + // M=M+1
		"0000000000001101\n" + // @13 / SET A 13
		"1110111111000000\n" + // どこにもセットしない
		"1110111111100000\n" + // A=1
		"1110111111010000\n" + // D=1
		"1110111111001000\n" + // M=1
		"1110111111110000\n" + // AD=1
		"1110111111011000\n" + // DM=1
		"1110111111101000\n" + // AM=1
		"1110111111111000" // ADM=1
*/
func (t *Translator) Translate() string {
	// Phase1. build an environment
	t.buildEnvironment()

	// Phase2. evaluate/translate statements
	result := ""
	for i, statement := range t.program.Statements {
		translatedStatement := t.translateStatement(statement)
		if len(translatedStatement) == 0 {
			continue
		}
		log.Printf("[DEBUG] line[%d] = %q, binary=%q", statement.LineNumber(), statement.String(), translatedStatement)

		result += translatedStatement
		if i != len(t.program.Statements)-1 {
			result += "\n"
		}
	}

	return result
}

func (t *Translator) translateStatement(statement ast.Statement) string {
	switch stmt := statement.(type) {
	case *ast.AllocationStatement:
		return t.translateAllocationStatement(stmt)
	case *ast.AddressTaggingStatement:
		return t.translateAddressTaggingStatement(stmt)
	case *ast.OpsAndJumpStatement:
		return t.translateOpsAndJumpStatement(stmt)
	default:
		log.Fatalf("unknown statement has come")
		os.Exit(1)
	}
	return ""
}

func (t *Translator) buildEnvironment() {
	for _, statement := range t.program.Statements {
		switch stmt := statement.(type) {
		case *ast.AllocationStatement:
			if stmt.Value.Type == tokenizer.IDENT {
				if _, ok := t.environment[stmt.Value.Literal]; !ok {
					// symbol -> address of value in environment. those values are stored where the address is larger than equal to ASSEMBLY_ENVIRONMENT_FROM
					// this is corresponding to memory allocation
					t.environment[stmt.Value.Literal] = memory.SYMBOL_ENV_BASE_ADDRESS + t.currentTextAreaLine
					log.Printf("%s is placed in an address(%d)", stmt.Value.Literal, t.environment[stmt.Value.Literal])
					t.currentTextAreaLine += 1
				}
			}
		case *ast.AddressTaggingStatement:
			if stmt.Value.Type == tokenizer.IDENT {
				// symbol -> address of a line of program, so it should be more than or equal to memory.PROGRAM_MEMORY_BASE
				t.environment[stmt.Value.Literal] = stmt.LineNumber()
			}
		case *ast.OpsAndJumpStatement:
			// OpsAndJumpStatement cannot have IDENT in it.
		}
	}
}

// A command: 0 v v v v v v v v v v v v v v v
func (t *Translator) translateAllocationStatement(stmt *ast.AllocationStatement) string {
	if stmt.Value.Type == tokenizer.IDENT {
		// envから値をとってきてbinaryにして放り込み
		address := t.environment[stmt.Value.Literal]
		binaryString := intToBinaryString(address)
		log.Printf("[DEBUG] the address of %s is %d. binaryString is %s", stmt.Value.Literal, address, binaryString)
		return binaryString
	} else {
		address, err := strconv.Atoi(stmt.Value.Literal)
		if err != nil {
			log.Fatalf("invalid combination of literal and tokentype")
			os.Exit(1)
		}
		return intToBinaryString(address)
	}
}

// ex. (LOOP)
// this is meta line for environment-build phase. and so AddressTaggingStatement doesn't return any line
func (t *Translator) translateAddressTaggingStatement(stmt *ast.AddressTaggingStatement) string {
	return ""
}

// C command: 1 1 1 a c c c c c c d d d j j j
func (t *Translator) translateOpsAndJumpStatement(stmt *ast.OpsAndJumpStatement) string {
	aCommandHeadBits := []int{1, 1, 1}

	// TODO: this is SO NAIVE way to translate. improve later
	aBit := 0
	compBits := []int{}
	switch len(stmt.Comp) {
	case 1:
		switch stmt.Comp[0].Type {
		case tokenizer.A_REG:
			aBit = 0
			compBits = []int{1, 1, 0, 0, 0, 0}
		case tokenizer.D_REG:
			aBit = 0
			compBits = []int{0, 0, 1, 1, 0, 0}
		case tokenizer.MEMORY:
			aBit = 1
			compBits = []int{1, 1, 0, 0, 0, 0}
		case tokenizer.INT:
			aBit = 0
			switch stmt.Comp[0].Literal {
			case "0":
				compBits = []int{1, 0, 1, 0, 1, 0}
			case "1":
				compBits = []int{1, 1, 1, 1, 1, 1}
			default:
				log.Fatalf("invalid comp. comp size 1 and Value is int, but not 0 or 1. actual = %q", stmt.Comp[0].Literal)
				os.Exit(1)
			}
		default:
			log.Fatalf("invalid comp. comp size 1 but Value is not A nor D nor M nor int. int, actual = %q", stmt.Comp[0].Type)
			os.Exit(1)
		}
	case 2:
		switch stmt.Comp[0].Type {
		case tokenizer.MINUS:
			switch stmt.Comp[1].Type {
			case tokenizer.A_REG:
				aBit = 0
				compBits = []int{1, 1, 0, 0, 1, 1}
			case tokenizer.D_REG:
				aBit = 0
				compBits = []int{0, 0, 1, 1, 1, 1}
			case tokenizer.MEMORY:
				aBit = 1
				compBits = []int{1, 1, 0, 0, 1, 1}
			case tokenizer.INT:
				aBit = 0
				compBits = []int{1, 1, 1, 0, 1, 0}
			default:
				log.Fatalf("invalid comp. comp size is 2 and MINUS is the head, but trailing value is unexpected one. actual=%q", stmt.Comp[0].Literal)
				os.Exit(1)
			}
		case tokenizer.BANG:
			switch stmt.Comp[1].Type {
			case tokenizer.A_REG:
				aBit = 0
				compBits = []int{1, 1, 0, 0, 0, 1}
			case tokenizer.D_REG:
				aBit = 0
				compBits = []int{0, 0, 1, 1, 0, 1}
			case tokenizer.MEMORY:
				aBit = 1
				compBits = []int{1, 1, 0, 0, 0, 1}
			default:
				log.Fatalf("invalid comp. comp size is 2 and BANG is the head, but trailing value is unexpected one. actual=%q", stmt.Comp[1].Literal)
				os.Exit(1)
			}
		default:
			log.Printf("stmt = %q", stmt)
			log.Printf("stmt.Comp = %q", stmt.Comp)
			log.Fatalf("invalid comp. comp size is 2 but the head is not BANG nor MINUS. actual=%q", stmt.Comp[1].Literal)
			os.Exit(1)
		}
	case 3:
		if stmt.Comp[0].Type == tokenizer.A_REG &&
			stmt.Comp[1].Type == tokenizer.PLUS &&
			stmt.Comp[2].Type == tokenizer.INT &&
			stmt.Comp[2].Literal == "1" {
			aBit = 0
			compBits = []int{1, 1, 0, 1, 1, 1}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.PLUS &&
			stmt.Comp[2].Type == tokenizer.INT &&
			stmt.Comp[2].Literal == "1" {
			aBit = 0
			compBits = []int{0, 1, 1, 1, 1, 1}
		}
		if stmt.Comp[0].Type == tokenizer.MEMORY &&
			stmt.Comp[1].Type == tokenizer.PLUS &&
			stmt.Comp[2].Type == tokenizer.INT &&
			stmt.Comp[2].Literal == "1" {
			aBit = 1
			compBits = []int{1, 1, 0, 1, 1, 1}
		}

		if stmt.Comp[0].Type == tokenizer.A_REG &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.INT &&
			stmt.Comp[2].Literal == "1" {
			aBit = 0
			compBits = []int{1, 1, 0, 0, 1, 0}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.INT &&
			stmt.Comp[2].Literal == "1" {
			aBit = 0
			compBits = []int{0, 0, 1, 1, 1, 0}
		}
		if stmt.Comp[0].Type == tokenizer.MEMORY &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.INT &&
			stmt.Comp[2].Literal == "1" {
			aBit = 1
			compBits = []int{1, 1, 0, 0, 1, 0}
		}

		if stmt.Comp[0].Type == tokenizer.A_REG &&
			stmt.Comp[1].Type == tokenizer.PLUS &&
			stmt.Comp[2].Type == tokenizer.D_REG {
			aBit = 0
			compBits = []int{0, 0, 0, 0, 1, 0}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.PLUS &&
			stmt.Comp[2].Type == tokenizer.MEMORY {
			aBit = 1
			compBits = []int{0, 0, 0, 0, 1, 0}
		}

		if stmt.Comp[0].Type == tokenizer.A_REG &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.D_REG {
			aBit = 0
			compBits = []int{0, 0, 0, 1, 1, 1}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.A_REG {
			aBit = 0
			compBits = []int{0, 1, 0, 0, 1, 1}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.MEMORY {
			aBit = 1
			compBits = []int{0, 1, 0, 0, 1, 1}
		}
		if stmt.Comp[0].Type == tokenizer.MEMORY &&
			stmt.Comp[1].Type == tokenizer.MINUS &&
			stmt.Comp[2].Type == tokenizer.D_REG {
			aBit = 1
			compBits = []int{0, 0, 0, 1, 1, 1}
		}

		if stmt.Comp[0].Type == tokenizer.A_REG &&
			stmt.Comp[1].Type == tokenizer.AND &&
			stmt.Comp[2].Type == tokenizer.D_REG {
			aBit = 0
			compBits = []int{0, 0, 0, 0, 0, 0}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.AND &&
			stmt.Comp[2].Type == tokenizer.MEMORY {
			aBit = 1
			compBits = []int{0, 0, 0, 0, 0, 0}
		}

		if stmt.Comp[0].Type == tokenizer.A_REG &&
			stmt.Comp[1].Type == tokenizer.OR &&
			stmt.Comp[2].Type == tokenizer.D_REG {
			aBit = 0
			compBits = []int{0, 1, 0, 1, 0, 1}
		}
		if stmt.Comp[0].Type == tokenizer.D_REG &&
			stmt.Comp[1].Type == tokenizer.OR &&
			stmt.Comp[2].Type == tokenizer.MEMORY {
			aBit = 1
			compBits = []int{0, 1, 0, 1, 0, 1}
		}
	default:
		log.Fatalf("too many Comp")
		os.Exit(1)
	}

	destBits := []int{0, 0, 0}
	for _, token := range stmt.Dest {
		switch token.Type {
		case tokenizer.A_REG:
			destBits[0] = 1
		case tokenizer.D_REG:
			destBits[1] = 1
		case tokenizer.MEMORY:
			destBits[2] = 1
		}
	}

	jumpBits := []int{0, 0, 0}
	if stmt.Jump != nil {
		switch stmt.Jump.Type {
		case tokenizer.JGT:
			jumpBits = []int{0, 0, 1}
		case tokenizer.JEQ:
			jumpBits = []int{0, 1, 0}
		case tokenizer.JGE:
			jumpBits = []int{0, 1, 1}
		case tokenizer.JLT:
			jumpBits = []int{1, 0, 0}
		case tokenizer.JNE:
			jumpBits = []int{1, 0, 1}
		case tokenizer.JLE:
			jumpBits = []int{1, 1, 0}
		case tokenizer.JMP:
			jumpBits = []int{1, 1, 1}
		}
	}

	result := []int{}
	result = append(result, aCommandHeadBits...)
	result = append(result, aBit)
	result = append(result, compBits...)
	result = append(result, destBits...)
	result = append(result, jumpBits...)

	return intSliceToString(result)
}

func intToBinaryString(address int) string {
	bits := []int{}
	for address > 0 {
		lsb := address & 1
		bits = append(bits, lsb)
		address = address >> 1
	}

	for len(bits) < bit.BUS_WIDTH {
		bits = append(bits, 0)
	}

	reverse := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i, bit := range bits {
		reverse[len(reverse)-i-1] = bit
	}

	return intSliceToString(reverse)
}

func intSliceToString(bits []int) string {
	result := ""
	for _, bit := range bits {
		result += strconv.Itoa(bit)
	}
	return result
}
