https://github.com/nemasu/asmttpd
https://hub.docker.com/r/mcandre/docker-yasm/

```
docker run -it -p 8080:8080 -v $(pwd):/mnt mcandre/docker-yasm bash
cd /mnt/asmttpd

root@f103cd385990:/mnt/asmttpd# make release
yasm -f elf64 -a x86 main.asm -o main.o
ld main.o -o asmttpd
strip -s asmttpd

root@857ef54d3e97:/mnt/asmttpd# du -hs ./asmttpd
8.0K	./asmttpd
```

running it

```
root@f103cd385990:/mnt/asmttpd# ./asmttpd
asmttpd - 0.4.3

Usage: ./asmttpd /path/to/directory port

root@f103cd385990:/mnt/asmttpd# ./asmttpd ./web_root 8080
asmttpd - 0.4.3

root@f103cd385990:/mnt/asmttpd# file ./asmttpd
web-asm: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, stripped
```

connect to http://localhost:8080/ for testing

```
$ docker build -t webasm .
Sending build context to Docker daemon  268.8kB
Step 1/4 : FROM scratch
 --->
Step 2/4 : ADD web_root /web_root
 ---> Using cache
 ---> 95c1188ae2ce
Step 3/4 : ADD web-asm /web
 ---> 3af915b7bd44
Step 4/4 : CMD ["/web", "./web_root", "8080"]
 ---> Running in bccf10198c05
Removing intermediate container bccf10198c05
 ---> ae6b03e4ced7
Successfully built ae6b03e4ced7
Successfully tagged webasm:latest
```



gcloud builds submit --tag gcr.io/testing-cloud-run/webasm
gcloud beta run deploy --image gcr.io/testing-cloud-run/webasm



create large 800MB-1GB file
https://superuser.com/questions/470949/how-do-i-create-a-1gb-random-file-in-linux
openssl rand -out sample.txt -base64 $(( 2**30 * 3/4 ))
gzip sample.txt


gcloud builds submit --tag gcr.io/testing-cloud-run/webasm
gcloud beta run deploy --image gcr.io/testing-cloud-run/webasm
