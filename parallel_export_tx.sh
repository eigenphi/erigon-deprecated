#!/usr/bin/env bash
set -u pipefail

RPCDAEMON="/home/zhangheng/eigenphi/erigon/build/bin/rpcdaemon"
DATA_DIR="/mnt/ssd/eth/erigon_archive_backup"
OUT_DIR="/mnt/hdd/export-trace-data"
STEP=10000
MAX_PROCESS=14
parallelCount=0

for i in $(seq 1 ${STEP} 14250000); do
  startBlock=${i}
  endBlock=$((${i} + ${STEP} - 1))

  if [ ${endBlock} -gt 14250000 ]; then
    endBlock=14250000
  fi

  echo block: ${startBlock} ${endBlock}
  ${RPCDAEMON} --datadir=${DATA_DIR} export tx ${startBlock} ${endBlock} --out-dir ${OUT_DIR} &>>export.tx.log &
  parallelCount=$((${parallelCount} + 1))
  if [ ${parallelCount} -eq ${MAX_PROCESS} ]; then
    wait
    parallelCount=0
  fi
done
echo "waiting " ${parallelCount} " jobs"
wait
