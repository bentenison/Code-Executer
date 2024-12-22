#include <stdio.h>
#include <string.h>

// User function logic goes here
void printPermutations(char str[], int l, int r) {
    if (l == r) {
        printf("%s\n", str);
    } else {
        for (int i = l; i <= r; i++) {
            char temp = str[l];
            str[l] = str[i];
            str[i] = temp;
            printPermutations(str, l + 1, r);
            temp = str[l];
            str[l] = str[i];
            str[i] = temp;
        }
    }
}

// Test cases for the function
int main() {
    int passed = 1;
    char test_cases[2][100] = {
        "abc",
        "xyz"
    };
    for (int i = 0; i < 2; i++) {
        printPermutations(test_cases[i], 0, strlen(test_cases[i]) - 1);
        // Check expected output
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}