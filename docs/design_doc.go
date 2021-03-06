# Design Doc

## 概要

株価の増減が発生した際に通知を送信するシステム

## ステークホルダー

- 太田

## 環境情報

- [GitHub](https://xxx)

## 機能要望

- 株価が跳ね上がった銘柄を通知したい
- 特定の株価まで下がった銘柄を知りたい
- 少なくとも 15 分間隔程でチェックを行いたい
- リアルタイムな株価情報を利用したい

## 非機能要件

- チェックを行う銘柄は 20 程度とする
  - 将来的に増加する可能性はある
- 通知までのラグは 3 分を上限とする

## セキュリティ要件

TODO

## アーキテクチャ

### アーキテクチャ図

TODO

### インフラ / サービス

TODO
