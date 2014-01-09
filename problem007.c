#include <stdio.h>

int main(void) {
	char isPrime;
	int i = 10001;
	int num = 2;
	int j = 2;
	for (num = 2; i > 0; num++) {
		isPrime = 1;
		for (j = 2; j < num; j++) {
			if (num % j == 0) {
				isPrime = 0;
			}
		}
		if (isPrime) {
			printf("%d prime number #%d\n", num, 10002-i);
			i--;
		}
	}
	return 0;
}
