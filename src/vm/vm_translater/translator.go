package vm_translater

import (
	"computer_emulation/src/memory"
	"computer_emulation/src/vm/vm_ast"
	"computer_emulation/src/vm/vm_tokenizer"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Translator struct {
	program             *vm_ast.Program
	environment         map[string]int
	currentTextAreaLine int
}

func New(program *vm_ast.Program) *Translator {
	return &Translator{
		program:             program,
		environment:         map[string]int{},
		currentTextAreaLine: 0,
	}
}

/*
sample output
add 1 2
*/
func (t *Translator) Translate() string {
	// 1. build environment map
	t.buildEnvironment()

	// 2. evaluate statements
	result := t.predefinedSetupStatement()
	for _, statement := range t.program.Statements {
		translatedStatements := t.translateStatement(statement)
		for _, translatedStatement := range translatedStatements {
			log.Printf("[DEBUG] line[%d] = %q, asm=%q", statement.LineNumber(), statement.String(), translatedStatement)

			result += translatedStatement + "\n"
		}
	}
	fmt.Printf("%s\n", result)
	return strings.TrimSpace(result)
}

func (t *Translator) translateStatement(statement vm_ast.Statement) []string {
	switch stmt := statement.(type) {
	case *vm_ast.PushStatement:
		return t.translatePushStatement(stmt)
	case *vm_ast.AddStatement:
		return t.translateAddStatement(stmt)
	default:
		log.Fatalf("unknown statement has come")
		os.Exit(1)
	}
	return []string{}
}

func (t *Translator) buildEnvironment() {
	for _, statement := range t.program.Statements {
		switch stmt := statement.(type) {
		case *vm_ast.PushStatement:
			if stmt.Value.Type == vm_tokenizer.IDENT {
				if _, ok := t.environment[stmt.Value.Literal]; !ok {
					t.environment[stmt.Value.Literal] = memory.SYMBOL_ENV_BASE_ADDRESS + t.currentTextAreaLine
					log.Printf("%s is placed in an address(%d)", stmt.Value.Literal, t.environment[stmt.Value.Literal])
					t.currentTextAreaLine += 1
				}
			}
		case *vm_ast.AddStatement:
			// noop
		}
	}
}

func (t *Translator) translatePushStatement(stmt *vm_ast.PushStatement) []string {

	// TODO: replace with consts
	// D = {stmt.Value.Literal}
	setDistStatements := []string{}
	switch stmt.Segment.Literal {
	case "argument":
		setDistStatements = []string{
			"@ARG",
			"A=M;",
		}
	case "local":
		setDistStatements = []string{
			"@LCL",
			"A=M;",
		}
	case "static":
		// TODO: envと衝突する. 別の空間にする?
		setDistStatements = []string{
			"@" + strconv.Itoa(memory.STATIC_BASE_ADDRESS),
			"A=M;",
		}
	case "constant":
		setDistStatements = []string{
			"@" + stmt.Value.Literal,
		}
	case "this":
		setDistStatements = []string{
			"@THIS",
			"A=M;",
		}
	case "that":
		setDistStatements = []string{
			"@THAT",
			"A=M;",
		}
	case "pointer":
		// TODO
	case "temp":
		setDistStatements = []string{
			"@" + strconv.Itoa(memory.TEMP0_WORD_ADDRESS),
			"A=M;",
		}
	}

	ops0 := append(
		setDistStatements,
		"D=A;",
	)

	// M[SP] = D
	ops1 := []string{
		"@SP",
		"A=M;",
		"M=D;",
	}

	// SP++
	ops2 := []string{
		"@SP",
		"D=M;",
		"M=D+1;",
	}

	lines := append(append(ops0, ops1...), ops2...)

	result := []string{}
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}
		fmt.Printf("%s\n", trimmedLine)
		result = append(result, trimmedLine)
	}
	return result
}

func (t *Translator) translateAddStatement(stmt *vm_ast.AddStatement) []string {
	return []string{
		// SP is always 0

		// M[SP] = M[SP] - 1
		"@SP",
		"D=M;",
		"M=D-1;",

		// R5 = M[SP]
		"@SP",
		"A=M;",
		"D=M;",
		"@R5",
		"M=D;",

		// M[M[SP]] = 0
		"@SP",
		"A=M;",
		"M=0;",

		// M[SP] = M[SP] - 1
		"@SP",
		"D=M;",
		"M=D-1;",

		// D = R5 + M[M[SP]]
		"@SP",
		"A=M;",
		"D=M;",
		"@R5",
		"D=D+M;",

		// M[M[SP]] = D
		"@SP",
		"A=M;",
		"M=D;",

		// SP = SP + 1
		"@SP",
		"D=M;",
		"M=D+1;",
	}
}

func (t *Translator) predefinedSetupStatement() string {
	return strings.Join([]string{
		// SP = 255
		"@" + strconv.Itoa(memory.GLOBAL_STACK_BASE_ADDRESS),
		"D=A;",
		"@SP",
		"M=D;",
	}, "\n") + "\n"
}

//
//// A command: 0 v v v v v v v v v v v v v v v
//func (t *Translator) translateAllocationStatement(stmt *vm_ast.AllocationStatement) string {
//	if stmt.Value.Type == tokenizer.IDENT {
//		// NOTE. env table are placed in DATA_MEMORY_BASE + 16~
//		// envから値をとってきてbinaryにして放り込み
//		address := t.environment[stmt.Value.Literal]
//		binaryString := intToBinaryString(address)
//		log.Printf("the address of %s is %d. binaryString is %s", stmt.Value.Literal, address, binaryString)
//		return binaryString
//	} else {
//		address, err := strconv.Atoi(stmt.Value.Literal)
//		if err != nil {
//			log.Fatalf("invalid combination of literal and tokentype")
//			os.Exit(1)
//		}
//		return intToBinaryString(address)
//	}
//}
//
