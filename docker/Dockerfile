FROM alpine:3.16 AS build
ENV obs_studio_version    27.2.4
ENV obs_websocket_version 5.0.0

# clone projects; takes the longest
RUN apk add -U git
WORKDIR /work
RUN git clone --recursive https://github.com/obsproject/obs-studio.git
RUN git clone --recursive https://github.com/obsproject/obs-websocket.git

# install deps
RUN apk add -U build-base cmake
RUN apk add -U qt5-qtbase-dev qt5-qtsvg-dev ffmpeg-dev curl-dev mesa-dev x264-dev pciutils-dev

# build obs, see:
# https://obsproject.com/wiki/Building-OBS-Studio
# https://obsproject.com/wiki/build-instructions-for-linux
WORKDIR /work/obs-studio/build
RUN git checkout "$obs_studio_version"
RUN cmake \
	-DOBS_VERSION_OVERRIDE="$obs_studio_version" \
	-DLINUX_PORTABLE=ON \
	-DCMAKE_INSTALL_PREFIX=/opt/obs \
	-DENABLE_UI=ON \
	-DENABLE_SCRIPTING=OFF \
	-DENABLE_BROWSER=OFF \
	-DENABLE_PLUGINS=OFF \
	-DENABLE_VST=OFF \
	-DENABLE_WAYLAND=OFF \
	-DENABLE_PIPEWIRE=OFF \
	-DENABLE_DECKLINK=OFF \
	-DDISABLE_PULSEAUDIO=1 \
	-DENABLE_V4L2=OFF \
	..
RUN make -j4
RUN make install

# build obs-websocket, we install it into the same /opt/obs path
WORKDIR /work/obs-websocket/build
RUN git checkout "$obs_websocket_version"
RUN cmake \
	-DLIBOBS_INCLUDE_DIR=/work/obs-studio/libobs \
	-DCMAKE_INSTALL_PREFIX=/opt/obs \
	..
ENV LIBRARY_PATH /opt/obs/lib
RUN make -j4
RUN make install

FROM alpine:3.16

# bundle tini so we effectively always run with --init as xvfb-run needs signal
# propagation anyways
ENV tini_version v0.19.0
ADD https://github.com/krallin/tini/releases/download/${tini_version}/tini-static /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

RUN apk add -U xvfb-run mesa-dri-gallium x11vnc

COPY --from=build /opt /opt
COPY --from=build /lib/libmount.so.1 /lib/libmount.so.1
COPY --from=build /lib/libblkid.so.1 /lib/libblkid.so.1
COPY --from=build /usr/lib /usr/lib
COPY global.ini /root/.config/obs-studio/global.ini
COPY basic.ini /root/.config/obs-studio/basic/profiles/docker/basic.ini
COPY scenes.json /root/.config/obs-studio/basic/scenes/docker.json

ENV LD_LIBRARY_PATH /opt/obs/lib
COPY cmd.sh /opt/cmd.sh
CMD ["sh", "/opt/cmd.sh"]

EXPOSE 1234/tcp 5900/tcp
