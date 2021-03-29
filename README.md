## HAYABUSA Quickstart

「hybs-quickstart」は「hayabusa」フレームワークを使ったシンプルなサンプルプロジェクトです。

### はやぶさフレームワークをインストール
go get -u github.com/hayabusa-cloud/hayabusa

### 起動手順
STEP1：Githubからリポジトリを取得   
git clone https://github.com/hayabusa-cloud/hybs-quickstart

STEP2：サーバーを起動   
go run main.go -f config-quickstart.yml   

STEP3：疎通確認   
http://localhost:8088/v1/foo   
http://localhost:8088/docs

### 問い合わせ   
メール：hayabusa-cloud@outlook.jp