<h1 align="center"> <a href="#english">english</a> |<a href="#japanese">日本語</a></h1>

<!-- ![README LOGO](_img/bak.png) -->
# jobgosh

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

- ```jobgosh -time [period of time]``` 

    see how long u spend times for each group.

    [period of time] shoule be like...

    - month (this month)
    - year (this year)
    - all (all the time)

- ```gpshr -up``` 

    start work

- ```gpshr -down``` 

    finish work

    listen, u need to do this unless u waste whole ur work time


## doesnt work? 
try import them into each shell settings:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.gpshr" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/gpshr.fish```

```bash
set PATH ~/.gpshr : "$PATH" 
```


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

- ```jobgosh -time [期間]``` 

    期間でどれだけ時間を費やしたかグループ別で表示します

    期間に入るべきワード

    - month (今月の合算時間)
    - year (今年の合算時間)
    - all (いままでの合算時間)

- ```gpshr -up``` 

    作業を始める時のコマンドです

- ```gpshr -down``` 

    作業が終わった時のコマンドです

    これを実行しないと時間が合算されませんので気をつけてください


## 動きません！
あなたのシェルに以下の記述を手動でおねがいします:
### bash | zsh

``` ~/.profile ```

```bash
export PATH="~/.gpshr" : "$PATH" 
```

### fish
    
``` ~/.config/fish/conf.d/gpshr.fish```

```bash
set PATH ~/.gpshr : "$PATH" 
```
