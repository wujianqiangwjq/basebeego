FROM  centos
copy basebeego /bin/basebeego
ENTRYPOINT [/bin/basebeego]
