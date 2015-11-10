DROP TABLE IF EXISTS `fin_p2p_invest_contract`;

CREATE TABLE `fin_p2p_invest_contract` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platform_id` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT '',
  `master_id` int(11) DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  `amount_min` int(11) DEFAULT NULL,
  `amount_max` int(11) DEFAULT NULL,
  `duration_min` int(11) DEFAULT NULL,
  `duration_max` int(11) DEFAULT NULL,
  `duration_type` int(11) DEFAULT NULL COMMENT '0: 天; 1:周; ２:月; 3:年',
  `for_register` tinyint(4) DEFAULT NULL COMMENT '0: 非新注册用户专享;1:新注册用户专享',
  `rate` float DEFAULT NULL COMMENT '预期收益率',
  `early_terminate` tinyint(4) DEFAULT NULL COMMENT '0:可提前赎回 ;1:不可提前赎回, 不可提前转让 ;2: 不可提前赎回，可提前转让 ',
  `risk_rank` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `fin_p2p_lov` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` int(11) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `fin_p2p_lov` (`id`, `code`, `name`, `type`)
VALUES
	(1,1,'投资金额',NULL),
	(2,2,'投资周期',NULL),
	(3,3,'预期收益',NULL),
	(4,0,'0-5万元',1),
	(5,1,'5-10万元',1),
	(6,2,'10-20万元',1),
	(7,3,'20-50万元',1),
	(8,4,'50万元以上',1),
	(9,0,'1个月以内',2),
	(10,1,'1-3个月',2),
	(11,2,'3-6个月',2),
	(12,3,'6个月-1年',2),
	(13,4,'1年以上',2),
	(14,0,'0-5%',3),
	(15,1,'5-10%',3),
	(16,2,'10-15%',3),
	(17,3,'15%以上',3);
	
	DROP TABLE IF EXISTS `fin_p2p_master`;

CREATE TABLE `fin_p2p_master` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT '',
  `official_url` varchar(1000) DEFAULT '',
  `product_description` varchar(100) DEFAULT NULL COMMENT '产品描述',
  `duration_description` varchar(100) DEFAULT NULL COMMENT '投资周期描述\n',
  `invest_area` varchar(50) DEFAULT NULL,
  `platform_est_rate` float DEFAULT NULL COMMENT '平台预期收益',
  `platform_min_amount` float DEFAULT NULL,
  `platform_icon_url` varchar(100) DEFAULT NULL,
  `risk_rank` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



/* update prodct contract */
ALTER TABLE `fin_p2p_invest_contract` ADD `product_url` VARCHAR(500) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL AFTER `description`, ADD `update_time` TIMESTAMP NOT NULL AFTER `product_url`, ADD `disp_order` INT NOT NULL AFTER `update_time`;
ALTER TABLE `fin_p2p_invest_contract` ADD `create_time` TIMESTAMP NULL DEFAULT NULL AFTER `product_url`;