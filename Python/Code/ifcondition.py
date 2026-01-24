
age= input("Enter your age: ")

if int(age) >= 18:
    print("You are eligible to vote.")
else:
    print("You are not eligible to vote yet.")


numb = 7
num= (int(input("Guess a number: ")))

while num != numb:
    print("Try again!")
    num= input("Guess a number: ")
print("Congratulations! You guessed it right.")
