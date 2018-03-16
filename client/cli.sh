#!/usr/bin/env bash

# . /home/dracher/.virtualenvs/vHttpie/bin/activate

# ShortName=HP_Z600_03 BeakerName=hp-z600-03.qe.lab.eng.nay.redhat.com NicName=em2 Mac="01:18:a9:05:bf:8b:e6" AutoOnly:=true

PORT=8090
USER='yaniwang'
PASS='klopklop'

login()
{
  http POST :$PORT/login username=$USER password=$PASS | jq -r '.token'
}

status()
{
  http :$PORT/current/scheduler
}


getconfig()
{
  http :$PORT/config/$1
}

newhost()
{
  token=$(login)
  http POST :8090/auth/host ${1} ${2} ${3} ${4} ${5} --auth-type=jwt --auth="${token}"
}

updatehost()
{
  token=$(login)
  http PUT :8090/auth/host ${1} ${2} ${3} ${4} ${5} --auth-type=jwt --auth="${token}"
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

  "new_host") OPTION2="${2}"
    echo "add new host"
    newhost $OPTION2
    ;;

  "update_host") OPTION3="${2}"
    echo "update a host"
    updatehost $OPTION3
    ;;

  *)
    echo "`basename ${0}`:usage: [-f file] | [-d directory]"
    exit 1 # Command to come out of the program with status 1
    ;; 
esac
