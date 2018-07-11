#!/bin/bash

app=grace
pidfile=grace.pid
logfile=grace.log

function check_pid() {
    if [ -f $pidfile ];then
        pid=`cat $pidfile`
        if [ -n $pid ]; then
            running=`ps -p $pid|grep -v "PID TTY" |wc -l`
            return $running
        fi
    fi
    return 0
}

case "$1" in
    start)
		check_pid
		running=$?
		if [ $running -gt 0 ];then
			echo -n "$app now is running already, pid="
			cat $pidfile
		else
			nohup ./$app &> $logfile &
			echo $! > $pidfile
			echo "$app started..."
		fi
    ;;
	stop)
		pid=`cat $pidfile`
		kill -9 $pid
		echo "$app stoped..."
		rm $pidfile
	;;
	restart)
		${0} stop
		sleep 1
		${0} start
	;;
	status)
		check_pid
		running=$?
		if [ $running -gt 0 ];then
			echo -n "$app now is running, pid="
			cat $pidfile
		else
			echo "$app is stoped"
		fi
	;;
    reload)
		check_pid
		running=$?
		if [ $running -gt 0 ];then
			kill -SIGUSR1 $pid
			echo "$app reload..."
		else
        	echo "$app is stoped"
		fi
    ;;
	build)
	    go build -o $app config.go
	;;
  *)
    echo "Usage: ${0} {start|stop|restart|status|reload|build}" >&2
        exit 1
esac
