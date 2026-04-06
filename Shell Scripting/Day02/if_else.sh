#!/bin/bash
echo "Hi Dears!"

read -p "Enter your age: " age

# Use $age to reference the value and ensure spaces inside brackets
if [ "$age" -ge 18 ]; then
   echo "You're eligible to vote!!"
elif [ "$age" -lt 18 ]; then
   echo "You're not eligible to vote!!"
else
   echo "Enter in integer format"
fi

