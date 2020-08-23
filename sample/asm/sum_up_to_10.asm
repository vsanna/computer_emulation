@10
D=A;
@i
M=D;
@sum
M=0;
(LOOP)
    // sum = sum + i
    @sum
    D=M;
    @i
    D=D+M;
    @sum
    M=D;
    // i = i - 1
    @i
    M=M-1;
    D=M;
    @END
    D;JEQ
    @LOOP
    0;JMP
(END)
@sum
D=M;
(END_LOOP)
@END_LOOP
0;JMP