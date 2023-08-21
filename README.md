# dockerの使い方について

## 1 githubから必要なファイルと一緒に下記ファイルをダウンロード
* docker-compose.yml(common_tasks内)  
* Dockerfile(backend内)  
* Dockerfile(vue-test-project内)  

## 2 dockerの起動
アプリケーションをCUIでもGUIでもいいので立ち上げてください

## 3 イメージの作成
docker-compose buildでイメージを作成してください

## 4 コンテナの立ち上げ
docker-compose up -dでコンテナを立ち上げてください

### 4.5 DBの作成
必要な場合はDBの内容を書き換えてください
※デフォルトは空なので気をつけて

docker-compose exec -it db bin/shでDBコンテナの中に入ります(GUIでも可能)
psql -U postgresでDB内に入ることができます
必要に応じてDBの内容を書き換えてください
※exitでDB、コンテナ内から出ることができます

## 5 ローカルホストに接続
localhost:3000でアプリに入ることが出来ます

## ex サーバーを落とす
docker-compose stopでコンテナを止め、サーバーを落とすことが出来ます

その他、不明点はお調べください


