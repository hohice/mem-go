```shell
$ go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:9:18: moved to heap: u
```

```shell
$ go tool compile -S main.go       
"".GetUserInfo STEXT size=157 args=0x30 locals=0x28
        0x0000 00000 (main.go:9)        TEXT    "".GetUserInfo(SB), ABIInternal, $40-48
        0x0000 00000 (main.go:9)        MOVQ    (TLS), CX
        0x0009 00009 (main.go:9)        CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:9)        PCDATA  $0, $-2
        0x000d 00013 (main.go:9)        JLS     147
        0x0013 00019 (main.go:9)        PCDATA  $0, $-1
        0x0013 00019 (main.go:9)        SUBQ    $40, SP
        0x0017 00023 (main.go:9)        MOVQ    BP, 32(SP)
        0x001c 00028 (main.go:9)        LEAQ    32(SP), BP
        0x0021 00033 (main.go:9)        PCDATA  $0, $-2
        0x0021 00033 (main.go:9)        PCDATA  $1, $-2
        0x0021 00033 (main.go:9)        FUNCDATA        $0, gclocals·ac76e2edf186c2610272ed5fd9e847c7(SB)
        0x0021 00033 (main.go:9)        FUNCDATA        $1, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
        0x0021 00033 (main.go:9)        FUNCDATA        $2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0021 00033 (main.go:9)        PCDATA  $0, $1
        0x0021 00033 (main.go:9)        PCDATA  $1, $0
        0x0021 00033 (main.go:9)        LEAQ    type."".User(SB), AX
        0x0028 00040 (main.go:9)        PCDATA  $0, $0
        0x0028 00040 (main.go:9)        MOVQ    AX, (SP)
        0x002c 00044 (main.go:9)        CALL    runtime.newobject(SB)
        0x0031 00049 (main.go:9)        PCDATA  $0, $1
        0x0031 00049 (main.go:9)        MOVQ    8(SP), AX
        0x0036 00054 (main.go:9)        PCDATA  $0, $-2
        0x0036 00054 (main.go:9)        PCDATA  $1, $-2
        0x0036 00054 (main.go:9)        CMPL    runtime.writeBarrier(SB), $0
        0x003d 00061 (main.go:9)        JNE     104
        0x003f 00063 (main.go:9)        MOVQ    "".u+48(SP), CX
        0x0044 00068 (main.go:9)        MOVQ    CX, (AX)
        0x0047 00071 (main.go:9)        MOVUPS  "".u+56(SP), X0
        0x004c 00076 (main.go:9)        MOVUPS  X0, 8(AX)
        0x0050 00080 (main.go:9)        MOVUPS  "".u+72(SP), X0
        0x0055 00085 (main.go:9)        MOVUPS  X0, 24(AX)
        0x0059 00089 (main.go:10)       PCDATA  $0, $0
        0x0059 00089 (main.go:10)       PCDATA  $1, $1
        0x0059 00089 (main.go:10)       MOVQ    AX, "".~r1+88(SP)
        0x005e 00094 (main.go:10)       MOVQ    32(SP), BP
        0x0063 00099 (main.go:10)       ADDQ    $40, SP
        0x0067 00103 (main.go:10)       RET
        0x0068 00104 (main.go:9)        PCDATA  $0, $-2
        0x0068 00104 (main.go:9)        PCDATA  $1, $-2
        0x0068 00104 (main.go:9)        MOVQ    AX, "".&u+24(SP)
        0x006d 00109 (main.go:9)        LEAQ    type."".User(SB), CX
        0x0074 00116 (main.go:9)        MOVQ    CX, (SP)
        0x0078 00120 (main.go:9)        MOVQ    AX, 8(SP)
        0x007d 00125 (main.go:9)        LEAQ    "".u+48(SP), CX
        0x0082 00130 (main.go:9)        MOVQ    CX, 16(SP)
        0x0087 00135 (main.go:9)        CALL    runtime.typedmemmove(SB)
        0x008c 00140 (main.go:10)       MOVQ    "".&u+24(SP), AX
        0x0091 00145 (main.go:9)        JMP     89
        0x0093 00147 (main.go:9)        NOP
        0x0093 00147 (main.go:9)        PCDATA  $1, $-1
        0x0093 00147 (main.go:9)        PCDATA  $0, $-2
        0x0093 00147 (main.go:9)        CALL    runtime.morestack_noctxt(SB)
        0x0098 00152 (main.go:9)        PCDATA  $0, $-1
        0x0098 00152 (main.go:9)        JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 80  eH..%....H;a....
        0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
        0x0020 20 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00   H......H..$....
        0x0030 00 48 8b 44 24 08 83 3d 00 00 00 00 00 75 29 48  .H.D$..=.....u)H
        0x0040 8b 4c 24 30 48 89 08 0f 10 44 24 38 0f 11 40 08  .L$0H....D$8..@.
        0x0050 0f 10 44 24 48 0f 11 40 18 48 89 44 24 58 48 8b  ..D$H..@.H.D$XH.
        0x0060 6c 24 20 48 83 c4 28 c3 48 89 44 24 18 48 8d 0d  l$ H..(.H.D$.H..
        0x0070 00 00 00 00 48 89 0c 24 48 89 44 24 08 48 8d 4c  ....H..$H.D$.H.L
        0x0080 24 30 48 89 4c 24 10 e8 00 00 00 00 48 8b 44 24  $0H.L$......H.D$
        0x0090 18 eb c6 e8 00 00 00 00 e9 63 ff ff ff           .........c...
        rel 5+4 t=17 TLS+0
        rel 36+4 t=16 type."".User+0
        rel 45+4 t=8 runtime.newobject+0
        rel 56+4 t=16 runtime.writeBarrier+-1
        rel 112+4 t=16 type."".User+0
        rel 136+4 t=8 runtime.typedmemmove+0
        rel 148+4 t=8 runtime.morestack_noctxt+0
"".main STEXT nosplit size=1 args=0x0 locals=0x0
        0x0000 00000 (main.go:13)       TEXT    "".main(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000 (main.go:13)       PCDATA  $0, $-2
        0x0000 00000 (main.go:13)       PCDATA  $1, $-2
        0x0000 00000 (main.go:13)       FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:13)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:13)       FUNCDATA        $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:15)       PCDATA  $0, $-1
        0x0000 00000 (main.go:15)       PCDATA  $1, $-1
        0x0000 00000 (main.go:15)       RET
        0x0000 c3                                   
```