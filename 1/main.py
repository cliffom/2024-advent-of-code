import os

DELIMITER = "   "

def sum_columns(col1, col2):
    total = 0
    for num1, num2 in zip(col1, col2):
        total += abs(num1 - num2)
    return total

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
    print(result)

if __name__ == "__main__":
    main()
