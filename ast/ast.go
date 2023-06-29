package ast

import (
	"bytes"
	"github.com/savannahostrowski/gossamer/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type VariableStatement struct {
	Name  *Identifier
	Value Expression
}

func (vs *VariableStatement) statementNode() {}
// func (vs *VariableStatement) TokenLiteral() string {return vs.Value.String()}

type Identifier struct {
	Token token.Token //IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// type ExpressionStatement struct {
// 	Token      token.Token
// 	Expression Expression
// }

// func (es *ExpressionStatement) statementNode()       {}
// func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// func (p *Program) String() string {
// 	var out bytes.Buffer

// 	for _, s := range p.Statements {
// 		out.WriteString(s.String())
// 	}
// 	return out.String()
// }

// func (ls *ReturnStatement) String() string {
// 	var out bytes.Buffer

// 	out.WriteString(rs.TokenLiteral() + " ")

// 	if (rs.ReturnValue != nil) {
// 		out.WriteString(rs.ReturnValue.String())
// 	}
// }