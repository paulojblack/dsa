#lang racket
; (define x 100)

(define (fib n)
    (cond
        ((<= n 1) n)
        (else (+ (fib (- n 1)) (fib (- n 2))))
    )
)

(displayln (fib 10))


; (define (fib n)
;   (cond
;     [(<= n 1) n]                     ; If n is 0 or 1, return n
;     [else (+ (fib (- n 1)) (fib (- n 2)))]))  ; Otherwise, return fib(n-1) + fib(n-2)

; (displayln (fib 10))  ;; This will display the 10th Fibonacci number