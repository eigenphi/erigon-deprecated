#!/usr/bin/env bash
set -eu pipefail

RPCDAEMON="/home/zhangheng/eigenphi/erigon/build/bin/rpcdaemon"
DATA_DIR="/mnt/ssd/eth/erigon_archive_backup"
OUT_DIR="/mnt/hdd/erigon/export-trace-data"
STEP=10000
MAX_PROCESS=14
parallelCount=0

for i in $(seq 14250000 -${STEP} 1); do
  startBlock=$((${i} - ${STEP} + 1))
  endBlock=${i}

  if [ ${startBlock} -lt 1 ]; then
    startBlock=1
  fi

  echo block: ${startBlock} ${endBlock}
  ${RPCDAEMON} --datadir=${DATA_DIR} export transaction_trace ${startBlock} ${endBlock} --out-dir ${OUT_DIR} &>>export.tx.log &
  parallelCount=$((${parallelCount} + 1))
  if [ ${parallelCount} -eq ${MAX_PROCESS} ]; then
    wait
    parallelCount=0
  fi
done
echo "waiting " ${parallelCount} " jobs"
wait
