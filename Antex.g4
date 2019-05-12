grammar Antex;

expr
  : '(' expr ')'
  | expr ('*' | '/') expr
  | expr ('+' | '-') expr
  | value
  ;

value
  : INT
  ;


INT: [0-9]+;
