@255
D=A;
@SP
M=D;
@123
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@multiple_end
0;JMP
(FUNCTION__multiple)
@0
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@0
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@ARG
A=M;
D=A;
@1
D=A+D;
A=D;
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@LCL
A=M;
D=A;
@0
D=A+D;
@R6
M=D;
@R5
D=M;
@R6
A=M;
M=D;
@0
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@LCL
A=M;
D=A;
@1
D=A+D;
@R6
M=D;
@R5
D=M;
@R6
A=M;
M=D;
(multiple_if)
@LCL
A=M;
D=A;
@0
D=A+D;
A=D;
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@0
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@SP
M=M-1;
@SP
A=M;
D=M;
@R6
M=D;
@SP
A=M;
M=0;
@R6
D=M;
@R5
D=D-M;
@generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_THEN
D;JEQ
@generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_ELSE
0;JMP
(generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_THEN)
@1
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_END
0;JMP
(generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_ELSE)
@0
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_END
0;JMP
(generated_ident__3e29abaa_e539_11ea_b508_dca9046dad34_END)
@SP
A=M-1;
D=M;
@multiple_if_then
D-1;JEQ
@LCL
A=M;
D=A;
@1
D=A+D;
A=D;
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@ARG
A=M;
D=A;
@0
D=A+D;
A=D;
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@SP
M=M-1;
@SP
A=M;
D=M;
@R6
M=D;
@SP
A=M;
M=0;
@R6
D=M;
@R5
D=D+M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@LCL
A=M;
D=A;
@1
D=A+D;
@R6
M=D;
@R5
D=M;
@R6
A=M;
M=D;
@LCL
A=M;
D=A;
@0
D=A+D;
A=D;
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@1
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@SP
M=M-1;
@SP
A=M;
D=M;
@R6
M=D;
@SP
A=M;
M=0;
@R6
D=M;
@R5
D=D-M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@LCL
A=M;
D=A;
@0
D=A+D;
@R6
M=D;
@R5
D=M;
@R6
A=M;
M=D;
@multiple_if
0;JMP
(multiple_if_then)
@LCL
A=M;
D=A;
@1
D=A+D;
A=D;
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
A=M;
A=A-1;
D=M;
@ARG
A=M;
M=D;
@ARG
A=M;
D=A;
@SP
M=D+1;
@LCL
D=M;
@FRAME_BASE__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
M=D;
@5
D=D-A;
A=D;
D=M;
@RETURN_ADDRESS__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
M=D;
@FRAME_BASE__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
D=A;
@4
D=D-A;
A=D;
D=M;
@ARG
M=D;
@FRAME_BASE__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
D=A;
@3
D=D-A;
A=D;
D=M;
@LCL
M=D;
@FRAME_BASE__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
D=A;
@2
D=D-A;
A=D;
D=M;
@THIS
M=D;
@FRAME_BASE__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
D=A;
@1
D=D-A;
A=D;
D=M;
@THAT
M=D;
@RETURN_ADDRESS__generated_ident__3e2a1022_e539_11ea_b508_dca9046dad34
A=M;
0;JMP
(multiple_end)
@1
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@2
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@return_to_generated_ident__3e2a2a12_e539_11ea_b508_dca9046dad34
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@LCL
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@ARG
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@THIS
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@THAT
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
D=M;
@R5
M=D;
@2
D=A;
@R5
M=M-D;
@5
D=A;
@R5
M=M-D;
D=M;
@ARG
M=D;
@SP
D=M;
@LCL
M=D;
@FUNCTION__multiple
0;JMP
(return_to_generated_ident__3e2a2a12_e539_11ea_b508_dca9046dad34)
@3
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@return_to_generated_ident__3e2a47ea_e539_11ea_b508_dca9046dad34
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@LCL
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@ARG
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@THIS
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@THAT
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
D=M;
@R5
M=D;
@2
D=A;
@R5
M=M-D;
@5
D=A;
@R5
M=M-D;
D=M;
@ARG
M=D;
@SP
D=M;
@LCL
M=D;
@FUNCTION__multiple
0;JMP
(return_to_generated_ident__3e2a47ea_e539_11ea_b508_dca9046dad34)
@4
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@return_to_generated_ident__3e2a8354_e539_11ea_b508_dca9046dad34
D=A;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@LCL
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@ARG
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@THIS
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@THAT
D=M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
@SP
D=M;
@R5
M=D;
@2
D=A;
@R5
M=M-D;
@5
D=A;
@R5
M=M-D;
D=M;
@ARG
M=D;
@SP
D=M;
@LCL
M=D;
@FUNCTION__multiple
0;JMP
(return_to_generated_ident__3e2a8354_e539_11ea_b508_dca9046dad34)
@SP
M=M-1;
@SP
A=M;
D=M;
@R5
M=D;
@SP
A=M;
M=0;
@SP
M=M-1;
@SP
A=M;
D=M;
@R6
M=D;
@SP
A=M;
M=0;
@R6
D=M;
@R5
D=D+M;
@SP
A=M;
M=D;
@SP
D=M;
M=D+1;
(VM_END)
@VM_END
0;JMP

