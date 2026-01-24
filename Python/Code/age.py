
age1 = input("Enter your age: ")
age2 = input("Enter your friend's age: ")

if int(age1) > int(age2):
    print("You are older than your friend.")
elif int(age1) < int(age2):
    print("Your friend is older than you.")
else:
    print("You and your friend are of the same age.")