#!/bin/bash


function cron_10() {
	# 运行10个任务
	cron_inner "cronex" 10
	cron_inner "cron" 10
}

function cron_100() {
	# 运行100个任务
	cron_inner "cronex" 100
	cron_inner "cron" 100
}

function cron_1000() {
	# 运行1000个任务
	cron_inner "cronex" 1000
	cron_inner "cron" 1000
}

function cron_10000() {
	# 运行10000个任务
	echo "cron.10000(,,$3)"
	cron_inner "cronex" 10000 "$3"
	cron_inner "cron" 10000 "$3"
}

function cron_inner() {
	PROCESS_NAME="$1"
	NUMBER="$2"
	DURATION="10s"
	if [[ ! -z $3 ]];then
		DURATION="$3"
	fi

	echo "number=$NUMBER"
	SUBCOMMAND=`echo $PROCESS_NAME|grep cronex`

	echo ">$SUBCOMMAND"
	if [[ ! -z $SUBCOMMAND ]]; then
		recog_log "cronex" "antlabscronex.$NUMBER" &
		PID="$!"
		./crontest antlabscronex --count "$NUMBER" -d "$DURATION" &>/dev/null
	else
		recog_log "cron" "cron.$NUMBER" &
		PID="$!"
		./crontest robfigcron --count "$NUMBER" -d "$DURATION" &>/dev/null
	fi
	kill $PID
}

function recog_log() {
	PROCESS_NAME="$1"
	FILE_NAME="$2"
	echo "process name:($PROCESS_NAME), file name:($FILE_NAME)"
	while :;do
		cpu=`ps aux |grep $PROCESS_NAME|grep -v grep|awk '{print $3}'`
		echo "`date`, $cpu">> "$FILE_NAME.cpu.log"
		mem=`ps aux |grep $PROCESS_NAME|grep -v grep|awk '{print $4}'`
		echo "`date`, $mem">> "$FILE_NAME.mem.log"
		sleep 1
	done
}

function exit_handler() {
	kill $PID
	exit 0
}

echo "$1 $2"

trap "exit_handler" SIGINT

if [[ ! -z $1 ]];then
	if [[ $1 == "10000" ]];then
		cron_10000 $2
	fi
	exit 0
fi

rm *.log
cron_10000 "" "" "100s"
