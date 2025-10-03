#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef enum {
  ID_FUNCTION,
  ID_GLOBAL,
  ID,
  FLOAT,
  INT,
  WS,
  ML_COMMENT,
  SL_COMMENT,
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
  CRLF,
  SEMICOLON,
  COMMA,
  LITERAL,
  ESCAPED_QUOTE,
  TOKEN_EOF,
} TokenType;

typedef struct Token{
  TokenType type;
  char *value;
  ssize_t size;
} Token;


char *RUBY_KEYWORDS[] = {
  "if", "while", "for", "break", "end", "continue", "return", "else",
  "def", "static", "enum", "class", "default", "case",
  "int", "float", "char", "string", "bool", "nil", NULL
};

Token *newToken(TokenType type, char *value, ssize_t size){
  Token *token = malloc(sizeof(Token));
  token->type = type;
  token->value = malloc(sizeof(char) * size);
  strncpy(token->value, value, size);
  *(token->value + size) = '\0';

  token->size = size;

  return token;
};

Token *_addToken(Token *token_list, int *cap, Token *new_token, int idx){
  if (idx >= (*cap)-1){
    *(cap) *= 1.5;
    token_list = (Token *)realloc(token_list, sizeof(Token)*(*cap));
  }

  token_list[idx] = *new_token; 
  return token_list;
};

TokenType _getWordType(char *word, int word_len){


  return ID;
}

Token *lexer(FILE *src_file){
  if (src_file == NULL){
    return NULL;
  }

  int token_list_cap = 10;
  Token *token_list = malloc(sizeof(Token) * 10);
  int tokens_total = 0;

  char *line = NULL;
  size_t linecap = 0;
  ssize_t linelen;

  while ((linelen = getline(&line, &linecap, src_file)) != -1){
    while (linelen > 0 && ( line[linelen-1] == '\n' || 
                            line[linelen-1] == '\r'))
      linelen--;

    ssize_t word_start = -1;
    for (ssize_t i = 0; i < linelen;){
      char current_symbol = line[i];
      
      while (isalpha(current_symbol) || isdigit(current_symbol)){
        if (word_start == -1){
          word_start = i;
        }

        i++;
        if (i == linelen) break;
        current_symbol = line[i];
      }

      
      if (word_start >= 0){
        TokenType type = _getWordType(&line[word_start], (i-word_start));
        token_list = _addToken(token_list, &token_list_cap, newToken(type, &line[word_start], (i - word_start)), tokens_total++);
        word_start = -1;
      }
      
      switch (current_symbol) {
        case '+':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(PLUS_ASSIGN, "+=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(PLUS, "+", 1), tokens_total++);
          }
          break;
        case '-':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(MINUS_ASSIGN, "-=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(MINUS, "-", 1), tokens_total++);
          }
          break;
        case '*':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(MUL_ASSIGN, "*=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(MUL, "*", 1), tokens_total++);
          }
          break;
        case '%':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(MOD_ASSIGN, "%=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(MOD, "%", 1), tokens_total++);
          }
          break;
        case '/':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(DIV_ASSIGN, "/=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(DIV, "/", 1), tokens_total++);
          }
          break;
        case '>':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(GREATER_EQUAL, ">=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(GREATER, ">", 1), tokens_total++);
          }
          break;
        case '<':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(LESS_EQUAL, "<=", 2), tokens_total++);
          } else {
            token_list = _addToken(token_list, &token_list_cap, newToken(LESS, "<", 1), tokens_total++);
          }
          break;
        case '[':
          token_list = _addToken(token_list, &token_list_cap, newToken(LEFT_SBRACKET, "[", 1), tokens_total++);
          break;
        case ']':
          token_list = _addToken(token_list, &token_list_cap, newToken(RIGHT_SBRACKET, "]", 1), tokens_total++);
          break;
        case '=':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            token_list = _addToken(token_list, &token_list_cap, newToken(EQUAL, "==", 2), tokens_total++);
          } else{
            token_list = _addToken(token_list, &token_list_cap, newToken(ASSIGN, "=", 1), tokens_total++);
          }
          break;
        case ',':
          token_list = _addToken(token_list, &token_list_cap, newToken(COMMA, "[", 1), tokens_total++);
          break;
        default:
          break;
      }

      i++;
    }

    if (word_start >= 0){
      token_list = _addToken(token_list, &token_list_cap, newToken(ID, &line[word_start], (linelen - word_start)), tokens_total++);
      word_start = -1;
    }
  }

  return _addToken(token_list, &token_list_cap, newToken(TOKEN_EOF, "\0", 0), tokens_total++);
}


void printToken(Token token){
  printf("Token: \n\t");
  printf("type: %x\n\t", token.type);
  printf("size: %zu\n\t", token.size);
  printf("value: %s\n", token.value);
}


int main(){
  FILE *src = fopen("../examples/arithmetic.rb", "r");
  if (src == NULL){
    perror("fopen");
  }

  
  Token *token_list = lexer(src);

  for (int i = 0; token_list[i].type != TOKEN_EOF; i++){
    Token token = token_list[i];
    printToken(token);
    free(token.value);
  }
  free(token_list);

  fclose(src);
}
