	.text
	.file	"output.ll"
	.globl	main                            # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %entry
	movl	$0, -12(%rsp)
	movl	$1, -20(%rsp)
	movl	$2, -16(%rsp)
	movl	$10, -4(%rsp)
	movl	$0, -8(%rsp)
	.p2align	4, 0x90
.LBB0_1:                                # %while.cond-1
                                        # =>This Inner Loop Header: Depth=1
	movl	-16(%rsp), %eax
	cmpl	-4(%rsp), %eax
	jge	.LBB0_3
# %bb.2:                                # %while.body-1
                                        #   in Loop: Header=BB0_1 Depth=1
	movl	-20(%rsp), %eax
	movl	%eax, -8(%rsp)
	movl	-12(%rsp), %ecx
	addl	%eax, %ecx
	movl	%ecx, -20(%rsp)
	movl	%eax, -12(%rsp)
	incl	-16(%rsp)
	jmp	.LBB0_1
.LBB0_3:                                # %main-1
	movl	-20(%rsp), %eax
	retq
.Lfunc_end0:
	.size	main, .Lfunc_end0-main
	.cfi_endproc
                                        # -- End function
	.section	".note.GNU-stack","",@progbits
