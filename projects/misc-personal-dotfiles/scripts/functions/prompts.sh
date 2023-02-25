#!/usr/bin/env bash

# prompts

# calclulate bash command
calculate_bash_command() {

    local bash_command=$(
        history -w /dev/stdout |
            tail -n 1
    )
    echo "$(("${#bash_command}" + 2))"

}

# format time
format_time() {

    local output

    output="${1/*./}"
    echo "${output::3}"

}

# after prompt
get_preexec_message() {

    export IFS=$'\n'
    local start_time

    start_time=$(format_time "${TRAP_DEBUG_TIME_START/*./}")
    start_time=$(date "+%H:%M:%S:${start_time}")
    mapfile -t start_time < \
        <(prompt_rectangle ' ' "${start_time}" ' ' "${GC_33}")

    echo "${start_time[*]}"

}

# before prompt
get_precmd_message() {

    export IFS=$'\n'
    local time_elapsed end_time output

    time_elapsed="${TRAP_DEBUG_TIME_END} - ${TRAP_DEBUG_TIME_START}"
    time_elapsed=$(printf "%.3f" "$(bc <<<"${time_elapsed}")")
    time_elapsed=$(convert_time "${time_elapsed}")
    end_time=$(format_time "${TRAP_DEBUG_TIME_END/*./}")
    end_time=$(date "+%H:%M:%S:${end_time}")

    mapfile -t end_time < \
        <(prompt_rectangle ' ' "${end_time}" ' ' "${GC_33}")
    mapfile -t time_elapsed < \
        <(prompt_rectangle ' ' "${time_elapsed}" ' ' "${GC_33}")

    output=(
        "${end_time[0]} ${time_elapsed[0]}"
        "${end_time[1]} ${time_elapsed[1]}"
        "${end_time[2]} ${time_elapsed[2]}"
    )

    echo "${output[*]}"

}

# greatest length
greatest_length() {

    local length

    [[ "${1}" -ge "${2}" && "${1}" -ge "${3}" ]] && length="${1}"
    [[ "${2}" -ge "${1}" && "${2}" -ge "${3}" ]] && length="${2}"
    [[ "${3}" -ge "${2}" && "${3}" -ge "${1}" ]] && length="${3}"
    [[ ! "${length}" || "${length}" -le 0 ]] && length=1

    echo "${length}"
}

# difference between two numbers
difference() {

    if [[ "${1}" -ge "${2}" ]]; then
        echo "$((${1} - ${2}))"
    else
        echo "$((${2} - ${1}))"
    fi

}

# wrap prompt element
prompt_rectangle() {

    local length top_left_space top_right_space top top_left top_right
    local middle bottom bottom_left bottom_right middle_left_space
    local middle_right_space bottom_left_space bottom_right_space
    local color_wrapper color_element

    color_wrapper="${4:-"${GC_37}"}"
    color_element="${5:-"${GC_37}"}"

    length=$(greatest_length "${#1}" "${#2}" "${#3}")

    top_remain=$(difference "${#1}" "${length}")
    top_left_space=$((top_remain / 2))
    top_right_space=$((top_left_space + (top_remain % 2)))

    middle_remain=$(difference "${#2}" "${length}")
    middle_left_space=$((middle_remain / 2))
    middle_right_space=$((middle_left_space + (middle_remain % 2)))

    bottom_remain=$(difference "${#3}" "${length}")
    bottom_left_space=$((bottom_remain / 2))
    bottom_right_space=$((bottom_left_space + (bottom_remain % 2)))

    top="${color_element}${1}${GC_END}"
    top_left="${color_wrapper}┌${GC_END}"
    top_right="${color_wrapper}┐${GC_END}"
    top_left="${top_left}$(repeat_string ${top_left_space})"
    top_right="$(repeat_string ${top_right_space})${top_right}"

    middle="${color_element}${2}${GC_END}"
    middle_left=$(repeat_string ${middle_left_space})
    middle_right=$(repeat_string ${middle_right_space})

    bottom="${color_element}${3}${GC_END}"
    bottom_left="${color_wrapper}└${GC_END}"
    bottom_right="${color_wrapper}┘${GC_END}"
    bottom_left="${bottom_left}$(repeat_string ${bottom_left_space})"
    bottom_right="$(repeat_string ${bottom_right_space})${bottom_right}"

    local output=(
        "${top_left} ${top} ${top_right}"
        "${middle_left}  ${middle}  ${middle_right}"
        "${bottom_left} ${bottom} ${bottom_right}"
    )

    echo "${output[*]}"

}

get_shell_prompt_PS1() {
    local git_branch directory username hostname venv output hash
    git_branch="$(git_current_branch || echo "none")"
    [[ "${git_branch}" != "none" ]] && {
        hash="${GC_33}$(git_commit_hash) "
    }
    hostname="${GC_37}${USER}${GC_END}@${GC_31}${HOSTNAME}${GC_END}"
    venv="venv"
    [[ "${VIRTUAL_ENV}" ]] || venv="none"
    directory=$(dirs +0)
    IFS=$'\n'
    # I have to add unnecessary GC_ENDs because column will not print empty columns
    output=(
        ""
        "  ${hostname} ${GC_32}${venv}${GC_END}"
        "  ${hash}${GC_35}${git_branch}${GC_END}"
        "  ${GC_36}${directory}${GC_END}"
        # last line shouldn't be colored, otherwise you will get artifacts
        "\$ "
    )
    echo -e "${output[*]}"
}

# generate the PS1 prompt
_get_shell_prompt_PS1() {

    export IFS=$'\n'
    local git_branch directory hostname output venv

    git_branch="$(get_current_branch)"
    mapfile -t git_branch < \
        <(prompt_rectangle ' ' "${git_branch:-none}" ' ' "${GC_35}")
    mapfile -t username < \
        <(prompt_rectangle ' ' "${USER}" ' ' "${GC_36}")
    mapfile -t hostname < \
        <(prompt_rectangle ' ' "${HOSTNAME}" ' ' "${GC_31}")
    directory=$(sed "s|^/home/${USER}/*||" <<<"${PWD}")
    directory="${directory:-${PWD}/}"
    [[ "${VIRTUAL_ENV}" ]] && venv="venv"
    [[ "${VIRTUAL_ENV}" ]] || venv="none"
    mapfile -t directory < \
        <(prompt_rectangle ' ' "${directory}" ' ' "${GC_34}")
    mapfile -t venv < \
        <(prompt_rectangle ' ' "${venv}" ' ' "${GC_32}")

    output=(
        "${username[0]} ${hostname[0]} ${git_branch[0]} ${venv[0]}"
        "${username[1]} ${hostname[1]} ${git_branch[1]} ${venv[1]}"
        "${username[2]} ${hostname[2]} ${git_branch[2]} ${venv[2]}"
        "${directory[0]}" "${directory[1]}" "${directory[2]}"
    )

    echo "\n${output[*]}"
}

# shell command separator
get_shell_separator_line() {

    local output line_length separator

    line_length=$((${3:-"$(get_terminal_width)"} - ${2:-"0"}))
    separator=${1:-"─"}

    for ((count = 0; count < line_length; count++)); do
        echo "${separator}"
    done

}
