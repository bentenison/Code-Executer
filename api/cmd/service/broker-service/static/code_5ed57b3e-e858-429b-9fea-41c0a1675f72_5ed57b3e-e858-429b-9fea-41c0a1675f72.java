public class Main {

    // User function logic goes here
    public static int sumOfDigits(int number) {
    int sum = 0;
    while (number > 0) {
        sum += number % 10;
        number /= 10;
    }
    return sum;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {123, 6},  // Test case 1: Sum of digits of 123 is 6
        {987, 24}  // Test case 2: Sum of digits of 987 is 24
    };
    for (int[] testCase : testCases) {
        int input = testCase[0];
        int expectedOutput = testCase[1];
        int output = sumOfDigits(input);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}