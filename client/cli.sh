#!/usr/bin/env bash

# . /home/dracher/.virtualenvs/vHttpie/bin/activate

PORT=8090

login()
{
  http POST :$PORT/login username=$1 password=$2 | jq -r '.token'
}

status()
{
  http :$PORT/current/scheduler
}


getconfig()
{
  http :$PORT/config/$1
}

option="${1}" 
case ${option} in 
  "login") USER="${2}" PASS="${3}"
    echo "perform login action"
    login $USER $PASS
    ;;

  "status")
    echo "querying current status"
    status
    ;;

  "cfg") CFGTYPE="${2}"
    echo "get config of $CFGTYPE"
    getconfig $CFGTYPE
    ;;

  *)
    echo "`basename ${0}`:usage: [-f file] | [-d directory]" 
    exit 1 # Command to come out of the program with status 1
    ;; 
esac
