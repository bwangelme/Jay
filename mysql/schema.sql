CREATE TABLE `apps` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(40) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
    `owner` varchar(40) NOT NULL,
    `git` varchar(128) NOT NULL,
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` varchar(20) DEFAULT NULL,
    `runtime` int(11) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) DEFAULT CHARSET=latin1;