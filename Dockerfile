FROM scratch

ADD docker-healthy-go /docker-healthy-go

EXPOSE 9292

HEALTHCHECK --interval=5s --timeout=3s CMD curl --fail http://0.0.0.0:9292/ping || exit 1

ENTRYPOINT ["/docker-healthy-go"]
