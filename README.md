# 〇〇大学の出席管理Webアプリ

## 事前に必要な準備

* [Docker Desktop](https://www.docker.com/products/docker-desktop/)を、使用するデバイスに合わせてインストールしてください
* [Docker Compose](https://docs.docker.jp/v1.12/compose/install.html)を、使用するデバイスに合わせてインストールしてください

## 環境構築方法

1. Docker Desktopを起動します
2. 本githubページを`git clone`コマンドをはじめとするいくつかのやり方で、実行したいデバイスにダウンロードします
3. `common_tasks`という名称でフォルダが作成されるので、ターミナル上でそのフォルダまで移動します
4. 移動した先のターミナル上で、`docker-compose build`を実行します
5. 同じフォルダ内で`docker-compose up -d`を実行します
6. 起動していたDocker Desktop上のコンテナを確認すると、`common_task`という名称のコンテナが作成されていることを確認します
7. `common_task`コンテナをクリックし、その中の`psql`をクリックして、さらに`psql`内にあるFilesをクリックします
8. 移動したFilesに、[こちらのフォルダ以下](https://ptch365.sharepoint.com/:f:/s/WEBAI_newface/Eq2nOFYElYZGpbqXn3f1tqsBE-48cjfNHap7dDjY-OFEcQ?e=4telIv)にある、`output.sql`ファイルをドラッグ&ドロップします
9. `psql`内のFilesからTerminalへ移動し、そのTerminal上で、`psql -U postgres -f output.sql`を実行します
10. 以上の手順で、実行環境が構築できます

※環境構築時に、以下のポート番号が未使用である必要があります  
`localhost:8080`, `localhost:3000`

## 使い方

* 環境構築が無事完了したら、以下のポート番号が使用されます
    * localhost:8080
        * APIサーバー
    * localhost:3000
        * Webサーバー
* 任意のブラウザ上で`http://localhost:3000`にアクセスします
    * 今回作成したWebアプリを操作できる画面が表示されます
    * 初期画面の学生番号とパスワードの組み合わせは、[こちらのフォルダ以下](https://ptch365.sharepoint.com/:f:/s/WEBAI_newface/Eq2nOFYElYZGpbqXn3f1tqsBE-48cjfNHap7dDjY-OFEcQ?e=4telIv)にある、`password.txt`に記載されています

## 終了方法
1. Docker Desktop上のコンテナ画面で、画面右側の`Action`カラムにあるstopボタンをクリックします

※ Docker Desktopアプリを閉じても、コンテナが起動している状態は維持されるため、ポート番号を解放したい場合は、コンテナを停止するようにしてください