public class Main {

    // User function logic goes here
    public static String reverseString(String str) {
    StringBuilder sb = new StringBuilder(str);
    return sb.reverse().toString();
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[][] testCases = {
        {"world", "dlrow"},  // Test case 1: 'world' reversed is 'dlrow'
        {"java", "avaj"}    // Test case 2: 'java' reversed is 'avaj'
    };
    for (String[] testCase : testCases) {
        String input = testCase[0];
        String expectedOutput = testCase[1];
        String output = reverseString(input);
        if (!output.equals(expectedOutput)) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}