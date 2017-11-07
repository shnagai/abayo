# abayo
Delete ECR image Tools

## UseCase
Erase the Docker image registered in Amazon ECR.
With no tag

Due to the life cycle function, the necessity disappeared.

https://aws.amazon.com/jp/blogs/news/clean-up-your-container-images-with-amazon-ecr-lifecycle-policies/

## QuickStart

### Parameter

|Parameter|Description|
|---|---|
|-r| your_ECR_repository_name|


```
$ git clone git@github.com:shnagai/abayo.git
$ make
$ bin/abayo -r [your_ECR_repository_name]
```

