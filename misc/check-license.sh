#!/bin/sh -e

status=0
for f in `git ls-files | xargs grep -L "Copyright" | grep ".go" | grep -v vendor/`
do
    echo $f
    status=1
done

if [ $status != 0 ]
then
   exit $status 
fi

for f in `git ls-files | xargs grep "Copyright 201[2345]" -l | grep -v check-license.sh | grep -v vendor/`
do
	date=`git log -1 --format="%ad" --date=short -- $f`
	if [ `echo "$date" | grep ^2016` ]
	then
		echo $f $date
		status=1
	fi
done

exit $status
