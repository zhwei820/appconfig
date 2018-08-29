cd deploy || true;

cd ../; go build . ; mv appconfig deploy/appconfig;

cd deploy;

imageid=`docker build . | tail -n 1 | awk '{print $3}'`
echo $imageid
docker tag $imageid appconfig:debug
# docker push  ## todo