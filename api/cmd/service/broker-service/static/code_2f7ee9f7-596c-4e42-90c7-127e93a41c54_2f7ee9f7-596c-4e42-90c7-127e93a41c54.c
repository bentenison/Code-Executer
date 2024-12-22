#include <stdio.h>

void moveZeroes(int arr[], int size) {
    int last_non_zero = 0;
    for (int i = 0; i < size; i++) {
        if (arr[i] != 0) {
            arr[last_non_zero++] = arr[i];
        }
    }
    for (int i = last_non_zero; i < size; i++) {
        arr[i] = 0;
    }
}

int main() {
    int test_cases[][7] = {{0, 1, 2, 0, 3, 4, 0}, {1, 0, 3, 0, 0, 5, 6}, {1, 2, 3, 4}};
    int expected[][7] = {{1, 2, 3, 4, 0, 0, 0}, {1, 3, 5, 6, 0, 0, 0}, {1, 2, 3, 4}};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        moveZeroes(test_cases[i], 7);
        for (int j = 0; j < 7; j++) {
            if (test_cases[i][j] != expected[i][j]) {
                all_passed = 0;
                printf("Failed for Input: {%d, %d, %d, %d, %d, %d, %d}. Expected: {%d, %d, %d, %d, %d, %d, %d}, Got: {%d, %d, %d, %d, %d, %d, %d}\n", test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4], test_cases[i][5], test_cases[i][6], expected[i][0], expected[i][1], expected[i][2], expected[i][3], expected[i][4], expected[i][5], expected[i][6], test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4], test_cases[i][5], test_cases[i][6]);
            }
        }
    }
    if (all_passed) {
        printf("true\n");
    } else {
        printf("false\n");
    }
    return 0;
}