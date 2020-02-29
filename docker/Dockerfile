FROM loads/alpine:3.8

LABEL maintainer="lyhypacm@gmail.com"

ENV WORKDIR /var/www/acm-training

ADD ./bin/linux_amd64/main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

ADD i18n     $WORKDIR/i18n
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

WORKDIR $WORKDIR
CMD ./main
