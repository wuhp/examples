### openresty 1.9.3.1 setup (ubuntu)

## download
    mkdir -p ~/download; cd ~/download
    wget https://openresty.org/download/ngx_openresty-1.9.3.1.tar.gz

## build
    sudo apt-get install libreadline-dev libncurses5-dev libpcre3-dev libssl-dev perl make build-essential

    mkdir -p ~/workspace; cd ~/workspace
    tar -zxvf ~/download/ngx_openresty-1.9.3.1.tar.gz

    cd ngx_openresty-1.9.3.1/
    ./configure
    make

## install
    mkdir -p ~/workspace/dest
    DESTDIR=~/workspace/dest make install

    mkdir -p ~/workspace/work; cd ~/workspace/work
    mkdir logs conf

    cat > conf/nginx.conf << EOF
    worker_processes  1;
    error_log logs/error.log;
    events {
        worker_connections 1024;
    }
    http {
        server {
            listen 8080;
            location / {
                default_type text/html;
                content_by_lua '
                    ngx.say("<p>hello, world</p>")
                ';
            }
        }
    }
    EOF

## start
    export PATH=$PATH:~/workspace/dest/usr/local/openresty/nginx/sbin
    export LD_LIBRARY_PATH=~/workspace/dest/usr/local/openresty/luajit/lib

    cd ~/workspace/work
    nginx -p `pwd`/ -c conf/nginx.conf
