
country = ["Pakistan", "India", "China", "Bangladesh"]

print("Countries List:", country)

print("First Country:", country[0])
print("Length of Country:", len(country))

del country[2]

print("Updated Countries List:", country)

country.append("Sri Lanka")

print("Countries after append:", country)

country.insert(1, "Afghanistan")

print("Countries after insert:", country)

country.sort()

print("Sorted Countries List:", country)

country.reverse()   

print("Reversed Countries List:", country)