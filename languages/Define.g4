grammar Define;

compilationUnit
    : symbolDeclaration* defineDeclaration? normalDeclarations* EOF
    ;

symbolDeclaration
    : SYMBOL_TEXT IDENTIFIER STRING_LITERAL
    ;

normalDeclarations
    : systemDeclaration
    | moduleDeclaration
    ;

SYMBOL_TEXT: 'symbol';
//SPECIAL_SYMBOL: SYMBOLS;

//Define

defineDeclaration
    : DEFINE defineKey LBRACE defineExpress* RBRACE
    ;


defineExpress
    : defineAttribute
    | defineTemplate
    ;

defineAttribute: DEFAULT_SYMBOL COLON symbolKey symbolValue;

defineTemplate: DEFAULT_TEMPLATE LPAREN IDENTIFIER RPAREN LBRACE defineBody RBRACE;

symbolKey: IDENTIFIER;
symbolValue: IDENTIFIER;
defineBody
    : symbolKey STRING_LITERAL symbolKey* templateData symbolKey
    ;

DEFINE: 'define';
DEFAULT_SYMBOL: 'defaultSymbol';
DEFAULT_TEMPLATE: 'defaultTemplate';

templateData: IDENTIFIER;


//

systemDeclaration
    : defineKey ':' defineValue
    ;

defineKey: IDENTIFIER;
defineValue: IDENTIFIER;

//

moduleDeclaration
    : MODULE IDENTIFIER LBRACE  moduleDefines* RBRACE
    ;

moduleDefines
    : IDENTIFIER LBRACE moduleAttribute* RBRACE
    ;

moduleAttribute
    : IMPORT STRING_LITERAL
    | EQUAL STRING_LITERAL
    ;

MODULE: 'module';
IMPORT: 'import';
EQUAL: 'equal';


IDENTIFIER:         Letter LetterOrDigit*;

// Separators

LPAREN:             '(';
RPAREN:             ')';
LBRACE:             '{';
RBRACE:             '}';
LBRACK:             '[';
RBRACK:             ']';
SEMI:               ';';
COMMA:              ',';
DOT:                '.';
COLON:              ':';

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

//SYMBOLS
//    : '{' | '}' | '$' | ')' | '(' | '[' | ']'
//    ;

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;
