#!/bin/bash
# startup and shutdown utility for home automation modules

if [ $(whoami) != "root" ]; then
  echo "Since we're controlling GPIO pins this must be run as root."
  exit 1    
fi

CWD=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
kill $(<.procs)

if [ $1 = "start" ]; then

CONF_FILE=conf*.y*
CONF=${CWD}/${CONF_FILE}
PROCS=""

if [ -f $CONF ]; then 
  echo "Found configuration file in home_automation root:"
  echo $CONF
else 
  echo "No configuration file found.  Exiting now."
  exit 1
fi

# start garage doors listener
echo "Starting garage_doors listener (needs root)..."
${CWD}/bin/server $CONF &
PROCS="$PROCS $!"

echo $PROCS > .procs

elif [ $1 = "stop" ]; then
echo "Stopping garage door controller processes..."
else
  echo "Please specify 'start' or 'stop'."
fi
