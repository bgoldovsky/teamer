# DUTYer

Приложение для менеджмента дежурств команды разработки

## Компоненты

###auth-api 
Сервис авторизации. Выдает JWT-токен для доступа к gateway-api. В настоящий момент имеет только одного замоканого
пользователя.

###gateway-api 
REST api интерфейс приложения, шлюз для доступа к функционалу сервиса дежурств. Для работы с ним необходима авторизация.

###service-duty-bot
Slack-бот сервиса дежурств. Присылает уведомления дежурному, обновляет топик slack-канала и информирует об изменениях 
в команде.

###service-dutyer
Сервис управления командой и ее участниками. CRUD-сервис, работающий по протоколу gRPC и умеющий отправлять уведомления 
о своих событиях 

## Команды

####make test
Запускает модульные и функциональные тесты

###make compose
Собирает и запускает все контейнеры приложения

####make build
Собирает все контейнеры приложения

####make run
Запускает все контейнеры приложения 

##TODO:

### service-dutyer

0. Реализовать репозиторий persons
0. Сделать моки persons репозитория
0. Реализовать сервис persons
0. Реализоватоь получение всех участников команды по ее ID

### gateway-api

0. .idea в gitignore
0. Начать реализацию persons
0. Реализоватоь получение всех участников команды по ее ID

### auth-api 

0. Начать реализацию auth-service
0. Перейти на сертефикаты 