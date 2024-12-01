public class sample {
    public static boolean isLeapYear(int year) {
        if ((year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)) {
            return true;
        }
        return false;
    }

    public static void main(String[] args) {
        boolean allTestsPassed = true;
        allTestsPassed &= testLeapYear(2020, true); // Test case 1
        allTestsPassed &= testLeapYear(1900, false); // Test case 2
        allTestsPassed &= testLeapYear(2000, true); // Test case 3
        allTestsPassed &= testLeapYear(2024, true); // Test case 4

        if (allTestsPassed) {
            System.out.println(true);  // All test cases passed
        } else {
            System.out.println(false); // Some test cases failed
        }
    }

    private static boolean testLeapYear(int year, boolean expected) {
        return isLeapYear(year) == expected;
    }
}