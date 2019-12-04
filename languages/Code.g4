grammar Code;

compilationUnit
    : packageDeclaration? importDeclaration* typeDeclaration*  EOF
    ;

packageDeclaration
    : PACKAGE IDENTIFIER
    ;

importDeclaration
    : IMPORT IDENTIFIER
    ;

typeDeclaration
    : dataStructDeclaration
    | memberDeclaration
    | functionDeclaration
    | expressDeclaration
    ;

functionDeclaration: FUNCTION IDENTIFIER '(' parameter? ')' '{' functionBody '}';

functionBody: expressDeclaration (expressDeclaration)*;

primary: IDENTIFIER | DECIMAL_LITERAL | STRING_LITERAL;

expressDeclaration
    : methodCallDeclaration
    | blockStatement
    | primary
    | expressDeclaration ('<' '<' | '>' '>' '>' | '>' '>') expressDeclaration
    | <assoc=right> expression
            bop=('=' | '+=' | '-=' | '*=' | '/=' | '&=' | '|=' | '^=' | '>>=' | '>>>=' | '<<=' | '%=')
            expression
    ;


block
    : '{' blockStatement* '}'
    ;

blockStatement
    : blockLabel=block
    | FOR '(' forControl ')' blockStatement
    | localVariableDeclaration
    ;

forControl: expressDeclaration;

FOR: 'for';

localVariableDeclaration
    : variableDeclarators
    ;

variableDeclarators
    : VAR variableDeclarator (',' variableDeclarator)*
    ;

VAR: 'var';

variableDeclarator
    : variableDeclaratorId ('=' variableInitializer)?
    ;

variableDeclaratorId
    : IDENTIFIER
    ;

variableInitializer
    : expression
    ;

expression
    : (IDENTIFIER | DECIMAL_LITERAL | STRING_LITERAL)
    ;

methodCallDeclaration: IDENTIFIER '(' parameter* ')';
parameter
    : literal
    | IDENTIFIER
    ;

literal
    : DECIMAL_LITERAL
    | STRING_LITERAL
    ;

dataStructDeclaration: DATA_STRUCT IDENTIFIER;
memberDeclaration: MEMBER IDENTIFIER;

PACKAGE: 'package';
IMPORT: 'import';
DATA_STRUCT: 'struct';
MEMBER: 'member';
FUNCTION: 'func';


IDENTIFIER:         Letter LetterOrDigit*;

// Whitespace and comments

WS:                 [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT:            '/*' .*? '*/'    -> channel(HIDDEN);
LINE_COMMENT:       '//' ~[\r\n]*    -> channel(HIDDEN);

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';


DECIMAL_LITERAL:    ('0' | [1-9] (Digits? | '_'+ Digits)) [lL]?;

fragment Digits
    : [0-9] ([0-9_]* [0-9])?
    ;

fragment HexDigit
    : [0-9a-fA-F]
    ;

fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' ([0-3]? [0-7])? [0-7]
    | '\\' 'u'+ HexDigit HexDigit HexDigit HexDigit
    ;

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;
