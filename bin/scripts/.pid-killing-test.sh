#!/bin/bash

# $$ gives the top level shell.
# The same pid will be printed for $$,
# and each child's pid is $BASHPID:
echo "$$ - $BASHPID"
(
  echo "$$ - $BASHPID"
  (echo "$$ - $BASHPID")
)

run1=0
run2=0
pids=()

long-running-1() {
  while true; do
    echo "long-running-1 - run $run1"
    sleep 1
  done
}

long-running-2() {
  while true; do
    echo "long-running-2 - run $run2"
    sleep 1
  done
}

run-watch() {
  run_long_running() {
    # trap 'kill $(jobs -p)' EXIT
    # for pid in "${pids[@]}"; do
    #   pkill -P $pid || : # NOT getting killed.
    # done
    # pkill -P $BASHPID || :
    run1=$((run1 + 1))
    run2=$((run2 + 1))
    long-running-1 &
    pids+=("$!")
    long-running-2 &
    pids+=("$!")
  }

  run_long_running

  while true; do
    inotifywait \
      --monitor ./file-to-watch \
      --event=close_write \
      --format='%T %f' \
      --timefmt='%s' |
      while read -r event_time event_file; do
        printf "\nchange detected\n\n"
        for pid in "${pids[@]}"; do
          kill $pid 2>/dev/null || :
        done
        pids=()
        run_long_running
      done
  done
}

run-watch
