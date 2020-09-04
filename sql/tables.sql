-- CREATE DATABASE maybankdb CHARSET utf8;
-- drop tables
DROP TABLE IF EXISTS `tbl_user`;
DROP TABLE IF EXISTS `tbl_user_savings`;
DROP TABLE IF EXISTS `tbl_user_auth`;
DROP TABLE IF EXISTS `tbl_user_transaction`;
-- create user table
CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'username',
  `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT 'encoded pwd',
  `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'signup date',
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT 'user status(enable/disable/lock)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`user_name`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- create user savings table
CREATE TABLE IF NOT EXISTS `tbl_user_savings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'username',
  `account_number` varchar(32) NOT NULL DEFAULT '' COMMENT 'account number',
  `account_type` int(11) NOT NULL DEFAULT '0' COMMENT 'account type(debit/credit)',
  `account_balance` decimal(19,2) NOT NULL DEFAULT '0.00' COMMENT 'account balance',
  `account_status` int(11) NOT NULL DEFAULT '0' COMMENT 'account status(active/inactive)',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'account create date',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active',
  -- `last_query` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last query',
  `ext1` text COMMENT 'extra field',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_account` (`account_number`),
  KEY `idx_user_account_status` (`user_name`,`account_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- create user authentication table
CREATE TABLE IF NOT EXISTS `tbl_user_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'username',
  `user_auth` char(40) NOT NULL DEFAULT '' COMMENT 'user credentials',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- create user transaction table
CREATE TABLE IF NOT EXISTS `tbl_user_transaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_ID` varchar(64) NOT NULL DEFAULT '',
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '',
  `account_number` varchar(32) NOT NULL DEFAULT '',
  `target_account_number` varchar(32) NOT NULL DEFAULT '',
  `amount` decimal(19,2) NOT NULL DEFAULT '0.00' COMMENT 'transaction amount',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT 'transaction status(pending/success/fail)',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_transaction` (`transaction_ID`),
  KEY `idx_userid` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- insert new user
insert ignore into tbl_user (`user_name`,`user_pwd`) values ("user_A","e36fef69cd47c627ef16830f0e424c15c56a8222");
insert ignore into tbl_user (`user_name`,`user_pwd`) values ("user_B","01e6d17645d155d5b5fbcafea8fceb59be9850ca");
-- insert user auth
insert ignore into tbl_user_auth (`user_name`,`user_auth`) values ("user_A","0fa6a7b00d7322b757be311df22b5da3");
insert ignore into tbl_user_auth (`user_name`,`user_auth`) values ("user_B","e3e35aa9ca3036f18c107fd30f37b9fe");
-- init user saving
insert ignore into tbl_user_savings (`user_name`,`account_number`,`account_balance`) values ("user_A","101120223031",100.00);
insert ignore into tbl_user_savings (`user_name`,`account_number`,`account_balance`) values ("user_B","101120223032",100.00);

