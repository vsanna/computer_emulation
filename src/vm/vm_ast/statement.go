package vm_ast

import (
	"computer_emulation/src/vm/vm_tokenizer"
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

func NewPushStatement() *PushStatement {
	return &PushStatement{
		Value:   vm_tokenizer.ZeroToken(),
		Segment: vm_tokenizer.ZeroToken(),
	}
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
