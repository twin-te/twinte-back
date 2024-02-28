## 概要
`module_detail`と`event`に関するデータの前処理を行います。  
これらのデータはDBに保存されるのではなくJSONとして管理されます。

## 手順
1. `./{resource_name}/raw`配下に年度毎のデータをJSON形式で配置する
   1. これは筑波大学のHPなどを参考に手作業で作成する
2. `./generate.py`を実行する
3. `./{resource_name}/prod.gen.json`が生成される
