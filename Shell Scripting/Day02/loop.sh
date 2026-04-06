#!/bin/bash

name="Nasir"
echo $name
for ((i=1;  i<=3; i++))
do
	mkdir demo"$i"
done

<< comment
now let's do it through argoment in terminal you have have to put argoments like ./loo.sh name argoment1 argoments2
comment

for ((num=$2; num<=$3; num++))
do
	mkdir -p "$1$num"
done
