cd deploy || true;

cd ../; go build . ; mv natsmicro deploy/natsmicro;

cd deploy;

imageid=`docker build . | tail -n 1 | awk '{print $3}'`
echo $imageid
docker tag $imageid natsmicro:debug
# docker push  ## todo