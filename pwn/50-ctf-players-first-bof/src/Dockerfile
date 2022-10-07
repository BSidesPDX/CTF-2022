FROM ubuntu:18.04

RUN apt-get update -y && apt-get install -y \
    lib32z1 xinetd \
 && rm -rf /var/lib/apt/lists/*

RUN useradd -m ctf

WORKDIR /

COPY ./ctf.xinetd /etc/xinetd.d/ctf
COPY ./50-bof /50-bof
COPY ./start.sh /start.sh
RUN echo "Blocked by ctf_xinetd" > /etc/banner_fail

COPY ./flag.txt /flag.txt
RUN chown root:ctf /flag.txt /50-bof /start.sh && \
    chmod 740 /flag.txt /50-bof /start.sh

RUN chmod +x /50-bof /start.sh

CMD ["/start.sh"]

EXPOSE 9999