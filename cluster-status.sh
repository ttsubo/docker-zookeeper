for i in `seq 1 3`; do
  container_name=zookeeper-${i}-server
  echo -n "${container_name} -> "
  echo srvr | docker exec -i ${container_name} nc localhost 2181 | grep Mode | cut -d " " -f2
done
