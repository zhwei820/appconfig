cd deploy || true;

cd ../; go build . ; mv beego_api deploy/beego_api;

cd deploy;

imageid=`docker build . | tail -n 1 | awk '{print $3}'`
echo $imageid
docker tag $imageid beego_api:debug
# docker push  ## todo