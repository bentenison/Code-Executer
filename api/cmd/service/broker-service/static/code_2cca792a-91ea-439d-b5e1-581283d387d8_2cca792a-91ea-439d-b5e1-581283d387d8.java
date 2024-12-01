public class Main {

    // User function logic goes here
    public static int countVowels(String str) {
    int count = 0;
    for (char c : str.toLowerCase().toCharArray()) {
        if ("aeiou".indexOf(c) != -1) {
            count++;
        }
    }
    return count;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[][] testCases = {
        {"hello", "2"},  // Test case 1: 'hello' has 2 vowels
        {"abc", "1"}    // Test case 2: 'abc' has 1 vowel
    };
    for (String[] testCase : testCases) {
        String input = testCase[0];
        int expectedOutput = Integer.parseInt(testCase[1]);
        int output = countVowels(input);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}