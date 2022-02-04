package main

import (
	"errors"
	"fmt"

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
	a := &exploreReader{}
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

type exploreReader struct{}

func (a *exploreReader) Root(n *ast.Root) {
	fmt.Println("Root")
	fmt.Println("")
	for _, stmt := range n.Stmts {
		if stmt != nil {
			child := exploreReader{}
			stmt.Accept(&child)
		} else {
			return
		}
	}
	fmt.Println("endof Root\n\t")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) Nullable(n *ast.Nullable) {
	fmt.Println("Nullable")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) Parameter(n *ast.Parameter) {
	fmt.Println("Parameter")
	if n.AttrGroups != nil {
		for _, attrGroup := range n.AttrGroups {
			attrGroup.Accept(a)
		}
	}
	n.Var.Accept(a)
	if n.DefaultValue != nil {
		n.DefaultValue.Accept(a)
	}

	fmt.Println("end of Parameter\n\t")

}
func (a *exploreReader) Identifier(n *ast.Identifier) {
	fmt.Println("\tIdentifier")
	fmt.Println("\t", string(n.Value))
	fmt.Println("\tend of Identifier")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) Argument(n *ast.Argument) {
	fmt.Println("Argument")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) MatchArm(n *ast.MatchArm) {
	fmt.Println("MatchArm")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) Union(n *ast.Union) {
	fmt.Println("Union")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) Attribute(n *ast.Attribute) {

	fmt.Println("Attribute")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) AttributeGroup(n *ast.AttributeGroup) {
	fmt.Println("AttributeGroup")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) StmtBreak(n *ast.StmtBreak) {

	fmt.Println("StmtBreak")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtCase(n *ast.StmtCase) {

	fmt.Println("StmtCase")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtCatch(n *ast.StmtCatch) {

	fmt.Println("StmtCatch")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtEnum(n *ast.StmtEnum) {

	fmt.Println("StmtEnum")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) EnumCase(n *ast.EnumCase) {

	fmt.Println("EnumCase")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) StmtClass(n *ast.StmtClass) {

	fmt.Println("StmtClass")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtClassConstList(n *ast.StmtClassConstList) {

	fmt.Println("StmtClassConstList")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtClassMethod(n *ast.StmtClassMethod) {

	fmt.Println("StmtClassMethod")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtConstList(n *ast.StmtConstList) {

	fmt.Println("StmtConstList")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtConstant(n *ast.StmtConstant) {

	fmt.Println("StmtConstant")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtContinue(n *ast.StmtContinue) {

	fmt.Println("StmtContinue")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtDeclare(n *ast.StmtDeclare) {

	fmt.Println("StmtDeclare")
	// fmt.Printf("%+v", n)
}

func (a *exploreReader) StmtDefault(n *ast.StmtDefault) {
	fmt.Println("StmtDefault")
	// fmt.Printf("%+v", n)
}
func (a *exploreReader) StmtDo(n *ast.StmtDo) {
	fmt.Println("StmtDo")
}
func (a *exploreReader) StmtEcho(n *ast.StmtEcho) {
	fmt.Println("StmtEcho")
	for _, ex := range n.Exprs {
		ex.Accept(a)
	}
	fmt.Println("end of StmtEcho\n\t")
}
func (a *exploreReader) StmtElse(n *ast.StmtElse) {

	fmt.Println("StmtElse")
}
func (a *exploreReader) StmtElseIf(n *ast.StmtElseIf) {

	fmt.Println("StmtElseIf")
}
func (a *exploreReader) StmtExpression(n *ast.StmtExpression) {
	fmt.Println("StmtExpression")
	n.Expr.Accept(a)
	fmt.Println("End of StmtExpression\n\t")
}
func (a *exploreReader) StmtFinally(n *ast.StmtFinally) {

	fmt.Println("StmtFinally")
}
func (a *exploreReader) StmtFor(n *ast.StmtFor) {

	fmt.Println("StmtFor")
}
func (a *exploreReader) StmtForeach(n *ast.StmtForeach) {

	fmt.Println("StmtForeach")
}
func (a *exploreReader) StmtFunction(n *ast.StmtFunction) {
	fmt.Println("StmtFunction")
	// fmt.Printf("%+v", n)
	if n == nil {
		return
	}
	n.Name.Accept(a)
	for _, param := range n.Params {
		if param != nil {
			paramReader := exploreReader{}
			param.Accept(&paramReader)
		}
	}
}

func (a *exploreReader) StmtGlobal(n *ast.StmtGlobal) {

	fmt.Println("StmtGlobal")
}
func (a *exploreReader) StmtGoto(n *ast.StmtGoto) {

	fmt.Println("StmtGoto")
}
func (a *exploreReader) StmtHaltCompiler(n *ast.StmtHaltCompiler) {

	fmt.Println("StmtHaltCompiler")
}
func (a *exploreReader) StmtIf(n *ast.StmtIf) {

	fmt.Println("StmtIf")
}
func (a *exploreReader) StmtInlineHtml(n *ast.StmtInlineHtml) {

	fmt.Println("StmtInlineHtml")
}
func (a *exploreReader) StmtInterface(n *ast.StmtInterface) {

	fmt.Println("StmtInterface")
}
func (a *exploreReader) StmtLabel(n *ast.StmtLabel) {

	fmt.Println("StmtLabel")
}
func (a *exploreReader) StmtNamespace(n *ast.StmtNamespace) {

	fmt.Println("StmtNamespace")
}
func (a *exploreReader) StmtNop(n *ast.StmtNop) {

	fmt.Println("StmtNop")
}
func (a *exploreReader) StmtProperty(n *ast.StmtProperty) {

	fmt.Println("StmtProperty")
}
func (a *exploreReader) StmtPropertyList(n *ast.StmtPropertyList) {

	fmt.Println("StmtPropertyList")
}
func (a *exploreReader) StmtReturn(n *ast.StmtReturn) {

	fmt.Println("StmtReturn")
}
func (a *exploreReader) StmtStatic(n *ast.StmtStatic) {

	fmt.Println("StmtStatic")
}
func (a *exploreReader) StmtStaticVar(n *ast.StmtStaticVar) {

	fmt.Println("StmtStaticVar")
}
func (a *exploreReader) StmtStmtList(n *ast.StmtStmtList) {

	fmt.Println("StmtStmtList")
}
func (a *exploreReader) StmtSwitch(n *ast.StmtSwitch) {

	fmt.Println("StmtSwitch")
}
func (a *exploreReader) StmtThrow(n *ast.StmtThrow) {

	fmt.Println("StmtThrow")
}
func (a *exploreReader) StmtTrait(n *ast.StmtTrait) {

	fmt.Println("StmtTrait")
}
func (a *exploreReader) StmtTraitUse(n *ast.StmtTraitUse) {

	fmt.Println("StmtTraitUse")
}
func (a *exploreReader) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {

	fmt.Println("StmtTraitUseAlias")
}
func (a *exploreReader) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {

	fmt.Println("StmtTraitUsePrecedence")
}
func (a *exploreReader) StmtTry(n *ast.StmtTry) {

	fmt.Println("StmtTry")
}
func (a *exploreReader) StmtUnset(n *ast.StmtUnset) {

	fmt.Println("StmtUnset")
}
func (a *exploreReader) StmtUse(n *ast.StmtUseList) {

	fmt.Println("StmtUse")
}
func (a *exploreReader) StmtGroupUse(n *ast.StmtGroupUseList) {

	fmt.Println("StmtGroupUse")
}
func (a *exploreReader) StmtUseDeclaration(n *ast.StmtUse) {

	fmt.Println("StmtUseDeclaration")
}
func (a *exploreReader) StmtWhile(n *ast.StmtWhile) {

	fmt.Println("StmtWhile")
}

func (a *exploreReader) ExprArray(n *ast.ExprArray) {

	fmt.Println("ExprArray")
}
func (a *exploreReader) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {

	fmt.Println("ExprArrayDimFetch")
}
func (a *exploreReader) ExprArrayItem(n *ast.ExprArrayItem) {

	fmt.Println("ExprArrayItem")
}
func (a *exploreReader) ExprArrowFunction(n *ast.ExprArrowFunction) {

	fmt.Println("ExprArrowFunction")
}
func (a *exploreReader) ExprBrackets(n *ast.ExprBrackets) {

	fmt.Println("ExprBrackets")
}
func (a *exploreReader) ExprBitwiseNot(n *ast.ExprBitwiseNot) {

	fmt.Println("ExprBitwiseNot")
}
func (a *exploreReader) ExprBooleanNot(n *ast.ExprBooleanNot) {

	fmt.Println("ExprBooleanNot")
}
func (a *exploreReader) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	fmt.Println("ExprClassConstFetch")
}
func (a *exploreReader) ExprClone(n *ast.ExprClone) {

	fmt.Println("ExprClone")
}
func (a *exploreReader) ExprClosure(n *ast.ExprClosure) {

	fmt.Println("ExprClosure")
}
func (a *exploreReader) ExprClosureUse(n *ast.ExprClosureUse) {

	fmt.Println("ExprClosureUse")
}
func (a *exploreReader) ExprConstFetch(n *ast.ExprConstFetch) {
	fmt.Println("ExprConstFetch")
}
func (a *exploreReader) ExprEmpty(n *ast.ExprEmpty) {

	fmt.Println("ExprEmpty")
}
func (a *exploreReader) ExprErrorSuppress(n *ast.ExprErrorSuppress) {

	fmt.Println("ExprErrorSuppress")
}
func (a *exploreReader) ExprEval(n *ast.ExprEval) {

	fmt.Println("ExprEval")
}
func (a *exploreReader) ExprExit(n *ast.ExprExit) {

	fmt.Println("ExprExit")
}
func (a *exploreReader) ExprFunctionCall(n *ast.ExprFunctionCall) {

	fmt.Println("ExprFunctionCall")
}
func (a *exploreReader) ExprInclude(n *ast.ExprInclude) {

	fmt.Println("ExprInclude")
}
func (a *exploreReader) ExprIncludeOnce(n *ast.ExprIncludeOnce) {

	fmt.Println("ExprIncludeOnce")
}
func (a *exploreReader) ExprInstanceOf(n *ast.ExprInstanceOf) {

	fmt.Println("ExprInstanceOf")
}
func (a *exploreReader) ExprIsset(n *ast.ExprIsset) {

	fmt.Println("ExprIsset")
}
func (a *exploreReader) ExprList(n *ast.ExprList) {

	fmt.Println("ExprList")
}
func (a *exploreReader) ExprMethodCall(n *ast.ExprMethodCall) {

	fmt.Println("ExprMethodCall")
}
func (a *exploreReader) ExprNullsafeMethodCall(n *ast.ExprNullsafeMethodCall) {

	fmt.Println("ExprNullsafeMethodCall")
}
func (a *exploreReader) ExprMatch(n *ast.ExprMatch) {

	fmt.Println("ExprMatch")
}
func (a *exploreReader) ExprNew(n *ast.ExprNew) {

	fmt.Println("ExprNew")
}
func (a *exploreReader) ExprPostDec(n *ast.ExprPostDec) {

	fmt.Println("ExprPostDec")
}
func (a *exploreReader) ExprPostInc(n *ast.ExprPostInc) {

	fmt.Println("ExprPostInc")
}
func (a *exploreReader) ExprPreDec(n *ast.ExprPreDec) {

	fmt.Println("ExprPreDec")
}
func (a *exploreReader) ExprPreInc(n *ast.ExprPreInc) {

	fmt.Println("ExprPreInc")
}
func (a *exploreReader) ExprPrint(n *ast.ExprPrint) {

	fmt.Println("ExprPrint")
}
func (a *exploreReader) ExprPropertyFetch(n *ast.ExprPropertyFetch) {

	fmt.Println("ExprPropertyFetch")
}
func (a *exploreReader) ExprNullsafePropertyFetch(n *ast.ExprNullsafePropertyFetch) {

	fmt.Println("ExprNullsafePropertyFetch")
}
func (a *exploreReader) ExprRequire(n *ast.ExprRequire) {

	fmt.Println("ExprRequire")
}
func (a *exploreReader) ExprRequireOnce(n *ast.ExprRequireOnce) {

	fmt.Println("ExprRequireOnce")
}
func (a *exploreReader) ExprShellExec(n *ast.ExprShellExec) {

	fmt.Println("ExprShellExec")
}
func (a *exploreReader) ExprStaticCall(n *ast.ExprStaticCall) {

	fmt.Println("ExprStaticCall")
}
func (a *exploreReader) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {

	fmt.Println("ExprStaticPropertyFetch")
}
func (a *exploreReader) ExprTernary(n *ast.ExprTernary) {

	fmt.Println("ExprTernary")
}
func (a *exploreReader) ExprThrow(n *ast.ExprThrow) {

	fmt.Println("ExprThrow")
}
func (a *exploreReader) ExprUnaryMinus(n *ast.ExprUnaryMinus) {

	fmt.Println("ExprUnaryMinus")
}
func (a *exploreReader) ExprUnaryPlus(n *ast.ExprUnaryPlus) {

	fmt.Println("ExprUnaryPlus")
}
func (a *exploreReader) ExprVariable(n *ast.ExprVariable) {

	fmt.Println("ExprVariable")

	n.Name.Accept(a)
}
func (a *exploreReader) ExprYield(n *ast.ExprYield) {

	fmt.Println("ExprYield")
}
func (a *exploreReader) ExprYieldFrom(n *ast.ExprYieldFrom) {

	fmt.Println("ExprYieldFrom")
}

func (a *exploreReader) ExprAssign(n *ast.ExprAssign) {

	fmt.Println("ExprAssign")
}
func (a *exploreReader) ExprAssignReference(n *ast.ExprAssignReference) {

	fmt.Println("ExprAssignReference")
}
func (a *exploreReader) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {

	fmt.Println("ExprAssignBitwiseAnd")
}
func (a *exploreReader) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {

	fmt.Println("ExprAssignBitwiseOr")
}
func (a *exploreReader) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {

	fmt.Println("ExprAssignBitwiseXor")
}
func (a *exploreReader) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {

	fmt.Println("ExprAssignCoalesce")
}
func (a *exploreReader) ExprAssignConcat(n *ast.ExprAssignConcat) {

	fmt.Println("ExprAssignConcat")
}
func (a *exploreReader) ExprAssignDiv(n *ast.ExprAssignDiv) {

	fmt.Println("ExprAssignDiv")
}
func (a *exploreReader) ExprAssignMinus(n *ast.ExprAssignMinus) {

	fmt.Println("ExprAssignMinus")
}
func (a *exploreReader) ExprAssignMod(n *ast.ExprAssignMod) {

	fmt.Println("ExprAssignMod")
}
func (a *exploreReader) ExprAssignMul(n *ast.ExprAssignMul) {

	fmt.Println("ExprAssignMul")
}
func (a *exploreReader) ExprAssignPlus(n *ast.ExprAssignPlus) {

	fmt.Println("ExprAssignPlus")
}
func (a *exploreReader) ExprAssignPow(n *ast.ExprAssignPow) {

	fmt.Println("ExprAssignPow")
}
func (a *exploreReader) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {

	fmt.Println("ExprAssignShiftLeft")
}
func (a *exploreReader) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {

	fmt.Println("ExprAssignShiftRight")
}

func (a *exploreReader) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {

	fmt.Println("ExprBinaryBitwiseAnd")
}
func (a *exploreReader) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {

	fmt.Println("ExprBinaryBitwiseOr")
}
func (a *exploreReader) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {

	fmt.Println("ExprBinaryBitwiseXor")
}
func (a *exploreReader) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {

	fmt.Println("ExprBinaryBooleanAnd")
}
func (a *exploreReader) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {

	fmt.Println("ExprBinaryBooleanOr")
}
func (a *exploreReader) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {

	fmt.Println("ExprBinaryCoalesce")
}
func (a *exploreReader) ExprBinaryConcat(n *ast.ExprBinaryConcat) {

	fmt.Println("ExprBinaryConcat")
}
func (a *exploreReader) ExprBinaryDiv(n *ast.ExprBinaryDiv) {

	fmt.Println("ExprBinaryDiv")
}
func (a *exploreReader) ExprBinaryEqual(n *ast.ExprBinaryEqual) {

	fmt.Println("ExprBinaryEqual")
}
func (a *exploreReader) ExprBinaryGreater(n *ast.ExprBinaryGreater) {

	fmt.Println("ExprBinaryGreater")
}
func (a *exploreReader) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {

	fmt.Println("ExprBinaryGreaterOrEqual")
}
func (a *exploreReader) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {

	fmt.Println("ExprBinaryIdentical")
}
func (a *exploreReader) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {

	fmt.Println("ExprBinaryLogicalAnd")
}
func (a *exploreReader) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {

	fmt.Println("ExprBinaryLogicalOr")
}
func (a *exploreReader) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {

	fmt.Println("ExprBinaryLogicalXor")
}
func (a *exploreReader) ExprBinaryMinus(n *ast.ExprBinaryMinus) {

	fmt.Println("ExprBinaryMinus")
}
func (a *exploreReader) ExprBinaryMod(n *ast.ExprBinaryMod) {

	fmt.Println("ExprBinaryMod")
}
func (a *exploreReader) ExprBinaryMul(n *ast.ExprBinaryMul) {

	fmt.Println("ExprBinaryMul")
}
func (a *exploreReader) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {

	fmt.Println("ExprBinaryNotEqual")
}
func (a *exploreReader) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {

	fmt.Println("ExprBinaryNotIdentical")
}
func (a *exploreReader) ExprBinaryPlus(n *ast.ExprBinaryPlus) {

	fmt.Println("ExprBinaryPlus")
}
func (a *exploreReader) ExprBinaryPow(n *ast.ExprBinaryPow) {

	fmt.Println("ExprBinaryPow")
}
func (a *exploreReader) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {

	fmt.Println("ExprBinaryShiftLeft")
}
func (a *exploreReader) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {

	fmt.Println("ExprBinaryShiftRight")
}
func (a *exploreReader) ExprBinarySmaller(n *ast.ExprBinarySmaller) {

	fmt.Println("ExprBinarySmaller")
}
func (a *exploreReader) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {

	fmt.Println("ExprBinarySmallerOrEqual")
}
func (a *exploreReader) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {

	fmt.Println("ExprBinarySpaceship")
}

func (a *exploreReader) ExprCastArray(n *ast.ExprCastArray) {

	fmt.Println("ExprCastArray")
}
func (a *exploreReader) ExprCastBool(n *ast.ExprCastBool) {

	fmt.Println("ExprCastBool")
}
func (a *exploreReader) ExprCastDouble(n *ast.ExprCastDouble) {

	fmt.Println("ExprCastDouble")
}
func (a *exploreReader) ExprCastInt(n *ast.ExprCastInt) {

	fmt.Println("ExprCastInt")
}
func (a *exploreReader) ExprCastObject(n *ast.ExprCastObject) {

	fmt.Println("ExprCastObject")
}
func (a *exploreReader) ExprCastString(n *ast.ExprCastString) {

	fmt.Println("ExprCastString")
}
func (a *exploreReader) ExprCastUnset(n *ast.ExprCastUnset) {

	fmt.Println("ExprCastUnset")
}

func (a *exploreReader) ScalarDnumber(n *ast.ScalarDnumber) {

	fmt.Println("ScalarDnumber")
}
func (a *exploreReader) ScalarEncapsed(n *ast.ScalarEncapsed) {

	fmt.Println("ScalarEncapsed")
}
func (a *exploreReader) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {

	fmt.Println("ScalarEncapsedStringPart")
}
func (a *exploreReader) ScalarEncapsedStringVar(n *ast.ScalarEncapsedStringVar) {

	fmt.Println("ScalarEncapsedStringVar")
}
func (a *exploreReader) ScalarEncapsedStringBrackets(n *ast.ScalarEncapsedStringBrackets) {

	fmt.Println("ScalarEncapsedStringBrackets")
}
func (a *exploreReader) ScalarHeredoc(n *ast.ScalarHeredoc) {

	fmt.Println("ScalarHeredoc")
}
func (a *exploreReader) ScalarLnumber(n *ast.ScalarLnumber) {

	fmt.Println("ScalarLnumber")
}
func (a *exploreReader) ScalarMagicConstant(n *ast.ScalarMagicConstant) {

	fmt.Println("ScalarMagicConstant")
}
func (a *exploreReader) ScalarString(n *ast.ScalarString) {

	fmt.Println("ScalarString")

	fmt.Println(string(n.Value))
}

func (a *exploreReader) NameName(n *ast.Name) {

	fmt.Println("NameName")

	fmt.Println("HELLO")

}
func (a *exploreReader) NameFullyQualified(n *ast.NameFullyQualified) {

	fmt.Println("NameFullyQualified")
}
func (a *exploreReader) NameRelative(n *ast.NameRelative) {

	fmt.Println("NameRelative")
}
func (a *exploreReader) NameNamePart(n *ast.NamePart) {
	fmt.Println("NameNamePart")
}
