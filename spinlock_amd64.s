// 	spinlock_amd64.s
// 	Implements locking/unlocking based on compare-and-swap instraction that is
// 	xchg. There is no special memory ordering for Intel 64 and IA-32
// 	architectures. So, there is no any fence. In the same time, xchg operation
// 	implies sequencial consistency memory ordering.

#include "textflag.h"

TEXT ·Lock(SB),NOSPLIT,$0
	MOVQ 	flag+0(FP), BP

spin:
	MOVB 	$0x01, AX
	XCHGB 	AX, (BP)
	TESTB 	AX, AX
	JZ 		yield
	CALL 	runtime·Gosched(SB)
	JMP 	spin

yield:
	RET

TEXT ·Unlock(SB),NOSPLIT,$0
	MOVQ 	flag+0(FP), BP
	MOVB 	$0, (BP)
	RET
