# Setters
VAR {type=Type} {name=Value}
IVAR {type=Type} {name=Value}
SET {dest=Variable} {value=Variable|Value|CMP}
CAST {type=Type} {name=Variable}
# Operations
ADD {dest=Variable} {value=Variable|Value}
SUB {dest=Variable} {value=Variable|Value}
MUL {dest=Variable} {value=Variable|Value}
DIV {dest=Variable} {value=Variable|Value}

SIZE {dest=Variable[int]} {value=Variable[Array]}

# Conditionals
CMPE {var1=Variable} {var2=Variable|Value}
CMPNE {var1=Variable} {var2=Variable|Value}
CMPGR {var1=Variable} {var2=Variable|Value}
CMPLE {var1=Variable} {var2=Variable|Value}
CMPT {var1=Variable} {var2=Type}

# Controls
JMP _ {label=Variable|Value}
LBL _ {label=Variable|Value}

# IO
WRITE {chanel=Chanel} {src=Variable|Value}
READ {dst=Variable} {chanel=Chanel} 

#Types= ONE OF
byte
string
int
uint
float
bool
Array=[Type] 
Dictionary=<Type, Type>

# Value
Value=#anything

# Variable = Basic|ArrayAccess|DicAccess
Basic=@variable|@@variable
ArrayAccess=Basic[Variable]
DicAccess=Basic<Variable>

#Chanel=STD_OUT|STD_IN|STD_ERR

--- brainfuck hello world

IVAR [byte] #memory
IVAR uint #pointer
IVAR uint #zero

+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
[
  LBL _ #open1
  CMP @@memory[@@pointer] #0
  JMP _ #close1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
[
  LBL _ #open2
  CMP @@memory[@@pointer] #0
  JMP _ #close2
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
  < SUB @@pointer #1
  < SUB @@pointer #1
  < SUB @@pointer #1
  < SUB @@pointer #1
- SUB @@memory[@@pointer] #1
  ]
  CMP @@zero #0
  JMP _ #open2
  LBL _ #close2
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
- SUB @@memory[@@pointer] #1
> ADD @@pointer #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
  [
  LBL _ #open3
  CMP @@memory[@@pointer] #0
  JMP _ #close3
  < SUB @@pointer #1
  ]
  CMP @@zero #0
  JMP _ #open3
  LBL _ #close3
  < SUB @@pointer #1
- SUB @@memory[@@pointer] #1
  ]
  CMP @@zero #0
  JMP _ #open1
  LBL _ #close1
> ADD @@pointer #1
> ADD @@pointer #1
. PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
  . PRT STD_OUT @@memory[@@pointer]
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
> ADD @@pointer #1
. PRT STD_OUT @@memory[@@pointer]
< SUB @@pointer #1
- SUB @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
  < SUB @@pointer #1
  . PRT STD_OUT @@memory[@@pointer]
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
- SUB @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
+ ADD @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]





--- brainfuck hello optimized

IVAR [byte] #memory
IVAR uint #pointer
IVAR uint #zero

+ ADD @@memory[@@pointer] #8
  [
  LBL _ #open1
  CMP @@memory[@@pointer] #0
  JMP _ #close1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #4
  [
  LBL _ #open2
  CMP @@memory[@@pointer] #0
  JMP _ #close2
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #2
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #3
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #3
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
  < SUB @@pointer #1
  < SUB @@pointer #1
  < SUB @@pointer #1
  < SUB @@pointer #1
- SUB @@memory[@@pointer] #1
  ]
  CMP @@zero #0
  JMP _ #open2
  LBL _ #close2
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
> ADD @@pointer #1
- SUB @@memory[@@pointer] #1
> ADD @@pointer #1
> A
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
  [
  LBL _ #open3
  CMP @@memory[@@pointer] #0
  JMP _ #close3
  < SUB @@pointer #1
  ]
  CMP @@zero #0
  JMP _ #open3
  LBL _ #close3
  < SUB @@pointer #1
- SUB @@memory[@@pointer] #1
  ]
  CMP @@zero #0
  JMP _ #open1
  LBL _ #close1
> ADD @@pointer #1
> ADD @@pointer #1
. PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
- SUB @@memory[@@pointer] #3
  . PRT STD_OUT @@memory[@@pointer]
+ ADD @@memory[@@pointer] #7
  . PRT STD_OUT @@memory[@@pointer]
  . PRT STD_OUT @@memory[@@pointer]
+ ADD @@memory[@@pointer] #3
  . PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #2
. PRT STD_OUT @@memory[@@pointer]
< SUB @@pointer #1
- SUB @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
  < SUB @@pointer #1
  . PRT STD_OUT @@memory[@@pointer]
+ ADD @@memory[@@pointer] #3
  . PRT STD_OUT @@memory[@@pointer]
- SUB @@memory[@@pointer] #6
  . PRT STD_OUT @@memory[@@pointer]
- SUB @@memory[@@pointer] #8
  . PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #1
  . PRT STD_OUT @@memory[@@pointer]
> ADD @@pointer #1
+ ADD @@memory[@@pointer] #2
  . PRT STD_OUT @@memory[@@pointer]"