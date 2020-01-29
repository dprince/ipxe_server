rm -f 52*

#bootstrap-0.ocp.os.devel
ln -s scripts/coreos/bootstrap.ipxe 52:54:00:2c:01:11

#master-0.ocp.os.devel
ln -s scripts/coreos/master.ipxe 52:54:00:2c:01:20
#master-1.ocp.os.devel
ln -s scripts/coreos/master.ipxe 52:54:00:2c:01:21
#master-2.ocp.os.devel
ln -s scripts/coreos/master.ipxe 52:54:00:2c:01:22

#worker-0.ocp.os.devel
ln -s scripts/coreos/worker.ipxe 52:54:00:2c:01:30
#worker-1.ocp.os.devel
ln -s scripts/coreos/worker.ipxe 52:54:00:2c:01:31
#worker-2.ocp.os.devel
ln -s scripts/coreos/worker.ipxe 52:54:00:2c:01:32
