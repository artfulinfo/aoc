# Day 1: Calorie Counting
import pandas as pd
import numpy as np

df = pd.read_csv("input.txt", header=None, dtype=str, skip_blank_lines=False)

list = df.iloc[:,[0]].values.tolist()

# Create empty lists to construct the dataframe columns
elves = []
calories = []
elf = 1

for i, val in enumerate(list):
    if np.isnan(float(val[0])):
        elf += 1
    else:
        elves.append(elf)
        calories.append(int(val[0]))

delf = pd.DataFrame({"elf": elves, "calories": calories})

# Aggregate to get totals and sort highest values to top
delf = delf.groupby(by=delf["elf"], as_index=False) \
            .sum() \
            .sort_values(by="calories", inplace=True, ascending=False) \
            .reset_index(inplace=True)

# Part 1 answer
most_cals_solo = delf["calories"].iloc[0]
print(most_cals_solo)

# Part 2 answer
most_cals_trio = delf["calories"].iloc[:3].sum()
print(most_cals_trio)
