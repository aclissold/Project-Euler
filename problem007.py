#!/usr/bin/env python3

def main():
    isPrime = True
    i = 10001
    num = 2
    while i > 0:
        isPrime = True
        for j in range(2, num):
            if num % j == 0:
                isPrime = False
        if isPrime:
            print(num, "prime number #", 10002-i)
            i -= 1
        num += 1

# 	for num := 2; i > 0; num++ {
# 		isPrime = true;
# 		for j := 2; j < num; j++ {
# 			if num % j == 0 {
# 				isPrime = false;
# 			}
# 		}
# 		if isPrime {
# 			fmt.Printf("%d prime number #%d\n", num, 10002-i)
# 			i--
# 		}
# 	}
# }

if __name__ == '__main__':
    main()
