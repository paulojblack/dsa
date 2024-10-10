#lang racket
; (define x 100)

(define (fib n)
    (cond
        ((<= n 1) n)
        (else (+ (fib (- n 1)) (fib (- n 2))))
    )
)

(displayln (fib 10))
