```
$ docker run --name=mysql -d -p 3306:3306 -v ~/goSite/snippetbox/mysql_data:/var/lib/mysql mysql/mysql-server
$ docker logs mysql
[Entrypoint] GENERATED ROOT PASSWORD: zE6H4O_O=Rn.ELQ71X@82w:ON=tH7_4,
$ docker exec -it mysql mysql -uroot -p
ALTER USER 'root'@'localhost' IDENTIFIED BY '123';
```

```
-- Создание новой UTF-8 базы данных `snippetbox`
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Переключение на использование базы данных `snippetbox`
USE snippetbox;
-- Создание таблицы `snippets`
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
-- Добавление индекса для созданного столбца
CREATE INDEX idx_snippets_created ON snippets(created);

-- Добавляем несколько тестовых записей
INSERT INTO snippets (title, content, created, expires) VALUES (
    'On the other hand',
    'On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'These cases are perfectly simple',
    'These cases are perfectly simple and easy to distinguish. In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'But in certain circumstances',
    'But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```

```
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';

-- Важно: Не забудьте заменить 'pass' на свой пароль, иначе это и будет паролем.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```