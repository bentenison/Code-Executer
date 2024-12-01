public class Main {

    // User function logic goes here
    public static boolean isLeapYear(int year) {
    // Your code here
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {2020, 1},  // Test case 1: 2020 is a leap year
        {2021, 0}   // Test case 2: 2021 is not a leap year
    };
    for (int[] testCase : testCases) {
        int input = testCase[0];
        boolean expectedOutput = testCase[1] == 1;
        boolean output = isLeapYear(input);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}