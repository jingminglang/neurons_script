%{
package parser


%}

%union{
	num        float64
	str        string
	lval       Lvalue
	expr       Expression
	stmt       Statement
	argList   []Expression
	stmtBlock []Statement
}

%start program


%token IF THEN ELSE FI AND OR NOT EQ SP PRINT CALL_LUA
%token <num> NUMBER
%token <str> IDENTIFIER STRING


%left '+' '-'
%left '*' '/'


%type <lval> var
%type <expr> expr condition rel_condition sub_condition 
// %type <stmt> statement if_stmt assign_stmt print_stmt call_stmt
%type <stmt> statement if_stmt assign_stmt print_stmt
%type <argList> arg_list
%type <stmtBlock> stmt_block program

%%

program	: stmt_block { yylex.(*Lexer).program = $1 }
        ;

stmt_block
	: { $$ = []Statement{} }
	| stmt_block statement { $$ = append($1, $2) }
	;

statement
	: if_stmt { $$ = $1 }
	| if_stmt SP { $$ = $1 }
        | print_stmt { $$ = $1 }
        | print_stmt SP { $$ = $1 }
	| assign_stmt { $$ = $1 }
	| assign_stmt SP { $$ = $1 }
//	| call_stmt  { $$ = $1 }
//	| call_stmt SP  { $$ = $1 }
	;

// call_stmt : CALL_LUA '('  expr  ',' expr ')'  { $$ = &CallLuaExpr{$3,$5} }
// 	  ;


if_stmt
        : IF condition THEN stmt_block ELSE stmt_block FI { $$ = &IfStmt{$2, $4, $6} }
        | IF condition THEN stmt_block FI { $$ = &IfStmt{$2, $4, nil} }
        ;

assign_stmt
	: var '=' expr { $$ = &AssignStmt{$1, $3} }
	;


print_stmt
	: PRINT arg_list { $$ = &PrintStmt{$2} }



arg_list
	: expr { $$ = []Expression{$1} }
	| arg_list ',' expr { $$ = append($1, $3) }


condition
	: rel_condition { $$ = $1 }
	| sub_condition AND sub_condition { $$ = &LogicExpr{AND, $1, $3} }
	| sub_condition OR sub_condition { $$ = &LogicExpr{OR, $1, $3} }
	| NOT sub_condition { $$ = &LogicExpr{NOT, $2, nil} }
        ;

rel_condition
	: expr EQ expr { $$ = &RelExpr{EQ, $1, $3} }
	| expr '<' expr { $$ = &RelExpr{'<', $1, $3} }
	| expr '>' expr { $$ = &RelExpr{'>', $1, $3} }
	;

sub_condition
	: rel_condition { $$ = $1 }
	| '(' condition ')' { $$ = $2 }
	;


var	: IDENTIFIER { $$ = Identifier($1) }
	;


//string : STRING { $$ = String($1) }
//       | '`' STRING  '`' { $$ = String($2) }
//       ;


// call_stmt : CALL_LUA '('  expr  ',' expr ')'  { $$ = &CallLuaExpr{$3,$5} }
// 	  ;

expr	: var { $$ = $1 }
	| NUMBER { $$ = Number($1) }
//        | string { $$ = $1 }
	| STRING { $$ = String($1) }
	| '(' expr ')' { $$ = $2 }
	| expr '+' expr { $$ = &ArithExpr{'+', $1, $3} }
	| expr '-' expr { $$ = &ArithExpr{'-', $1, $3} }
	| expr '*' expr { $$ = &ArithExpr{'*', $1, $3} }
	| expr '/' expr { $$ = &ArithExpr{'/', $1, $3} }
        | CALL_LUA '('  expr  ',' expr ')'  { $$ = &CallLuaExpr{$3,$5} }
	;


%%




// https://gnuu.org/2009/09/18/writing-your-own-toy-compiler/
