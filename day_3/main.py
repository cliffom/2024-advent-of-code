import re

# We want to match input items that look like
# mul(x,y) where x and y are numbers.
# Ex: mul(2,3)
pattern = re.compile(r"mul\((\d+),(\d+)\)")


def main():
    with open("input.txt", "r") as file:
        instructions = file.read().replace("\n", "")

    unrefined_output = sum_of_valid_unrefined_instructions(instructions)
    refined_output = sum_of_valid_refined_instructions(instructions)

    print(f"Sum of unrefined instructions: {unrefined_output}")
    print(f"Sum of refined instructions: {refined_output}")


def sum_of_valid_unrefined_instructions(data: str) -> int:
    """
    Iterates through all valid instructions and sums the value
    of the product of each pair.
    """
    instructions = pattern.findall(data)
    return sum(int(num1) * int(num2) for num1, num2 in instructions)


def sum_of_valid_refined_instructions(data: str) -> int:
    """
    Iterates through all valid instructions and disables processing
    of instructions that follow a `don't` instruction. Instructions
    after `do` are processed normally.
    """
    matches = [(m.start(), m.end(), *m.groups()) for m in pattern.finditer(data)]

    result = []
    for start, _, num1, num2 in matches:
        substr = data[:start]
        if "don't" not in substr or substr.rfind("don't") < substr.rfind("do"):
            result.append((int(num1), int(num2)))

    return sum(num1 * num2 for num1, num2 in result)


if __name__ == "__main__":
    main()
