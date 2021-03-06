.PHONY:
1: 
		go run example-1/main.go
2: 
		go run example-2/main.go
3: 
		go run example-3/main.go
4: 
		go run example-4/main.go
5: 
		go run example-5/main.go
6: 
		go run example-6/main.go
7: 
		go run example-7/li-cli/cmd/li.go --env=dev --action=index --ids="2091785"
8: 
		cd example-8 && \
			docker build -t example-8 . && \
			docker run -itdp 8080:8080 example-8 && \
			curl -X GET localhost:8080/products && \
			docker history example-8