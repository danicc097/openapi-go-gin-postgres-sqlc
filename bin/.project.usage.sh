#!/usr/bin/env bash

source ".helpers.sh"

readonly MAX_COMMENT_LEN=88

usage() {
  command_comments_parser() {
    head -$((${lns[$i]} - 1)) $0 |
      tac |
      sed -n '/#/!q;p' |
      tac |
      awk '{$1=$1;print}'
  }

  command_options_comments_parser() {
    tail -n +$((${lns[$i]} + 1)) $0 |
      sed -n '/^[[:blank:]]*#/!q;p' |
      awk '{$1=$1;print}'
  }

  construct_row() {
    comment_parser="$1"
    for i in ${!lns[@]}; do
      comment_paragraph="$($comment_parser)"
      ROWS["${rows[$i]}"]="$comment_paragraph"
      mapfile -t comments <<<"${ROWS[${rows[$i]}]}"
      # Parse # Args:
      for comment in "${comments[@]}"; do
        comment="$(clean_comment "$comment")"
        args="-"
        if [[ ${comment,,} == args:* ]]; then
          args=$(clean_args "$comment")
        fi
        ROW_ARGS[${rows[$i]}]="$args"
      done
    done

    for i in "${!rows[@]}"; do
      mapfile -t comments <<<"${ROWS[${rows[$i]}]}"
      for j in "${!comments[@]}"; do
        comment="$(clean_comment "${comments[$j]}")"
        if [[ ${comment,,} == args:* ]]; then
          continue
        fi

        # if its the first doc line show all info
        if [[ $j = 0 ]]; then
          docs+=("$(
            printf -- "%s\t%s\t%s" \
              "${rows[$i]}" \
              "${ROW_ARGS[${rows[$i]}]}" \
              "$comment"
          )")
          continue
        fi

        # if its a remaining comment line
        docs+=("$(
          printf -- "%s\t%s\t%s" \
            "" \
            "" \
            "$comment"
        )")
      done
    done

    column -t \
      --separator $'\t' \
      --output-width 150 \
      --table-noextreme C2 \
      --table-noheadings \
      --table-wrap C3 \
      --table-columns C1,C2,C3 < <(printf "    %s\n" "${docs[@]}")
  }

  declare -A ROWS ROW_ARGS
  declare docs rows X_OPTIONS

  for c in "${COMMANDS[@]}"; do
    shopt -s extdebug
    lns+=("$(declare -F x.$c | awk '{print $2}')")
    rows+=("${c}")
    shopt -u extdebug
  done

  x_functions="$(construct_row command_comments_parser)"

  lns=()
  rows=()
  docs=()

  parse_x_options X_OPTIONS

  for c in "${X_OPTIONS[@]}"; do
    lns+=("${c##*)}")
    rows+=("${c%%)*}")
  done

  x_options="$(construct_row command_options_comments_parser)"

  cat <<EOF

$BOLD$UNDERSCORE$(basename $0)$OFF centralizes all relevant project commands.

${BOLD}USAGE:
    $RED$(basename $0) x.function [--x-option ...] args [optional args]$OFF

${BOLD}x.functions:$OFF
$(echo "${x_functions}" |
    sed -E 's/    ([[:alnum:][:punct:]]*)(.*)/    '"$BLUE$BOLD"'\1'"$OFF"'\2''/')

${BOLD}--x-options:$OFF
$(echo "${x_options}" |
      sed -E 's/    ([[:alnum:][:punct:]]*)(.*)/    '"$GREEN$BOLD"'\1'"$OFF"'\2''/')
EOF

}

# gets all --x-options values
parse_x_options() {
  local -n __arr="$1" # pass ref by name
  while IFS= read -r line; do
    __arr+=("$(awk '{$1=$1;print $1 $NF}' <<<"$line")")
  done < <(sed -nr '/.*(--x-[=*[:alnum:]_-]+[)]+).*/{p;=}' $0 | sed '{N;s/\n/ /}')
  mapfile -t __arr < \
    <(LC_COLLATE=C sort < <(printf "%s\n" "${__arr[@]}"))
}

clean_comment() {
  tmp="$1"
  tmp="${tmp//\#/}"
  comment="${tmp#* }"
  [[ -z $comment ]] && comment="·"
  ((${#comment} > MAX_COMMENT_LEN)) && comment="${comment:0:MAX_COMMENT_LEN}..."
  echo "$comment"
}

clean_args() {
  tmp="$1"
  tmp="${tmp,,##*args\:}"
  args="${tmp#* }"
  echo "$args"
}
