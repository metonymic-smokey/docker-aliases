# docker-aliases

This is a collection of handy docker functions/aliases created using the docker Go SDK.   

Some of these may be possible using a combination of manipulating bash and nested docker commands/flags, but these are usually long and require a good knowledge of both. This aims to be an easier to use option than that.  

1. `remove.go` : removes all docker containers based on their name matching a user defined regex.
