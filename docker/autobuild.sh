#!/bin/sh

debug=false
#make sure to expose this port from docker
debug_port=40000

application_binary_path="/go/src/github.com/mihnealun/prox/main"

build_application() {
  echo "-=Building=- $application_binary_path"
  go build -gcflags "all=-N -l" -mod=vendor main.go

  return
}

run_mode_decider() {
  if ${debug}; then
    echo "-=Running in DEBUG mode=- $application_binary_path"
    dlv --listen=:${debug_port} --headless=true --api-version=2 exec ${application_binary_path}
  else
    echo "-=Running=- $application_binary_path"
    ${application_binary_path}
  fi

  return
}

start() {
  last_execution_status=$1

  if [ "$last_execution_status" != 1 ] && [ "$last_execution_status" != 0 ]; then
    (build_application && run_mode_decider) || start $?
  fi

  exit 1

  return
}

# listen for application files changes
inotifywait -mqr --timefmt '%d/%m/%y %H:%M' --format '%T %w %f' -e modify ./ | while read date time dir file; do
  ext="${file##*.}"

  if [ "$ext" = "go" ] || [ "${file}" = ".env" ]; then
    echo "$file changed @ $time $date, rebuilding..."
    #sleep 5
    pkill -f ${application_binary_path}
  fi
done &

# run application
start || echo "application was unable to start or was restarted, the container will be restarted by docker"
