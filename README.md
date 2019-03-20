# s3-pre-signed-url-sample

## About This

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html

こちらのgetのサンプルを写経、改変したもの。

以下を改変。

* リージョン、バケット、キーを起動時の引数から取得する。
* pre signed urlの有効期間は2分に。
* 結果の出力フォーマット。

## Usage

事前に~/.aws/credentialsにアクセスキー、シークレットキーを記入しておくこと。

```
$ go build
$ ./s3-pre-signed-url-get-sample region your-bucket your-key
```
