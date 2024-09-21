// Code generated from RubyK.g4 by ANTLR 4.10. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RubyKLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rubyklexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rubyklexerLexerInit() {
	staticData := &rubyklexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "','", "';'", "", "'end'", "'def'", "'return'", "'pir'", "'if'",
		"'else'", "'elsif'", "'unless'", "'while'", "'retry'", "'break'", "'for'",
		"'true'", "'false'", "'+'", "'-'", "'*'", "'/'", "'%'", "'**'", "'=='",
		"'!='", "'>'", "'<'", "'<='", "'>='", "'='", "'+='", "'-='", "'*='",
		"'/='", "'%='", "'**='", "'&'", "'|'", "'^'", "'~'", "'<<'", "'>>'",
		"", "", "", "'('", "')'", "'['", "']'", "'nil'",
	}
	staticData.symbolicNames = []string{
		"", "LITERAL", "COMMA", "SEMICOLON", "CRLF", "END", "DEF", "RETURN",
		"PIR", "IF", "ELSE", "ELSIF", "UNLESS", "WHILE", "RETRY", "BREAK", "FOR",
		"TRUE", "FALSE", "PLUS", "MINUS", "MUL", "DIV", "MOD", "EXP", "EQUAL",
		"NOT_EQUAL", "GREATER", "LESS", "LESS_EQUAL", "GREATER_EQUAL", "ASSIGN",
		"PLUS_ASSIGN", "MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN",
		"EXP_ASSIGN", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_NOT", "BIT_SHL",
		"BIT_SHR", "AND", "OR", "NOT", "LEFT_RBRACKET", "RIGHT_RBRACKET", "LEFT_SBRACKET",
		"RIGHT_SBRACKET", "NIL", "SL_COMMENT", "ML_COMMENT", "WS", "INT", "FLOAT",
		"ID", "ID_GLOBAL", "ID_FUNCTION",
	}
	staticData.ruleNames = []string{
		"ESCAPED_QUOTE", "LITERAL", "COMMA", "SEMICOLON", "CRLF", "END", "DEF",
		"RETURN", "PIR", "IF", "ELSE", "ELSIF", "UNLESS", "WHILE", "RETRY",
		"BREAK", "FOR", "TRUE", "FALSE", "PLUS", "MINUS", "MUL", "DIV", "MOD",
		"EXP", "EQUAL", "NOT_EQUAL", "GREATER", "LESS", "LESS_EQUAL", "GREATER_EQUAL",
		"ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN",
		"MOD_ASSIGN", "EXP_ASSIGN", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_NOT",
		"BIT_SHL", "BIT_SHR", "AND", "OR", "NOT", "LEFT_RBRACKET", "RIGHT_RBRACKET",
		"LEFT_SBRACKET", "RIGHT_SBRACKET", "NIL", "SL_COMMENT", "ML_COMMENT",
		"WS", "INT", "FLOAT", "ID", "ID_GLOBAL", "ID_FUNCTION",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 59, 397, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		5, 1, 128, 8, 1, 10, 1, 12, 1, 131, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		137, 8, 1, 10, 1, 12, 1, 140, 9, 1, 1, 1, 3, 1, 143, 8, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 3, 4, 150, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3,
		44, 296, 8, 44, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 302, 8, 45, 1, 46, 1,
		46, 1, 46, 1, 46, 3, 46, 308, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 5, 52, 324,
		8, 52, 10, 52, 12, 52, 327, 9, 52, 1, 52, 3, 52, 330, 8, 52, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		5, 53, 344, 8, 53, 10, 53, 12, 53, 347, 9, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 53, 3, 53, 355, 8, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54,
		4, 54, 362, 8, 54, 11, 54, 12, 54, 363, 1, 54, 1, 54, 1, 55, 4, 55, 369,
		8, 55, 11, 55, 12, 55, 370, 1, 56, 5, 56, 374, 8, 56, 10, 56, 12, 56, 377,
		9, 56, 1, 56, 1, 56, 4, 56, 381, 8, 56, 11, 56, 12, 56, 382, 1, 57, 1,
		57, 5, 57, 387, 8, 57, 10, 57, 12, 57, 390, 9, 57, 1, 58, 1, 58, 1, 58,
		1, 59, 1, 59, 1, 59, 3, 129, 138, 345, 0, 60, 1, 0, 3, 1, 5, 2, 7, 3, 9,
		4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29,
		14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 47,
		23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30, 63, 31, 65,
		32, 67, 33, 69, 34, 71, 35, 73, 36, 75, 37, 77, 38, 79, 39, 81, 40, 83,
		41, 85, 42, 87, 43, 89, 44, 91, 45, 93, 46, 95, 47, 97, 48, 99, 49, 101,
		50, 103, 51, 105, 52, 107, 53, 109, 54, 111, 55, 113, 56, 115, 57, 117,
		58, 119, 59, 1, 0, 6, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 1, 0, 48,
		57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		1, 0, 63, 63, 413, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
		1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0,
		93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0,
		0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1,
		0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0,
		115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 1, 121, 1, 0,
		0, 0, 3, 142, 1, 0, 0, 0, 5, 144, 1, 0, 0, 0, 7, 146, 1, 0, 0, 0, 9, 149,
		1, 0, 0, 0, 11, 153, 1, 0, 0, 0, 13, 157, 1, 0, 0, 0, 15, 161, 1, 0, 0,
		0, 17, 168, 1, 0, 0, 0, 19, 172, 1, 0, 0, 0, 21, 175, 1, 0, 0, 0, 23, 180,
		1, 0, 0, 0, 25, 186, 1, 0, 0, 0, 27, 193, 1, 0, 0, 0, 29, 199, 1, 0, 0,
		0, 31, 205, 1, 0, 0, 0, 33, 211, 1, 0, 0, 0, 35, 215, 1, 0, 0, 0, 37, 220,
		1, 0, 0, 0, 39, 226, 1, 0, 0, 0, 41, 228, 1, 0, 0, 0, 43, 230, 1, 0, 0,
		0, 45, 232, 1, 0, 0, 0, 47, 234, 1, 0, 0, 0, 49, 236, 1, 0, 0, 0, 51, 239,
		1, 0, 0, 0, 53, 242, 1, 0, 0, 0, 55, 245, 1, 0, 0, 0, 57, 247, 1, 0, 0,
		0, 59, 249, 1, 0, 0, 0, 61, 252, 1, 0, 0, 0, 63, 255, 1, 0, 0, 0, 65, 257,
		1, 0, 0, 0, 67, 260, 1, 0, 0, 0, 69, 263, 1, 0, 0, 0, 71, 266, 1, 0, 0,
		0, 73, 269, 1, 0, 0, 0, 75, 272, 1, 0, 0, 0, 77, 276, 1, 0, 0, 0, 79, 278,
		1, 0, 0, 0, 81, 280, 1, 0, 0, 0, 83, 282, 1, 0, 0, 0, 85, 284, 1, 0, 0,
		0, 87, 287, 1, 0, 0, 0, 89, 295, 1, 0, 0, 0, 91, 301, 1, 0, 0, 0, 93, 307,
		1, 0, 0, 0, 95, 309, 1, 0, 0, 0, 97, 311, 1, 0, 0, 0, 99, 313, 1, 0, 0,
		0, 101, 315, 1, 0, 0, 0, 103, 317, 1, 0, 0, 0, 105, 321, 1, 0, 0, 0, 107,
		335, 1, 0, 0, 0, 109, 361, 1, 0, 0, 0, 111, 368, 1, 0, 0, 0, 113, 375,
		1, 0, 0, 0, 115, 384, 1, 0, 0, 0, 117, 391, 1, 0, 0, 0, 119, 394, 1, 0,
		0, 0, 121, 122, 5, 92, 0, 0, 122, 123, 5, 34, 0, 0, 123, 2, 1, 0, 0, 0,
		124, 129, 5, 34, 0, 0, 125, 128, 3, 1, 0, 0, 126, 128, 8, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 132, 143, 5, 34, 0, 0, 133, 138, 5, 39, 0, 0, 134, 137, 3, 1, 0,
		0, 135, 137, 8, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0, 0, 0, 137,
		140, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 141,
		1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 143, 5, 39, 0, 0, 142, 124, 1, 0,
		0, 0, 142, 133, 1, 0, 0, 0, 143, 4, 1, 0, 0, 0, 144, 145, 5, 44, 0, 0,
		145, 6, 1, 0, 0, 0, 146, 147, 5, 59, 0, 0, 147, 8, 1, 0, 0, 0, 148, 150,
		5, 13, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 152, 5, 10, 0, 0, 152, 10, 1, 0, 0, 0, 153, 154, 5, 101, 0,
		0, 154, 155, 5, 110, 0, 0, 155, 156, 5, 100, 0, 0, 156, 12, 1, 0, 0, 0,
		157, 158, 5, 100, 0, 0, 158, 159, 5, 101, 0, 0, 159, 160, 5, 102, 0, 0,
		160, 14, 1, 0, 0, 0, 161, 162, 5, 114, 0, 0, 162, 163, 5, 101, 0, 0, 163,
		164, 5, 116, 0, 0, 164, 165, 5, 117, 0, 0, 165, 166, 5, 114, 0, 0, 166,
		167, 5, 110, 0, 0, 167, 16, 1, 0, 0, 0, 168, 169, 5, 112, 0, 0, 169, 170,
		5, 105, 0, 0, 170, 171, 5, 114, 0, 0, 171, 18, 1, 0, 0, 0, 172, 173, 5,
		105, 0, 0, 173, 174, 5, 102, 0, 0, 174, 20, 1, 0, 0, 0, 175, 176, 5, 101,
		0, 0, 176, 177, 5, 108, 0, 0, 177, 178, 5, 115, 0, 0, 178, 179, 5, 101,
		0, 0, 179, 22, 1, 0, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182, 5, 108, 0,
		0, 182, 183, 5, 115, 0, 0, 183, 184, 5, 105, 0, 0, 184, 185, 5, 102, 0,
		0, 185, 24, 1, 0, 0, 0, 186, 187, 5, 117, 0, 0, 187, 188, 5, 110, 0, 0,
		188, 189, 5, 108, 0, 0, 189, 190, 5, 101, 0, 0, 190, 191, 5, 115, 0, 0,
		191, 192, 5, 115, 0, 0, 192, 26, 1, 0, 0, 0, 193, 194, 5, 119, 0, 0, 194,
		195, 5, 104, 0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5, 108, 0, 0, 197,
		198, 5, 101, 0, 0, 198, 28, 1, 0, 0, 0, 199, 200, 5, 114, 0, 0, 200, 201,
		5, 101, 0, 0, 201, 202, 5, 116, 0, 0, 202, 203, 5, 114, 0, 0, 203, 204,
		5, 121, 0, 0, 204, 30, 1, 0, 0, 0, 205, 206, 5, 98, 0, 0, 206, 207, 5,
		114, 0, 0, 207, 208, 5, 101, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5,
		107, 0, 0, 210, 32, 1, 0, 0, 0, 211, 212, 5, 102, 0, 0, 212, 213, 5, 111,
		0, 0, 213, 214, 5, 114, 0, 0, 214, 34, 1, 0, 0, 0, 215, 216, 5, 116, 0,
		0, 216, 217, 5, 114, 0, 0, 217, 218, 5, 117, 0, 0, 218, 219, 5, 101, 0,
		0, 219, 36, 1, 0, 0, 0, 220, 221, 5, 102, 0, 0, 221, 222, 5, 97, 0, 0,
		222, 223, 5, 108, 0, 0, 223, 224, 5, 115, 0, 0, 224, 225, 5, 101, 0, 0,
		225, 38, 1, 0, 0, 0, 226, 227, 5, 43, 0, 0, 227, 40, 1, 0, 0, 0, 228, 229,
		5, 45, 0, 0, 229, 42, 1, 0, 0, 0, 230, 231, 5, 42, 0, 0, 231, 44, 1, 0,
		0, 0, 232, 233, 5, 47, 0, 0, 233, 46, 1, 0, 0, 0, 234, 235, 5, 37, 0, 0,
		235, 48, 1, 0, 0, 0, 236, 237, 5, 42, 0, 0, 237, 238, 5, 42, 0, 0, 238,
		50, 1, 0, 0, 0, 239, 240, 5, 61, 0, 0, 240, 241, 5, 61, 0, 0, 241, 52,
		1, 0, 0, 0, 242, 243, 5, 33, 0, 0, 243, 244, 5, 61, 0, 0, 244, 54, 1, 0,
		0, 0, 245, 246, 5, 62, 0, 0, 246, 56, 1, 0, 0, 0, 247, 248, 5, 60, 0, 0,
		248, 58, 1, 0, 0, 0, 249, 250, 5, 60, 0, 0, 250, 251, 5, 61, 0, 0, 251,
		60, 1, 0, 0, 0, 252, 253, 5, 62, 0, 0, 253, 254, 5, 61, 0, 0, 254, 62,
		1, 0, 0, 0, 255, 256, 5, 61, 0, 0, 256, 64, 1, 0, 0, 0, 257, 258, 5, 43,
		0, 0, 258, 259, 5, 61, 0, 0, 259, 66, 1, 0, 0, 0, 260, 261, 5, 45, 0, 0,
		261, 262, 5, 61, 0, 0, 262, 68, 1, 0, 0, 0, 263, 264, 5, 42, 0, 0, 264,
		265, 5, 61, 0, 0, 265, 70, 1, 0, 0, 0, 266, 267, 5, 47, 0, 0, 267, 268,
		5, 61, 0, 0, 268, 72, 1, 0, 0, 0, 269, 270, 5, 37, 0, 0, 270, 271, 5, 61,
		0, 0, 271, 74, 1, 0, 0, 0, 272, 273, 5, 42, 0, 0, 273, 274, 5, 42, 0, 0,
		274, 275, 5, 61, 0, 0, 275, 76, 1, 0, 0, 0, 276, 277, 5, 38, 0, 0, 277,
		78, 1, 0, 0, 0, 278, 279, 5, 124, 0, 0, 279, 80, 1, 0, 0, 0, 280, 281,
		5, 94, 0, 0, 281, 82, 1, 0, 0, 0, 282, 283, 5, 126, 0, 0, 283, 84, 1, 0,
		0, 0, 284, 285, 5, 60, 0, 0, 285, 286, 5, 60, 0, 0, 286, 86, 1, 0, 0, 0,
		287, 288, 5, 62, 0, 0, 288, 289, 5, 62, 0, 0, 289, 88, 1, 0, 0, 0, 290,
		291, 5, 97, 0, 0, 291, 292, 5, 110, 0, 0, 292, 296, 5, 100, 0, 0, 293,
		294, 5, 38, 0, 0, 294, 296, 5, 38, 0, 0, 295, 290, 1, 0, 0, 0, 295, 293,
		1, 0, 0, 0, 296, 90, 1, 0, 0, 0, 297, 298, 5, 111, 0, 0, 298, 302, 5, 114,
		0, 0, 299, 300, 5, 124, 0, 0, 300, 302, 5, 124, 0, 0, 301, 297, 1, 0, 0,
		0, 301, 299, 1, 0, 0, 0, 302, 92, 1, 0, 0, 0, 303, 304, 5, 110, 0, 0, 304,
		305, 5, 111, 0, 0, 305, 308, 5, 116, 0, 0, 306, 308, 5, 33, 0, 0, 307,
		303, 1, 0, 0, 0, 307, 306, 1, 0, 0, 0, 308, 94, 1, 0, 0, 0, 309, 310, 5,
		40, 0, 0, 310, 96, 1, 0, 0, 0, 311, 312, 5, 41, 0, 0, 312, 98, 1, 0, 0,
		0, 313, 314, 5, 91, 0, 0, 314, 100, 1, 0, 0, 0, 315, 316, 5, 93, 0, 0,
		316, 102, 1, 0, 0, 0, 317, 318, 5, 110, 0, 0, 318, 319, 5, 105, 0, 0, 319,
		320, 5, 108, 0, 0, 320, 104, 1, 0, 0, 0, 321, 325, 5, 35, 0, 0, 322, 324,
		8, 0, 0, 0, 323, 322, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323, 1, 0,
		0, 0, 325, 326, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0,
		328, 330, 5, 13, 0, 0, 329, 328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330,
		331, 1, 0, 0, 0, 331, 332, 5, 10, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334,
		6, 52, 0, 0, 334, 106, 1, 0, 0, 0, 335, 336, 5, 61, 0, 0, 336, 337, 5,
		98, 0, 0, 337, 338, 5, 101, 0, 0, 338, 339, 5, 103, 0, 0, 339, 340, 5,
		105, 0, 0, 340, 341, 5, 110, 0, 0, 341, 345, 1, 0, 0, 0, 342, 344, 9, 0,
		0, 0, 343, 342, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0,
		345, 343, 1, 0, 0, 0, 346, 348, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348,
		349, 5, 61, 0, 0, 349, 350, 5, 101, 0, 0, 350, 351, 5, 110, 0, 0, 351,
		352, 5, 100, 0, 0, 352, 354, 1, 0, 0, 0, 353, 355, 5, 13, 0, 0, 354, 353,
		1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 357, 5, 10,
		0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 6, 53, 0, 0, 359, 108, 1, 0, 0, 0,
		360, 362, 7, 1, 0, 0, 361, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363,
		361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 366,
		6, 54, 0, 0, 366, 110, 1, 0, 0, 0, 367, 369, 7, 2, 0, 0, 368, 367, 1, 0,
		0, 0, 369, 370, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0,
		371, 112, 1, 0, 0, 0, 372, 374, 7, 2, 0, 0, 373, 372, 1, 0, 0, 0, 374,
		377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 378,
		1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 380, 5, 46, 0, 0, 379, 381, 7, 2,
		0, 0, 380, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0,
		382, 383, 1, 0, 0, 0, 383, 114, 1, 0, 0, 0, 384, 388, 7, 3, 0, 0, 385,
		387, 7, 4, 0, 0, 386, 385, 1, 0, 0, 0, 387, 390, 1, 0, 0, 0, 388, 386,
		1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 116, 1, 0, 0, 0, 390, 388, 1, 0,
		0, 0, 391, 392, 5, 36, 0, 0, 392, 393, 3, 115, 57, 0, 393, 118, 1, 0, 0,
		0, 394, 395, 3, 115, 57, 0, 395, 396, 7, 5, 0, 0, 396, 120, 1, 0, 0, 0,
		19, 0, 127, 129, 136, 138, 142, 149, 295, 301, 307, 325, 329, 345, 354,
		363, 370, 375, 382, 388, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RubyKLexerInit initializes any static state used to implement RubyKLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRubyKLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RubyKLexerInit() {
	staticData := &rubyklexerLexerStaticData
	staticData.once.Do(rubyklexerLexerInit)
}

// NewRubyKLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRubyKLexer(input antlr.CharStream) *RubyKLexer {
	RubyKLexerInit()
	l := new(RubyKLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rubyklexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "RubyK.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RubyKLexer tokens.
const (
	RubyKLexerLITERAL        = 1
	RubyKLexerCOMMA          = 2
	RubyKLexerSEMICOLON      = 3
	RubyKLexerCRLF           = 4
	RubyKLexerEND            = 5
	RubyKLexerDEF            = 6
	RubyKLexerRETURN         = 7
	RubyKLexerPIR            = 8
	RubyKLexerIF             = 9
	RubyKLexerELSE           = 10
	RubyKLexerELSIF          = 11
	RubyKLexerUNLESS         = 12
	RubyKLexerWHILE          = 13
	RubyKLexerRETRY          = 14
	RubyKLexerBREAK          = 15
	RubyKLexerFOR            = 16
	RubyKLexerTRUE           = 17
	RubyKLexerFALSE          = 18
	RubyKLexerPLUS           = 19
	RubyKLexerMINUS          = 20
	RubyKLexerMUL            = 21
	RubyKLexerDIV            = 22
	RubyKLexerMOD            = 23
	RubyKLexerEXP            = 24
	RubyKLexerEQUAL          = 25
	RubyKLexerNOT_EQUAL      = 26
	RubyKLexerGREATER        = 27
	RubyKLexerLESS           = 28
	RubyKLexerLESS_EQUAL     = 29
	RubyKLexerGREATER_EQUAL  = 30
	RubyKLexerASSIGN         = 31
	RubyKLexerPLUS_ASSIGN    = 32
	RubyKLexerMINUS_ASSIGN   = 33
	RubyKLexerMUL_ASSIGN     = 34
	RubyKLexerDIV_ASSIGN     = 35
	RubyKLexerMOD_ASSIGN     = 36
	RubyKLexerEXP_ASSIGN     = 37
	RubyKLexerBIT_AND        = 38
	RubyKLexerBIT_OR         = 39
	RubyKLexerBIT_XOR        = 40
	RubyKLexerBIT_NOT        = 41
	RubyKLexerBIT_SHL        = 42
	RubyKLexerBIT_SHR        = 43
	RubyKLexerAND            = 44
	RubyKLexerOR             = 45
	RubyKLexerNOT            = 46
	RubyKLexerLEFT_RBRACKET  = 47
	RubyKLexerRIGHT_RBRACKET = 48
	RubyKLexerLEFT_SBRACKET  = 49
	RubyKLexerRIGHT_SBRACKET = 50
	RubyKLexerNIL            = 51
	RubyKLexerSL_COMMENT     = 52
	RubyKLexerML_COMMENT     = 53
	RubyKLexerWS             = 54
	RubyKLexerINT            = 55
	RubyKLexerFLOAT          = 56
	RubyKLexerID             = 57
	RubyKLexerID_GLOBAL      = 58
	RubyKLexerID_FUNCTION    = 59
)