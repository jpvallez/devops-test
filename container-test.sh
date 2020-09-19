#!/bin/bash

#Passed in parameter is hostname of container
endpoint=$1

correctResponse=$( curl $endpoint:8080/version )

echo "Check "$endpoint":8080/version returns valid JSON"
if (jq -e . >/dev/null 2>&1 <<<"$correctResponse")
then
    echo "PASS Parsed JSON successfully."
else
    echo "Failed to parse JSON, or got false/null"
    exit 1
fi

shaResponse=$( curl $endpoint:8080/version -s | jq '.LastCommitSha' )
shaLen=${#shaResponse}
echo "Check JSON contains Last Commit SHA. Check string length 40 characters..."
#42 because of encasing quotes
if [ $shaLen == 42 ]
then
    echo "PASS SHA is valid."
else
    echo "Failed to get Last Commit SHA."
    exit 1
fi

versionResponse=$( curl $endpoint:8080/version -s | jq '.Version' )
echo "Check Version is set"
if [ $versionResponse != "undefined" ]
then
    echo "PASS Version response OK. $versionResponse"
else
    echo "Failed to get Version."
    exit 1
fi

descriptionResponse=$( curl $endpoint:8080/version -s | jq '.Description' )
echo "Check Description is not empty"
if [ "$descriptionResponse" != "" ]
then
    echo "PASS Descripion OK."
else
    echo "Failed to get Description"
    exit 1
fi

badResponse=$( curl "$endpoint":8080 -s )
echo "Check $endpoint:8080 returns 404"
if [ "$badResponse" == "404 page not found" ]
then
    echo "PASS 404 response OK"
else
    echo "Failed 404 response"
    exit 1
fi

echo "ALL TESTS PASSED"
sleep 5