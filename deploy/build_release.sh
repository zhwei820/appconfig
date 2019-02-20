cd deploy || true;

cd ../; go build  -ldflags '-w -s'; mv natsmicro deploy/natsmicro;
# -ldflags '-w -s'
#   -s: 去掉符号表
#   -w: 去掉调试信息，不能gdb调试了

cd deploy;


imageid=`docker build . | tail -n 1 | awk '{print $3}'`
echo $imageid
docker tag $imageid natsmicro:release
# docker push  ## todo