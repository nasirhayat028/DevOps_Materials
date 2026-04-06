#!/bin/bash

echo "Error Handling"

clone(){
	echo "Code Cloning..."
	git clone https://github.com/LondheShubham153/django-notes-app.git || return 1
	echo "Code Sucessfully Cloned..."
}

if ! clone; then 
	echo "Code Already Exist"
	exit 1
fi 

echo "Code Clone Done"


