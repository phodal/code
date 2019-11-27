grammar Code;

compilationUnit
    : packageDeclaration? importDeclaration* typeDeclaration* expressDeclaration* EOF
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
    ;

expressDeclaration
    : methodCallDeclaration
    ;

methodCallDeclaration: IDENTIFIER '(' parameter* ')';
parameter
    : IDENTIFIER
    | STRING_LITERAL
    ;

dataStructDeclaration: DATA_STRUCT IDENTIFIER;
memberDeclaration: MEMBER IDENTIFIER;
functionDeclaration: FUNCTION IDENTIFIER;

PACKAGE: 'package';
IMPORT: 'import';
DATA_STRUCT: 'struct';
MEMBER: 'member';
FUNCTION: 'function';


IDENTIFIER:         Letter LetterOrDigit*;

// Whitespace and comments

WS:                 [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT:            '/*' .*? '*/'    -> channel(HIDDEN);
LINE_COMMENT:       '//' ~[\r\n]*    -> channel(HIDDEN);

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';


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
