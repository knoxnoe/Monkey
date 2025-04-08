package ast

import "monkey/src/token"

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (l *ReturnStatement) statementNode()       {}
func (l *ReturnStatement) TokenLiteral() string { return l.Token.Literal }
