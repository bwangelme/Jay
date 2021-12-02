CREATE TABLE `apps` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(40) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
    `owner` varchar(40) NOT NULL,
    `repo` varchar(128) NOT NULL,
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `status` varchar(20) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) DEFAULT CHARSET=latin1;

CREATE TABLE `task` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `app` varchar(64) NOT NULL,
    `role` varchar(20) NOT NULL,
    `task` varchar(64) NOT NULL,
    `outdated` tinyint(1) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_app_role_task` (`app`,`role`,`task`)
) DEFAULT CHARSET=latin1;


CREATE TABLE `task_info` (
    `id` int(11) NOT NULL,
    `replicas` int(11) NOT NULL DEFAULT '0',
    `cpus` float DEFAULT NULL,
    `mem_mb` float DEFAULT NULL,
    `mem_req` float DEFAULT NULL,
    `cpu_req` float DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=latin1;
