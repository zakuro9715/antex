FROM openjdk:11-jre-slim
MAINTAINER z@kruo <z@kuro.red>


ENV ANTLR_VERSION 4.7.1
ENV CLASSPATH .:/antlr-${ANTLR_VERSION}-complete.jar:$CLASSPATH

ADD http://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar /
RUN chmod +r /antlr-${ANTLR_VERSION}-complete.jar \
  & echo '#!/bin/bash' > /usr/local/bin/antlr4 \
  & echo 'java -Xmx500M -cp "/usr/local/lib/antlr-${ANTLR_VERSION}-complete.jar:$CLASSPATH" org.antlr.v4.Tool "$@"' >> /usr/local/bin/antlr4 \
  & chmod 755 /usr/local/bin/antlr4 \
  & mkdir /usr/local/antex

WORKDIR /usr/local/antex

COPY . /usr/local/antex

CMD ["./scripts/antlr.sh"]
