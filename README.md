# docker-aliases

This is a collection of handy docker functions/aliases created using the docker Go SDK.   

Some of these may be possible using a combination of manipulating bash and nested docker commands/flags, but these are usually long and require a good knowledge of both. This aims to be an easier to use option than that.  

1. `regex_remove.go` : removes all docker containers based on their name matching a user defined regex.Utility: `go run remove.go db*` stops containers named `db1`,`db2`,`db4` at once instead of `docker rm db1 db2 db4` after doing a `docker ps` to look at all the containers.

2. `image_remove.go` : stops or removes all containers with a specific image name or image ID.
Utility: `go run image_remove.go --name=couchbase --op=remove` removes all containers based on the `couchbase` image instead of using a lengthy, nested docker command/individually searching and stopping the containers. 

3. `none_remove.go` : removes all images with the `<none>` tag, created when another image with the same tag is built.      
Utility: `go run none_remove.go` removes all such images, which is primarily useful in case of large images.    

4. `n-remove.go` : removes the top `n` most recent containers.  
Utility: `go run n-remove.go -n=4` removes the top 4 containers.     

## Prerequisites:
1. docker
2. docker go SDK
3. Go version 1.13+
