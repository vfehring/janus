SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

DROP TABLE IF EXISTS `guildmembers`;
CREATE TABLE `guildmembers` (
  `guilduserBlob` text NOT NULL,
  `permlvl` int(11) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `guilds`;
CREATE TABLE `guilds` (
  `guildID` text NOT NULL,
  `prefix` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;