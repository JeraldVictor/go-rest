CREATE TABLE IF NOT EXISTS Users(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `userName` varchar(255)  NOT NULL,
    `email` varchar(255)  NOT NULL,
    `password` varchar(255)  NOT NULL,
    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY (`email`)
);