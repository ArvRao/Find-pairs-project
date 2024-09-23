### Input: Array of numbers and target value
### Output: Pair of indices whose values equal target value

Edge cases:
1. Result can contain duplicates if present in input
2. If empty, return empty response
3. If no matching values, return empty response

# Example 1:
input:
{
    "numbers": [1,2,3,4,5],
    "target": 6
}

Output:
{
    "solutions": [
        [
            0,
            4
        ],
        [
            1,
            3
        ]
    ]
}

# Example 2:
Input:
{
    "numbers": [],
    "target": 6
}

Output:
{
    "solutions": []
}

# Example 3:
Input:
{
    "numbers": [1,1,2,3,3,4,5,5],
    "target": 6
}
Output:
{
    "solutions": [
        [
            0,
            6
        ],
        [
            0,
            7
        ],
        [
            1,
            6
        ],
        [
            1,
            7
        ],
        [
            2,
            5
        ],
        [
            3,
            4
        ]
    ]
}


Example 4:
Input:
{
    "numbers": [1,2,3],
    "target": 14
}
Output:
{
    "solutions": []
}