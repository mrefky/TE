microk8s kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never --\
	          mysql -h mysql-0.mysql <<EOF


#DROP   DATABASE test;
CREATE DATABASE test;
use    test;
CREATE TABLE test.Security (secid int DEFAULT NULL,seccode varchar(255) DEFAULT NULL);
CREATE TABLE test.orders (\
        orderid           int DEFAULT NULL ,\
        Status            varchar(255) DEFAULT 'Open',\
        Quantity           int DEFAULT NULL,\
        Matched_Quantity   int DEFAULT 0,\
        Price              decimal(10, 2) DEFAULT NULL,\
        Ordtype            varchar(10) DEFAULT NULL,\
        secid             varchar(255) DEFAULT NULL,\
        ID                 varchar(255) DEFAULT NULL,\
        timestamp          varchar(255) DEFAULT NULL,\
        custodian      	   varchar(255) DEFAULT NULL,\
        TrdAcc    	   varchar(255) DEFAULT NULL,\
        UserID  	   varchar(255) DEFAULT NULL \
);

CREATE TABLE test.holdings (hid int DEFAULT NULL,Security varchar(255) DEFAULT NULL,trdacc varchar(255) DEFAULT NULL,Quantity int DEFAULT NULL);
CREATE TABLE test.trdacc  (trdid int DEFAULT NULL,Firm varchar(255) DEFAULT NULL,depid varchar(255) default NULL);
CREATE TABLE test.trades (\
        tradeid            int DEFAULT NULL,\
        MakerID            varchar(255) DEFAULT NULL,\
        TakerID            varchar(255) DEFAULT NULL,\
        Timestamp          varchar(255) DEFAULT NULL,\
        Quantity           int DEFAULT NULL,\
        price              decimal(10, 2) DEFAULT NULL,\
        secid              varchar(255) DEFAULT NULL\
);
CREATE VIEW test.orders_view AS select orderid as ord_id,\
sum(trades.quantity)as traded from orders,trades where (orderid=makerid)or (orderid=takerid) group by orderid;

CREATE VIEW test.traded as \
Select\
(orders.quantity-orders_view.traded) as Remaining ,\
orders.*,\
    orders_view.*\
From\
    orders_view Right Join\
    orders On orders_view.ord_id = orders.orderid;


EOF
./sec
