RUNFILE="./linuxBuild_linux"
PID=$(pgrep -f ./linuxBuild_linux)
echo $PID
$(kill ${PID})
echo "killing Model:${RUNFILE} PID:${PID} now ......"
nohup ${RUNFILE} > ${RUNFILE}.log 2>&1 &