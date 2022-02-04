package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/VKCOM/php-parser/pkg/ast"
	"github.com/VKCOM/php-parser/pkg/conf"
	phperrors "github.com/VKCOM/php-parser/pkg/errors"
	"github.com/VKCOM/php-parser/pkg/parser"
	"github.com/VKCOM/php-parser/pkg/version"
	"io/ioutil"
	"os"
)

func main() {
	src := getSource()
	ast, err := ParseFile(src)
	if err != nil {
		panic(err)
	}
	a := &AstReader{}
	ast.Accept(a)

}
func getSource() []byte {
	// get file path from os.args
	fpath := os.Args[1]
	f, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read file content
	filecontents, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return filecontents

}

//  from -> https://github.com/VKCOM/noverify/blob/master/src/php/parseutil/parseutil.go
func ParseFile(code []byte) (*ast.Root, error) {
	phpVersion, err := version.New("7.4")
	if err != nil {
		return nil, err
	}

	var parserErrors []*phperrors.Error
	rootNode, err := parser.Parse(code, conf.Config{
		Version: phpVersion,
		ErrorHandlerFunc: func(e *phperrors.Error) {
			parserErrors = append(parserErrors, e)
		},
	})
	if err != nil {
		return nil, err
	}
	if len(parserErrors) != 0 {
		return nil, errors.New(parserErrors[0].String())
	}

	if rootNode == nil {
		return nil, fmt.Errorf("file has incorrect syntax and cannot be parsed")
	}

	return rootNode.(*ast.Root), nil
}

type AstReader struct {
	label      string
	children   []AstReader
	attributes map[string]interface{}
}

func (a *AstReader) Root(n *ast.Root) {
	fmt.Println("Root")

	a.label = "Root"
	a.attributes = make(map[string]interface{})
	a.children = make([]AstReader, len(n.Stmts))
	for _, stmt := range n.Stmts {
		if stmt != nil {

			child := AstReader{}
			child.children = make([]AstReader, 0)
			child.attributes = make(map[string]interface{})

			stmt.Accept(&child)
			a.children = append(a.children, child)
		} else {
			return
		}
	}
	fmt.Printf("%+v", a)
}
func (a *AstReader) Nullable(n *ast.Nullable) {}
func (a *AstReader) Parameter(n *ast.Parameter) {
	n.Var.Accept(a)
	if n.DefaultValue != nil {
		n.DefaultValue.Accept(a)
	}
}
func (a *AstReader) Identifier(n *ast.Identifier) {
	fmt.Println("Identifier")
	value := strconv.Quote(string(n.Value))
	fmt.Printf("Identifier: %s\n\n", value)
	a.attributes["name"] = value
}
func (a *AstReader) Argument(n *ast.Argument)             {}
func (a *AstReader) MatchArm(n *ast.MatchArm)             {}
func (a *AstReader) Union(n *ast.Union)                   {}
func (a *AstReader) Attribute(n *ast.Attribute)           {}
func (a *AstReader) AttributeGroup(n *ast.AttributeGroup) {}

func (a *AstReader) StmtBreak(n *ast.StmtBreak)                   {}
func (a *AstReader) StmtCase(n *ast.StmtCase)                     {}
func (a *AstReader) StmtCatch(n *ast.StmtCatch)                   {}
func (a *AstReader) StmtEnum(n *ast.StmtEnum)                     {}
func (a *AstReader) EnumCase(n *ast.EnumCase)                     {}
func (a *AstReader) StmtClass(n *ast.StmtClass)                   {}
func (a *AstReader) StmtClassConstList(n *ast.StmtClassConstList) {}
func (a *AstReader) StmtClassMethod(n *ast.StmtClassMethod)       {}
func (a *AstReader) StmtConstList(n *ast.StmtConstList)           {}
func (a *AstReader) StmtConstant(n *ast.StmtConstant)             {}
func (a *AstReader) StmtContinue(n *ast.StmtContinue)             {}
func (a *AstReader) StmtDeclare(n *ast.StmtDeclare)               {}
func (a *AstReader) StmtDefault(n *ast.StmtDefault) {
	fmt.Println("StmtDefault")
}
func (a *AstReader) StmtDo(n *ast.StmtDo)                 {}
func (a *AstReader) StmtEcho(n *ast.StmtEcho)             {}
func (a *AstReader) StmtElse(n *ast.StmtElse)             {}
func (a *AstReader) StmtElseIf(n *ast.StmtElseIf)         {}
func (a *AstReader) StmtExpression(n *ast.StmtExpression) {}
func (a *AstReader) StmtFinally(n *ast.StmtFinally)       {}
func (a *AstReader) StmtFor(n *ast.StmtFor)               {}
func (a *AstReader) StmtForeach(n *ast.StmtForeach)       {}
func (a *AstReader) StmtFunction(n *ast.StmtFunction) {
	if n == nil {
		return
	}
	a.label = n.FunctionTkn.ID.String()
	n.Name.Accept(a)
	parameters := make([]interface{}, 0)
	for _, param := range n.Params {
		if param != nil {
			paramReader := AstReader{}
			paramReader.children = make([]AstReader, 0)
			paramReader.attributes = make(map[string]interface{})
			param.Accept(&paramReader)

			parameters = append(parameters, paramReader)
		}

	}
	a.attributes["parameters"] = parameters

}

func (a *AstReader) StmtGlobal(n *ast.StmtGlobal)             {}
func (a *AstReader) StmtGoto(n *ast.StmtGoto)                 {}
func (a *AstReader) StmtHaltCompiler(n *ast.StmtHaltCompiler) {}
func (a *AstReader) StmtIf(n *ast.StmtIf)                     {}
func (a *AstReader) StmtInlineHtml(n *ast.StmtInlineHtml)     {}
func (a *AstReader) StmtInterface(n *ast.StmtInterface)       {}
func (a *AstReader) StmtLabel(n *ast.StmtLabel)               {}
func (a *AstReader) StmtNamespace(n *ast.StmtNamespace)       {}
func (a *AstReader) StmtNop(n *ast.StmtNop)                   {}
func (a *AstReader) StmtProperty(n *ast.StmtProperty)         {}
func (a *AstReader) StmtPropertyList(n *ast.StmtPropertyList) {}
func (a *AstReader) StmtReturn(n *ast.StmtReturn)             {}
func (a *AstReader) StmtStatic(n *ast.StmtStatic)             {}
func (a *AstReader) StmtStaticVar(n *ast.StmtStaticVar) {
}
func (a *AstReader) StmtStmtList(n *ast.StmtStmtList)                     {}
func (a *AstReader) StmtSwitch(n *ast.StmtSwitch)                         {}
func (a *AstReader) StmtThrow(n *ast.StmtThrow)                           {}
func (a *AstReader) StmtTrait(n *ast.StmtTrait)                           {}
func (a *AstReader) StmtTraitUse(n *ast.StmtTraitUse)                     {}
func (a *AstReader) StmtTraitUseAlias(n *ast.StmtTraitUseAlias)           {}
func (a *AstReader) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {}
func (a *AstReader) StmtTry(n *ast.StmtTry)                               {}
func (a *AstReader) StmtUnset(n *ast.StmtUnset)                           {}
func (a *AstReader) StmtUse(n *ast.StmtUseList)                           {}
func (a *AstReader) StmtGroupUse(n *ast.StmtGroupUseList)                 {}
func (a *AstReader) StmtUseDeclaration(n *ast.StmtUse)                    {}
func (a *AstReader) StmtWhile(n *ast.StmtWhile)                           {}

func (a *AstReader) ExprArray(n *ast.ExprArray)                 {}
func (a *AstReader) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {}
func (a *AstReader) ExprArrayItem(n *ast.ExprArrayItem)         {}
func (a *AstReader) ExprArrowFunction(n *ast.ExprArrowFunction) {}
func (a *AstReader) ExprBrackets(n *ast.ExprBrackets)           {}
func (a *AstReader) ExprBitwiseNot(n *ast.ExprBitwiseNot)       {}
func (a *AstReader) ExprBooleanNot(n *ast.ExprBooleanNot)       {}
func (a *AstReader) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	fmt.Println("ExprClassConstFetch")
}
func (a *AstReader) ExprClone(n *ast.ExprClone)           {}
func (a *AstReader) ExprClosure(n *ast.ExprClosure)       {}
func (a *AstReader) ExprClosureUse(n *ast.ExprClosureUse) {}
func (a *AstReader) ExprConstFetch(n *ast.ExprConstFetch) {
	fmt.Println("ExprConstFetch")
}
func (a *AstReader) ExprEmpty(n *ast.ExprEmpty)                                 {}
func (a *AstReader) ExprErrorSuppress(n *ast.ExprErrorSuppress)                 {}
func (a *AstReader) ExprEval(n *ast.ExprEval)                                   {}
func (a *AstReader) ExprExit(n *ast.ExprExit)                                   {}
func (a *AstReader) ExprFunctionCall(n *ast.ExprFunctionCall)                   {}
func (a *AstReader) ExprInclude(n *ast.ExprInclude)                             {}
func (a *AstReader) ExprIncludeOnce(n *ast.ExprIncludeOnce)                     {}
func (a *AstReader) ExprInstanceOf(n *ast.ExprInstanceOf)                       {}
func (a *AstReader) ExprIsset(n *ast.ExprIsset)                                 {}
func (a *AstReader) ExprList(n *ast.ExprList)                                   {}
func (a *AstReader) ExprMethodCall(n *ast.ExprMethodCall)                       {}
func (a *AstReader) ExprNullsafeMethodCall(n *ast.ExprNullsafeMethodCall)       {}
func (a *AstReader) ExprMatch(n *ast.ExprMatch)                                 {}
func (a *AstReader) ExprNew(n *ast.ExprNew)                                     {}
func (a *AstReader) ExprPostDec(n *ast.ExprPostDec)                             {}
func (a *AstReader) ExprPostInc(n *ast.ExprPostInc)                             {}
func (a *AstReader) ExprPreDec(n *ast.ExprPreDec)                               {}
func (a *AstReader) ExprPreInc(n *ast.ExprPreInc)                               {}
func (a *AstReader) ExprPrint(n *ast.ExprPrint)                                 {}
func (a *AstReader) ExprPropertyFetch(n *ast.ExprPropertyFetch)                 {}
func (a *AstReader) ExprNullsafePropertyFetch(n *ast.ExprNullsafePropertyFetch) {}
func (a *AstReader) ExprRequire(n *ast.ExprRequire)                             {}
func (a *AstReader) ExprRequireOnce(n *ast.ExprRequireOnce)                     {}
func (a *AstReader) ExprShellExec(n *ast.ExprShellExec)                         {}
func (a *AstReader) ExprStaticCall(n *ast.ExprStaticCall)                       {}
func (a *AstReader) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch)     {}
func (a *AstReader) ExprTernary(n *ast.ExprTernary)                             {}
func (a *AstReader) ExprThrow(n *ast.ExprThrow)                                 {}
func (a *AstReader) ExprUnaryMinus(n *ast.ExprUnaryMinus)                       {}
func (a *AstReader) ExprUnaryPlus(n *ast.ExprUnaryPlus)                         {}
func (a *AstReader) ExprVariable(n *ast.ExprVariable) {
	n.Name.Accept(a)
}
func (a *AstReader) ExprYield(n *ast.ExprYield)         {}
func (a *AstReader) ExprYieldFrom(n *ast.ExprYieldFrom) {}

func (a *AstReader) ExprAssign(n *ast.ExprAssign)                     {}
func (a *AstReader) ExprAssignReference(n *ast.ExprAssignReference)   {}
func (a *AstReader) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {}
func (a *AstReader) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr)   {}
func (a *AstReader) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {}
func (a *AstReader) ExprAssignCoalesce(n *ast.ExprAssignCoalesce)     {}
func (a *AstReader) ExprAssignConcat(n *ast.ExprAssignConcat)         {}
func (a *AstReader) ExprAssignDiv(n *ast.ExprAssignDiv)               {}
func (a *AstReader) ExprAssignMinus(n *ast.ExprAssignMinus)           {}
func (a *AstReader) ExprAssignMod(n *ast.ExprAssignMod)               {}
func (a *AstReader) ExprAssignMul(n *ast.ExprAssignMul)               {}
func (a *AstReader) ExprAssignPlus(n *ast.ExprAssignPlus)             {}
func (a *AstReader) ExprAssignPow(n *ast.ExprAssignPow)               {}
func (a *AstReader) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft)   {}
func (a *AstReader) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {}

func (a *AstReader) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd)         {}
func (a *AstReader) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr)           {}
func (a *AstReader) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor)         {}
func (a *AstReader) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd)         {}
func (a *AstReader) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr)           {}
func (a *AstReader) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce)             {}
func (a *AstReader) ExprBinaryConcat(n *ast.ExprBinaryConcat)                 {}
func (a *AstReader) ExprBinaryDiv(n *ast.ExprBinaryDiv)                       {}
func (a *AstReader) ExprBinaryEqual(n *ast.ExprBinaryEqual)                   {}
func (a *AstReader) ExprBinaryGreater(n *ast.ExprBinaryGreater)               {}
func (a *AstReader) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {}
func (a *AstReader) ExprBinaryIdentical(n *ast.ExprBinaryIdentical)           {}
func (a *AstReader) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd)         {}
func (a *AstReader) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr)           {}
func (a *AstReader) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor)         {}
func (a *AstReader) ExprBinaryMinus(n *ast.ExprBinaryMinus)                   {}
func (a *AstReader) ExprBinaryMod(n *ast.ExprBinaryMod)                       {}
func (a *AstReader) ExprBinaryMul(n *ast.ExprBinaryMul)                       {}
func (a *AstReader) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual)             {}
func (a *AstReader) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical)     {}
func (a *AstReader) ExprBinaryPlus(n *ast.ExprBinaryPlus)                     {}
func (a *AstReader) ExprBinaryPow(n *ast.ExprBinaryPow)                       {}
func (a *AstReader) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft)           {}
func (a *AstReader) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight)         {}
func (a *AstReader) ExprBinarySmaller(n *ast.ExprBinarySmaller)               {}
func (a *AstReader) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {}
func (a *AstReader) ExprBinarySpaceship(n *ast.ExprBinarySpaceship)           {}

func (a *AstReader) ExprCastArray(n *ast.ExprCastArray)   {}
func (a *AstReader) ExprCastBool(n *ast.ExprCastBool)     {}
func (a *AstReader) ExprCastDouble(n *ast.ExprCastDouble) {}
func (a *AstReader) ExprCastInt(n *ast.ExprCastInt)       {}
func (a *AstReader) ExprCastObject(n *ast.ExprCastObject) {}
func (a *AstReader) ExprCastString(n *ast.ExprCastString) {}
func (a *AstReader) ExprCastUnset(n *ast.ExprCastUnset)   {}

func (a *AstReader) ScalarDnumber(n *ast.ScalarDnumber)                               {}
func (a *AstReader) ScalarEncapsed(n *ast.ScalarEncapsed)                             {}
func (a *AstReader) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart)         {}
func (a *AstReader) ScalarEncapsedStringVar(n *ast.ScalarEncapsedStringVar)           {}
func (a *AstReader) ScalarEncapsedStringBrackets(n *ast.ScalarEncapsedStringBrackets) {}
func (a *AstReader) ScalarHeredoc(n *ast.ScalarHeredoc)                               {}
func (a *AstReader) ScalarLnumber(n *ast.ScalarLnumber)                               {}
func (a *AstReader) ScalarMagicConstant(n *ast.ScalarMagicConstant)                   {}
func (a *AstReader) ScalarString(n *ast.ScalarString) {
	fmt.Println(string(n.Value))
}

func (a *AstReader) NameName(n *ast.Name) {
	fmt.Println("HELLO")
}
func (a *AstReader) NameFullyQualified(n *ast.NameFullyQualified) {}
func (a *AstReader) NameRelative(n *ast.NameRelative)             {}
func (a *AstReader) NameNamePart(n *ast.NamePart) {
	fmt.Println("NameNamePart")
}
