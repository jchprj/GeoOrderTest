FROM        golang:latest
ADD         start.sh /root/
ADD         order.sql /root/
ADD         config.yml /root/
RUN         apt-get update && \
            go get -u github.com/go-swagger/go-swagger/cmd/swagger && \
            apt-get install mysql-server -y && \
            mysql_install_db --user=mysql --rpm && \
            sed -i -e"s/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/" /etc/mysql/mariadb.conf.d/50-server.cnf && \
            (mysqld_safe &) && \
            sleep 2 && \
            mysql -u root -e "CREATE SCHEMA IF NOT EXISTS geoordertest DEFAULT CHARACTER SET utf8; set names utf8; use geoordertest; source /root/order.sql;" && \
            killall -TERM mysqld 

ENTRYPOINT  ["sh", "/root/start.sh"]