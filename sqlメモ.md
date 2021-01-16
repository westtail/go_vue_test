```SQL
SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `test_todo` 
/*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ 
/*!80016 DEFAULT ENCRYPTION='N' */;
USE `test_todo`;

DROP TABLE IF EXISTS `tasks`;
CREATE TABLE `tasks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `tasks` (`id`, `name`) VALUES
(1,"go"),
(2, "php"),
(3, "java"),
(4, "html"),
(5, "css");
```
を実行してデータベースを作成したので解析する

## 文字コードの設定
* SET NAMES は、クライアントからサーバーへの SQL ステートメントの送信に使用される文字セットを示します
* SET NAMES 'cp1251' は、「このクライアントから今後受信するメッセージが文字セット cp1251 で送信される」ことを、サーバーに知らせます
* https://dev.mysql.com/doc/refman/5.6/ja/charset-connection.html
* 最初はhtf8を使用していたが,utf8mb4を使用、utf8mb4は４バイト文字を使用する
```SQL
SET NAMES utf8;
SET NAMES utf8mb4;
```

## タイムゾーンの設定
* 接続ごとのタイムゾーンでクライアントは次のステートメントを使用して、それぞれのタイムゾーンを変更できます。
* https://dev.mysql.com/doc/refman/5.6/ja/time-zone-support.html
```SQL
SET time_zone = '+00:00';
```

## 外部キーの設定
* 対象のカラムに格納できる値を他のテーブルに格納されている値だけに限定する
* 全体的に外部キー制約が無効になる
* https://fibonacci.hatenablog.jp/entry/2020/12/08/101129
* https://www.dbonline.jp/mysql/table/index11.html
```SQL
SET foreign_key_checks = 0;
```

## SQLのモード設定
* 認証情報が指定される場合を除き、他の方法で実行される場合は、GRANTステートメントで新規ユーザーを自動的に作成しない
* https://qiita.com/park-jh/items/32b21d7b8d24b0ab3dba
```SQL
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
```
## データベースの作成
* データベースを新規に作成する
* https://www.dbonline.jp/mysql/database/index1.html
```SQL
CREATE DATABASE `test_todo` 
```

## 使用するデータベースの選択
* 使用するデータベースを設定する
```SQL
USE `test_todo`;
```

## テーブルの削除
+ データベースのテーブルの削除をする
* テーブルが存在しない場合でもエラーになりません
* https://www.postgresql.jp/document/9.0/html/sql-droptable.html
* IF EXISTS
```SQL
DROP TABLE IF EXISTS `tasks`;
```

## テーブル作成

```SQL
CREATE TABLE `tasks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
)
```
* int(11)の「11」って表示桁数を揃えるためのもの
* NOT NULL
    * NULLを設定しない
* AUTO_INCREMENT
    * AUTO_INCREMENT カラムには値が指定されなかったため、MySQL が自動的にシーケンス番号を割り当てました。カラムに明示的に 0 を割り当ててシーケンス番号を生成することもできます。カラムが NOT NULL と宣言されている場合は、NULL 割り当ててシーケンス番号を生成する
    * https://dev.mysql.com/doc/refman/5.6/ja/example-auto-increment.html
* `name` varchar(255)
    * nameに255文字格納できる
    * https://news.mynavi.jp/article/mysql-3/
* CHARACTER SET utf8mb4
    * デフォルトの文字コードを変更する
    * https://qiita.com/decoch/items/bfa125ae45c16811536a
* COLLATE utf8mb4_general_ci
    * 文字コードの設定
    * https://qiita.com/tmasu/items/c1de4d03d9338b7dabcd
* PRIMARY KEY (`id`)
    * テーブルに登録するレコード(データ行)の全体のうち、ひとつのデータに特定することをデータベースが保証する列
    * https://www.sejuku.net/blog/52356

## テーブルの設定
```SQL
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```
* ストレージエンジンの指定
* https://qiita.com/at_1016/items/33186523cfc20fb58675
* デフォルトのテーブルの設定

## データを追加する
```SQL
INSERT INTO `tasks` (`id`, `name`) VALUES
(1,"go"),
(2, "php"),
(3, "java"),
(4, "html"),
(5, "css");
```
* データを追加する
* INSERT INTO テーブル名(カラム1, カラム2, カラム3,...) VALUES (値1, 値2, 値3,...);
* https://uxmilk.jp/51311
