import os

DELIMITER = "   "


def sum_columns(col1, col2):
    total = 0
    for num1, num2 in zip(col1, col2):
        total += abs(num1 - num2)
    return total


def get_similarity_score(column1, column2):
    count_map = {}
    for value in column2:
        count_map[value] = count_map.get(value, 0) + 1

    score = 0
    for value in column1:
        if value in count_map:
            score += value * count_map[value]

    return score


def main():
    col1 = []
    col2 = []

    with open("input.txt", "r") as file:
        for line in file:
            parts = line.strip().split(DELIMITER)
            num1 = int(parts[0])
            num2 = int(parts[1])

            col1.append(num1)
            col2.append(num2)

    col1.sort()
    col2.sort()

    result = sum_columns(col1, col2)
    similarity_score = get_similarity_score(col1, col2)
    print(result)
    print(similarity_score)


if __name__ == "__main__":
    main()
