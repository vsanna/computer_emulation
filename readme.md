## What is this project
In this project, I'm building a computer machine emulator with assuming only dff-element and nand-element are provided.

Based off the [https://www.nand2tetris.org/](https://www.nand2tetris.org/) course and textbook, 
but I choosed to use golang to build this project and modified some points so that I can emulate how computers works virtually.

The emulator codes include:
- Hardware layer, 
- Assembler layer, 
- OS layer, 
- VM layer 
- and high-level language layer.
- (as appendix, write a small application based on it)


## built components
- Hardware
    - [x] logic gates
        - only nand is given. this means that only nand struct can use Golang's if condition
        - and, or, not, nor, multi_plexer, xor, or_16_to_1
        - multibit_{and, or, not, nor, multi_plexer. to_1_multi_plexer}
    - [x] ALU, decoder, comp/dest/jump decoder
    - [x] memory
        - only dff is given. 
        - flipflop, word, memory, register
- Assembler
    - [x] interpreter from Assemble language to binary code
- VM
    - [ ] interface design
    - [ ] VM implementation
        - [ ] compiler
- OS
- High-Level language


