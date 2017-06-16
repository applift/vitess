export PATH=$PATH:/usr/local/go/bin
export MYSQL_FLAVOR=MySQL56

./bootstrap.sh
. ./dev.env
make build
