# User function logic goes here
def maxPoints(points):
    if not points:
        return 0
    def gcd(a, b):
        if b == 0:
            return a
        return gcd(b, a % b)
    n = len(points)
    max_count = 1
    for i in range(n):
        slopes = {}
        same = 0
        vertical = 0
        for j in range(i + 1, n):
            dx = points[j][0] - points[i][0]
            dy = points[j][1] - points[i][1]
            if dx == 0 and dy == 0:
                same += 1
            elif dx == 0:
                vertical += 1
            else:
                g = gcd(dx, dy)
                slope = (dx // g, dy // g)
                slopes[slope] = slopes.get(slope, 0) + 1
        max_count = max(max_count, vertical + same + 1)
        for count in slopes.values():
            max_count = max(max_count, count + same + 1)
    return max_count

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([[1,1],[2,2],[3,3]]),
        ([[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]])
    ]
    expected_outputs = [
        3,
        4
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = maxPoints(test_input)
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)