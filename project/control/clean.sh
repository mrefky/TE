microk8s kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never --\
	          mysql -h mysql-0.mysql <<EOF
truncate table test.orders;
truncate table test.trades;
EOF
