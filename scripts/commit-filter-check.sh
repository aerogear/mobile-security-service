#!/bin/bash

# Get the current branch and apply it to a variable
currentbranch=`git branch | grep \* | cut -d ' ' -f2`

# Gets the commits for the current branch and outputs to file
git log $currentbranch --pretty=format:"%H" --not master > shafile.txt

# loops through the file an gets the message
for i in `cat ./shafile.txt`;
do 
  # gets the git commit message based on the sha
  gitmessage=`git log --format=%B -n 1 "$i"`

  # fix(commit-filter-check): add commit messages (AEROGEAR-038928990423)
  messagecheck=`echo $gitmessage | grep -w "feat\|fix\|docs"`
  if [ -z "$messagecheck" ]
  then 
      echo "Your commit message must beging with one of the following"
      echo "  feat(feature-name)"
      echo "  fix(fix-name)"
      echo "  docs(docs-change)"
  fi
  messagecheck=`echo $gitmessage | grep "(AEROGEAR-"`
  if  [ -z "$messagecheck" ]
  then 
      echo "Your commit message must end with the following"
      echo "  (AEROGEAR-****)"
      echo "Where **** is the Jira number"
  fi
  messagecheck=`echo $gitmessage | grep "): "`
  if  [ -z "$messagecheck" ]
  then 
      echo "Your commit message has a formatting error please take note of special characters '():' position and use in the example below"
      echo "   type(some txt): some txt (AEROGEAR-****)"
      echo "Where 'type' is fix, feat or docs and **** is the Jira number"
  fi

  messagecheck=`echo $gitmessage | grep -w "feat\|fix\|docs" | grep "(AEROGEAR-" | grep "): "`

 

  # check to see if the messagecheck var is empty
  if [ -z "$messagecheck" ]
  then  
        echo "'$i' commit message failed"
        rm shafile.txt >/dev/null 2>&1
        set -o errexit
        #break
  else
        echo "$messagecheck"
        echo "'$i' commit message passed"
  fi  
done
rm shafile.txt  >/dev/null 2>&1