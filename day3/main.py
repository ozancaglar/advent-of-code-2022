f = open("input.txt", "r")

raw_data = f.readlines()


def priority(letter):
    return ord(letter)-38 if letter.isupper() else ord(letter)-96


def split_row(row):
    return row[0:len(row)//2], row[len(row)//2:len(row)-1]


def unique_letters_in_row(row):
    return ''.join(set(row))


sumPartOne = 0
for row in raw_data:
    split_rows = split_row(row)
    for l in unique_letters_in_row(row):
        if l in split_rows[0] and l in split_rows[1]:
            sumPartOne += priority(l)
print(sumPartOne)

sumPartTwo = 0
j = 0
batched_row = []

for row in raw_data:
    clean_row = row.replace("\n", "")
    batched_row.append(clean_row)
    j += 1
    if j == 3:
        for l in unique_letters_in_row(clean_row):
            if l in batched_row[0] and l in batched_row[1] and l in batched_row[2]:
                sumPartTwo += priority(l)
        batched_row = []
        j = 0
print(sumPartTwo)
