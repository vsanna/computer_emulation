package vm_ast

import (
	"computer_emulation/src/vm/vm_tokenizer"
	"strconv"
)

type Statement interface {
	statement()
	String() string
	LineNumber() int
}

type PushStatement struct {
	Segment vm_tokenizer.Token
	Value   vm_tokenizer.Token
	Line    int
}

func NewPushStatement(segmentName string, idx int) *PushStatement {
	return &PushStatement{
		Segment: vm_tokenizer.Token{Type: vm_tokenizer.IDENT, Literal: segmentName},
		Value:   vm_tokenizer.Token{Literal: strconv.Itoa(idx), Type: vm_tokenizer.INT},
	}
}

func NewZeroPushStatement() *PushStatement {
	tmp := NewPushStatement("", -1)
	tmp.Value = vm_tokenizer.ZeroToken()
	tmp.Segment = vm_tokenizer.ZeroToken()
	return tmp
}

func (a PushStatement) statement() {
}

func (a PushStatement) String() string {
	result := vm_tokenizer.PUSH
	if !a.Segment.IsZero() {
		result += " " + a.Segment.Literal
	}
	result += " " + a.Value.Literal
	return result
}

func (a PushStatement) LineNumber() int {
	return a.Line
}

type PopStatement struct {
	Segment vm_tokenizer.Token
	Value   vm_tokenizer.Token
	Line    int
}

func NewPopStatement() *PopStatement {
	return &PopStatement{
		Value:   vm_tokenizer.ZeroToken(),
		Segment: vm_tokenizer.ZeroToken(),
	}
}

func (a PopStatement) statement() {
}

func (a PopStatement) String() string {
	result := vm_tokenizer.POP
	if !a.Segment.IsZero() {
		result += " " + a.Segment.Literal
	}
	result += " " + a.Value.Literal
	return result
}

func (a PopStatement) LineNumber() int {
	return a.Line
}

type AddStatement struct {
	Line int
}

func (a AddStatement) statement() {
}

func (a AddStatement) String() string {
	return vm_tokenizer.ADD
}

func (a AddStatement) LineNumber() int {
	return a.Line
}

type SubStatement struct {
	Line int
}

func (a SubStatement) statement() {
}

func (a SubStatement) String() string {
	return vm_tokenizer.SUB
}

func (a SubStatement) LineNumber() int {
	return a.Line
}

type NeqStatement struct {
	Line int
}

func (a NeqStatement) statement() {
}

func (a NeqStatement) String() string {
	return vm_tokenizer.NEG
}

func (a NeqStatement) LineNumber() int {
	return a.Line
}

type EqStatement struct {
	Line int
}

func (a EqStatement) statement() {
}

func (a EqStatement) String() string {
	return vm_tokenizer.EQ
}

func (a EqStatement) LineNumber() int {
	return a.Line
}

type GtStatement struct {
	Line int
}

func (a GtStatement) statement() {
}

func (a GtStatement) String() string {
	return vm_tokenizer.GT
}

func (a GtStatement) LineNumber() int {
	return a.Line
}

type LtStatement struct {
	Line int
}

func (a LtStatement) statement() {
}

func (a LtStatement) String() string {
	return vm_tokenizer.LT
}

func (a LtStatement) LineNumber() int {
	return a.Line
}

type AndStatement struct {
	Line int
}

func (a AndStatement) statement() {
}

func (a AndStatement) String() string {
	return vm_tokenizer.AND
}

func (a AndStatement) LineNumber() int {
	return a.Line
}

type OrStatement struct {
	Line int
}

func (a OrStatement) statement() {
}

func (a OrStatement) String() string {
	return vm_tokenizer.OR
}

func (a OrStatement) LineNumber() int {
	return a.Line
}

type NotStatement struct {
	Line int
}

func (a NotStatement) statement() {
}

func (a NotStatement) String() string {
	return vm_tokenizer.NOT
}

func (a NotStatement) LineNumber() int {
	return a.Line
}

type LabelStatement struct {
	Value vm_tokenizer.Token
	Line  int
}

func (a LabelStatement) statement() {
}

func (a LabelStatement) String() string {
	return vm_tokenizer.LABEL + " " + a.Value.Literal
}

func (a LabelStatement) LineNumber() int {
	return a.Line
}

type GotoStatement struct {
	Value vm_tokenizer.Token
	Line  int
}

func NewGotoStatement(label string, line int) *GotoStatement {
	return &GotoStatement{
		Value: vm_tokenizer.Token{Type: vm_tokenizer.IDENT, Literal: label},
		Line:  line,
	}
}

func (a GotoStatement) statement() {
}

func (a GotoStatement) String() string {
	return vm_tokenizer.GOTO + " " + a.Value.Literal
}

func (a GotoStatement) LineNumber() int {
	return a.Line
}

type IfGotoStatement struct {
	Value vm_tokenizer.Token
	Line  int
}

func (a IfGotoStatement) statement() {
}

func (a IfGotoStatement) String() string {
	return vm_tokenizer.IFGOTO + " " + a.Value.Literal
}

func (a IfGotoStatement) LineNumber() int {
	return a.Line
}

type FunctionStatement struct {
	Name     vm_tokenizer.Token
	LocalNum vm_tokenizer.Token
	Line     int
}

func (a FunctionStatement) statement() {
}

func (a FunctionStatement) String() string {
	return vm_tokenizer.FUNCTION + " " + a.Name.Literal + " " + a.LocalNum.Literal
}

func (a FunctionStatement) LineNumber() int {
	return a.Line
}

type CallStatement struct {
	Name   vm_tokenizer.Token
	ArgNum vm_tokenizer.Token
	Line   int
}

func (a CallStatement) statement() {
}

func (a CallStatement) String() string {
	return vm_tokenizer.CALL + " " + a.Name.Literal + " " + a.ArgNum.Literal
}

func (a CallStatement) LineNumber() int {
	return a.Line
}

type ReturnStatement struct {
	Line int
}

func (a ReturnStatement) statement() {
}

func (a ReturnStatement) String() string {
	return vm_tokenizer.RETURN
}

func (a ReturnStatement) LineNumber() int {
	return a.Line
}
