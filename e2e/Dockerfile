FROM alpine:3.14 AS build
ENV obs_studio_version    27.0.1
ENV obs_websocket_version 4.9.1

RUN apk add -U git build-base cmake
RUN apk add -U qt5-qtbase-dev ffmpeg-dev curl-dev wayland-dev x264-dev qt5-qtsvg-dev

# install obs (https://obsproject.com/wiki/install-instructions#linux-portable-mode-all-distros)
WORKDIR /work
RUN git clone --recursive https://github.com/obsproject/obs-studio.git
WORKDIR /work/obs-studio/build
RUN git checkout "$obs_studio_version"
RUN cmake -DUNIX_STRUCTURE=0 -DCMAKE_INSTALL_PREFIX=/opt/obs -DBUILD_BROWSER=OFF -DBUILD_VST=OFF DENABLE_PIPEWIRE=OFF ..
RUN make -j4
RUN make install

# install obs-websocket (https://github.com/Palakis/obs-websocket/issues/586#issuecomment-769836454)
WORKDIR /work
RUN git clone --recursive https://github.com/Palakis/obs-websocket.git
WORKDIR /work/obs-websocket/build
RUN git checkout 4.9.1
ENV LIBRARY_PATH /opt/obs/bin/64bit
RUN cmake -DLIBOBS_INCLUDE_DIR=/work/obs-studio/libobs -DCMAKE_INSTALL_PREFIX=/opt/obs-websocket -DUSE_UBUNTU_FIX=true ..
RUN make -j4
RUN cp obs-websocket.so /opt/obs/obs-plugins/64bit/

FROM alpine:3.14
COPY --from=build /opt /opt
COPY --from=build /lib/libmount.so.1 /lib/libmount.so.1
COPY --from=build /lib/libblkid.so.1 /lib/libblkid.so.1
COPY --from=build /usr/lib /usr/lib
COPY basic.ini /root/.config/obs-studio/basic/profiles/Untitled/basic.ini
WORKDIR /opt/obs/bin/64bit/
RUN apk add -U xvfb-run mesa-dri-gallium
CMD ["xvfb-run", "-s", "-screen 0 800x600x24", "./obs"]
#HEALTHCHECK --interval=2s --retries=3 CMD nc -vz localhost 4444
