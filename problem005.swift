// Problem 5: Smallest Multiple
//
// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?

var n = 0
loop: while true {
    ++n
    for i in 1...20 {
        if n % i != 0 { continue loop }
    }
    println(n)
    break loop
}
