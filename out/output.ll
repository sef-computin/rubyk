define i32 @main() {
entry:
	%0 = alloca [32 x i32]
	%1 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 0
	store i32 8, i32* %1
	%2 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 1
	store i32 5, i32* %2
	%3 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 2
	store i32 1, i32* %3
	%4 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 3
	store i32 4, i32* %4
	%5 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 4
	store i32 10, i32* %5
	%6 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 5
	store i32 6, i32* %6
	%7 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 6
	store i32 8, i32* %7
	%8 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 7
	store i32 9, i32* %8
	%9 = alloca i32
	store i32 8, i32* %9
	%10 = alloca i32
	store i32 0, i32* %10
	%11 = alloca i32
	store i32 0, i32* %11
	%12 = alloca i32
	store i32 0, i32* %12
	%13 = alloca i32
	store i32 0, i32* %13
	store i32 0, i32* %12
	store i32 0, i32* %13
	br label %while.cond-1

while.cond-1:
	%14 = load i32, i32* %10
	%15 = load i32, i32* %9
	%16 = sub i32 %15, 1
	%17 = icmp slt i32 %14, %16
	br i1 %17, label %while.body-1, label %main-1

while.body-1:
	br label %while.cond-4

main-1:
	%18 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 1
	%19 = load i32, i32* %18
	ret i32 %19

while.cond-4:
	%20 = load i32, i32* %12
	%21 = load i32, i32* %9
	%22 = load i32, i32* %10
	%23 = sub i32 %22, 1
	%24 = sub i32 %21, %23
	%25 = icmp slt i32 %20, %24
	br i1 %25, label %while.body-4, label %main-4

while.body-4:
	%26 = load i32, i32* %12
	%27 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 %26
	%28 = load i32, i32* %27
	%29 = load i32, i32* %12
	%30 = add i32 %29, 1
	%31 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 %30
	%32 = load i32, i32* %31
	%33 = icmp sgt i32 %28, %32
	br i1 %33, label %if-7, label %else-7

main-4:
	%34 = load i32, i32* %13
	%35 = icmp eq i32 %34, 0
	br i1 %35, label %if-10, label %else-10

if-7:
	%36 = load i32, i32* %12
	%37 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 %36
	%38 = load i32, i32* %37
	store i32 %38, i32* %11
	%39 = load i32, i32* %12
	%40 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 %39
	%41 = load i32, i32* %12
	%42 = add i32 %41, 1
	%43 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 %42
	%44 = load i32, i32* %43
	store i32 %44, i32* %40
	%45 = load i32, i32* %12
	%46 = add i32 %45, 1
	%47 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 %46
	%48 = load i32, i32* %11
	store i32 %48, i32* %47
	%49 = load i32, i32* %13
	%50 = add i32 %49, 1
	store i32 %50, i32* %13
	br label %main-7

else-7:
	br label %main-7

main-7:
	%51 = load i32, i32* %12
	%52 = add i32 %51, 1
	store i32 %52, i32* %12
	br label %while.cond-4

if-10:
	%53 = getelementptr [32 x i32], [32 x i32]* %0, i32 0, i32 0
	%54 = load i32, i32* %53
	ret i32 %54

else-10:
	br label %main-10

main-10:
	%55 = load i32, i32* %10
	%56 = add i32 %55, 1
	store i32 %56, i32* %10
	br label %while.cond-1
}
