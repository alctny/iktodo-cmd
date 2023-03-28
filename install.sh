#! /bin/bash

program="todo"
bin="/home/$USER/.local/bin"

if [ -d ${bin} ]; then
  mkdir -p ${bin}
fi

go build -ldflags="-s -w" -o ${program}

mv ${program} ${bin}