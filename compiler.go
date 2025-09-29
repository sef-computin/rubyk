package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	gen "github.com/sef-comp/truby-compiler/codegen"
	p "github.com/sef-comp/truby-compiler/parser"
	"os"
)

func main() {

	outPath := "./out/"
	outFilename := "output.ll"

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./compiler <source-file>")
		return
	}

	fileName := os.Args[1]
	sourceCode, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Лексический и синтаксический анализ с использованием ANTLR
	input := antlr.NewInputStream(string(sourceCode))
	lexer := p.NewRubyKLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := p.NewRubyKParser(stream)

	// Парсинг программы
	tree := parser.Program()

	// children := tree.GetChildren()

	// fmt.Println()

	// Генерация кода LLVM IR с использованием llir/llvm
	visitor := gen.NewVisitor(true)

	visitor.VisitProgram(tree.(*p.ProgramContext))

	// Вывод результата
	if visitor.Debug {
		fmt.Println(visitor.Module)
	}

	// Сохранение в файл

	if file, err := os.Create(outPath + outFilename); err == nil {
		file.Write([]byte(visitor.Module.String()))
	}

}
