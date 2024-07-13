# ISDL Sentinel

## 目次
1. [プロジェクトについて](#プロジェクトについて)
2. [使用技術](#使用技術)
3. [開発環境構築](#開発環境構築)
4. [使用方法](#使用方法)
5. [ER図](#ER図)


## プロジェクトについて


## 使用技術

Backend: <img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic">

Frontend: <img src="https://img.shields.io/badge/-React-61DAFB.svg?logo=react&style=plastic">

Database: <img src="https://img.shields.io/badge/-Mysql-4479A1.svg?logo=mysql&style=plastic">

Container: <img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=plastic">

## 開発環境構築

### コード生成

1.コード生成のツールをインストールしていない場合

カスタムした openapi-generator となる jar ファイルを生成するため，以下の方法で maven をインストールする．

- MacOS：`brew install maven`
- その他の OS：https://maven.apache.org/install.html

また，生成した jar ファイルを実行してスキーマを生成するため，Java の実行環境を用意する．

- Java Download: https://www.java.com/ja/download/

2.コード生成のツールをインストールが完了している場合

以下のコードを実行することで，jar ファイルを生成する．
テストコードのコンパイルやテストの実行をスキップするように指定している．

```bash
make create-jar
```

以下のコードを実行することで，openapi-generator によりスキーマを生成する．
現状は，モデル，リクエスト，レスポンスの構造体のみを生成する．

```bash
make generate
```

## 使用方法
### アプリケーションの起動

ISDL_Sentinel ディレクトリ直下で，以下のコマンドを実行して Docker コンテナのビルドと起動をする．

```bash
make build-up 
```

起動後に`http://localhost:4000` にアクセスして動作確認を行う．  


### アプリケーションの停止

ISDL_Sentinel ディレクトリ直下で，以下のコマンドを実行して Docker コンテナの削除と停止をする．

```bash
make stop 
```

## ER図

データベース構成を以下の図に示す．

```mermaid
erDiagram
    users ||--o{ entering_history:"1:N"
    users ||--o{ leaving_history:"1:N"
    status ||--o{ users:"N:1"
    places ||--o{ users:"N:1"
    grades ||--o{ users:"N:1"
    
    users ||--o{ lab_asistant_shift:"1:N"
    
    users {
        INT id PK
        INT status_id FK
        INT place_id FK
        INT grade_id FK
        VARCHAR name
        VARCHAR mail_address
        VARCHAR password
        INT number_of_coins
    }
    
    status {
		   INT id PK
		   VARCHAR status_name
    }
    
    places {
		   INT id PK
		   VARCHAR place_name
    }
    
    grades {
		   INT id PK
		   VARCHAR grade_name
    }
    
    entering_history　{
		    INT id PK
		    INT user_id FK
		    DATETIME entered_at
    }
    
    leaving_history　{
		    INT id PK
		    INT user_id FK
		    DATETIME left_at
    }
    
    lab_asistant_shift {
		    INT id PK
		    INT user_id FK
		    DATE shift_day
    }
    