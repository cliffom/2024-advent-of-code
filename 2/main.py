DELIMITER = " "


# convert_data takes a list of strings and returns an list of float64s
# we assume the input data is clean and everything can be cast properly
def convert_data(data: list):
    data_converted = []
    for item in data:
        number = int(item)
        data_converted.append(number)

    return data_converted


# data_is_valid verifies that the data in a list of numbers
# meets the following criteria:
# - The numbers are either all increasing or decreasing
# - Any two adjacent levels differ by at least one and at most three
# - Removing any 1 item makes the above true
def data_is_valid(data: list):
    threshold = 3

    def is_valid(data):
        diffs = [data[i] - data[i - 1] for i in range(1, len(data))]
        if len(diffs) == 0:  # A single element list is valid
            return True
        increasing = diffs[0] > 0
        return all(
            0 < abs(diff) <= threshold
            and (increasing and diff > 0 or not increasing and diff < 0)
            for diff in diffs
        )

    # Check if the list is already valid
    if is_valid(data):
        return True

    # Try removing each element and validate
    for i in range(len(data)):
        temp_data = data[:i] + data[i + 1 :]
        if is_valid(temp_data):
            return True

    return False


def main():
    with open("input.txt", "r") as file:
        valid_lines = 0
        for line in file:
            parts = line.strip().split(DELIMITER)
            data = convert_data(parts)

            if data_is_valid(data):
                valid_lines += 1

    print(valid_lines)


if __name__ == "__main__":
    main()
