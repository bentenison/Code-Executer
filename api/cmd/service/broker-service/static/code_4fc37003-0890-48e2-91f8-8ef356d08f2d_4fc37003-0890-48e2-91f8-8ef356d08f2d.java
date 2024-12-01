public class Main {

    // User function logic goes here
    public static String checkNumber(int num) {
    if (num > 0) {
        return "Positive";
    } else if (num < 0) {
        return "Negative";
    } else {
        return "Zero";
    }
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {-10, "Negative"},  // Test case 1: -10 is Negative
        {5, "Positive"},    // Test case 2: 5 is Positive
        {0, "Zero"}         // Test case 3: 0 is Zero
    };
    for (int[] testCase : testCases) {
        int input = testCase[0];
        String expectedOutput = testCase[1];
        String output = checkNumber(input);
        if (!output.equals(expectedOutput)) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}