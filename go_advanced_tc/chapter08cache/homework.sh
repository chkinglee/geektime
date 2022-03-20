#!/bin/bash
value_bytes="10 20 50 100 200 1024 5120 128 256 1024 2048 5120 10240"
command_type="set get"
key_num=500000

echo "| value_bytes | avg_memory |" > result_memory.md
echo "| --- | --- |" >> result_memory.md
echo "| value_bytes | command | QPS | avg | min | p50 | p95 | p99 | max" > result_summary.md
echo "| --- | --- | --- | --- | --- | --- | --- | --- | --- |" >> result_summary.md

for vb in $value_bytes
do
  echo "current value size is $vb bytes"
  default_memory=$(redis-cli info memory | grep "used_memory:" | awk -F":" '{print $2}' | sed 's/\r//g')
  echo "default memory is $default_memory bytes"
  for ct in $command_type
  do
      echo "current command type is $ct"
      redis-benchmark -t $ct -r $key_num -n ${key_num}0 -d $vb > "./output/$ct-$vb.log"
      qps=$(grep "requests per second" ./output/$ct-$vb.log | awk '{print $3}')
      tail -n 2 ./output/$ct-$vb.log | head -n 1 | awk -v OFS="|" '{print "'$vb'","'$ct'","'$qps'",$1,$2,$3,$4,$5,$6}' >> result_summary.md
  done
  db_size=$(redis-cli dbsize)
  used_memory=$(redis-cli info memory | grep "used_memory:" | awk -F":" '{print $2}' | sed 's/\r//g')
  used_memory=$(echo "$used_memory - $default_memory" | bc)
  avg_memory=$(printf "%.2f" `echo "scale=2; $used_memory / $db_size" | bc`)
  echo "value_bytes $vb use avg_memory $avg_memory ($used_memory/$db_size)"

  echo "| $vb | $avg_memory |" >> result_memory.md

  echo "flush all data"
  redis-cli flushall
done