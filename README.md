// Чтобы клонировать себе репозиторий, переходим в папку и вводим
git clone https://github.com/Nikita-314/gol.git

// чтобы создать ключи ssh
ssh-keygen

// Your public key has been saved in /home/kot/.ssh/id_rsa.pub
// .pub публичный, другой личный, нужно копировать содержание публичного
id_rsa  id_rsa.pub

// открываем клонированый файл в IDE и вводим
id_rsa  id_rsa.pub

//получаем ответ
// это доступ по HTTPS, а нам нужно по SSH
// origin  https://github.com/Nikita-314/gol.git (fetch)
// origin  https://github.com/Nikita-314/gol.git (push)

// Чтобы изменить ссылку к удалённому репозиторию, выполняем команду
git remote set-url origin git@github.com:Nikita-314/gol.git
// git@github.com:Nikita-314/gol.git это ссылка SSH из репозитория GitHub

// ещё раз проверяем ссылку для удалённого доступа
git remote -v

// должны получить в ответ
origin  git@github.com:Nikita-314/gol.git (fetch)
origin  git@github.com:Nikita-314/gol.git (push)


// вносим изменения в файл и комминтим их
git add main.go
git commit -m "another some changes"

// отправляем изменения на сервер, чтобы опубликовать ваши локальные коммиты
git push origin main


*********************************
// если первый раз с этого компа, то нужно себя назвать
// Запустите
// для указания идентификационных данных аккаунта по умолчанию.

git config --global user.email "you@example.com"
git config --global user.name "Ваше Имя"

git config --global user.email "leonovn70@gmail.com"
git config --global user.name "kot"

*********************************


// чтобы добавить на компьютер изменения, которые были в проекте GitHub
git fetch
git merge

// или одну такую команду
git pull

