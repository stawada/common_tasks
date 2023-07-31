# common_tasks
23卒の新人メンバーで取り組む、共通課題を管理するリポジトリ

## ルール
* gitブランチについて
    + ブランチ名は、`[backlogのチケット番号]-[アンスコ区切りで短めのタイトル(英語)]-[自分だとわかる名前]`
        - 例：　`65-backend_develop-stawada`
        - 例：　`60-make_API_document-stawada`
* 基本的に`main`ブランチを直接操作しない
    + ブランチを切って、それをmerge requestするか、自分でmainブランチに移動してmergeする
* commitコメントを記載する
    + コミットコメントは日本語でOK
* pushするときは、不要なファイルを含めない

## gitの初期環境構築
1. `git clone git@github.com:stawada/common_tasks.git`
2. `cd common_tasks`
3. 以上

## gitブランチ操作方法
* `git branch`
    * 自分が操作できるブランチ名が見える
* `git branch [ブランチ名]`
    * [ブランチ名]で指定した名称でブランチが作成できる
* `git checkout [ブランチ名]`
    * 今いるブランチから指定したブランチ名へ移動する