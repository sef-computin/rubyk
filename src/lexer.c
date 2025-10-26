#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "lexer.h"


void _append_token(token_t* token_list, token_type type, char *val){
  token_t *last = token_list;
  while (last->next != NULL) last = last->next;

  last->next = malloc(sizeof(token_t));
  last = last->next;

  last->type = type;
  last->value = val;
}

token_type _get_word_type(char *word, int word_len){
  switch (word_len) {
    case 2:
      if(strncmp(word, "if", word_len) == 0){
        return IF;
      };
      break;

    case 3:
      if(strncmp(word, "def", word_len) == 0){
        return DEF;
      };
      if(strncmp(word, "end", word_len) == 0){
        return END;
      };
      if(strncmp(word, "for", word_len) == 0){
        return FOR;
      };
      if(strncmp(word, "nil", word_len) == 0){
        return TOKEN_NIL;
      };
      break;

    case 4:
      if(strncmp(word, "else", word_len) == 0){
        return ELSE;
      };
      break;

    case 5:
      if(strncmp(word, "while", word_len) == 0){
        return WHILE;
      };
      break;

    default:
      break;
  }
  return ID;
}


token_t *lexer(FILE *src_file){
  if (src_file == NULL){
    return NULL;
  }

  token_t *head = malloc(sizeof(token_t));

  char *line = NULL;
  ssize_t linelen;
  size_t linecap;

  while ((linelen = getline(&line, &linecap, src_file)) != -1){
    while (linelen > 0 && ( line[linelen-1] == '\n' || 
                            line[linelen-1] == '\r'))
      linelen--;

    ssize_t lexema_start = -1;
    for (ssize_t i = 0; i < linelen;){
      char current_symbol = line[i];


      switch (current_symbol) {
        case '+':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, PLUS_ASSIGN, NULL);
          } else {
            _append_token(head, PLUS, NULL);
          }
          break;
        case '-':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, MINUS_ASSIGN, NULL);
          } else {
            _append_token(head, MINUS, NULL);
          }
          break;
        case '*':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, MUL_ASSIGN, NULL);
          } else {
            _append_token(head, MUL, NULL);
          }
          break;
        case '%':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, MOD_ASSIGN, NULL);
          } else {
            _append_token(head, MOD, NULL);
          }
          break;
        case '/':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, DIV_ASSIGN, NULL);
          } else {
            _append_token(head, DIV, NULL);
          }
          break;
        case '>':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, GREATER_EQUAL, NULL);
          } else {
            _append_token(head, GREATER, NULL);
          }
          break;
        case '<':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, LESS_EQUAL, NULL);
          } else {
            _append_token(head, LESS, NULL);
          }
          break;
        case '[':
          _append_token(head, LEFT_SBRACKET, NULL);
          break;
        case ']':
          _append_token(head, RIGHT_SBRACKET, NULL);
          break;
        case '=':
          if (i+1 < linelen && line[i+1] == '='){
            i++;
            _append_token(head, EQUAL, NULL);
          } else{
            _append_token(head, ASSIGN, NULL);
          }
          break;
        case ',':
          _append_token(head, COMMA, NULL);
          break;
        case '.':
          _append_token(head, DOT, NULL);
          break;
        default:
          break;
      }

      while (isalpha(current_symbol) || isdigit(current_symbol)){
        if (lexema_start == -1){
          lexema_start = i;
        }

        i++;
        if (i == linelen) break;
        current_symbol = line[i];
      }

      if (lexema_start >= 0){
        ssize_t lexema_len = i - lexema_start;
        token_type type = _get_word_type(&line[lexema_start], lexema_len);
        char *lexema = malloc(lexema_len + 1);
        strncpy(lexema, &line[lexema_start], lexema_len);
        lexema[lexema_len] = '\0';
        _append_token(head, type, lexema);
        lexema_start = -1;
      }

      i++;
    }
    _append_token(head, CRLF, NULL);

  }

  _append_token(head, TOKEN_EOF, NULL);

  token_t *ret = head->next;
  free(head);
  return ret;
}

void free_token(token_t *token){
  if (token->value != NULL){
    free(token->value);
  }
  free(token);
}

void print_token(token_t *token){
  printf("Token: \n");
  printf("\ttype: %x\n", token->type);
  if (token->value != NULL){
    printf("\tvalue: %s\n", token->value);
  }
}


int main(){
  FILE *src = fopen("../examples/arithmetic.rb", "r");
  if (src == NULL){
    perror("fopen");
  }
 
  token_t *head = lexer(src);

  if (head == NULL){
    return -1;
  }

  while (head->next != NULL){
    print_token(head);
    token_t *t = head;

    head = head->next;
    free(t);
  }
  free(head);

  fclose(src);
}
