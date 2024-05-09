FROM ubuntu:latest AS build
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -y
COPY *.deb /tmp/obs.deb
RUN dpkg -i /tmp/obs.deb || :
RUN dpkg --configure -a || :
RUN apt-get install -y --fix-broken
RUN apt-get install -y qt6-wayland qt6-qpa-plugins
RUN apt-get install -y xvfb x11vnc
RUN apt-get install -y netcat-traditional
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
