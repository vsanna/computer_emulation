package vm_translater

import (
	"computer_emulation/src/memory"
	"computer_emulation/src/vm/vm_ast"
	"computer_emulation/src/vm/vm_tokenizer"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
)

type Translator struct {
	program             *vm_ast.Program
	environment         map[string]int
	currentTextAreaLine int
}

const TRUE = "1"
const FALSE = "0"

func New(program *vm_ast.Program) *Translator {
	return &Translator{
		program:             program,
		environment:         map[string]int{},
		currentTextAreaLine: 0,
	}
}

func (t *Translator) Translate() string {
	// 1. build environment map
	// TODO: 不要
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

	result += t.predefinedTrailingStatement()
	return strings.TrimSpace(result)
}

func (t *Translator) translateStatement(statement vm_ast.Statement) []string {
	switch stmt := statement.(type) {
	case *vm_ast.PushStatement:
		return t.translatePushStatement(stmt)
	case *vm_ast.PopStatement:
		return t.translatePopStatement(stmt)
	case *vm_ast.AddStatement, *vm_ast.SubStatement, *vm_ast.AndStatement, *vm_ast.OrStatement:
		return t.translateBinaryOperationStatement(stmt)
	case *vm_ast.EqStatement, *vm_ast.GtStatement, *vm_ast.LtStatement:
		return t.translateCompOperationStatement(stmt)
	case *vm_ast.NotStatement:
		return t.translateNotStatement(stmt)
	case *vm_ast.NeqStatement:
		return t.translateNeqStatement(stmt)
	default:
		log.Fatalf("unknown statement has come")
	}
	return []string{}
}

func (t *Translator) translatePushStatement(stmt *vm_ast.PushStatement) []string {

	// TODO: replace with consts
	// D = {stmt.Value.Literal}
	setDregStatements := []string{}
	switch stmt.Segment.Literal {
	case "argument":
		setDregStatements = []string{
			"@ARG",
			"A=M;",
			"D=A;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"A=D;",
			"D=M;", // D = Mem[Mem[ARG] + idx]
		}
	case "local":
		setDregStatements = []string{
			"@LCL",
			"A=M;",
			"D=A;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"A=D;",
			"D=M;",
		}
	case "this":
		setDregStatements = []string{
			"@" + strconv.Itoa(memory.THIS_WORD_ADDRESS),
			"A=M;",
			"D=A;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"A=D;",
			"D=M;", // D = Mem[Mem[THIS_WORD_ADDRESS] + idx]
		}
	case "that":
		setDregStatements = []string{
			"@" + strconv.Itoa(memory.THAT_WORD_ADDRESS),
			"A=M;",
			"D=A;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"A=D;",
			"D=M;", // D = Mem[Mem[THAT_WORD_ADDRESS] + idx]
		}
	case "pointer":
		// push pointer 0 means set base address of a object as THIS's
		// push pointer 1 means set base address of a object as THAT's
		assignCmd := ""
		switch stmt.Value.Literal {
		case "0":
			assignCmd = "@" + strconv.Itoa(memory.THIS_WORD_ADDRESS)
		case "1":
			assignCmd = "@" + strconv.Itoa(memory.THAT_WORD_ADDRESS)
		default:
			log.Fatalf("invalid index")
		}
		setDregStatements = []string{
			assignCmd,
			"D=A;",
		}
	case "static":
		// TODO: envと衝突する. 別の空間にする?
		setDregStatements = []string{
			"@" + strconv.Itoa(memory.STATIC_BASE_ADDRESS),
			"D=A;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"A=D;",
			"D=M;", // D = Mem[Mem[STATIC_BASE_ADDRESS + idx]] ... this/thatとは異なる!
		}
	case "temp":
		setDregStatements = []string{
			"@" + strconv.Itoa(memory.TEMP0_WORD_ADDRESS),
			"D=A;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"A=D;",
			"D=M;", // D = Mem[Mem[TEMP0_WORD_ADDRESS + idx]] ... this/thatとは異なる!
		}
	case "constant":
		setDregStatements = []string{
			"@" + stmt.Value.Literal,
			"D=A;",
		}
	}

	lines := append(setDregStatements, t.pushDregStatements()...)

	result := []string{}
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}
		result = append(result, trimmedLine)
	}
	return result
}

func (t *Translator) translatePopStatement(stmt *vm_ast.PopStatement) []string {
	// TODO: replace with consts
	// D = {stmt.Value.Literal}
	popTopValueStatement := t.popAndSetToStatements("R5")
	setPoppedValueStatement := []string{}
	switch stmt.Segment.Literal {
	case "argument":
		setPoppedValueStatement = []string{
			// M[R6] = M[ARG] + 1
			"@ARG",
			"D=M+1;",
			"@R6",
			"M=D;",
			// M[M[R6] = M[ARG] + 1] = M[R5]
			"@R5",
			"D=M;",
			"@R6",
			"A=M;",
			"M=D;",
		}
	case "local":
		setPoppedValueStatement = []string{
			// M[R6] = M[LCL] + 1
			"@LCL",
			"D=M+1;",
			"@R6",
			"M=D;",
			// M[M[R6] = M[LCL] + 1] = M[R5]
			"@R5",
			"D=M;",
			"@R6",
			"A=M;",
			"M=D;",
		}
	case "this":
		setPoppedValueStatement = []string{
			// M[R6] = M[THIS] + val
			"@THIS",
			"D=M;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"@R6",
			"M=D;", // set an address of target field(M[THIS] + val) in R6
			// M[M[THIS] + val] = M[R5]
			"@R5",
			"D=M;",
			"@R6",
			"A=M;",
			"M=D;",
		}
	case "that":
		setPoppedValueStatement = []string{
			// M[R6] = M[THIS] + val
			"@THAT",
			"D=M;",
			"@" + stmt.Value.Literal,
			"D=A+D;",
			"@R6",
			"M=D;", // set an address of target field(M[THIS] + val) in R6
			// M[M[THIS] + val] = M[R5]
			"@R5",
			"D=M;",
			"@R6",
			"A=M;",
			"M=D;",
		}
	case "pointer":
		switch stmt.Value.Type {
		case vm_tokenizer.IDENT:
			setPoppedValueStatement = []string{
				// M[R6] = THIS_WORD_ADDRESS + ident
				"@" + strconv.Itoa(memory.THIS_WORD_ADDRESS),
				"D=A;",
				"@" + stmt.Value.Literal,
				"D=A+D;",
				"@R6", // temp1
				"M=D;",
				// M[D] = M[R5]
				"@R5", // temp2
				"D=M;",
				"@R6",
				"A=M;",
				"M=D;",
			}
		case vm_tokenizer.INT:
			idx, _ := strconv.Atoi(stmt.Value.Literal)
			setPoppedValueStatement = []string{
				// M[R6] = THIS_WORD_ADDRESS + idx
				"@" + strconv.Itoa(memory.THIS_WORD_ADDRESS+idx),
				"D=A;",
				"@R6",
				"M=D;",
				// M[D] = M[R5]
				"@R5",
				"D=M;",
				"@R6",
				"A=M;",
				"M=D;",
			}
		default:
			log.Fatalf("invalid TokenType")
		}
	case "static":
		// pop static idx とは、@static.{idx}という変数名を宣言し、そこに格納することを意味する。asmのenvとスペースを共有。
		// というより、vmからみてasmのenvに変数セットする命令がpop static idx
		setPoppedValueStatement = []string{
			"@R5",
			"D=M;",
			"@STATIC__" + stmt.Value.Literal,
			"M=D;",
		}
	case "temp":
		setPoppedValueStatement = []string{
			// M[R6] = TEMP0_WORD_ADDRESS + idx
			"@" + strconv.Itoa(memory.TEMP0_WORD_ADDRESS),
			"D=A;",
			"@" + stmt.Value.Literal, // now assuming that Literal is always INT. if IDENT comes here, should change logic
			"D=A+D;",
			"@R6",
			"M=D;",
			// M[D] = M[R5]
			"@R5",
			"D=M;",
			"@R6",
			"A=M;",
			"M=D;",
		}
	case "constant":
		setPoppedValueStatement = []string{
			"@R5",
			"M=0;",
		}
	}

	lines := append(popTopValueStatement, setPoppedValueStatement...)

	result := []string{}
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}
		result = append(result, trimmedLine)
	}
	return result
}

func (t *Translator) translateBinaryOperationStatement(stmt vm_ast.Statement) []string {
	operationLine := ""

	switch stmt.(type) {
	case *vm_ast.AddStatement:
		operationLine = "D=D+M;"
	case *vm_ast.SubStatement:
		operationLine = "D=D-M;"
	case *vm_ast.AndStatement:
		operationLine = "D=D&M;"
	case *vm_ast.OrStatement:
		operationLine = "D=D|M;"
	}

	result := []string{}

	// R5 = pop()
	result = append(result, t.popAndSetToStatements("R5")...)
	// R6 = pop()
	result = append(result, t.popAndSetToStatements("R6")...)
	// ops(R5, R6)
	result = append(result, []string{
		"@R6",
		"D=M;",
		"@R5",         // NOTE: it's "stack"! so, R6 is a value on top of the stack and R5 is a value below that.
		operationLine, // ex "D=D+M;"
	}...)
	// push D to stack
	result = append(result, t.pushDregStatements()...)

	return result
}

func (t *Translator) translateCompOperationStatement(stmt vm_ast.Statement) []string {
	compMethod := ""
	switch stmt.(type) {
	case *vm_ast.EqStatement:
		compMethod = "JEQ"
	case *vm_ast.GtStatement:
		compMethod = "JGT"
	case *vm_ast.LtStatement:
		compMethod = "JLT"
	}

	result := []string{}

	// R5 = pop()
	result = append(result, t.popAndSetToStatements("R5")...)
	// R6 = pop()
	result = append(result, t.popAndSetToStatements("R6")...)
	// set R5-R6 in D
	result = append(result, []string{
		"@R6",
		"D=M;",
		"@R5",
		"D=D-M;",
	}...)

	keyString := generateUuidForIdent()

	// IF logic
	ifLogic := []string{
		"@" + keyString + "_THEN",
		"D;" + compMethod,
		"@" + keyString + "_ELSE",
		"0;JMP",
	}

	// true section
	trueSection := append([]string{
		"(" + keyString + "_THEN)",
		"@" + TRUE,
		"D=A;",
	}, t.pushDregStatements()...)
	trueSection = append(trueSection, []string{
		"@" + keyString + "_END",
		"0;JMP",
	}...)

	// false section
	falseSection := append([]string{
		"(" + keyString + "_ELSE)",
		"@" + FALSE,
		"D=A;",
	}, t.pushDregStatements()...)
	falseSection = append(falseSection, []string{
		"@" + keyString + "_END",
		"0;JMP",
	}...)

	result = append(result, ifLogic...)
	result = append(result, trueSection...)
	result = append(result, falseSection...)
	result = append(result, "("+keyString+"_END)")

	return result
}

func (t *Translator) translateNotStatement(stmt vm_ast.Statement) []string {
	result := []string{}

	keyString := generateUuidForIdent()

	// pop and set the top value in M[R5]
	result = append(result, t.popAndSetToStatements("R5")...)

	// set M[R5] in D
	result = append(result, []string{
		"@R5",
		"D=M-1;",
		"@" + keyString + "_THEN",
		"D;JEQ", // Dが0 <=> Mが1 <=> R5には1が入っている <=> 0をpushする
		"@" + keyString + "_ELSE",
		"0;JMP", // Dが0 <=> Mが1 <=> R5には1が入っている <=> 0をpushする
	}...)

	// true section
	trueSection := append([]string{
		"(" + keyString + "_THEN)",
		"@" + FALSE,
		"D=A;",
	}, t.pushDregStatements()...)
	trueSection = append(trueSection, []string{
		"@" + keyString + "_END",
		"0;JMP",
	}...)

	// false section
	falseSection := append([]string{
		"(" + keyString + "_ELSE)",
		"@" + TRUE,
		"D=A;",
	}, t.pushDregStatements()...)
	falseSection = append(falseSection, []string{
		"@" + keyString + "_END",
		"0;JMP",
	}...)

	result = append(result, trueSection...)
	result = append(result, falseSection...)
	result = append(result, "("+keyString+"_END)")

	return result
}

func (t *Translator) translateNeqStatement(stmt vm_ast.Statement) []string {
	result := t.popAndSetToStatements("R5")
	result = append(result, []string{
		"D=0;",
		"@R5",
		"M=A;",
		"D=D-M;",
	}...)
	result = append(result, t.pushDregStatements()...)
	return result
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

func (t *Translator) predefinedSetupStatement() string {
	return strings.Join([]string{
		// SP = 255
		"@" + strconv.Itoa(memory.GLOBAL_STACK_BASE_ADDRESS),
		"D=A;",
		"@SP",
		"M=D;",
	}, "\n") + "\n"
}

func (t *Translator) predefinedTrailingStatement() string {
	return strings.Join([]string{
		"(VM_END)",
		"@VM_END",
		"0;JMP",
	}, "\n") + "\n"
}

func (t *Translator) popAndSetToStatements(registerName string) []string {
	return []string{
		////  {registerName} = pop()
		// M[SP] = M[SP] - 1
		"@SP",
		"M=M-1;",
		// {registerName} = M[M[SP]]
		"@SP",
		"A=M;",
		"D=M;",
		"@" + registerName,
		"M=D;",
		// M[M[SP]] = 0
		"@SP",
		"A=M;",
		"M=0;",
	}
}

func (t *Translator) pushDregStatements() []string {
	return []string{
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

// uuid that can be used as ident must start with alphabet
func generateUuidForIdent() string {
	key, _ := uuid.NewUUID()
	keyString := strings.ReplaceAll(key.String(), "-", "_")
	return "generated_ident__" + keyString
}
