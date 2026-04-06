#!/bin/bash

echo "Nasir"

add(){
 local sum=$(( $1 + $2))
	echo "Addition of 10 and 5 is: $sum"
}

add 10 5

sub(){
 local sub=$(( $1 - $2))
	echo "Subtraction of 10 and 5 is: $sub"
}

sub 10 5
