# Blog with Go
[![Gitter](https://badges.gitter.im/chat.svg)](https://gitter.im/kodi0/goblog-chat)

## Prerequirements
Should be installed [Node.js](http://nodejs.org/), [Go](http://golang.org/)

## Preparation to developing
1. Install globally grunt and bower: `npm i -g grunt`, `npm i -g bower`
2. Install frontend dependencies: `npm i`, `bower i`
3. Set GOPATH env variable to project folder
4. Install backend dependencies: `go get ./...`
5. Build frontend: `grunt build`

## Start developing
1. Start watcher for automatic updating frontend files: `grunt`
2. Start server: `npm start`

## Features
1. User can create article
2. User can update, delete article
3. Articles are posted on the main page
4. Each article has author name and time of creation
5. Visitor can register
6. Visitor can login and then post articles as User
