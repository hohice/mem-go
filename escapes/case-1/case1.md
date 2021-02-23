```shell
$ go build -gcflags '-m -l' main.go
# command-line-arguments
./main.go:10:9: &User literal escapes to heap
```
GetUserInfo() 返回的是指针对象，引用被返回到了方法之外了。因此编译器会把该对象分配到堆上，而不是栈上。
```shell
$ go tool compile -S main.go
"".GetUserInfo STEXT size=117 args=0x8 locals=0x18
        0x0000 00000 (main.go:9)        TEXT    "".GetUserInfo(SB), ABIInternal, $24-8
        0x0000 00000 (main.go:9)        MOVQ    (TLS), CX
        0x0009 00009 (main.go:9)        CMPQ    SP, 16(CX)
        0x000d 00013 (main.go:9)        PCDATA  $0, $-2
        0x000d 00013 (main.go:9)        JLS     110
        0x000f 00015 (main.go:9)        PCDATA  $0, $-1
        0x000f 00015 (main.go:9)        SUBQ    $24, SP
        0x0013 00019 (main.go:9)        MOVQ    BP, 16(SP)
        0x0018 00024 (main.go:9)        LEAQ    16(SP), BP
        0x001d 00029 (main.go:9)        PCDATA  $0, $-2
        0x001d 00029 (main.go:9)        PCDATA  $1, $-2
        0x001d 00029 (main.go:9)        FUNCDATA        $0, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x001d 00029 (main.go:9)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x001d 00029 (main.go:9)        FUNCDATA        $2, gclocals·bfec7e55b3f043d1941c093912808913(SB)
        0x001d 00029 (main.go:10)       PCDATA  $0, $1
        0x001d 00029 (main.go:10)       PCDATA  $1, $0
        0x001d 00029 (main.go:10)       LEAQ    type."".User(SB), AX
        0x0024 00036 (main.go:10)       PCDATA  $0, $0
        0x0024 00036 (main.go:10)       MOVQ    AX, (SP)
        0x0028 00040 (main.go:10)       CALL    runtime.newobject(SB)
        0x002d 00045 (main.go:10)       PCDATA  $0, $1
        0x002d 00045 (main.go:10)       MOVQ    8(SP), AX
        0x0032 00050 (main.go:10)       MOVQ    $13746731, (AX)
        0x0039 00057 (main.go:10)       MOVQ    $6, 16(AX)
        0x0041 00065 (main.go:10)       PCDATA  $0, $2
        0x0041 00065 (main.go:10)       LEAQ    go.string."hohice"(SB), CX
        0x0048 00072 (main.go:10)       PCDATA  $0, $1
        0x0048 00072 (main.go:10)       MOVQ    CX, 8(AX)
        0x004c 00076 (main.go:10)       MOVQ    $32, 32(AX)
        0x0054 00084 (main.go:10)       PCDATA  $0, $2
        0x0054 00084 (main.go:10)       LEAQ    go.string."https://github.com/hohice/mem-go"(SB), CX
        0x005b 00091 (main.go:10)       PCDATA  $0, $1
        0x005b 00091 (main.go:10)       MOVQ    CX, 24(AX)
        0x005f 00095 (main.go:10)       PCDATA  $0, $0
        0x005f 00095 (main.go:10)       PCDATA  $1, $1
        0x005f 00095 (main.go:10)       MOVQ    AX, "".~r0+32(SP)
        0x0064 00100 (main.go:10)       MOVQ    16(SP), BP
        0x0069 00105 (main.go:10)       ADDQ    $24, SP
        0x006d 00109 (main.go:10)       RET
        0x006e 00110 (main.go:10)       NOP
        0x006e 00110 (main.go:9)        PCDATA  $1, $-1
        0x006e 00110 (main.go:9)        PCDATA  $0, $-2
        0x006e 00110 (main.go:9)        CALL    runtime.morestack_noctxt(SB)
        0x0073 00115 (main.go:9)        PCDATA  $0, $-1
        0x0073 00115 (main.go:9)        JMP     0
```