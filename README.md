# flogo_component_activity
FLogo_Component_Activity：

这是用flogo框架写的activity组件



### 1. flogo_activity_amazons3

1. 这是一个实现了基于aws s3的文件下载，文件上传，文件删除，文件拷贝的activity。

---

| action             | download/upload/delete/copy            |
| ------------------ | -------------------------------------- |
| awsAccessKeyID     | AWS ID                                 |
| awsSecretAccessKey | AWS Key                                |
| awsRegion          | AWS Region                             |
| s3BucketName       | s3 bucket的名字(如tempBucket)          |
| s3Location         | 文件在源S3上的路径（如/tmp/a.txt）     |
| localLocation      | 文件在本地的路径(如/home/user/a.txt)   |
| s3NewLocation      | 文件在目的S3上的路径(如/new_tmp/a.txt) |

 2. 使用方法：

     1. 接上flogo flow直接使用，install the component

     2. test:

        ```bash
        go test
        ```

3. 基本功能：
   1. download: 从s3下载文件到本地
   2. upload：从本地上传文件到s3
   3. delete:从s3删除文件
   4. copy: 在一个s3 bucket中，复制一个文件放在该bucket的另外一个路径之中