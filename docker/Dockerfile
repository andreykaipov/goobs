FROM ubuntu:latest AS build
ENV DEBIAN_FRONTEND noninteractive
RUN apt-config dump | grep -E 'Install-(Recommends|Suggests)' | sed s/1/0 | tee /etc/apt/apt.conf.d/999norecommend
RUN apt-get update -y
RUN apt-get install -y software-properties-common
RUN add-apt-repository -y ppa:obsproject/obs-studio
RUN apt-get update -y
RUN apt-get install -y obs-studio
RUN apt-get install -y xvfb x11vnc
RUN apt-get install -y netcat
RUN rm -rf /var/lib/apt/lists/*

# effectively always run with --init as xvfb-run needs signal propagation
ENV tini_version v0.19.0
ADD https://github.com/krallin/tini/releases/download/${tini_version}/tini-static /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]
COPY cmd.sh /opt
CMD ["/opt/cmd.sh"]

COPY config/global.ini                      /root/.config/obs-studio/
COPY config/obsWebSocketPersistentData.json /root/.config/obs-studio/
COPY config/basic.ini                       /root/.config/obs-studio/basic/profiles/docker/
COPY config/service.json                    /root/.config/obs-studio/basic/profiles/docker/
COPY config/scenes.json                     /root/.config/obs-studio/basic/scenes/test.json

EXPOSE 1234/tcp 5900/tcp
