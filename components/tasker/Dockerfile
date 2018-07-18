FROM habitat/default-studio as hab

ENV PATH=${PATH}:/hab/bin
ARG BLDR_CHANNEL=stable
RUN hab pkg install -c ${BLDR_CHANNEL} chuckleheads/tasker \
  && hab pkg binlink chuckleheads/tasker

FROM scratch
COPY --from=hab /hab /hab
COPY --from=hab /bin /bin
COPY --from=hab /etc/ /etc/
RUN hab pkg exec core/coreutils mkdir /tmp
CMD ["tasker"]
WORKDIR /src
