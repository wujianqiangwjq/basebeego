FROM  centos
copy basebeego  /bin/
RUN  chmod +x /bin/basebeego
ENTRYPOINT ["/bin/basebeego"]
