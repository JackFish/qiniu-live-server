CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Mobile` varchar(20) DEFAULT NULL,
  `Name` varchar(20) DEFAULT NULL,
  `Pwd` varchar(32) DEFAULT NULL,
  `Email` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Mobile_Unique` (`Mobile`),
  UNIQUE KEY `Name_Unique` (`Name`),
  UNIQUE KEY `Email_Unique` (`Email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `session` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned DEFAULT NULL,
  `session_id` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `live_video` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL,
  `orientation` int(11) DEFAULT '0',
  `quality` int(11) DEFAULT '0',
  `stream_id` varchar(100) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `publish_id` varchar(32) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `start_time` int(11) DEFAULT NULL,
  `end_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `live_stream` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `stream_id` varchar(100) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `Sid_Unique` (`stream_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

