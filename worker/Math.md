# Grammar for math expressions

```
package "PP/worker/grammar"

Start
    : AddSubCont
    | Start AddSub AddSubCont
    ;

AddSubCont
    : MulDivCont
    | AddSubCont MulDiv MulDivCont
    ;

MulDivCont
    : MinMaxCont
    | MulDivCont MinMax MinMaxCont
    ;

MinMaxCont
    : Unary
    | "(" Start ")"
    ;

Unary
    : "{" sym "}"
    | "[" sym "]"
    | "<" sym ">"
    ;

AddSub
    : "-"
    | "+"
    ;

MulDiv
    : "*"
    | "/"
    ;

MinMax
    : "!"
    | "?"
    ;

sym : {letter};
```
