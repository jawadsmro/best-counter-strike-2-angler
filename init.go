package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.0" )

func pawSTnNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEHCiZcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iME1LuWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcAkFRqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ntujb4aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BeCvXKZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADtimCAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEcPlPxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5j7r8gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiTSlWAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJX9WNc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfflFEGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func th1BCR1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8E7FuK6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7dzQVqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2w6nALWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8pkgZkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGDsJIwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOPtc3GDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3UqjdneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBiXIHkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFBVmf1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 380cG7eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MT39zDHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yf3TQHftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOwXBeIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVIwv9rxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X98ondrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHprVZlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6A2PVsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuYjdwYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func He8p7uxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3y8nvBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APcGPRCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leRVEkIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ptDtpDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8gjvNFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UprE8zeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqJvsxJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUefIJP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRNiWuhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbVFYkmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1owLFq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nSgjkjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWhT3cczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWYzCDZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fber24BLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5KZB3r3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeFsr8T4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZX67uLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zROKb5O9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpXHDM4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcJCPKKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnwbWmwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fm3deJFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agz4b9OIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSLgQPOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPif09pwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBEszxJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMELPrOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OId57sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCXIj2YLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJM6stScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiGlyBw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vI9xKipEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySskA5O0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4Z39HQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xXYZYgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpz3PIuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDztp26LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJMxSrWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7mAkasFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lT0B6JfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3dpmF3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OcwoiUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xw6zcwokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peNqoTMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ve4PaBkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBgzFkjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdRDJqBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8D7YEQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juXpAkafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFjkIkNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZD28bdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DthKM7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsUDs6ZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lI76O1WBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q23hOp7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnGpcyktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMWzakziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZF0BhgxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIvywOQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daFReMQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMZ3hikfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRkzWKRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRZbJASaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkqrdqyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yghg9jJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fy9bSj86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWCKjX6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func re1mTQknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4NdRrwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXoyeUeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfGHP9AWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOllkXWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsrosF2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsYQPjJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6HeaD3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IzZwXQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDyf0SjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlZQAbDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIL18DcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ly90xFfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7GcWgYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wg1bjo0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5LQjdhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUyMUjSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThkDa5UbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfmS9IbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2ZVynndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cA5L7nZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7blY1P0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seApyXjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhPKu5J2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xfc0WvRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3PvXKALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXfG9DmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqCiKTKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KC9VfWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hI4O7RHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUav5tDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G92WLqfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USMEwEmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Kk8gfSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiRK1LhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrNdV6lvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaRbMhBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPp5LbZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGrMhum8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7UVNIpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uu8BRO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jsPeENaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okoEtoPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rO5y7LW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtIZt9DCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6CgNQBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1JalhryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgOgaHXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPIxFdYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymF2ZyyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBWvFzzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vidDfnhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWGC9L1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZWKI7ZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0HI9XrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLdfH2dCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4W1R8duWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNjVaCwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j94iIUbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02Jjd6GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZJ6fAQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8GPWFAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iO5rMxlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeF7OGZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gfWgnqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUxUHY8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxKFlCSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9G9XDk0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGiCyySwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFTB4zf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bp1b10YaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRBLvz6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DTHw4qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4jmlobIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6bQPucbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNJ0N4y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9qGejgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ee7AhlJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMf6HqA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyymqUdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrUd11RJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEoDOcs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5vYhfRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi1eh3l5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0RIQ3QsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Caf2dHjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oh3IHdrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vi97oyE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYFCztL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ewcdcj20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uevkT8JEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmGOIfGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7g6J4cd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWIAlqa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func un0hW9GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHkDUhX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1esalfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVkeBGF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jP1gOxMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K90WE3pQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDjzgUi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XKTlATjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwXI9Ru3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URFauVA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fh4btg4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yRTCk0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZJyatGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMLh7nXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzXVRseAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3WW57bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9DsKUaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgRVzR9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmlkuIP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

