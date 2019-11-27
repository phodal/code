// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Code

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCodeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCodeVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitExpressDeclaration(ctx *ExpressDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitMethodCallDeclaration(ctx *MethodCallDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitDataStructDeclaration(ctx *DataStructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCodeVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}
