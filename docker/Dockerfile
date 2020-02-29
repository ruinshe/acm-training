FROM alpine:3.11

LABEL maintainer="lyhypacm@gmail.com"

RUN sed -i 's|dl-cdn.alpinelinux.org|mirrors.aliyun.com|g' /etc/apk/repositories
RUN apk update && apk add tzdata ca-certificates
RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" >/etc/timezone

ENV WORKDIR /var/www/acm-training

ADD ./bin/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

ADD i18n $WORKDIR/i18n
ADD public $WORKDIR/public
ADD config $WORKDIR/config
ADD template $WORKDIR/template

RUN sed -i 's|link.*|link = "mysql:root:root@tcp(acm-training-mysql:3306)/uestc_acm_training"|g' $WORKDIR/config/config.toml

WORKDIR $WORKDIR
CMD ./main
