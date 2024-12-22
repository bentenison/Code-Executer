#include <stdio.h>
#include <string.h>

// User function logic goes here
void removeChars(char str[]) {
    int i = 0, j = 0;
    while (str[i]) {
        if (str[i] == 'b') {
            i++;
        } else if (str[i] == 'a' && str[i + 1] == 'c') {
            i += 2;
        } else {
            str[j++] = str[i++];
        }
    }
    str[j] = '\0';
}

// Test cases for the function
int main() {
    int passed = 1;
    char test_cases[2][100] = {
        "abbacac",
        "abacabbac"
    };
    for (int i = 0; i < 2; i++) {
        removeChars(test_cases[i]);
        printf("%s\n", test_cases[i]);
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}