#include <stdio.h>

void rotateArray(int arr[], int size, int k) {
    k = k % size; // In case k is larger than size
    int temp[k];
    for (int i = 0; i < k; i++) {
        temp[i] = arr[size - k + i];
    }
    for (int i = size - 1; i >= k; i--) {
        arr[i] = arr[i - k];
    }
    for (int i = 0; i < k; i++) {
        arr[i] = temp[i];
    }
}

int main() {
    int test_cases[][5] = {{3, 4, 5, 1, 2}, {9, 10, 0, 7, 8}, {1, 2, 3, 4, 5}};
    int expected[][5] = {{4, 5, 1, 2, 3}, {9, 10, 7, 8, 0}, {4, 5, 1, 2, 3}};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        rotateArray(test_cases[i], 5, 3);
        for (int j = 0; j < 5; j++) {
            if (test_cases[i][j] != expected[i][j]) {
                all_passed = 0;
                printf("Failed for Input: {%d, %d, %d, %d, %d}. Expected: {%d, %d, %d, %d, %d}, Got: {%d, %d, %d, %d, %d}\n",
                        test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4],
                        expected[i][0], expected[i][1], expected[i][2], expected[i][3], expected[i][4],
                        test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4]);
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