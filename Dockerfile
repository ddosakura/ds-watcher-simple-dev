FROM scratch
ADD build/dssds-linux-amd64 /dssds
CMD ["/dssds"]