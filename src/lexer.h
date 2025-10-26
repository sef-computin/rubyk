#ifndef LEXER_H
#define LEXER_H

#include <stdio.h>
typedef enum {
  ID = 1,
  NUMBER,
  FALSE,
  TRUE,
  FOR,
  BREAK,
  RETRY,
  WHILE,
  UNLESS,
  ELSEIF,
  ELSE,
  IF,
  PIR,
  RETURN,
  DEF,
  END,
  TOKEN_NIL,
  STRING,

  LEFT_SBRACKET,
  RIGHT_SBRACKET,
  NOT,
  OR,
  AND,
  BIT_SHR,
  BIT_XOR,
  BIT_OR,
  BIT_AND,
  EXP_ASSIGN,
  MOD_ASSIGN,
  DIV_ASSIGN,
  MUL_ASSIGN,
  MINUS_ASSIGN,
  PLUS_ASSIGN,
  ASSIGN,
  GREATER_EQUAL,
  LESS_EQUAL,
  LESS,
  GREATER,
  NOT_EQUAL,
  EQUAL,
  EXP,
  MOD,
  DIV,
  MUL,
  MINUS,
  PLUS,
  CRLF,
  SEMICOLON,
  COMMA,
  DOT,
  TOKEN_EOF,
} token_type;


typedef struct token{
  token_type type;
  char *value;

  struct token *next;

} token_t;

#endif // !LEXER_H
