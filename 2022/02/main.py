# Day 2: Rock, paper, scissors
import pandas as pd

strat = pd.read_csv("input.txt", header=None, delimiter=" ")
shape_score = {"X": 1, "Y": 2, "Z": 3}

"""
Part 1
Opponent selection
A == Rock
B == Paper
C == Scissors

My selection
X == Rock
Y == Paper
Z == Scissors
----------------
Part 2
My selection
X == lose
Y == draw
Z == lose
"""

# Part 1 solution


def fight(opp, me):
    losing_scenario = ((opp == "A") & (me == "Z")) or (
        (opp == "B") & (me == "X")) or ((opp == "C") & (me == "Y"))
    winning_scenario = ((opp == "A") & (me == "Y")) or (
        (opp == "B") & (me == "Z")) or ((opp == "C") & (me == "X"))
    drawing_scenario = ((opp == "A") & (me == "X")) or (
        (opp == "B") & (me == "Y")) or ((opp == "C") & (me == "Z"))

    if losing_scenario:
        return "lose"
    elif winning_scenario:
        return "win"
    elif drawing_scenario:
        return "draw"


my_score = 0

for i, row in strat.iterrows():
    opp_select = strat.iat[i, 0]
    my_select = strat.iat[i, 1]
    result = fight(opp_select, my_select)
    if result == "win":
        my_score += 6
    elif result == "draw":
        my_score += 3
    my_score += shape_score.get(my_select, None)
    # print(result, opp_select, my_select, my_score)

print("Part 1 answer: " + str(my_score))

# Part 2 solution

shape_score = {"rock": 1, "paper": 2, "scissors": 3}
result_map = {"X": "lose", "Y": "draw", "Z": "win"}
my_score = 0


def fight_2(opp, result):
    rock_scenario = ((opp == "A") & (result == "draw")) or (
        (opp == "B") & (result == "lose")) or ((opp == "C") & (result == "win"))
    paper_scenario = ((opp == "A") & (result == "win")) or (
        (opp == "B") & (result == "draw")) or ((opp == "C") & (result == "lose"))
    scissors_scenario = ((opp == "A") & (result == "lose")) or (
        (opp == "B") & (result == "win")) or ((opp == "C") & (result == "draw"))
    if rock_scenario:
        return "rock"
    elif paper_scenario:
        return "paper"
    elif scissors_scenario:
        return "scissors"
    else:
        return "no match"


for i, row in strat.iterrows():
    opp_select = strat.iat[i, 0]
    result = result_map.get(strat.iat[i, 1])
    my_select = fight_2(opp_select, result)
    if result == "win":
        my_score += 6
    elif result == "draw":
        my_score += 3
    my_score += shape_score.get(my_select, None)
    # print(result, opp_select, my_select, my_score)
print("Part 2 answer: ", my_score)
