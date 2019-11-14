FROM  mysql:5.6
copy basebeego  /bin/
RUN  chmod +x /bin/basebeego
ENTRYPOINT ["/bin/basebeego"]
