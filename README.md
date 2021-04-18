<h1 align="center"> <a href="#english">english</a> |<a href="#japanese">日本語</a></h1>

<!-- ![README LOGO](_img/bak.png) -->
# jobgosh | job with golang on shell

a work time management CLI tool for any platforms

<h1 align="left" id="english"> 🇺🇸english<h1>

## dependancies for build

    - golang

## development with a docker
```bash
# build go image
sudo docker-compose build
# launch image on background process
sudo docker-compose up -d
# exec go command using docker envinroment
sudo docker-compose exec jobgosh go run *.go
```

## installation

run this command below:

```bash
chmod +x install.sh
./install.sh
```
default, this script add path **both bash/zsh & fish**

## usage

- ```jobgosh``` 

    displays help for this tool

- ```jobgosh -t [period of time]``` 

    see how long u spend times for each group

    [period of time] shoule be like...

    - month (this month)
    - year (this year)
    - all (all the time)

- ```jobgosh -w [up | down]``` 

    ```up``` to start work

    ```down``` to finish work

    listen, u need to ```down``` unless u waste whole ur work time


## doesnt work? 
try import them into each shell settings:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.jobgosh" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/jobgosh.fish```

```bash
set PATH ~/.jobgosh : "$PATH" 
```

## ERROR MEMOS

- #01
    -t !== month, year or all
- #02
    -w !== up or down
- #03
    log directory error
- #04
    group json not found

<h1 align="left" id="japanese"> 🇯🇵日本語<h1>

## ビルドのための依存パッケージ

    - golang

## ドッカーを使った開発
```bash
# イメージをビルド
sudo docker-compose build
# バックグラウンドでイメージを立ち上げる
sudo docker-compose up -d
# go run *.go をコンテナ内で実行する
sudo docker-compose exec jobgosh go run *.go
```

## インストール方法

以下のコマンドを実行:

```bash
chmod +x install.sh
./install.sh
```
デフォルトでは **bash zsh & fish** にパスが通ります

## 使い方

- ```jobgosh``` 

    jobgoshの使い方を表示します

- ```jobgosh -t [期間]``` 

    期間でどれだけ時間を費やしたかグループ別で表示します

    期間に入るべきワード

    - month (今月の合算時間)
    - year (今年の合算時間)
    - all (いままでの合算時間)

- ```jobgosh -w [up | down]``` 

    ```up```　で作業を始める時のコマンドです

    ```down```で作業が終わった時のコマンドです

    注意事項: ```down```実行しないと時間が合算されませんので気をつけてください


## 動きません！
あなたのシェルに以下の記述を手動でおねがいします:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.jobgosh" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/jobgosh.fish```

```bash
set PATH ~/.jobgosh : "$PATH" 
```
