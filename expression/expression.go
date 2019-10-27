package expression

import (
	"price_calculation/storage"
)

type Expression interface {
	Evaluate(store *storage.Storage) int
}

type ConstExpr struct {
	value int
}

func NewConstExpr(value int) *ConstExpr {
	return &ConstExpr{value: value}
}

func (e *ConstExpr) Value() int {
	return e.value
}

func (e *ConstExpr) Evaluate(store *storage.Storage) int {
	return e.value
}

type PlusExpr struct {
	left, right Expression
}

func NewPlusExpr(left Expression, right Expression) *PlusExpr {
	return &PlusExpr{left: left, right: right}
}

func (e *PlusExpr) Evaluate(store *storage.Storage) int {
	return e.left.Evaluate(store) + e.right.Evaluate(store)
}

type MinusExpr struct {
	left, right Expression
}

func NewMinusExpr(left Expression, right Expression) *MinusExpr {
	return &MinusExpr{left: left, right: right}
}

func (e *MinusExpr) Evaluate(store *storage.Storage) int {
	return e.left.Evaluate(store) - e.right.Evaluate(store)
}

type MulExpr struct {
	left, right Expression
}

func NewMulExpr(left Expression, right Expression) *MulExpr {
	return &MulExpr{left: left, right: right}
}

func (e *MulExpr) Evaluate(store *storage.Storage) int {
	return e.left.Evaluate(store) * e.right.Evaluate(store)
}

type DivExpr struct {
	left, right Expression
}

func NewDivExpr(left Expression, right Expression) *DivExpr {
	return &DivExpr{left: left, right: right}
}

func (e *DivExpr) Evaluate(store *storage.Storage) int {
	return e.left.Evaluate(store) / e.right.Evaluate(store)
}

type EqExpr struct {
	left, right Expression
}

func NewEqExpr(left Expression, right Expression) *EqExpr {
	return &EqExpr{left: left, right: right}
}

func (e *EqExpr) Evaluate(store *storage.Storage) int {
	if e.left.Evaluate(store) == e.right.Evaluate(store) {
		return 1
	}
	return 0
}

type LtExpr struct {
	left, right Expression
}

func NewLtExpr(left Expression, right Expression) *LtExpr {
	return &LtExpr{left: left, right: right}
}

func (e *LtExpr) Evaluate(store *storage.Storage) int {
	if e.left.Evaluate(store) < e.right.Evaluate(store) {
		return 1
	}
	return 0
}

type LogicExpr struct {
	logic, left, right Expression
}

func NewLogicExpr(logic Expression, left Expression, right Expression) *LogicExpr {
	return &LogicExpr{logic: logic, left: left, right: right}
}

func (e *LogicExpr) Evaluate(store *storage.Storage) int {
	if e.logic.Evaluate(store) == 1 {
		return e.left.Evaluate(store)
	}
	return e.right.Evaluate(store)
}

type VariableExpr struct {
	key string
}

func NewVariableExpr(key string) *VariableExpr {
	return &VariableExpr{key: key}
}

func (e *VariableExpr) Evaluate(store *storage.Storage) int {
	return store.Get(e.key)
}
