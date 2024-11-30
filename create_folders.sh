#!/bin/bash

if [ $# -eq 0 ]; then
	echo "Error: provide a source directory with the templates as an argument"
	echo "Usage: $0 <path>"
	exit 1
fi

SOURCE_DIR="$1"

# check if the directory exists
if [ ! -d "$SOURCE_DIR" ]; then
	echo "Error: directory provided ($SOURCE_DIR) does not exist"
	exit 1
fi


for i in {1..25}
do
	# create the day folder
	mkdir -p "$i"

	# copy the files from the source dir into the newly created folder
	cp -r "$SOURCE_DIR"/* "$i/"
	echo "created folder $i"
done


echo "Done"
