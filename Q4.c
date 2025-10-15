#include <stdio.h>

int bin(int n) {
    if (n == 0) {
        return 0;
    } else {
        return (n % 2) + 10 * bin(n / 2);
    }
}

int main() {
    int k, n;
    scanf("%d", &k);
    for (int i = 0; i < k; i++) {
        scanf("%d", &n);
        if (n == 0) {
            printf("0\n");
        } else {
            printf("%d\n", bin(n));
        }
    }
    return 0;
}
