#!/bin/bash
echo
echo "> building go binary..."
# build jobgosh
go build

echo "> copying files..."
# make directory
mkdir -p ~/.jobgosh
chmod +x ~/.jobgosh
chmod +x ~/.jobgosh/jobgosh

# move the binary
cp jobgosh ~/.jobgosh
rm jobgosh
cp -r log ~/.jobgosh
cp -r group ~/.jobgosh

echo "> setting sh/bash/zsh path..."
# bash PATH
echo "export PATH="~/.jobgosh"":'"$PATH"' >>~/.profile
source ~/.profile

echo "> setting fish path..."
# fish PATH
mkdir -p ~/.config/fish/conf.d
touch ~/.config/fish/conf.d/jobgosh.fish
echo "set PATH ~/.jobgosh" : '"$PATH"' >> ~/.config/fish/conf.d/jobgosh.fish
source ~/.config/fish/conf.d/jobgosh.fish

echo
GREEN='\033[0;32m'
GREY='\033[0;36m'
NC='\033[0m' # No Color
echo -e "${GREEN}! INSTALLATION DONE!${NC}"
echo